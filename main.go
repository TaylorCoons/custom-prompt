package main

import (
	"flag"
	"fmt"
	"os"
	"path"
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
	Time   Opt
}

var funcMap = template.FuncMap{
	"ellipse": format.Ellipse,
	"format":  format.FormatChar,
}

func Initialize() *PromptOpts {
	opts := PromptOpts{}
	opts.Git = new(git.Git)
	opts.Format = new(format.Format)
	opts.Venv = new(venv.Venv)
	opts.Git.Initialize()
	opts.Format.Initialize()
	opts.Venv.Initialize()
	return &opts
}

func run() error {
	promptPtr := flag.String("prompt", "ubuntu", "prompt template to load")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home dir %v", err)
	}
	promptBase := ".custom-prompt/prompts"
	flag.Parse()

	promptDir := path.Join(homeDir, promptBase, *promptPtr, "*")

	opts := Initialize()

	t, err := template.New("ps1.tmpl").Funcs(funcMap).ParseGlob(promptDir)
	if err != nil {
		return err
	}
	t.ExecuteTemplate(os.Stdout, "ps1.tmpl", opts)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("[%sError: %ssee log in ~/.custom-prompt/log]\\$ ", format.FormatChar("31"), format.FormatChar("0"))
		os.Exit(1)
	}
}
