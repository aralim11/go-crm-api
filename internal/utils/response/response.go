package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
	Data       any `json:"data,omitempty"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	resData := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	return json.NewEncoder(w).Encode(resData)
}
