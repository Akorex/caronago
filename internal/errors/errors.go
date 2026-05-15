package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Raw     error  `json:"-"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Unwrap() error {
	return e.Raw
}

func NewApiError(code int, message string, raw error) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Raw:     raw,
	}
}

func ErrNotFound(resource string) *ApiError {
	return &ApiError{
		Code: http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", resource),
	}
}


func ErrBadRequest(message string) *ApiError{
	return &ApiError{
		Code: http.StatusBadRequest,
		Message: message,
	}
}

func ErrConflict(resource string) *ApiError{
	return &ApiError{
		Code: http.StatusConflict,
		Message: fmt.Sprintf("%s already exists", resource),
	}
}

func ErrUnauthorized(message string) *ApiError{
	return &ApiError{
		Code: http.StatusUnauthorized,
		Message: message,
	}

}

func ErrInternalServerError(err error) *ApiError{
	return &ApiError{
		Code: http.StatusInternalServerError,
		Message: "Internal Server Error",
		Raw: err,
	
	}
}