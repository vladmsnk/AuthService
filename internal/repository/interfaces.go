package repository

import (
	"auth/vladmsnk/internal/entity"
	"context"
	"github.com/google/uuid"
)

type AuthRepo interface {
	SaveUser(ctx context.Context, user *entity.User) (uuid.UUID, error)
	FindUserUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserUserByUsernameAndPassword(ctx context.Context, username, password string) (*entity.User, error)
}
