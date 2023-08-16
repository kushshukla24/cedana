package cmd

import (
	"os"
	"testing"

	"github.com/cedana/cedana/utils"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/afero"
)

// we are skipping ci for now as we are using dump which requires criu, need to build criu on gh action
func skipCIT(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestClient_WriteOnlyFds(t *testing.T) {
	skipCIT(t)
	openFds := []process.OpenFilesStat{
		{Fd: 1, Path: "/path/to/file1"},
		{Fd: 2, Path: "/path/to/file2 (deleted)"},
		{Fd: 3, Path: "/path/to/file3"},
	}

	fs := afero.NewMemMapFs()
	contents := map[string]string{
		"/proc/1/fdinfo/1": "flags: 010002",
		"/proc/1/fdinfo/2": "flags: 0100000", //readonly - should not pass
		"/proc/1/fdinfo/3": "flags: 0100004", // readonly & append - should not pass
	}

	mockFS := &afero.Afero{Fs: fs}
	mockFS.MkdirAll("/proc/1/fdinfo", 0755)

	for k, v := range contents {
		mockFS.WriteFile(k, []byte(v), 0644)
	}

	logger := utils.GetLogger()
	c := &Client{
		fs:     mockFS,
		logger: &logger,
	}

	paths := c.WriteOnlyFds(openFds, 1)

	// Test case 1: Check if the path of the first file is included in the output
	if !contains(paths, "/path/to/file1") {
		t.Errorf("expected path '/path/to/file1' to be included in the output, but it was not")
	}

	// Test case 2: Check if the path of the second file (with '(deleted)' suffix removed) is included in the output
	if contains(paths, "/path/to/file2") {
		t.Errorf("expected path '/path/to/file2' to not be included in the output, but it was")
	}

	// Test case 3: Check if the path of the third file is included in the output
	if contains(paths, "/path/to/file3") {
		t.Errorf("expected path '/path/to/file3' to not be included in the output, but it was")
	}
}

func contains(paths []string, path string) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}
