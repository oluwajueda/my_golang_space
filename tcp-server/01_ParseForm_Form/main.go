package main

import (
	"html/template"
	"log"
	"net/http"
)

type formtype int

func (m formtype) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles())
}

func main() {
	var d formtype
	http.ListenAndServe(":8080", d)
}
