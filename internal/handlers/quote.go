package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Dilgo-dev/quotesly/internal/config"
	"github.com/Dilgo-dev/quotesly/internal/models"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var quotes []models.Quote

	db.Find(&quotes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if quote.Quote == "" || quote.Author == "" {
		http.Error(w, "Quote and author are required", http.StatusBadRequest)
		return
	}

	db := config.GetDB()
	if err := db.Create(&quote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(quote)
}