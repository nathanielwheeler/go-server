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
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
