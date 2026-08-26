package memory

import (
	"context"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(
	ctx context.Context,
	user domain.User,
) error {
	r.users = append(r.users, user)

	return nil
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	for i := range r.users {
		if r.users[i].Email == email {
			return &r.users[i], nil
		}
	}

	return nil, nil
}
