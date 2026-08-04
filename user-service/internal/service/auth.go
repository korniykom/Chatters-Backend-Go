package service

import (
	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
	"github.com/korniykom/Chatters-Backend-Go/internal/repository/memory"
)

type AuthService struct {
	repository *memory.UserRepository
}

func NewAuthService(repository *memory.UserRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Register(req domain.RegisterRequest) error {

	savedUser, err := s.repository.FindByEmail(req.Email)

	if err != nil {
		return err
	}

	if savedUser != nil {
		return ErrEmailAlreadyExists
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
	}

	return s.repository.Save(user)
}
