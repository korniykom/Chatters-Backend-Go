package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/korniykom/Chatters-Backend-Go/internal/config"
	"github.com/korniykom/Chatters-Backend-Go/internal/database"
	"github.com/korniykom/Chatters-Backend-Go/internal/handler"
	"github.com/korniykom/Chatters-Backend-Go/internal/repository/memory"
	"github.com/korniykom/Chatters-Backend-Go/internal/service"
)

func main() {
	cfg := config.Load()

	if err := database.Migrate(cfg); err != nil {
		panic(err)
	}

	db, err := database.New(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := memory.NewUserRepository()
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	r := chi.NewRouter()
	r.Get("/health", handler.Health)
	r.Get("/version", handler.Version)
	r.Get("/ping", handler.Ping)
	r.Post("/register", authHandler.Register)

	addr := ":" + cfg.Port

	fmt.Printf("Loaded configuration: %+v\n", cfg)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}

}
