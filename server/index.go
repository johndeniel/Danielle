package server

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/johndeniel/danielle/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
	tmpl.Execute(w, models.Perfumes)
}
