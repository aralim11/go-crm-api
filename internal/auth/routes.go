package auth

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterModule(router *http.ServeMux, db *sqlx.DB) {
	// router.Handle("/api/login", http.HandlerFunc(han))
}
