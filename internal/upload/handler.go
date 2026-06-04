package upload

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aralim11/go-crm-api/internal/utils/response"
	"github.com/aralim11/go-crm-api/internal/utils/validator"
	"github.com/google/uuid"
)

type Handler struct {
	service UploadService
}

func NewUploadHandler(service UploadService) *Handler {
	return &Handler{service: service}
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

func (h *Handler) CsvUpload(w http.ResponseWriter, r *http.Request) {
	// check request method
	if r.Method != http.MethodPost {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	// perse file from request
	file, header, err := r.FormFile("file")
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, "Error retrieving file", err.Error())
		return
	}
	defer file.Close()

	// get extension and validation
	extension := strings.ToLower(filepath.Ext(header.Filename))
	allowedExtensions := map[string]bool{
		".csv": true,
	}

	if !allowedExtensions[extension] {
		response.JsonResponse(w, http.StatusUnsupportedMediaType, "Invalid file format. Only csv allowed.", nil)
		return
	}

	// blank buffer for next usage
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Error resetting file pointer", nil)
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

	err = h.service.ProcessCSV(dstPath)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Processing error", err.Error())
		return
	}

}

func (h *Handler) PeopleList(w http.ResponseWriter, r *http.Request) {
	// check request method
	if r.Method != http.MethodGet {
		response.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	// get limit, page for pagination
	page := 1
	limit := 10

	// page
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	// limit
	if l := r.URL.Query().Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	// max limit protection
	if limit > 100 {
		limit = 100
	}

	// get peoples
	peoples, err := h.service.PeopleList(page, limit)
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to fetch peoples", err.Error())
		return
	}

	// respond with no peoples found if empty
	if len(peoples) == 0 {
		response.JsonResponse(w, http.StatusOK, "No peoples found", nil)
		return
	}

	// count total items
	count, err := h.service.Count()
	if err != nil {
		response.JsonResponse(w, http.StatusInternalServerError, "Failed to count peoples", err.Error())
		return
	}

	// process data
	peopleData := PeopleResponse{
		Data: peoples,
		Pagination: Pagination{
			Limit:       limit,
			CurrentPage: page,
			TotalItems:  count,
			TotalPages:  count / limit,
		},
	}

	// respond with peoples
	response.JsonResponse(w, http.StatusOK, "Peoples fetched successfully", peopleData)
}
