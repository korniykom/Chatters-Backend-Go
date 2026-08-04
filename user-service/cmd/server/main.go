package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/korniykom/Chatters-Backend-Go/internal/handler"
	"github.com/korniykom/Chatters-Backend-Go/internal/repository/memory"
	"github.com/korniykom/Chatters-Backend-Go/internal/service"
)

func main() {
	userRepository := memory.NewUserRepository()
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	r := chi.NewRouter()
	r.Get("/health", handler.Health)
	r.Get("/version", handler.Version)
	r.Get("/ping", handler.Ping)
	r.Post("/register", authHandler.Register)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
