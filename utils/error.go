package utils

type ErrorBadRequest struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorNotFound struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorUnauthorized struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}