package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Dilgo-dev/quotesly/internal/config"
	"github.com/Dilgo-dev/quotesly/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("PORT is not set, using default port 3000")
	}

	_, err := config.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/quotes", handlers.GetQuote)
	r.Get("/api/quotes/random", handlers.GetRandomQuote)
	r.Post("/api/quotes", handlers.CreateQuote)
	r.Get("/api/quotes/category/{category}", handlers.GetQuoteByCategory)

	log.Println("Starting server on port http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}