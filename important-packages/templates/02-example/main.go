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

func main() {
	var err error
	course := Course{"Go", 40}
	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Duration: {{.Duration}}"))
	err = t.Execute(os.Stdout, course)
	if err != nil {
		log.Fatal(err)
	}
}
