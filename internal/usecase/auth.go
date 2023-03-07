package usecase

import (
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/util"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type AuthUseCase struct {
	authRepo AuthRepo
}

func NewAuthUseCase(ar AuthRepo) *AuthUseCase {
	return &AuthUseCase{
		authRepo: ar,
	}
}

func (uc *AuthUseCase) CreateUser(ctx context.Context,
	request dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {

	_, err := uc.authRepo.FindUserUserByEmail(ctx, request.Email)
	if err != nil && err != pgx.ErrNoRows {
		return dto.UserRegisterResponse{}, err
	}

	userID := uuid.New()
	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	userEntity := request.FromDTO(userID, hashedPassword)

	_, err = uc.authRepo.SaveUser(ctx, userEntity)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	return dto.ToDTO(userEntity), nil

}

func (uc *AuthUseCase) GenerateToken(ctx context.Context, request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, err := uc.authRepo.FindUserUserByEmail(ctx, request.Email)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	credErr := util.CheckPassword(request.Password, user.PasswordHash)
	if credErr != nil {
		return dto.UserLoginResponse{}, credErr
	}
	tokenString, err := util.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{Token: tokenString}, nil
}
