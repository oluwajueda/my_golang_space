package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Fname string
	Lname string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	first := sage{
		Fname: "olamide",
		Lname: "osuolale",
	}
	second := sage{
		Fname: "ola",
		Lname: "osu",
	}

	third := sage{
		Fname: "olamy",
		Lname: "osuolale",
	}

	fourth := sage{
		Fname: "olamy",
		Lname: "osuolale",
	}

	sages := []sage{first, second, third, fourth}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}

}
