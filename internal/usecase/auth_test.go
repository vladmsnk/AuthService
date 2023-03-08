package usecase

import (
	"auth/vladmsnk/config"
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/entity"
	mock_repository "auth/vladmsnk/internal/mocks"
	"auth/vladmsnk/internal/util"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

var (
	testUserRequest = dto.UserRegisterRequest{
		Email:    "vyumoiseenkov@gmail.com",
		Password: "password",
		Number:   "943293392",
		Username: "vladmsnk",
	}

	testLoginRequest = dto.UserLoginRequest{
		Username: "vladmsnk",
		Password: "password",
	}
	testContext = context.Background()
)

func TestRegisterUserAlreadyExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testUser := &entity.User{}

	mockRepo := mock_repository.NewMockAuthRepo(mockCtrl)

	mockRepo.EXPECT().FindUserUserByEmail(gomock.Any(), "vyumoiseenkov@gmail.com").Return(testUser, nil).
		Times(1)
	testUseCase := NewAuthUseCase(mockRepo, config.Auth{})

	_, err := testUseCase.CreateUser(testContext, testUserRequest)

	if !errors.Is(err, util.ErrUserAlreadyExists) {
		t.Fail()
	}
}

func TestRegisterUserRegistered(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	testUUID := uuid.UUID{}
	mockRepo := mock_repository.NewMockAuthRepo(mockCtrl)

	mockRepo.EXPECT().FindUserUserByEmail(gomock.Any(), "vyumoiseenkov@gmail.com").Return(nil, nil).
		Times(1)

	mockRepo.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return(testUUID, nil).Times(1)

	testUseCase := NewAuthUseCase(mockRepo, config.Auth{HashSalt: "sugar"})

	_, err := testUseCase.CreateUser(testContext, testUserRequest)
	if err != nil {
		t.Fail()
	}
}

func TestGenerateTokenUserNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_repository.NewMockAuthRepo(mockCtrl)

	mockRepo.EXPECT().FindUserUserByUsernameAndPassword(gomock.Any(), "vladmsnk", gomock.Any()).
		Return(nil, nil).Times(1)

	testUseCase := NewAuthUseCase(mockRepo, config.Auth{HashSalt: "sugar"})

	_, err := testUseCase.GenerateToken(testContext, testLoginRequest)
	if !errors.Is(err, util.ErrUserNotFound) {
		t.Fail()
	}
}

func TestGenerateTokenUserWithoutErrors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testUser := &entity.User{}

	mockRepo := mock_repository.NewMockAuthRepo(mockCtrl)
	mockRepo.EXPECT().FindUserUserByUsernameAndPassword(gomock.Any(), "vladmsnk", gomock.Any()).
		Return(testUser, nil).Times(1)
	testUseCase := NewAuthUseCase(mockRepo, config.Auth{HashSalt: "sugar", SigningKey: "key"})
	_, err := testUseCase.GenerateToken(testContext, testLoginRequest)
	if err != nil {
		t.Fail()
	}
}
