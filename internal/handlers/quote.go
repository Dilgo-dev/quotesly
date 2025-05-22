package handlers

import "net/http"

func GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🦆 Quotesly API"))
}