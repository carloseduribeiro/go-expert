package main

import (
	"log"
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

type CoursesList []Course

func main() {
	var err error
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err = t.Execute(os.Stdout, CoursesList{
		{"Go", 40},
		{"Java", 60},
		{"Python", 30},
	})
	if err != nil {
		log.Fatal(err)
	}
}
