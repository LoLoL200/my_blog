package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	//DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Getting ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/articles/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("articles/article%d.json", id)
	// Cheking for void directory
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "article not found", http.StatusNotFound)
		return
	}
	// Delete files
	err = os.Remove(filePath)
	// Cheking for void directory
	if err != nil {
		http.Error(w, "cannot delete file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
