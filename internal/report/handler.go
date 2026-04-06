package report

import (
	"encoding/json"
	"net/http"

	"github.com/aralim11/go-crm-api/internal/utils/response"
)

type Handler struct {
	service ReportService
}

func NewReportHandler(service ReportService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SearchData(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed!!", nil)
		return
	}

	// decode JSON
	var req SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	// validate
	// hash password
	// check email exists
	// search from DB
	result := h.service.SearchData()
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to search data", nil)
		return
	}

	response.JsonResponse(w, http.StatusOK, "Data found", result)

}
