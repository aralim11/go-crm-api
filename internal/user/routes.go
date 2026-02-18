package user

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *Handler) {
	router.Handle("/users", http.HandlerFunc(handler.CreateData))
}
