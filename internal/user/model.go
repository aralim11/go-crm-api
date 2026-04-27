package user

import "time"

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Mobile    string    `db:"mobile"`
	Address   *string   `db:"address"`
	Status    int8      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
}

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
