package errs

import "net/http"

type AppErr struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e *AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
}

func NewAppErr(code int, message string) *AppErr {
	return &AppErr{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return NewAppErr(http.StatusNotFound, message)
}

func NewBadRequestError(message string) *AppErr {
	return NewAppErr(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) *AppErr {
	return NewAppErr(http.StatusInternalServerError, message)
}

func NewConflictError(message string) *AppErr {
	return NewAppErr(http.StatusConflict, message)
}

func NewForbiddenError(message string) *AppErr {
	return NewAppErr(http.StatusForbidden, message)
}

func NewUnauthorizedError(message string) *AppErr {
	return NewAppErr(http.StatusUnauthorized, message)
}
