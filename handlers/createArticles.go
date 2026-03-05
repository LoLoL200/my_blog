package handlers

import (
	"encoding/json"
	"fmt"
	"my_blog/models"
	"net/http"
	"os"
	"strings"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	//GET
	if r.Method == http.MethodGet {
		tmpl := models.ParseTemplate("new_article.html")
		tmpl.Execute(w, nil)
		return
	}

	var a models.Article

	// Reading data from a form
	r.ParseForm()
	a.Title = r.FormValue("title")
	a.Content = r.FormValue("content")
	a.Date = r.FormValue("date")

	//os.MkdirAll("articles", os.ModePerm)
	files, err := os.ReadDir("articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search for Max ID
	maxID := 0
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			var id int
			fmt.Sscanf(file.Name(), "article%d.json", &id)
			if id > maxID {
				maxID = id
			}
		}
	}

	// Max ID additions +1
	a.ID = maxID + 1

	//Generating a file path for a new article
	filePath := fmt.Sprintf("articles/article%d.json", a.ID)

	//Create new file for directory(disk)
	file, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	//recording in JSON file
	json.NewEncoder(file).Encode(a)

	//Send a response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}
