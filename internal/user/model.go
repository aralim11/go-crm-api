package user

import "time"

type User struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Mobile    string    `db:"mobile" json:"mobile"`
	Address   *string   `db:"address" json:"address,omitempty"`
	Status    int8      `db:"status" json:"status"`
	Password  string    `db:"password" json:"password,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `db:"password"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
}

type UserResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Address string `json:"address,omitempty"`
}

type UpdateUserRequest struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Mobile  string  `json:"mobile"`
	Address *string `json:"address,omitempty"`
}
