package model

type ErrorResponse struct {
	ErrorDetails string `json:"errorDetails,omitempty" example:"An error occurred"`
	Code         int    `json:"code" example:"500"`
}
