package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aralim11/go-crm-api/internal/utils/response"
	"github.com/aralim11/go-crm-api/internal/utils/validator"
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

	// validate request
	if validator.IsBlank(req.Name) {
		response.JsonResponse(w, http.StatusBadRequest, "Name is required", nil)
		return
	}

	if validator.IsBlank(req.Email) {
		response.JsonResponse(w, http.StatusBadRequest, "Email is required", nil)
		return
	}

	if validator.IsBlank(req.Mobile) {
		response.JsonResponse(w, http.StatusBadRequest, "Mobile is required", nil)
		return
	}

	if validator.IsBlank(req.Password) {
		response.JsonResponse(w, http.StatusBadRequest, "Password is required", nil)
		return
	}

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Password validation failed", err.Error())
		return
	}

	// create user
	user, err := h.service.Create(req)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	// respond with created user
	response.JsonResponse(w, http.StatusCreated, "User created successfully", user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// get users
	users, err := h.service.List()
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	// respond with no users found if empty
	if len(users) == 0 {
		response.JsonResponse(w, http.StatusOK, "No users found", nil)
		return
	}

	// respond with users
	response.JsonResponse(w, http.StatusOK, "Users fetched successfully", users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// extract user ID from URL
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 4 {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid URL format", nil)
		return
	}

	// validate user ID
	id := parts[3]
	if !validator.IsInteger(id) {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// convert user ID to int64
	idInt, err := validator.StrToInt64(id)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// get user by ID
	user, err := h.service.GetUserByID(idInt)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to fetch user", err.Error())
		return
	}

	// respond with no users found if empty
	if user == nil {
		response.JsonResponse(w, http.StatusOK, "No user found", nil)
		return
	}

	// respond with fetched user
	response.JsonResponse(w, http.StatusOK, "User fetched successfully", user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	// check method
	if r.Method != http.MethodPut {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// extract user ID from URL
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 5 {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid URL format", nil)
		return
	}

	// validate user ID
	id := parts[3]
	if !validator.IsInteger(id) {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// convert user ID to int64
	idInt, err := validator.StrToInt64(id)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// decode JSON
	var req UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	// validate request
	if validator.IsBlank(req.Name) {
		response.JsonResponse(w, http.StatusBadRequest, "Name is required", nil)
		return
	}

	if validator.IsBlank(req.Email) {
		response.JsonResponse(w, http.StatusBadRequest, "Email is required", nil)
		return
	}

	if validator.IsBlank(req.Mobile) {
		response.JsonResponse(w, http.StatusBadRequest, "Mobile is required", nil)
		return
	}

	// update user
	user, err := h.service.UpdateUser(&req, idInt)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	// respond with created user
	response.JsonResponse(w, http.StatusCreated, "User updated successfully", user)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodDelete {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// extract user ID from URL
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 5 {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid segment format", nil)
		return
	}

	// validate user ID
	id := parts[3]
	if !validator.IsInteger(id) {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// convert user ID to int64
	idInt, err := validator.StrToInt64(id)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// delete data
	err = h.service.DeleteUser(idInt)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	response.JsonResponse(w, http.StatusOK, "Delete successful", nil)
}
