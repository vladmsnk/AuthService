package usecase

import (
	"auth/vladmsnk/config"
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/util"
	"context"
	"github.com/google/uuid"
)

type AuthUseCase struct {
	authRepo AuthRepo
	cfg      config.Auth
}

func NewAuthUseCase(ar AuthRepo, authCfg config.Auth) *AuthUseCase {
	return &AuthUseCase{
		authRepo: ar,
		cfg:      authCfg,
	}
}

func (uc *AuthUseCase) CreateUser(ctx context.Context,
	request dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {

	user, err := uc.authRepo.FindUserUserByEmail(ctx, request.Email)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}
	if err == nil && user != nil {
		return dto.UserRegisterResponse{}, util.ErrUserAlreadyExists
	}

	userID := uuid.New()

	hashedPassword := util.HashPassword(request.Password, uc.cfg.HashSalt)

	userEntity := request.FromDTO(userID, hashedPassword)
	_, err = uc.authRepo.SaveUser(ctx, &userEntity)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	return dto.ToDTO(userEntity), nil

}

func (uc *AuthUseCase) GenerateToken(ctx context.Context, request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	hashedPassword := util.HashPassword(request.Password, uc.cfg.HashSalt)

	user, err := uc.authRepo.FindUserUserByUsernameAndPassword(ctx, request.Username, hashedPassword)

	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	if err == nil && user == nil {
		return dto.UserLoginResponse{}, util.ErrUserNotFound
	}

	tokenString, err := util.GenerateJWT(user.Email, user.Username, uc.cfg.SigningKey, uc.cfg.TokenTTL)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{Token: tokenString}, nil
}
