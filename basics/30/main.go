package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {

	course1 := Course{Name: "Go", Hours: 40}
	course2 := Course{Name: "Java", Hours: 40}
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(os.Stdout, Courses{course1, course2})
	if err != nil {
		panic(err)
	}
}
