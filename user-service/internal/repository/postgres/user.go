package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(
	ctx context.Context,
	user domain.User,
) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (username, email)
		 VALUES ($1, $2)`,
		user.Username,
		user.Email,
	)

	return err
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(
		ctx,
		`SELECT id, username, email
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
