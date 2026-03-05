package main

import (
	"fmt"
	api "my_blog/handlers"
	"my_blog/handlers/middleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warninig: .env file not found")
	}
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		os.Mkdir("articles", 0755)
	}
}

func main() {
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/dashboard", api.DashboardArticleWithAuthI())
	http.HandleFunc("/new", api.CreateArticleWithAuthI())
	http.HandleFunc("/edit/", api.UpdateArticleWithAuthI())
	http.HandleFunc("/articles/", api.DeleteArticleWithAuthI())
	http.HandleFunc("/login", middleware.LoginHandler)
	http.HandleFunc("/logout", middleware.LogoutHandler)

	fmt.Println("Server started on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
