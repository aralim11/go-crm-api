package user

import (
	"fmt"
	"net/http"
)

type Handler struct {
	service UserService
}

func NewUserHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateData(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {

	}
	// decode JSON
	// validate
	// hash password
	// check email exists
	// insert into DB

	fmt.Println("From User Handler")
	h.service.Create("abdul alim", "aralim@gmail")
}
