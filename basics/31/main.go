package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	course1 := Course{Name: "Go", Hours: 40}
	course2 := Course{Name: "Java", Hours: 40}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.New("template.html")
		tmp.Funcs(template.FuncMap{"toUpper": toUpper})
		tmp = template.Must(tmp.ParseFiles("template.html"))
		err := tmp.Execute(w, Courses{course1, course2})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)

}
