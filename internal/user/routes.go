package user

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterModule(router *http.ServeMux, db *sqlx.DB) {
	// handlers
	userRepository := NewUserRepo(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	router.Handle("/users", http.HandlerFunc(userHandler.CreateData))
}
