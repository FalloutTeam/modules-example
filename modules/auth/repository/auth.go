package repository

import (
	"database/sql"
	"errors"
)

type AuthRepository interface {
	Authenticate(username, password string) (string, error)
}

type authRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Authenticate(username, password string) (string, error) {
	var token string
	err := r.db.QueryRow("SELECT token FROM users WHERE username = $1 AND password = $2", username, password).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}
	return token, nil
}
