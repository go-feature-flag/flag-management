package api

import (
	"errors"
	"net/http"

	"github.com/go-feature-flag/flag-management/server/model"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		_ = c.JSON(he.Code, model.ErrorResponse{ErrorDetails: err.Error(), Code: he.Code})
		return
	}
	_ = c.JSON(http.StatusInternalServerError, model.ErrorResponse{ErrorDetails: err.Error(), Code: http.StatusInternalServerError})
}
