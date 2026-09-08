package service

import (
	"context"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
	"github.com/korniykom/Chatters-Backend-Go/internal/validation"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, user domain.User) (*domain.User, error)
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
) (*domain.User, error) {
	req = validation.NormalizeRegistrationRequest(req)

	if !validation.ValidateUsername(req.Username) {
		return nil, ErrInvalidUsername
	}

	if !validation.ValidateEmail(req.Email) {
		return nil, ErrInvalidEmail
	}

	savedUser, err := s.repository.FindByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if savedUser != nil {
		return nil, ErrUserAlreadyExists
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
	}

	createdUser, err := s.repository.Save(ctx, user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
