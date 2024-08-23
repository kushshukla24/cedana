package jobservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"

	"github.com/cedana/cedana/pkg/api/services/task"

	_ "github.com/mattn/go-sqlite3"
)

func New() (*JobQueueService, error) {
	// sqlite queue
	db, err := sql.Open("sqlite3", ":memory:?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := goqite.Setup(context.Background(), db); err != nil {
		return nil, err
	}
	q := goqite.New(goqite.NewOpts{DB: db, Name: "jobs"})
	r := jobs.NewRunner(jobs.NewRunnerOpts{
		Limit:        1,
		PollInterval: 10 * time.Millisecond,
		Queue:        q,
	})
	runtime, err := figureOutHostRuntime()
	if err != nil {
		// couldn't figure out runtime
		return nil, err
	}
	jqs := &JobQueueService{
		runner:  r,
		queue:   q,
		runtime: runtime,
	}
	jqs.RegisterJobs()

	return jqs, nil
}

func figureOutHostRuntime() (string, error) {
	fs, err := os.Open("/etc/crictl.yaml")
	if err != nil {
		return "", err
	}
	text, err := io.ReadAll(fs)
	if err != nil {
		return "", err
	}
	if strings.Contains(string(text), "containerd") {
		return "containerd", nil
	} else if strings.Contains(string(text), "crio") {
		return "crio", nil
	} else {
		return "runc", nil
	}
}

type JobQueueService struct {
	queue   *goqite.Queue
	runner  *jobs.Runner
	runtime string
}

func (jqs *JobQueueService) RegisterJobs() {
	jqs.runner.Register("checkpoint", func(ctx context.Context, b []byte) error {
		cj := &task.QueueJobCheckpointRequest{}
		err := json.Unmarshal(b, cj)
		if err != nil {
			return err
		}
		const ImageRefName = ""
		log.Printf("checkpointing (%s) %s %s\n", cj.PodName, cj.ContainerName, ImageRefName)
		// check runtime
		switch jqs.runtime {
		case "crio":
			// // TODO: should we auto-detect this too??
			// // I think this can never change with containerd
			// runcRoot := "/run/runc"
			// ctrByNameArgs := &task.RuncQueryArgs{
			// 	Root:           runcRoot,
			// 	Namespace:      cj.Namespace,
			// 	ContainerNames: []string{cj.ContainerName},
			// 	SandboxNames:   []string{cj.PodName},
			// }
			// ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			// if runcDiscoveryErr != nil {
			// 	return runcDiscoveryErr
			// }
			// // note: we assume we have atleast 1 container
			// runcId := ctrResp.Containers[0].ID
			// if runcId == "" {
			// 	return fmt.Errorf("empty runcId")
			// }
			// crioDumpArgs := &task.CRIORootfsDumpArgs{
			// 	ContainerID: runcId,
			// 	// storage containers
			// 	ContainerStorage: "/var/lib/containers/storage/overlay-containers",
			// 	// if this is same for two concurrent requests there would be issues
			// 	Dest: "/tmp/test",
			// }
			// _, err := c.CRIORootfsDump(ctx, crioDumpArgs)
			// if err != nil {
			// 	return err
			// }
			// TODO: push the root image for both crio and containerd rootfs
		case "containerd":
			// TODO: should we auto-detect this too??
			// I think this can never change with containerd
			// runcRoot := "/run/containerd/runc"
			// ctrByNameArgs := &task.RuncQueryArgs{
			// 	Root:           runcRoot,
			// 	Namespace:      cj.Namespace,
			// 	ContainerNames: []string{cj.ContainerName},
			// 	SandboxNames:   []string{cj.PodName},
			// }
			// ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			// if runcDiscoveryErr != nil {
			// 	return runcDiscoveryErr
			// }
			// // note: we assume we have atleast 1 container
			// runcId := ctrResp.Containers[0].ID
			// if runcId == "" {
			// 	return fmt.Errorf("empty runcId")
			// }
			// dumpArgs := &task.ContainerdRootfsDumpArgs{
			// 	ContainerID: runcId,
			// 	Namespace:   cj.Namespace,
			// 	// we use the default one
			// 	// TODO: add sock detection code
			// 	// cat /etc/crictl.yaml | cut -d ' ' -f 2
			// 	// then trim prefix `unix://`
			// 	Address:  "/run/containerd/containerd.sock",
			// 	ImageRef: cj.ImageRefName,
			// }
			// resp, err := c.ContainerdRootfsDump(ctx, dumpArgs)
			// if err != nil {
			// 	return err
			// }
			//
			// respBody, err := json.Marshal(resp)
			// if err != nil {
			// 	// resp is an error, we've likely failed somewhere upstream
			// 	respBody = []byte(fmt.Sprintf("error marshalling response: %v", err))
			// }
			//
			// completeReq := CheckpointComplete{DumpResp: string(respBody)}
			// _, _ = json.Marshal(completeReq)
		case "runc":
			// runcRoot := "/default/runc"
			// ctrByNameArgs := &task.RuncQueryArgs{
			// 	Root:           runcRoot,
			// 	Namespace:      cj.Namespace,
			// 	ContainerNames: []string{cj.ContainerName},
			// 	SandboxNames:   []string{cj.PodName},
			// }
			// ctrResp, runcDiscoveryErr := c.RuncQuery(ctx, ctrByNameArgs)
			// if runcDiscoveryErr != nil {
			// 	return runcDiscoveryErr
			// }
			// // note: we assume we have atleast 1 container
			// runcId := ctrResp.Containers[0].ID
			// if runcId == "" {
			// 	return fmt.Errorf("empty runcId")
			// }
			// criuOpts := &task.CriuOpts{
			// 	ImagesDirectory: cj.CheckpointPath,
			// 	WorkDirectory:   "",
			// 	LeaveRunning:    true,
			// 	TcpEstablished:  false,
			// }
			// var dumpType task.CRType
			// if cj.CheckpointPath != "" {
			// 	dumpType = task.CRType_LOCAL
			// } else {
			// 	dumpType = task.CRType_REMOTE
			// }
			// dumpArgs := &task.RuncDumpArgs{
			// 	Root:           runcRoot,
			// 	ContainerID:    runcId,
			// 	CheckpointPath: cj.CheckpointPath,
			// 	CriuOpts:       criuOpts,
			// 	Type:           dumpType,
			// }
			// _, err := c.RuncDump(ctx, dumpArgs)
			// if err != nil {
			// 	return err
			// }
		}
		return nil
	})

	jqs.runner.Register("restore", func(ctx context.Context, b []byte) error {
		rj := &task.QueueJobRestoreRequest{}
		err := json.Unmarshal(b, rj)
		if err != nil {
			return err
		}
		log.Printf("restoring (%s) %s\n", rj.PodName, rj.ContainerName)
		// check runtime
		// TODO: handle runtime
		switch jqs.runtime {
		case "crio":
		case "runc":
			// runcRoot := "/default/runc"
			// ctrByNameArgs := &task.RuncQueryArgs{
			// 	Root:           runcRoot,
			// 	Namespace:      rj.Namespace,
			// 	ContainerNames: []string{rj.ContainerName},
			// 	SandboxNames:   []string{rj.PodName},
			// }
			// ctrResp, err := c.RuncQuery(ctx, ctrByNameArgs)
			// if err != nil {
			// 	return err
			// }
			// runcId := ctrResp.Containers[0].ID
			// pausePidArgs := &task.RuncGetPausePidArgs{
			// 	BundlePath: ctrResp.Containers[0].BundlePath,
			// }
			// pidResp, _ := c.RuncGetPausePid(ctx, pausePidArgs)
			// netPid := pidResp.PausePid
			// if runcId == "" {
			// 	return fmt.Errorf("Failed to locate runc container.")
			// }
			// runcOpts := &task.RuncOpts{
			// 	Root:          runcRoot,
			// 	Bundle:        ctrResp.Containers[0].BundlePath,
			// 	ConsoleSocket: "",
			// 	Detach:        true,
			// 	NetPid:        int32(netPid),
			// }
			//
			// var restoreType task.CRType
			//
			// if rj.CheckpointPath != "" {
			// 	restoreType = task.CRType_LOCAL
			// } else {
			// 	restoreType = task.CRType_REMOTE
			// }
			//
			// id := uuid.New()
			// runcRestoreArgs := &task.RuncRestoreArgs{
			// 	ImagePath:   rj.CheckpointPath,
			// 	ContainerID: id.String(),
			// 	IsK3S:       true,
			// 	Opts:        runcOpts,
			// 	Type:        restoreType,
			// }
			// _, err = c.RuncRestore(ctx, runcRestoreArgs)
		case "containerd":
		}
		return nil
	})
}

func (jqs *JobQueueService) Restore(c *task.QueueJobRestoreRequest) error {
	log.Println("Restore job enqueue")
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := jobs.Create(context.Background(), jqs.queue, "restore", b); err != nil {
		return err
	}
	return nil
}

func (jqs *JobQueueService) Checkpoint(c *task.QueueJobCheckpointRequest) error {
	log.Println("Checkpoint job enqueue")
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := jobs.Create(context.Background(), jqs.queue, "checkpoint", b); err != nil {
		return err
	}
	return nil
}
