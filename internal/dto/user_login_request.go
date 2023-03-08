package dto

type UserLoginRequest struct {
	Username string `json:"username"` //todo login by Username or Email or Password
	Password string `json:"password"`
}
