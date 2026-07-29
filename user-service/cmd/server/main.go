package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/korniykom/Chatters-Backend-Go/internal/handler"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", handler.Health)
	r.Get("/version", handler.Version)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
