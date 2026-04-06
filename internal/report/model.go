package report

type SearchRequest struct {
	Barcode    string `json:"barcode"`
	ExpiryDate string `json:"expiry_date"`
}
