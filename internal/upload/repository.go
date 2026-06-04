package upload

import (
	"github.com/jmoiron/sqlx"
)

type UploadRepository interface {
	PeopleList(page int, limit int) ([]*People, error)
	CsvUpload(people *People) error
	Count() (int, error)
}

type uploadRepo struct {
	db *sqlx.DB
}

func NewUploadRepo(db *sqlx.DB) UploadRepository {
	return &uploadRepo{db: db}
}

func (r *uploadRepo) PeopleList(page int, limit int) ([]*People, error) {
	var peoples []*People
	offset := (page - 1) * limit

	err := r.db.Select(&peoples, "SELECT * FROM peoples limit ? offset ?", limit, offset)
	if err != nil {
		return nil, err
	}

	return peoples, nil
}

func (r *uploadRepo) Count() (int, error) {
	var total int

	err := r.db.QueryRow("SELECT COUNT(*) FROM peoples").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *uploadRepo) CsvUpload(people *People) error {
	_, err := r.db.Exec(
		"INSERT INTO peoples (user_id, first_name, last_name, sex, email, phone, date_of_birth, job_title) VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		people.UserID, people.FirstName, people.LastName, people.Sex, people.Email, people.Phone, people.DateOfBirth, people.JobTitle,
	)
	return err
}
