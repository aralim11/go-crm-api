package report

import "net/http"

type Handler struct {
	service ReportService
}

func NewReportHandler(service ReportService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SearchData(r http.ResponseWriter, req *http.Request) {
	h.service.Search()
}
