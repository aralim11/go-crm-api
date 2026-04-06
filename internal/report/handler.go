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

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
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
	result, err := h.service.SearchData(req.Barcode, req.ExpiryDate)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JsonResponse(w, http.StatusOK, "Data found", result)

}
