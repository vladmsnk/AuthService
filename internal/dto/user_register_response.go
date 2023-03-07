package dto

import "auth/vladmsnk/internal/entity"

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Number   string `json:"number"`
}

func ToDTO(u entity.User) UserRegisterResponse {
	return UserRegisterResponse{Username: u.Username, Email: u.Email, Number: u.Number}
}
