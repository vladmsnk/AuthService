package usecase

import (
	"auth/vladmsnk/internal/dto"
	"context"
)

type GreetUseCase struct {
}

func (gu *GreetUseCase) Greet(ctx context.Context) (dto.GreetResponse, error) {
	return dto.GreetResponse{Greeting: ""}, nil
}
