package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestGetBranch(t *testing.T) {
	var out bytes.Buffer
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Error("failed to get current branch")
		t.Error(err)
	}

	expected := strings.Trim(out.String(), "\n")
	fmt.Println(expected)

	actual := getBranch()
	if actual != expected {
		t.Errorf("got: %s, want: %s", actual, expected)
	}
}
