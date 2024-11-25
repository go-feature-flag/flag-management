package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_customHTTPErrorHandler(t *testing.T) {
	t.Run("should return a 500 with correct format if error is not a echo.HTTPError", func(t *testing.T) {
		e := echo.New()
		e.HTTPErrorHandler = customHTTPErrorHandler
		e.GET("/", func(c echo.Context) error {
			return errors.New("toto")
		})
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"errorDetails":"Internal server error: toto","code":500}`, rec.Body.String())
	})
	t.Run("should return an error with correct code if echo.HTTPError", func(t *testing.T) {
		e := echo.New()
		e.HTTPErrorHandler = customHTTPErrorHandler
		e.GET("/", func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		})
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, `{"errorDetails":"not found","code":404}`, rec.Body.String())
	})
}
