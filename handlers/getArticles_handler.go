package handlers

import (
	"encoding/json"
	"my_blog/models"
	"net/http"
	"os"
	"path/filepath"
)

// GET
func GetArticle(w http.ResponseWriter, r *http.Request) {
	//GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		tmpl := models.ParseTemplate("new_article.html")
		tmpl.Execute(w, nil)
		return
	}

}

// GET
func GetArticles() []models.Article {
	// Reading directory
	files, _ := os.ReadDir("articles")
	var articles []models.Article

	//Checking files in a folder and assembling them into a slice
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".json" {
			continue
		}
		data, _ := os.ReadFile(filepath.Join("articles", f.Name()))
		var art models.Article
		json.Unmarshal(data, &art)
		articles = append(articles, art)
	}
	return articles
}
