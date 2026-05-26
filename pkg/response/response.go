package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, payload ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Error(w http.ResponseWriter, status int, errMsg string) {
	JSONResponse(w, status, ApiResponse{
		Success: false,
		Error:   errMsg,
	})
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	JSONResponse(w, status, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
