package jobservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"

	cts "github.com/cedana/cedana/pkg/api/services"
	"github.com/cedana/cedana/pkg/api/services/task"
)

type Runtime string

const (
	CrioRuntime       Runtime = "crio"
	RuncRuntime       Runtime = "runc"
	ContainerdRuntime Runtime = "containerd"
)

type CheckpointComplete struct {
	DumpResp string `json:"dump_resp"`
}

type CheckpointJob struct {
	PodName        string  `json:"sandbox_name" validate:"required"`
	ContainerName  string  `json:"container_name" validate:"required"`
	Namespace      string  `json:"namespace" validate:"required"`
	ImageRefName   string  `json:"image_ref" validate:"required"`
	Runtime        Runtime `json:"runtime" validate:"required"`
	CheckpointPath string  `json:"checkpoint_path,omitempty"`
}

type RestoreJob struct {
	PodName        string  `json:"sandbox_name" validate:"required"`
	ContainerName  string  `json:"container_name" validate:"required"`
	ImageRefName   string  `json:"image_ref" validate:"required"`
	Namespace      string  `json:"namespace" validate:"required"`
	Runtime        Runtime `json:"runtime" validate:"required"`
	CheckpointPath string  `json:"checkpoint_path,omitempty"`
}

func RegisterJobs(r *jobs.Runner, c *cts.ServiceClient) {
	r.Register("checkpoint", func(ctx context.Context, b []byte) error {
		cj := CheckpointJob{}
		err := json.Unmarshal(b, &cj)
		if err != nil {
			return err
		}
		log.Printf("checkpointing (%s) %s %s\n", cj.PodName, cj.ContainerName, cj.ImageRefName)
		// check runtime
		switch cj.Runtime {
		case "crio":
			// TODO: should we auto-detect this too??
			// I think this can never change with containerd
			runcRoot := "/run/runc"
			ctrByNameArgs := &task.RuncQueryArgs{
				Root:           runcRoot,
				Namespace:      cj.Namespace,
				ContainerNames: []string{cj.ContainerName},
				SandboxNames:   []string{cj.PodName},
			}
			ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			if runcDiscoveryErr != nil {
				return runcDiscoveryErr
			}
			// note: we assume we have atleast 1 container
			runcId := ctrResp.Containers[0].ID
			if runcId == "" {
				return fmt.Errorf("empty runcId")
			}
			crioDumpArgs := &task.CRIORootfsDumpArgs{
				ContainerID: runcId,
				// storage containers
				ContainerStorage: "/var/lib/containers/storage/overlay-containers",
				// if this is same for two concurrent requests there would be issues
				Dest: "/tmp/test",
			}
			_, err := c.CRIORootfsDump(ctx, crioDumpArgs)
			if err != nil {
				return err
			}
			// TODO: push the root image for both crio and containerd rootfs
		case "containerd":
			// TODO: should we auto-detect this too??
			// I think this can never change with containerd
			runcRoot := "/run/containerd/runc"
			ctrByNameArgs := &task.RuncQueryArgs{
				Root:           runcRoot,
				Namespace:      cj.Namespace,
				ContainerNames: []string{cj.ContainerName},
				SandboxNames:   []string{cj.PodName},
			}
			ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			if runcDiscoveryErr != nil {
				return runcDiscoveryErr
			}
			// note: we assume we have atleast 1 container
			runcId := ctrResp.Containers[0].ID
			if runcId == "" {
				return fmt.Errorf("empty runcId")
			}
			dumpArgs := &task.ContainerdRootfsDumpArgs{
				ContainerID: runcId,
				Namespace:   cj.Namespace,
				// we use the default one
				// TODO: add sock detection code
				// cat /etc/crictl.yaml | cut -d ' ' -f 2
				// then trim prefix `unix://`
				Address:  "/run/containerd/containerd.sock",
				ImageRef: cj.ImageRefName,
			}
			resp, err := c.ContainerdRootfsDump(ctx, dumpArgs)
			if err != nil {
				return err
			}

			respBody, err := json.Marshal(resp)
			if err != nil {
				// resp is an error, we've likely failed somewhere upstream
				respBody = []byte(fmt.Sprintf("error marshalling response: %v", err))
			}

			completeReq := CheckpointComplete{DumpResp: string(respBody)}
			_, _ = json.Marshal(completeReq)
		case "runc":
			runcRoot := "/default/runc"
			ctrByNameArgs := &task.RuncQueryArgs{
				Root:           runcRoot,
				Namespace:      cj.Namespace,
				ContainerNames: []string{cj.ContainerName},
				SandboxNames:   []string{cj.PodName},
			}
			ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			if runcDiscoveryErr != nil {
				return runcDiscoveryErr
			}
			// note: we assume we have atleast 1 container
			runcId := ctrResp.Containers[0].ID
			if runcId == "" {
				return fmt.Errorf("empty runcId")
			}
			criuOpts := &task.CriuOpts{
				ImagesDirectory: cj.CheckpointPath,
				WorkDirectory:   "",
				LeaveRunning:    true,
				TcpEstablished:  false,
			}
			var dumpType task.CRType
			if cj.CheckpointPath != "" {
				dumpType = task.CRType_LOCAL
			} else {
				dumpType = task.CRType_REMOTE
			}
			dumpArgs := &task.RuncDumpArgs{
				Root:           runcRoot,
				ContainerID:    runcId,
				CheckpointPath: cj.CheckpointPath,
				CriuOpts:       criuOpts,
				Type:           dumpType,
			}
			_, err := c.RuncDump(ctx, dumpArgs)
			if err != nil {
				return err
			}
		}
		return nil
	})

	r.Register("restore", func(ctx context.Context, b []byte) error {
		rj := RestoreJob{}
		err := json.Unmarshal(b, &rj)
		if err != nil {
			return err
		}
		log.Printf("restoring (%s) %s %s\n", rj.PodName, rj.ContainerName, rj.ImageRefName)
		// check runtime
		// TODO: handle runtime
		switch rj.Runtime {
		case "crio":
		case "runc":
			runcRoot := "/default/runc"
			ctrByNameArgs := &task.RuncQueryArgs{
				Root:           runcRoot,
				Namespace:      rj.Namespace,
				ContainerNames: []string{rj.ContainerName},
				SandboxNames:   []string{rj.PodName},
			}
			ctrResp, err := c.RuncQuery(ctx, ctrByNameArgs)
			if err != nil {
				return err
			}
			runcId := ctrResp.Containers[0].ID
			pausePidArgs := &task.RuncGetPausePidArgs{
				BundlePath: ctrResp.Containers[0].BundlePath,
			}
			pidResp, _ := c.RuncGetPausePid(ctx, pausePidArgs)
			netPid := pidResp.PausePid
			if runcId == "" {
				return fmt.Errorf("Failed to locate runc container.")
			}
			runcOpts := &task.RuncOpts{
				Root:          runcRoot,
				Bundle:        ctrResp.Containers[0].BundlePath,
				ConsoleSocket: "",
				Detach:        true,
				NetPid:        int32(netPid),
			}

			var restoreType task.CRType

			if rj.CheckpointPath != "" {
				restoreType = task.CRType_LOCAL
			} else {
				restoreType = task.CRType_REMOTE
			}

			id := uuid.New()
			runcRestoreArgs := &task.RuncRestoreArgs{
				ImagePath:   rj.CheckpointPath,
				ContainerID: id.String(),
				IsK3S:       true,
				Opts:        runcOpts,
				Type:        restoreType,
			}
			_, err = c.RuncRestore(ctx, runcRestoreArgs)

		case "containerd":
		}
		return nil
	})
}

func Restore(q *goqite.Queue) func(echo.Context) error {
	return func(e echo.Context) error {
		log.Println("Restore job enqueue")
		var rj RestoreJob
		if err := e.Bind(&rj); err != nil {
			return fmt.Errorf("error parsing body: %v", err)
		}
		if err := e.Validate(&rj); err != nil {
			return err
		}
		b, err := json.Marshal(rj)
		if err != nil {
			return err
		}
		if err := jobs.Create(context.Background(), q, "restore", b); err != nil {
			return err
		}
		return nil
	}
}

func Checkpoint(q *goqite.Queue) func(echo.Context) error {
	return func(e echo.Context) error {
		log.Println("Checkpoint job enqueue")
		var cj CheckpointJob
		if err := e.Bind(&cj); err != nil {
			return fmt.Errorf("error parsing body: %v", err)
		}
		if err := e.Validate(&cj); err != nil {
			return err
		}
		b, err := json.Marshal(cj)
		if err != nil {
			return err
		}
		if err := jobs.Create(context.Background(), q, "checkpoint", b); err != nil {
			return err
		}
		return nil
	}
}

func StartService(q *goqite.Queue, port uint64) (*echo.Echo, error) {
	e := echo.New()
	e.POST("/api/alpha1v1/checkpoint", Checkpoint(q))
	e.POST("/api/alpha1v1/restore", Restore(q))
	return e, nil
}
