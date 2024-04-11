package crit

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/checkpoint-restore/go-criu/v7/crit"
)

type Crit struct {
	critter crit.Critter
}

func New(imageDir string) *Crit {
	c := crit.New(nil, nil, imageDir, false, false)

	return &Crit{critter: c}
}

func (c *Crit) ReadCgroupFds() (map[string]string, error) {

	fds, err := c.critter.ExploreFds()
	if err != nil {
		return nil, fmt.Errorf("failed to explore fds: %w", err)
	}

	result := map[string]string{}

	for _, fd := range fds {
		for _, file := range fd.Files {
			if !strings.HasPrefix(file.Path, "/sys/fs/cgroup/kubepods.slice") ||
				file.Type != "REG" {
				continue
			}
			result[filepath.Base(file.Path)] = file.Path
		}
	}
	return result, nil
}
