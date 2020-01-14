package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type cat struct {
	Name string
	Nickname string
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	sophie := cat{
		Name: "Sophie",
		Nickname: "Fluff",
	}

	cecy := cat{
		Name: "Cecily",
		Nickname: "Fuzz",
	}

	cats := []cat{sophie, cecy}

	err := tpl.Execute(os.Stdout, cats)
	if err != nil {
		log.Fatalln(err)
	}
}
