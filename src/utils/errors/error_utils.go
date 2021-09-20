package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Error   string `json:"error"`
}

func NewHTTPInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}

func NewHTTPBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad request to the server",
	}
}

func NewHTTPNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not found error",
	}
}
