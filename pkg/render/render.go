package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// TemplateRenderer parses the template file and Executes it using the ResponseWriter
func TemplateRenderer(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
