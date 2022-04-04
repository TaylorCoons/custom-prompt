package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func (g *Git) Initialize() {
	g.Branch = getBranch()
	g.Dirty = isDirty()
}

func getBranch() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = wd

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

func isDirty() bool {

	wd, err := os.Getwd()
	if err != nil {
		return false
	}

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = wd

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return false
	}
	modified := out.String()
	return len(modified) > 0
}
