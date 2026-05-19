package upload

import "github.com/jmoiron/sqlx"

type UploadRepository interface {
	ImageUpload()
}

type uploadRepo struct {
	db *sqlx.DB
}

func NewUploadRepo(db *sqlx.DB) UploadRepository {
	return &uploadRepo{db: db}
}

func (r *uploadRepo) ImageUpload() {

}
