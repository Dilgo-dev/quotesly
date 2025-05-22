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

	r.Get("/", handlers.GetQuote)

	log.Println("Starting server on port http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}