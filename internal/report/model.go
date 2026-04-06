package report

type SearchRequest struct {
	Barcode    string `db:"barcode" json:"barcode"`
	ExpiryDate string `db:"expiry_date" json:"expiry_date"`
}

type Product struct {
	ID         int    `db:"id" json:"id"`
	Barcode    string `db:"barcode" json:"barcode"`
	ExpiryDate string `db:"expiry_date" json:"expiry_date"`
}
