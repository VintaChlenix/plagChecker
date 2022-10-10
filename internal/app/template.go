package app

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("internal/templates/upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}
