package user

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *User) (*User, error)
	List() ([]*UserResponse, error)
	GetUserByID(id int64) (*UserResponse, error)
	UpdateUser(user *UpdateUserRequest, id int64) (*UserResponse, error)
	DeleteUser(id int64) error
	FindByEmail(email string) (*UserResponse, error)
	FindByMobile(mobile string) (*UserResponse, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *User) (*User, error) {
	result, err := r.db.Exec(`INSERT INTO users (name, email, mobile, address, status)
		VALUES (?, ?, ?, ?, ?)`, user.Name, user.Email, user.Mobile, user.Address, user.Status,
	)

	// 🔥 check for error
	if err != nil {
		return nil, err
	}

	// 🔥 get auto increment ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

func (r *userRepo) List() ([]*UserResponse, error) {
	var users []*UserResponse

	err := r.db.Select(&users, "SELECT id, name, email, mobile, address FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) FindByEmail(email string) (*UserResponse, error) {
	var user UserResponse
	err := r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE email = ?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByMobile(mobile string) (*UserResponse, error) {
	var user UserResponse
	err := r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE mobile = ?", mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserByID(id int64) (*UserResponse, error) {
	var user UserResponse
	err := r.db.QueryRow("SELECT id, name, email, mobile, address FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Mobile, &user.Address)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(user *UpdateUserRequest, id int64) (*UserResponse, error) {
	var userReq UserResponse
	result, err := r.db.Exec("UPDATE users SET name=?, email=?, mobile=? WHERE id=?", &userReq.Name, &userReq.Email, &userReq.Mobile, id)
	if err != nil {
		return nil, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &userReq, nil
}

func (r *userRepo) DeleteUser(id int64) error {

	result, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}

	row, _ := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("User not found")
	}

	return nil
}
