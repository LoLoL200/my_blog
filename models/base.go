package models

// Account
type Account struct {
	AccountID int    `json:"accountID"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Article
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Author  string `json:"author,omitempty"`
}
