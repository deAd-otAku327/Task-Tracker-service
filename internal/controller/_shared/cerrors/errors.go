package cerrors

import "errors"

var (
	ErrInvalidRequestBody         = errors.New("invalid request body")
	ErrInvalidRequestParams       = errors.New("invalid request params")
	ErrInvalidRequestParamsFormat = errors.New("invalid format of request params")
)
