package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", blog{title: "Test"})
	http.ListenAndServe(":8080", mux)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
