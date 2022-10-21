package main

import (
	"log"
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

type CoursesList []Course

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var err error
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err = t.Execute(w, CoursesList{
		{"Go", 40},
		{"Java", 40},
		{"Python", 30},
	})
	if err != nil {
		log.Fatal(err)
	}
}
