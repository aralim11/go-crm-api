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
	UpdateUser(user *UpdateUserRequest, id int64) error
	DeleteUser(id int64) error
	FindByEmail(email string, id ...int64) (*UserResponse, error)
	FindByMobile(mobile string, id ...int64) (*UserResponse, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *User) (*User, error) {
	result, err := r.db.Exec(`INSERT INTO users (name, email, mobile, address, status, password)
		VALUES (?, ?, ?, ?, ?, ?)`, user.Name, user.Email, user.Mobile, user.Address, user.Status, user.Password,
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

func (r *userRepo) FindByEmail(email string, id ...int64) (*UserResponse, error) {
	var user UserResponse
	var err error

	if len(id) > 0 && id[0] != 0 {
		err = r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE email = ? AND id != ?", email, id[0])
	} else {
		err = r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE email = ?", email)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByMobile(mobile string, id ...int64) (*UserResponse, error) {
	var user UserResponse
	var err error

	if len(id) > 0 && id[0] != 0 {
		err = r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE mobile = ? AND id != ?", mobile, id[0])
	} else {
		err = r.db.Get(&user, "SELECT id, name, email, mobile, address FROM users WHERE mobile = ?", mobile)
	}

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

func (r *userRepo) UpdateUser(user *UpdateUserRequest, id int64) error {
	_, err := r.db.Exec("UPDATE users SET name=?, email=?, mobile=? WHERE id=?", user.Name, user.Email, user.Mobile, id)
	if err != nil {
		return err
	}

	return nil
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
