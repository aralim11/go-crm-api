package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *User) error
	FindAll() error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *User) error {
	fmt.Println("User Repository", user)
	return nil
}

func (r *userRepo) FindAll() error {
	return nil
}
