package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type customErr struct {
	Code         int    `json:"code"`
	ErrorDetails string `json:"errorDetails"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		_ = c.JSON(he.Code, customErr{
			Code:         he.Code,
			ErrorDetails: fmt.Sprintf("%v", he.Message),
		})
		return
	}

	_ = c.JSON(http.StatusInternalServerError, customErr{
		Code:         http.StatusInternalServerError,
		ErrorDetails: "Internal server error: " + err.Error(),
	})
}
