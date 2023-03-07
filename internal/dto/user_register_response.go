package dto

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Number   string `json:"number"`
}
