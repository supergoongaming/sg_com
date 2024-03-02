package pages

import (
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("static/*.html"))
}

func GetTemplates() *template.Template {
	return templates
}