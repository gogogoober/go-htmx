package main

import (
	"os"
	"text/template"
)

type Task struct {
	Title   string
	Done    bool
	Message string
}

func main() {

	tasks := []Task{
		{
			Title:   "Task 1",
			Done:    false,
			Message: "This is my first task",
		},
		{
			Title:   "Task 2",
			Done:    true,
			Message: "This is my second task",
		},
	}

	filePath := "task.tmpl"
	templ, err := template.New(filePath).ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	err = templ.Execute(os.Stdout, tasks)
	if err != nil {
		panic(err)
	}

}
