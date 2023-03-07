package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password"`
	Number       string    `db:"number"`
}
