package handlers

import (
	"my_blog/models"
	"sort"
	"time"
)

func GetArticlesSorted() []models.Article {
	articles := GetArticles()

	sort.Slice(articles, func(i, j int) bool {
		t1, _ := time.Parse("2006-01-02", articles[i].Date)
		t2, _ := time.Parse("2006-01-02", articles[j].Date)
		return t1.After(t2) // новые сверху
	})

	return articles
}
