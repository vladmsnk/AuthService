package repository

const (
	InsertUser = `
	INSERT INTO users (id, username, email, password, number) VALUES
		$1, $2. $3, $4, $5 ON CONFLICT (id) DO NOTHING;
	`

	FindUserByEmail = `SELECT * FROM users WHERE email = $1;`
)
