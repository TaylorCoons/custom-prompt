package venv

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	currEnv := os.Getenv("VIRTUAL_ENV")
	defer func() {
		err := os.Setenv("VIRTUAL_ENV", currEnv)
		if err != nil {
			t.Error("failed to restore current environment.")
			t.Error(err)
		}
	}()

	expected := "test_env"
	venv := fmt.Sprintf("path/to/%s", expected)
	err := os.Setenv("VIRTUAL_ENV", venv)
	if err != nil {
		t.Error("failed to set VIRTUAL_ENV")
		t.Error(err)
	}

	actual := getVenv()
	if actual != expected {
		t.Errorf("wrong venv - got: %s, want %s", actual, expected)
	}

	err = os.Unsetenv("VIRTUAL_ENV")
	if err != nil {
		t.Error("failed to unset VIRTUAL_ENV")
		t.Error(err)
	}

	actual = getVenv()
	fmt.Println(actual)
	expected = ""
	if actual != expected {
		t.Errorf("wrong venv - got: %s, want: %s", actual, expected)
	}
}

func TestIsVenvActive(t *testing.T) {
	actual := isVenvActive("Some/venv")
	expected := true
	if actual != expected {
		t.Errorf("venv wrong status - got: %t, want: %t", actual, expected)
	}
	actual = isVenvActive("")
	expected = false
	if actual != expected {
		t.Errorf("venv wrong status - got: %t, want: %t", actual, expected)
	}
}
