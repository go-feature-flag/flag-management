package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpTest(t *testing.T) api.Server {
	// init the in-memory mock dao
	dbImpl, err := dao.NewInMemoryMockDao()
	require.NoError(t, err)

	// init the API handlers
	flagHandlers := handler.NewFlagAPIHandler(dbImpl, nil)
	healthHandlers := handler.NewHealth(dbImpl)

	// port is not important since we are not really starting the server in the tests
	apiServer := api.New(":0", flagHandlers, healthHandlers)
	require.NotNil(t, apiServer)
	return *apiServer
}

func TestHealthRouteExist(t *testing.T) {
	apiServer := setUpTest(t)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	apiServer.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"API is up and running","code":200}`, rec.Body.String())
}
