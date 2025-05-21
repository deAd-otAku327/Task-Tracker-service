package responser

import (
	"encoding/json"
	"net/http"
	"task-tracker-service/internal/types/dto"
)

const (
	contentTypeHeader = "Content-Type"
	setCookieHeader   = "Set-Cookie"
	contentTypeJSON   = "application/json"
)

func MakeResponseJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data) //nolint:errcheck
	}
}

func MakeErrorResponseJSON(w http.ResponseWriter, err *dto.ErrorResponse) {
	MakeResponseJSON(w, err.Code, err)
}
