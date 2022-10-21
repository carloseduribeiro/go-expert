package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Course struct {
	Name     string
	Duration int
}

type CoursesList []Course

var templates = []string{
	"header.html",
	"content.html",
	"footer.html",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var err error
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err = t.Execute(w, CoursesList{
		{"Go", 40},
		{"Java", 40},
		{"Python", 30},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
