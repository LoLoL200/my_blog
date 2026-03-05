package handlers

import (
	"my_blog/models"
	"net/http"
)

// Home Handler
func HomeHandler(write http.ResponseWriter, request *http.Request) {
	// Data retrieval
	articles := GetArticlesSorted() // <-- здесь мы получаем статьи уже отсортированные по дате

	// Connect with HTML
	tmpl := models.ParseTemplate("home.html")

	// Data display
	err := tmpl.Execute(write, articles)
	if err != nil {
		http.Error(write, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
