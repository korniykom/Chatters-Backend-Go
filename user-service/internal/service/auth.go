package service

import (
	"context"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, user domain.User) error
}

type AuthService struct {
	repository UserRepository
}

func NewAuthService(repository UserRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Register(
	ctx context.Context,
	req domain.RegisterRequest,
) error {
	savedUser, err := s.repository.FindByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	if savedUser != nil {
		return ErrUserAlreadyExists
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
	}

	return s.repository.Save(ctx, user)
}
