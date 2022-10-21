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
	course := Course{"Go", 40}
	tmp := template.New("courseTemplate")
	tmp, _ = tmp.Parse("Course: {{.Name}} - Duration: {{.Duration}}\n")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		log.Fatal(err)
	}
}
