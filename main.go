package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type PromptOpts struct {
	Branch string
}

func format(color string) string {
	c := strings.ToLower(color)
	if c == "red" {
		return "\\e[31m"
	} else if c == "nc" {
		return "\\e[0m"
	}
	return ""
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

// TODOS:
// Add python virtual environment
// Add colors
// Add text modifiers
// Add optional braces if option doesn't exit
// Add default ps1 params (\u, \h, \w, etc...)

func ellipse(text string, length int) string {
	if len(text) > length {
		return text[:(length-3)] + "..."
	}
	return text
}

func surroundWith(text string, open string, close string) string {
	if len(text) != 0 {
		text = open + text + close
	}
	return text
}

func appendWith(text string, postfix string) string {
	if len(text) != 0 {
		text = text + postfix
	}
	return text
}

var funcMap = template.FuncMap{
	"surroundWith": surroundWith,
	"appendWith":   appendWith,
	"ellipse":      ellipse,
	"format":       format,
}

func main() {
	// td := Todo{"Test templates", "Let's test a template to see the magic."}

	// t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	// if err != nil {
	// 	panic(err)
	// }
	// err = t.Execute(os.Stdout, td)
	// if err != nil {
	// 	panic(err)
	// }
	prompt_template := "\\u@\\h:\\w-'{{ format \"red\" }}{{ surroundWith (ellipse .Branch 10) \"[\" \"]\" }}{{ format \"nc\" }}'\\$ \n"
	opts := PromptOpts{}
	opts.Branch = getBranch()
	t, err := template.New("ps1").Funcs(funcMap).Parse(prompt_template)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, opts)
}
