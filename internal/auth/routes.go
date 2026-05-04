package auth

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterModule(router *http.ServeMux, db *sqlx.DB) {
	// handlers
	authRepository := NewAuthRepository(db)
	authService := NewAuthService(authRepository)
	authHandler := NewAuthHandler(authService)

	// route lists
	router.Handle("/api/login", http.HandlerFunc(authHandler.Login))
}
