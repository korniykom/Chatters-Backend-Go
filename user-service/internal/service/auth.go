package service

import (
	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Save(user domain.User) error
}

type AuthService struct {
	repository UserRepository
}

func NewAuthService(repository UserRepository) *AuthService {
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
