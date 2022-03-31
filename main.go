package main

import (
	"os"
	"text/template"
	"time"
)

// type Todo struct {
// 	Name        string
// 	Description string
// }

type PromptOpts struct {
	Time string
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
	prompt_template := "\\u@\\h:\\w-[{{ .Time }}]\\$ "
	opts := PromptOpts{}
	opts.Time = time.Now().String()
	t, err := template.New("ps1").Parse(prompt_template)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, opts)
}