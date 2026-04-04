package report

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterModule(router *http.ServeMux, db *sqlx.DB) {
	reportRepository := NewReportRepo(db)
	reportService := NewReportService(reportRepository)
	reportHandler := NewReportHandler(reportService)

	router.Handle("/api/search", http.HandlerFunc(reportHandler.SearchData))

}
