package transport

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	template, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
