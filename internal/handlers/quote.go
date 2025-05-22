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
	json.NewEncoder(w).Encode(quotes)
}