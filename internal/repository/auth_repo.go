package repository

import (
	"auth/vladmsnk/internal/entity"
	"auth/vladmsnk/pkg/postgres"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Repository struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg}
}

func (r *Repository) SaveUser(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	_, err := r.Pool.Exec(ctx, InsertUser, user.Id, user.Username, user.Email, user.PasswordHash, user.Number)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.Id, nil
}

func (r *Repository) FindUserUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)

	err := r.Pool.QueryRow(ctx, FindUserByEmail, email).Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Number)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}
	return user, nil
}

func (r *Repository) FindUserUserByUsernameAndPassword(ctx context.Context, username, password string) (*entity.User, error) {
	user := new(entity.User)

	err := r.Pool.QueryRow(ctx, FindUserByUsernameAndPassword, username, password).
		Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Number)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}
	return user, nil
}
