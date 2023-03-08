package usecase

import (
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/entity"
	"context"
	"github.com/google/uuid"
)

type (
	Greet interface {
		Greet(_ context.Context) (dto.GreetResponse, error)
	}

	Auth interface {
		CreateUser(ctx context.Context, request dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
		GenerateToken(ctx context.Context, request dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	AuthRepo interface {
		SaveUser(ctx context.Context, user *entity.User) (uuid.UUID, error)
		FindUserUserByEmail(ctx context.Context, email string) (*entity.User, error)
		FindUserUserByUsernameAndPassword(ctx context.Context, username, password string) (*entity.User, error)
	}
)
