package auth

import "net/http"

type Handler struct {
	handler AuthService
}

func NewAuthHandler(handler AuthService) *Handler {
	return &Handler{
		handler: handler,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}
