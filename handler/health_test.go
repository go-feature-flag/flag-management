package handler

import (
	"github.com/go-feature-flag/app-api/dao"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Health(t *testing.T) {
	e := echo.New()

	t.Run("API is up and running", func(t *testing.T) {
		mockDao, err := dao.NewInMemoryMockDao()
		require.NoError(t, err)
		mockDao.OnPingReturnError(false)

		h := NewHealth(mockDao)
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		require.NoError(t, h.Health(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"message":"API is up and running","code":200}`, rec.Body.String())
	})

	t.Run("Database not available", func(t *testing.T) {
		mockDao, err := dao.NewInMemoryMockDao()
		require.NoError(t, err)
		mockDao.OnPingReturnError(true)

		h := NewHealth(mockDao)
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		require.NoError(t, h.Health(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"errorDetails":"error on ping","code":500}`, rec.Body.String())
	})
}
