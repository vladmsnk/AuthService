package usecase

import (
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/util"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	_, found, err := uc.authRepo.FindUserUserByEmail(ctx, request.Email)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}
	if err == nil && found {
		return dto.UserRegisterResponse{}, util.ErrUserAlreadyExists
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
	user, found, err := uc.authRepo.FindUserUserByEmail(ctx, request.Email)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	if err == nil && !found {
		return dto.UserLoginResponse{}, util.ErrUserNotFound
	}
	credErr := util.CheckPassword(request.Password, user.PasswordHash)
	if credErr != nil {
		switch {
		case errors.Is(credErr, bcrypt.ErrMismatchedHashAndPassword):
			return dto.UserLoginResponse{}, util.ErrInvalidPassword
		default:
			return dto.UserLoginResponse{}, err
		}
	}
	tokenString, err := util.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{Token: tokenString}, nil
}
