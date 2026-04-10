package user

import (
	"encoding/json"
	"net/http"

	"github.com/aralim11/go-crm-api/internal/utils/response"
)

type Handler struct {
	service UserService
}

func NewUserHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// decode JSON
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	// create user
	user, err := h.service.Create(req)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	// respond with created user
	response.JsonResponse(w, http.StatusCreated, "User created successfully", user)
}
