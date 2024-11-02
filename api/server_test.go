package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpTest(t *testing.T) (api.Server, int) {
	// init the in-memory mock dao
	dbImpl, err := dao.NewInMemoryMockDao()
	dbImpl.SetFlags(testutils.DefaultInMemoryFlags())
	require.NoError(t, err)

	// init the API handlers
	options := &handler.FlagAPIHandlerOptions{
		Clock: testutils.ClockMock{},
	}

	flagHandlers := handler.NewFlagAPIHandler(dbImpl, options)
	healthHandlers := handler.NewHealthHandler(dbImpl)
	h := handler.Handlers{
		FlagAPIHandler: &flagHandlers,
		HealthHandler:  &healthHandlers,
	}

	port, err := testutils.GetFreePort()
	require.NoError(t, err)
	apiServer, err := api.New(fmt.Sprintf(":%d", port), h)
	require.NoError(t, err)
	require.NotNil(t, apiServer)
	return *apiServer, port
}

func TestHealthRouteExist(t *testing.T) {
	apiServer, _ := setUpTest(t)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	apiServer.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"API is up and running","code":200}`, rec.Body.String())
}

func TestRouteExist(t *testing.T) {
	apiServer, _ := setUpTest(t)
	tests := []struct {
		name     string
		method   string
		path     string
		body     *string
		wantCode int
		wantBody string
	}{
		{
			name:     "GET /v1/flags/:id",
			method:   http.MethodGet,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}`,
		},
		{
			name:     "GET /health",
			method:   http.MethodGet,
			path:     "/health",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `{"message":"API is up and running","code":200}`,
		},
		{
			name:     "GET /v1/flags",
			method:   http.MethodGet,
			path:     "/v1/flags",
			body:     nil,
			wantCode: http.StatusOK,
			wantBody: `[{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a111","name":"flagr6w8","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a222","name":"flagr576987209","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}]`,
		},
		{
			name:     "PATCH /v1/flags/:id/status",
			method:   http.MethodPatch,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932/status",
			body:     testutils.String(`{"disable":true}`),
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "POST /v1/flags",
			method:   http.MethodPost,
			path:     "/v1/flags",
			body:     testutils.String(`{"id":"926214f3-80c1-46e6-a913-b2d40b92a933","name":"flag2","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`),
			wantCode: http.StatusCreated,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a933","name":"flag2","createdDate":"2020-01-01T00:00:00Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"toto","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "PATCH /v1/flags/:id",
			method:   http.MethodPut,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     testutils.String(`{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`),
			wantCode: http.StatusOK,
			wantBody: `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
		},
		{
			name:     "DELETE /v1/flags/:id",
			method:   http.MethodDelete,
			path:     "/v1/flags/926214f3-80c1-46e6-a913-b2d40b92a932",
			body:     nil,
			wantCode: http.StatusNoContent,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body string = ""
			if tt.body != nil {
				body = *tt.body
			}
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			apiServer.ServeHTTP(rec, req)
			assert.Equal(t, tt.wantCode, rec.Code)
			fmt.Println(rec.Body.String())
			if tt.wantBody == "" {
				assert.Empty(t, rec.Body.String())
				return
			}
			assert.JSONEq(t, tt.wantBody, rec.Body.String())
		})
	}
}

func TestServerIsStartingAndStopping(t *testing.T) {
	apiServer, port := setUpTest(t)
	require.NotNil(t, apiServer)

	go apiServer.Start()
	// wait for the server to start or fail after 4 seconds
	time.Sleep(2 * time.Second)

	url := fmt.Sprintf("http://localhost:%d/health", port)
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	errStop := apiServer.Stop()
	require.NoError(t, errStop)
	time.Sleep(2 * time.Second)
	_, err = http.DefaultClient.Do(req)
	assert.Error(t, err)
}

func TestNoValidHandlers(t *testing.T) {
	daomock, _ := dao.NewInMemoryMockDao()
	mockHandlers, err := handler.InitHandlers(daomock)
	require.NoError(t, err)

	tests := []struct {
		name        string
		handlers    handler.Handlers
		wantErr     assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "no health handler",
			handlers: handler.Handlers{
				FlagAPIHandler: mockHandlers.FlagAPIHandler,
			},
			wantErr:     assert.Error,
			expectedErr: handler.ErrMissingHealthHandler,
		},
		{
			name: "no flag handler",
			handlers: handler.Handlers{
				HealthHandler: mockHandlers.HealthHandler,
			},
			wantErr:     assert.Error,
			expectedErr: handler.ErrMissingFlagAPIHandler,
		},
		{
			name: "all handlers provided",
			handlers: handler.Handlers{
				HealthHandler:  mockHandlers.HealthHandler,
				FlagAPIHandler: mockHandlers.FlagAPIHandler,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := api.New(":8080", tt.handlers)
			tt.wantErr(t, err)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
