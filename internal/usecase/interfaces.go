package usecase

import (
	"auth/vladmsnk/internal/dto"
	"context"
)

type (
	Greet interface {
		Greet(_ context.Context) (dto.GreetResponse, error)
	}

	Auth interface {
		CreateUser(ctx context.Context, request dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
		GenerateToken(ctx context.Context, request dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}
)
