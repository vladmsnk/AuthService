package dto

import (
	"auth/vladmsnk/internal/entity"
	"github.com/google/uuid"
)

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Number   string `json:"number"`
}

func (u UserRegisterRequest) FromDTO(userID uuid.UUID, passwordHash string) entity.User {
	return entity.User{Id: userID, Username: u.Username, Email: u.Email, PasswordHash: passwordHash, Number: u.Number}
}
