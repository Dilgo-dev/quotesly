package main

import (
	"net/http"

	"github.com/Dilgo-dev/quotesly/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", handlers.GetQuote)

	http.ListenAndServe(":3000", r)
}