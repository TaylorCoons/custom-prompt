package venv

import (
	"os"
	"path/filepath"
)

func (v *Venv) Initialize() {
	venv_path := os.Getenv("VIRTUAL_ENV")
	if len(venv_path) > 0 {
		v.Env = filepath.Base(venv_path)
		v.Active = true
	} else {
		v.Env = ""
		v.Active = false
	}
}
