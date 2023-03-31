package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate parses the template file and Executes it using the ResponseWriter
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
