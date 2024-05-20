package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Rlimit struct {
	Type int    `json:"type"`
	Hard uint64 `json:"hard"`
	Soft uint64 `json:"soft"`
}

// IDMap represents UID/GID Mappings for User Namespaces.
type IDMap struct {
	ContainerID int `json:"container_id"`
	HostID      int `json:"host_id"`
	Size        int `json:"size"`
}

// Seccomp represents syscall restrictions
// By default, only the native architecture of the kernel is allowed to be used
// for syscalls. Additional architectures can be added by specifying them in
// Architectures.
type Seccomp struct {
	DefaultAction    Action                   `json:"default_action"`
	Architectures    []string                 `json:"architectures"`
	Flags            []specs.LinuxSeccompFlag `json:"flags"`
	Syscalls         []*Syscall               `json:"syscalls"`
	DefaultErrnoRet  *uint                    `json:"default_errno_ret"`
	ListenerPath     string                   `json:"listener_path,omitempty"`
	ListenerMetadata string                   `json:"listener_metadata,omitempty"`
}

// Action is taken upon rule match in Seccomp
type Action int

const (
	Kill Action = iota + 1
	Errno
	Trap
	Allow
	Trace
	Log
	Notify
	KillThread
	KillProcess
)

// Operator is a comparison operator to be used when matching syscall arguments in Seccomp
type Operator int

const (
	EqualTo Operator = iota + 1
	NotEqualTo
	GreaterThan
	GreaterThanOrEqualTo
	LessThan
	LessThanOrEqualTo
	MaskEqualTo
)

// Arg is a rule to match a specific syscall argument in Seccomp
type Arg struct {
	Index    uint     `json:"index"`
	Value    uint64   `json:"value"`
	ValueTwo uint64   `json:"value_two"`
	Op       Operator `json:"op"`
}

// Syscall is a rule to match a syscall in Seccomp
type Syscall struct {
	Name     string `json:"name"`
	Action   Action `json:"action"`
	ErrnoRet *uint  `json:"errnoRet"`
	Args     []*Arg `json:"args"`
}

type HookName string
type HookList []Hook
type Hooks map[HookName]HookList

const (
	// Prestart commands are executed after the container namespaces are created,
	// but before the user supplied command is executed from init.
	// Note: This hook is now deprecated
	// Prestart commands are called in the Runtime namespace.
	Prestart HookName = "prestart"

	// CreateRuntime commands MUST be called as part of the create operation after
	// the runtime environment has been created but before the pivot_root has been executed.
	// CreateRuntime is called immediately after the deprecated Prestart hook.
	// CreateRuntime commands are called in the Runtime Namespace.
	CreateRuntime = "createRuntime"

	// CreateContainer commands MUST be called as part of the create operation after
	// the runtime environment has been created but before the pivot_root has been executed.
	// CreateContainer commands are called in the Container namespace.
	CreateContainer = "createContainer"

	// StartContainer commands MUST be called as part of the start operation and before
	// the container process is started.
	// StartContainer commands are called in the Container namespace.
	StartContainer = "startContainer"

	// Poststart commands are executed after the container init process starts.
	// Poststart commands are called in the Runtime Namespace.
	Poststart = "poststart"

	// Poststop commands are executed after the container init process exits.
	// Poststop commands are called in the Runtime Namespace.
	Poststop = "poststop"
)

type Capabilities struct {
	// Bounding is the set of capabilities checked by the kernel.
	Bounding []string
	// Effective is the set of capabilities checked by the kernel.
	Effective []string
	// Inheritable is the capabilities preserved across execve.
	Inheritable []string
	// Permitted is the limiting superset for effective capabilities.
	Permitted []string
	// Ambient is the ambient set of capabilities that are kept.
	Ambient []string
}

func (hooks HookList) RunHooks(state *specs.State) error {
	for i, h := range hooks {
		if err := h.Run(state); err != nil {
			return errors.Wrapf(err, "Running hook #%d:", i)
		}
	}

	return nil
}

func (hooks *Hooks) UnmarshalJSON(b []byte) error {
	var state map[HookName][]CommandHook

	if err := json.Unmarshal(b, &state); err != nil {
		return err
	}

	*hooks = Hooks{}
	for n, commandHooks := range state {
		if len(commandHooks) == 0 {
			continue
		}

		(*hooks)[n] = HookList{}
		for _, h := range commandHooks {
			(*hooks)[n] = append((*hooks)[n], h)
		}
	}

	return nil
}

func (hooks *Hooks) MarshalJSON() ([]byte, error) {
	serialize := func(hooks []Hook) (serializableHooks []CommandHook) {
		for _, hook := range hooks {
			switch chook := hook.(type) {
			case CommandHook:
				serializableHooks = append(serializableHooks, chook)
			default:
				logrus.Warnf("cannot serialize hook of type %T, skipping", hook)
			}
		}

		return serializableHooks
	}

	return json.Marshal(map[string]interface{}{
		"prestart":        serialize((*hooks)[Prestart]),
		"createRuntime":   serialize((*hooks)[CreateRuntime]),
		"createContainer": serialize((*hooks)[CreateContainer]),
		"startContainer":  serialize((*hooks)[StartContainer]),
		"poststart":       serialize((*hooks)[Poststart]),
		"poststop":        serialize((*hooks)[Poststop]),
	})
}

type Hook interface {
	// Run executes the hook with the provided state.
	Run(*specs.State) error
}

// NewFunctionHook will call the provided function when the hook is run.
func NewFunctionHook(f func(*specs.State) error) FuncHook {
	return FuncHook{
		run: f,
	}
}

type FuncHook struct {
	run func(*specs.State) error
}

func (f FuncHook) Run(s *specs.State) error {
	return f.run(s)
}

type Command struct {
	Path    string         `json:"path"`
	Args    []string       `json:"args"`
	Env     []string       `json:"env"`
	Dir     string         `json:"dir"`
	Timeout *time.Duration `json:"timeout"`
}

// NewCommandHook will execute the provided command when the hook is run.
func NewCommandHook(cmd Command) CommandHook {
	return CommandHook{
		Command: cmd,
	}
}

type CommandHook struct {
	Command
}

func (c Command) Run(s *specs.State) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	var stdout, stderr bytes.Buffer
	cmd := exec.Cmd{
		Path:   c.Path,
		Args:   c.Args,
		Env:    c.Env,
		Stdin:  bytes.NewReader(b),
		Stdout: &stdout,
		Stderr: &stderr,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	errC := make(chan error, 1)
	go func() {
		err := cmd.Wait()
		if err != nil {
			err = fmt.Errorf("error running hook: %v, stdout: %s, stderr: %s", err, stdout.String(), stderr.String())
		}
		errC <- err
	}()
	var timerCh <-chan time.Time
	if c.Timeout != nil {
		timer := time.NewTimer(*c.Timeout)
		defer timer.Stop()
		timerCh = timer.C
	}
	select {
	case err := <-errC:
		return err
	case <-timerCh:
		cmd.Process.Kill()
		cmd.Wait()
		return fmt.Errorf("hook ran past specified timeout of %.1fs", c.Timeout.Seconds())
	}
}
