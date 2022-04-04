package venv

import (
	"os"
	"path/filepath"
)

func (v *Venv) Initialize() {
	v.Env = getVenv()
	v.Active = isVenvActive(v.Env)
}

func getVenv() string {
	venvPath := os.Getenv("VIRTUAL_ENV")
	if len(venvPath) > 0 {
		return filepath.Base(venvPath)
	}
	return ""
}

func isVenvActive(v string) bool {
	return len(v) > 0
}
