package handlers

import (
	"encoding/json"
	"fmt"
	"my_blog/handlers/middleware"
	"my_blog/models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetArticlesByID(id int) *models.Article {
	filesPath := fmt.Sprintf("articles/article%d.json", id)
	date, err := os.ReadFile(filesPath)
	if err != nil {
		return nil
	}
	var a models.Article
	json.Unmarshal(date, &a)
	return &a
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/edit/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	// GET
	case http.MethodGet:
		article := GetArticlesByID(id)

		if article == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		tmpl := models.ParseTemplate("update_article.html")
		tmpl.Execute(w, article)

		//POST
	case http.MethodPost:
		// Getting the current authorized user
		user := middleware.GetUserFromContext(r)
		if user == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		// article := GetArticlesByID(id)
		// if article == nil {
		// 	http.Error(w, "not found", http.StatusNotFound)
		// 	return
		// }
		// // Checking if the current user is the author
		// if article.Author != user.Username {
		// 	http.Error(w, "forbidden", http.StatusForbidden)
		// 	return
		// }

		// Reading date form
		r.ParseForm()

		//Getting values ​​from a form
		title := r.FormValue("title")
		content := r.FormValue("content")
		published := r.FormValue("date")

		// Examination title
		if title == "" || len(title) > 100 {
			http.Error(w, "invalid title", http.StatusBadRequest)
			return
		}

		// Examination for void title
		if content == "" {
			http.Error(w, "content required", http.StatusBadRequest)
			return
		}

		// Examination for right
		if _, err := time.Parse("2006-01-02", published); err != nil {
			http.Error(w, "invalid date", http.StatusBadRequest)
			return
		}

		// Forming a path to a file
		filePath := fmt.Sprintf("articles/article%d.json", id)
		// Examination file Path for void derection
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		a := models.Article{
			ID:      id,
			Title:   title,
			Content: content,
			Date:    published,
			Author:  user.Username,
		}

		file, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		if err := json.NewEncoder(file).Encode(a); err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
