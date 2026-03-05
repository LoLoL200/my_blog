package models

import (
	"fmt"
	"text/template"
)

// Parsing tamplate
func ParseTemplate(templateName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s", templateName))
	if err != nil {
		panic(fmt.Sprintf("error parsing template %s,%v", templateName, err))
	}
	return tmpl
}
