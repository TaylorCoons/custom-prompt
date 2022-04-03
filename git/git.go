package git

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (g *Git) Initialize() {
	g.Branch = getBranch()
}

func getBranch() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	gitRepo := filepath.Join(wd, ".git")
	cmd := exec.Command("git", "--git-dir", gitRepo, "branch", "--show-current")

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return ""
	}
	name := out.String()
	name = strings.TrimSuffix(name, "\n")
	return name
}
