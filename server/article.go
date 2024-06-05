package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/johndeniel/danielle/models"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/article/")
	title = strings.ReplaceAll(title, "%20", " ") // Optional: Decode URL-encoded spaces
	log.Printf("GET request received for title: %s", title)

	var matchedArticle *models.Article
	for _, articleList := range models.Articles {
		for _, a := range articleList {
			if a.Title == title {
				matchedArticle = &a
				break
			}
		}
		if matchedArticle != nil {
			break
		}
	}

	if matchedArticle == nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "article.html")))
	err := tmpl.Execute(w, matchedArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
