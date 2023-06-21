package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	FirstName string
	Surname   string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	name := sage{
		FirstName: "olamide",
		Surname:   "osuolale",
	}

	err := tpl.Execute(os.Stdout, name)

	if err != nil {
		log.Fatalln(err)
	}
}
