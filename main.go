package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// This parses all of the files in the templates folder
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "<h1>Hello World!</h1>")
	if err != nil {
		log.Fatalln(err)
	}
}
