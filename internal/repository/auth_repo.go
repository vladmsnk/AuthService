package repository

import (
	"auth/vladmsnk/internal/entity"
	"auth/vladmsnk/pkg/postgres"
	"context"
	"github.com/gofrs/uuid"
)

type Repository struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg}
}

func (r *Repository) SaveUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	_, err := r.Pool.Exec(ctx, InsertUser, user)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.Id, nil
}

func (r *Repository) FindUserUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.Pool.QueryRow(ctx, FindUserByEmail, email).Scan(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
