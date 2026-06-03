package upload

import (
	"github.com/jmoiron/sqlx"
)

type UploadRepository interface {
	PeopleList() ([]*People, error)
	CsvUpload(people *People) error
}

type uploadRepo struct {
	db *sqlx.DB
}

func NewUploadRepo(db *sqlx.DB) UploadRepository {
	return &uploadRepo{db: db}
}

func (r *uploadRepo) PeopleList() ([]*People, error) {
	var peoples []*People

	err := r.db.Select(&peoples, "SELECT * FROM peoples limit 1000")
	if err != nil {
		return nil, err
	}

	return peoples, nil
}

func (r *uploadRepo) CsvUpload(people *People) error {
	_, err := r.db.Exec(
		"INSERT INTO peoples (user_id, first_name, last_name, sex, email, phone, date_of_birth, job_title) VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		people.UserID, people.FirstName, people.LastName, people.Sex, people.Email, people.Phone, people.DateOfBirth, people.JobTitle,
	)
	return err
}
