package auth

import (
	"encoding/json"
	"net/http"

	"github.com/aralim11/go-crm-api/internal/utils/response"
	"github.com/aralim11/go-crm-api/internal/utils/validator"
)

type Handler struct {
	handler AuthService
}

func NewAuthHandler(handler AuthService) *Handler {
	return &Handler{
		handler: handler,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// decode JSON
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	// validate request
	if validator.IsBlank(req.Email) {
		response.JsonResponse(w, http.StatusBadRequest, "Name is required", nil)
		return
	}

	if validator.IsBlank(req.Password) {
		response.JsonResponse(w, http.StatusBadRequest, "Password is required", nil)
		return
	}

	res, err := h.handler.LoginCheck(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response.JsonResponse(w, http.StatusCreated, "Login successfully", res)
}
