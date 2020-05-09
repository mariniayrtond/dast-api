package presenter

import (
	"fmt"
	"net/http"
)

type apiError struct {
	ErrorCode  string `json:"error_code"`
	StatusCode int    `json:"status_code"`
	Cause      string `json:"cause"`
}

func NewBadRequest(errorCode string, cause error) apiError {
	return apiError{
		ErrorCode:  errorCode,
		StatusCode: http.StatusBadRequest,
		Cause:      cause.Error(),
	}
}

func NewInternalServerError(errorCode string, cause error) apiError {
	return apiError{
		ErrorCode:  errorCode,
		StatusCode: http.StatusInternalServerError,
		Cause:      cause.Error(),
	}
}

func NewNotFound(resource string) apiError {
	return apiError{
		ErrorCode:  fmt.Sprintf("resource:%s not found", resource),
		StatusCode: http.StatusNotFound,
	}
}

func NewUnauthorized(resource string) apiError {
	return apiError{
		ErrorCode:  "unauthorized",
		StatusCode: http.StatusUnauthorized,
		Cause:      fmt.Sprintf("%s is not logged in", resource),
	}
}
