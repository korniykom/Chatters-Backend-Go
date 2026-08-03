package service

import (
	"fmt"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type AuthService struct {
	users []domain.User
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(req domain.RegisterRequest) error {
	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
	}

	s.users = append(s.users, user)

	fmt.Println(s.users)

	return nil
}
