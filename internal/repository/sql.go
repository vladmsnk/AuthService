package repository

const (
	InsertUser = `
	INSERT INTO users (id, username, email, password, number) VALUES ($1, $2, $3, $4, $5);
	`
	FindUserByEmail = `SELECT * FROM users WHERE email = $1;`

	FindUserByUsernameAndPassword = `SELECT * FROM users WHERE username = $1 AND password = $2;`
)
