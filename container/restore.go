package container

import (
	gocontext "context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/cedana/cedana/utils"
	"github.com/containerd/console"
	containerd "github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/cmd/ctr/commands/tasks"
	"github.com/containerd/containerd/log"
	"github.com/docker/docker/errdefs"
	rspec "github.com/opencontainers/runtime-spec/specs-go"
)

type RuncOpts struct {
	Root            string
	ContainerId     string
	Bundle          string
	SystemdCgroup   bool
	NoPivot         bool
	NoMountFallback bool
	NoNewKeyring    bool
	Rootless        string
	NoSubreaper     bool
	Keep            bool
	ConsoleSocket   string
	Detatch         bool
	PidFile         string
	PreserveFds     int
	Pid             int
	NetPid          int
	Notify          utils.Notify
}

func Restore(imgPath string, containerID string) error {

	err := containerdRestore(containerID, imgPath)

	if err != nil {
		return err
	}

	return nil

}

func containerdRestore(id string, ref string) error {
	ctx := gocontext.Background()
	logger := utils.GetLogger()

	logger.Info().Msgf("restoring container %s from %s", id, ref)

	containerdClient, ctx, cancel, err := newContainerdClient(ctx)
	if err != nil {
		return err
	}
	defer cancel()

	checkpoint, err := containerdClient.GetImage(ctx, ref)
	if err != nil {
		if !errdefs.IsNotFound(err) {
			return err
		}
		ck, err := containerdClient.Fetch(ctx, ref)
		if err != nil {
			return err
		}
		checkpoint = containerd.NewImage(containerdClient, ck)
	}

	opts := []containerd.RestoreOpts{
		containerd.WithRestoreImage,
		containerd.WithRestoreSpec,
		containerd.WithRestoreRuntime,
		containerd.WithRestoreRW,
	}

	ctr, err := containerdClient.Restore(ctx, id, checkpoint, opts...)
	if err != nil {
		return err
	}
	topts := []containerd.NewTaskOpts{}
	topts = append(topts, containerd.WithTaskCheckpoint(checkpoint))
	spec, err := ctr.Spec(ctx)
	if err != nil {
		return err
	}

	useTTY := spec.Process.Terminal
	// useTTY := true

	var con console.Console
	if useTTY {
		con = console.Current()
		defer con.Reset()
		if err := con.SetRaw(); err != nil {
			return err
		}
	}

	task, err := tasks.NewTask(ctx, containerdClient, ctr, "", con, false, "", []cio.Opt{}, topts...)
	if err != nil {
		return err
	}

	var statusC <-chan containerd.ExitStatus
	if useTTY {
		if statusC, err = task.Wait(ctx); err != nil {
			return err
		}
	}

	if err := task.Start(ctx); err != nil {
		return err
	}

	if !useTTY {
		return nil
	}

	if err := tasks.HandleConsoleResize(ctx, task, con); err != nil {
		log.G(ctx).WithError(err).Error("console resize")
	}

	status := <-statusC
	code, _, err := status.Result()
	if err != nil {
		return err
	}
	if _, err := task.Delete(ctx); err != nil {
		return err
	}
	if code != 0 {
		return errors.New("exit code not 0")
	}

	return nil
}

// TODO Temp solution to looping over external nvidia mounts
var nvidiaExternalMounts = map[string]string{
	"/proc/driver/nvidia/gpus/0000:01:00.0":                        "/proc/driver/nvidia/gpus/0000:01:00.0",
	"/run/nvidia-persistenced/socket":                              "/run/nvidia-persistenced/socket",
	"/usr/lib/firmware/nvidia/550.67/gsp_tu10x.bin":                "/usr/lib/firmware/nvidia/550.67/gsp_tu10x.bin",
	"/usr/lib/firmware/nvidia/550.67/gsp_ga10x.bin":                "/usr/lib/firmware/nvidia/550.67/gsp_ga10x.bin",
	"/usr/lib/x86_64-linux-gnu/libnvidia-nvvm.so.550.67":           "/usr/lib/x86_64-linux-gnu/libnvidia-nvvm.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libnvidia-allocator.so.550.67":      "/usr/lib/x86_64-linux-gnu/libnvidia-allocator.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libnvidia-ptxjitcompiler.so.550.67": "/usr/lib/x86_64-linux-gnu/libnvidia-ptxjitcompiler.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libnvidia-opencl.so.550.67":         "/usr/lib/x86_64-linux-gnu/libnvidia-opencl.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libcudadebugger.so.550.67":          "/usr/lib/x86_64-linux-gnu/libcudadebugger.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libcuda.so.550.67":                  "/usr/lib/x86_64-linux-gnu/libcuda.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libnvidia-cfg.so.550.67":            "/usr/lib/x86_64-linux-gnu/libnvidia-cfg.so.550.67",
	"/usr/lib/x86_64-linux-gnu/libnvidia-ml.so.550.67":             "/usr/lib/x86_64-linux-gnu/libnvidia-ml.so.550.67",
	"/usr/bin/nvidia-cuda-mps-server":                              "/usr/bin/nvidia-cuda-mps-server",
	"/usr/bin/nvidia-cuda-mps-control":                             "/usr/bin/nvidia-cuda-mps-control",
	"/usr/bin/nvidia-persistenced":                                 "/usr/bin/nvidia-persistenced",
	"/usr/bin/nvidia-debugdump":                                    "/usr/bin/nvidia-debugdump",
	"/usr/bin/nvidia-smi":                                          "/usr/bin/nvidia-smi",
}

var nvidiaExternalDevMounts = map[string]string{
	"/dev/nvidia0":          "/dev/nvidia0",
	"/dev/nvidia-uvm-tools": "/dev/nvidia-uvm-tools",
	"/dev/nvidia-uvm":       "/dev/nvidia-uvm",
	"/dev/nvidiactl":        "/dev/nvidiactl",
	"/dev/shm":              "/dev/shm",
}

func RuncRestore(imgPath string, containerId string, opts RuncOpts) error {
	var spec rspec.Spec

	configPath := opts.Bundle + "/config.json"

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		return err
	}

	if err := json.Unmarshal(data, &spec); err != nil {
		fmt.Println("Error decoding config.json:", err)
		return err
	}

	//Find where to mount to
	externalMounts := []string{}
	for _, m := range spec.Mounts {
		if m.Type == "bind" {
			externalMounts = append(externalMounts, fmt.Sprintf("mnt[%s]:%s", m.Destination, m.Source))
		}
	}

	// TODO need better way to find if process is gpu -> from our kv store higher up in stack
	if os.Getenv("CEDANA_GPU_ENABLED") == "true" {
		for _, m := range nvidiaExternalMounts {
			externalMounts = append(externalMounts, fmt.Sprintf("mnt[%s]:%s", m, m))
		}
		for _, d := range nvidiaExternalDevMounts {
			externalMounts = append(externalMounts, fmt.Sprintf("mnt[%s]:%s", d, d))
		}
	}

	criuOpts := CriuOpts{
		ImagesDirectory: imgPath,
		WorkDirectory:   "",
		External:        externalMounts,
		MntnsCompatMode: true,
		TcpClose:        true,
		Notify:          opts.Notify,
	}

	runcOpts := &RuncOpts{
		Root:          opts.Root,
		ContainerId:   containerId,
		Bundle:        opts.Bundle,
		ConsoleSocket: opts.ConsoleSocket,
		PidFile:       "",
		Detatch:       opts.Detatch,
		NetPid:        opts.NetPid,
	}

	_, err = StartContainer(runcOpts, CT_ACT_RESTORE, &criuOpts)
	if err != nil {
		return err
	}
	return nil
}
