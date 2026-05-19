package upload

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterModule(router *http.ServeMux, db *sqlx.DB) {
	uploadRepository := NewUploadRepo(db)
	uploadService := NewUploadService(uploadRepository)
	uploadHandler := NewUploadHandler(uploadService)

	router.Handle("/api/upload", http.HandlerFunc(uploadHandler.Upload))

}
