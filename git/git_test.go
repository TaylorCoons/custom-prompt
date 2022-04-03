package git

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGetBranch(t *testing.T) {
	expected := "temp-test-branch"
	var out bytes.Buffer
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		t.Error("failed to get current branch")
	}
	current := out.String()
	cmd = exec.Command("git", "checkout", "-b", expected)

	err = cmd.Run()

	if err != nil {
		t.Error("failed to checkout test branch")
	}

	actual := getBranch()
	if actual != expected {
		t.Errorf("got %s, want: %s", actual, expected)
	}
	cmd = exec.Command("git", "checkout", current)
	err = cmd.Run()

	if err != nil {
		t.Error("failed to checkout current branch")
	}
}
