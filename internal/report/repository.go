package report

import "github.com/jmoiron/sqlx"

type ReportRepository interface {
	SearchData(barcode string, expiryDate string) (*Product, error)
}

type reportRepo struct {
	db *sqlx.DB
}

func NewReportRepo(db *sqlx.DB) ReportRepository {
	return &reportRepo{db: db}
}

func (r *reportRepo) SearchData(barcode string, expiryDate string) (*Product, error) {
	query := "SELECT id FROM products WHERE name = ?"

	var product Product
	err := r.db.QueryRow(query, barcode).Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
