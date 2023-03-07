package dto

type UserLoginRequest struct {
	Email    string `json:"email"` //todo login by Username or Email or Password
	Password string `json:"password"`
}
