package upload

import "time"

type People struct {
	ID          int64     `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Sex         string    `db:"sex" json:"sex"`
	Email       string    `db:"email" json:"email"`
	Phone       string    `db:"phone" json:"phone"`
	DateOfBirth time.Time `db:"date_of_birth" json:"date_of_birth"`
	JobTitle    string    `db:"job_title" json:"job_title"`
}
