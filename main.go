package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.gohtml"))
}

func yearMonthDay(t time.Time) string {
	return t.Format("2006-01-02-Mon")
}

var fm = template.FuncMap{
	"fdateYMD": yearMonthDay,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
