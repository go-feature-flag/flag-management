package handler_test

import (
	"github.com/go-feature-flag/flag-management/server/api"
	"github.com/go-feature-flag/flag-management/server/config"
	"github.com/go-feature-flag/flag-management/server/dao"
	"github.com/go-feature-flag/flag-management/server/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthHandler_Health(t *testing.T) {
	t.Run("API is up and running", func(t *testing.T) {
		mockDao, err := dao.NewInMemoryMockDao()
		require.NoError(t, err)
		mockDao.OnPingReturnError(false)

		hh := handler.NewHealthHandler(mockDao)
		hf := handler.NewFlagAPIHandler(mockDao, nil)
		s, err := api.New(&config.Configuration{
			Mode: "development",
		}, handler.Handlers{
			HealthHandler:  &hh,
			FlagAPIHandler: &hf,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		s.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"message":"API is up and running","code":200}`, rec.Body.String())
	})

	t.Run("Database not available", func(t *testing.T) {
		mockDao, err := dao.NewInMemoryMockDao()
		require.NoError(t, err)
		mockDao.OnPingReturnError(true)

		hh := handler.NewHealthHandler(mockDao)
		hf := handler.NewFlagAPIHandler(mockDao, nil)
		s, err := api.New(&config.Configuration{
			Mode: "development",
		}, handler.Handlers{
			HealthHandler:  &hh,
			FlagAPIHandler: &hf,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		s.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"errorDetails":"error on ping","code":500}`, rec.Body.String())
	})
}
