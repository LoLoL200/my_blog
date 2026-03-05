package handlers

import (
	"my_blog/handlers/middleware"
	"net/http"
)

// Wraps
func CreateArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(CreateArticle)
}
func DashboardArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(DashboardHandler)
}
func UpdateArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(UpdateArticle)
}

func DeleteArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(DeleteArticle)
}
