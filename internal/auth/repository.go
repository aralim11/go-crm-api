package auth

import "github.com/jmoiron/sqlx"

type AuthRepository interface {
	FindByEmail(email string) (bool, error)
}

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepository {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) FindByEmail(email string) (bool, error) {
	return true, nil
}
