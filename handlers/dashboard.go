package handlers

import (
	"my_blog/models"
	"net/http"
)

func DashboardHandler(write http.ResponseWriter, request *http.Request) {
	// Data retrieval
	articles := GetArticlesSorted()

	// Connect with HTML
	tmpl := models.ParseTemplate("dashboard.html")

	// Data display
	err := tmpl.Execute(write, articles)
	if err != nil {
		http.Error(write, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
