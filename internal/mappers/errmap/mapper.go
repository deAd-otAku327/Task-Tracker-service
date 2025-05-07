package errmap

import (
	"task-tracker-service/internal/types/dto"
)

func MapToErrorResponse(err error, code int) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Message: err.Error(),
		Code:    code,
	}
}
