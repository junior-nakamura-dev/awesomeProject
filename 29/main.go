package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {

	course := Course{Name: "Go", Hours: 40}
	tmp := template.Must(template.New("Go Course").Parse("Course: {{ .Name }}, Hours: {{ .Hours }}"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
