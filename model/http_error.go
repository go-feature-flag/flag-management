package model

func NewHTTPError(code int, err error) (int, HTTPError) {
	if err == nil {
		return code, HTTPError{ErrorDetails: "No error passed, please report the issue", Code: code}
	}
	return code, HTTPError{ErrorDetails: err.Error(), Code: code}
}

type HTTPError struct {
	ErrorDetails string `json:"errorDetails,omitempty" example:"An error occurred"`
	Code         int    `json:"code" example:"500"`
}
