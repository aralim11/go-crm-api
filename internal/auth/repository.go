package auth

import (
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindByEmail(email string) (*LoginRequest, error)
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindByEmail(email string) (*LoginRequest, error) {
	var user LoginRequest
	err := r.db.Get(&user, "SELECT email, password FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
