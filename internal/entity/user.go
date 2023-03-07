package entity

import (
	"auth/vladmsnk/internal/dto"
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password"`
	Number       string    `db:"number"`
}

func (u User) ToDTO() dto.UserRegisterResponse {
	return dto.UserRegisterResponse{Username: u.Username, Email: u.Email, Number: u.Number}
}
