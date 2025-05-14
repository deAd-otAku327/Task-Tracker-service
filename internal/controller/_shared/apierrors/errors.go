package apierrors

import "errors"

var (
	ErrAuthenticationFailed = errors.New("authentication failed")

	ErrInvalidRequestBody         = errors.New("invalid request body")
	ErrInvalidRequestParams       = errors.New("invalid request params")
	ErrInvalidRequestParamsFormat = errors.New("invalid format of request params")
)
