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

	router.Handle("/api/user-create", http.HandlerFunc(userHandler.CreateUser))
	router.Handle("/api/users", http.HandlerFunc(userHandler.GetUsers))

}
