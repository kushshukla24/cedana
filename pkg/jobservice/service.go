package jobservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"

	"github.com/cedana/cedana/pkg/api/containerd"
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
		log.Printf("checkpointing (%s) %s %s\n", cj.PodName, cj.ContainerName, cj.ImageName)
		// check runtime
		switch jqs.runtime {
		case "crio":
		case "containerd":
			sockAddr := "/run/containerd/runc"
			containerdService, err := containerd.New(ctx, sockAddr, nil)
			if err != nil {
				return err
			}

			containerId, err := containerdService.GetContainerdID(ctx, cj.ContainerName, cj.Namespace)
			if err != nil {
				return err
			}

			ref, err := containerdService.DumpRootfs(ctx, containerId, cj.ImageName, cj.Namespace)
			if err != nil {
				return err
			}
			_ = ref
		case "runc":
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
		switch jqs.runtime {
		case "crio":
		case "runc":
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
