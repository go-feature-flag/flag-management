package model

func NewHTTPError(code int, err error) (int, HTTPError) {
	return code, HTTPError{ErrorDetails: err.Error(), Code: code}
}

type HTTPError struct {
	ErrorDetails string `json:"errorDetails" example:"An error occurred"`
	Code         int    `json:"code" example:"500"`
}
