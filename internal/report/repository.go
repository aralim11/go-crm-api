package report

import "github.com/jmoiron/sqlx"

type ReportRepository interface {
	SearchData()
}

type reportRepo struct {
	db *sqlx.DB
}

func NewReportRepo(db *sqlx.DB) ReportRepository {
	return &reportRepo{db: db}
}

func (r *reportRepo) SearchData() {

}
