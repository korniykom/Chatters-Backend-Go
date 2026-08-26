package memory

import (
	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user domain.User) error {

	r.users = append(r.users, user)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	for i := range r.users {
		if r.users[i].Email == email {
			return &r.users[i], nil
		}
	}

	return nil, nil
}
