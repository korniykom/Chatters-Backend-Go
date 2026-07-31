package service

import "github.com/korniykom/Chatters-Backend-Go/internal/domain"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(req domain.RegisterRequest) error {
	return nil
}
