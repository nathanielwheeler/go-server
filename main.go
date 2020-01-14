package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	cats := []string{"Sophie", "Cecily", "Lulu"}

	err := tpl.Execute(os.Stdout, cats)
	if err != nil {
		log.Fatalln(err)
	}
}
