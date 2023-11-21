package transport

import (
	"day06/internal/entity"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	template, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, entity.ErrFailedRenderTemplate.Error(), http.StatusInternalServerError)
		return
	}

	template.Execute(w, data)
}
