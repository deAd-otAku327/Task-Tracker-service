package errmap

import (
	"task-tracker-service/internal/types/dto"
)

func MapToErrorResponse(msg string, code int) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Message: msg,
		Code:    code,
	}
}
