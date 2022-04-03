package main

import (
	"os"
	"text/template"

	"github.com/taylorcoons/custom-prompt/format"
	"github.com/taylorcoons/custom-prompt/git"
	"github.com/taylorcoons/custom-prompt/venv"
)

type Opt interface {
	Initialize()
}

type PromptOpts struct {
	Git    Opt
	Format Opt
	Venv   Opt
}

// TODOS:
// Add python virtual environment
// Add colors
// Add text modifiers
// Add optional braces if option doesn't exit
// Add default ps1 params (\u, \h, \w, etc...)

var funcMap = template.FuncMap{
	"ellipse": format.Ellipse,
	"format":  format.FormatChar,
}

func main() {
	// prompt_template := "colors-{{ format .Fg.Red }}red{{ format .Fg.LightRed }}light_red{{ format .Fg.NC }}\\$ "
	// prompt_template := "\\u@\\h:\\w-'{{ format \"red\" }}{{ surroundWith (ellipse .Branch 10) \"[\" \"]\" }}{{ format \"nc\" }}'\\$ \n"
	// prompt_template := "test-{{ format .Format.Fg.Red .Format.Em.Underlined .Format.Em.Bold }}{{ .Git.Branch }}{{ format .Format.Fg.NC }}$"
	opts := PromptOpts{}
	opts.Git = new(git.Git)
	opts.Git.Initialize()
	opts.Format = new(format.Format)
	opts.Format.Initialize()
	opts.Venv = new(venv.Venv)
	opts.Venv.Initialize()

	paths := []string{
		"/home/taylor/code/custom-prompt/kali.tmpl",
	}

	t, err := template.New("kali.tmpl").Funcs(funcMap).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, opts)
	if err != nil {
		panic(err)
	}
}
