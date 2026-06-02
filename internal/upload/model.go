package upload

import "time"

type People struct {
	ID          int64
	UserID      string
	FirstName   string
	LastName    string
	Sex         string
	Email       string
	Phone       string
	DateOfBirth time.Time
	JobTitle    string
}
