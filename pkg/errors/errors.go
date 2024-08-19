package errors

import (
    "fmt"
    "net/http"
)

type HTTPError struct {
    StatusCode int
    Message    string
}

func (e *HTTPError) Error() string {
    return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

func NewHTTPError(statusCode int, message string) *HTTPError {
    return &HTTPError{
        StatusCode: statusCode,
        Message:    message,
    }
}

// Predefined errors
var (
    ErrBadRequest          = NewHTTPError(http.StatusBadRequest, "Bad Request")
    ErrUnauthorized        = NewHTTPError(http.StatusUnauthorized, "Unauthorized")
    ErrNotFound            = NewHTTPError(http.StatusNotFound, "Resource Not Found")
    ErrInternalServerError = NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
)
