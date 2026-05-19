package report

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aralim11/go-crm-api/internal/utils/response"
	"github.com/aralim11/go-crm-api/internal/utils/validator"
	"github.com/google/uuid"
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

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	// check request method
	if r.Method != http.MethodPost {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	// max request and file size validation
	const maxRequestBytesReader = 10 << 20
	const maxFileSize = 6 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBytesReader)
	err := r.ParseMultipartForm(maxRequestBytesReader)
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Request is too large", nil)
		return
	}

	// perse text field and validate
	title := r.FormValue("title")
	if validator.IsBlank(title) {
		response.JsonResponse(w, http.StatusBadRequest, "title is required", nil)
		return
	}

	// perse file from request
	file, header, err := r.FormFile("image")
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Error retrieving profile picture", err.Error())
		return
	}
	defer file.Close()

	// exact file size validation
	if header.Size > maxFileSize {
		response.JsonResponse(w, http.StatusBadRequest, "File size exceeds the 2MB limit", nil)
		return
	}

	// get extension and validation
	extension := strings.ToLower(filepath.Ext(header.Filename))
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".pdf":  true,
	}

	if !allowedExtensions[extension] {
		response.JsonResponse(w, http.StatusUnsupportedMediaType, "Invalid file format. Only JPG, JPEG, PNG, and GIF allowed.", nil)
		return
	}

	// validate (MIME check)
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil && err != io.EOF {
		response.JsonResponse(w, http.StatusInternalServerError, "Error reading file contents", err.Error())
		return
	}

	// blank buffer for next usage
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Error resetting file pointer", nil)
		return
	}

	// check content type
	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		response.JsonResponse(w, http.StatusUnsupportedMediaType, "The uploaded file is not a valid image binary", nil)
		return
	}

	// check directory
	err = os.MkdirAll("./storage", os.ModePerm)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to create uploads directory", err.Error())
		return
	}

	// create destination file
	newName := uuid.New().String() + extension
	dstPath := filepath.Join("./storage", newName)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to create destination file", nil)
		return
	}
	defer dstFile.Close()

	// copy file to the destination file
	if _, err := io.Copy(dstFile, file); err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to save file", err.Error())
		return
	}

	// return response
	response.JsonResponse(w, http.StatusOK, "Image successfully uploaded", nil)
}
