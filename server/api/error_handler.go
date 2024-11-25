package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomErr struct {
	Code         int    `json:"code"`
	ErrorDetails string `json:"errorDetails"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		_ = c.JSON(he.Code, CustomErr{
			Code:         he.Code,
			ErrorDetails: fmt.Sprintf("%v", he.Message),
		})
		return
	}

	_ = c.JSON(http.StatusInternalServerError, CustomErr{
		Code:         http.StatusInternalServerError,
		ErrorDetails: "Internal server error: " + err.Error(),
	})
}
