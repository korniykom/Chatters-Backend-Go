package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
	"github.com/korniykom/Chatters-Backend-Go/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Register(r.Context(), req); err != nil {
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	response := domain.RegisterResponse{Message: "User registered successfully"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
