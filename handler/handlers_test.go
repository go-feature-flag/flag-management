package handler_test

import (
	"fmt"
	"testing"

	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitHandlers(t *testing.T) {
	mockDao, err := dao.NewInMemoryMockDao()
	require.NoError(t, err)
	expectedFlagAPIHandler := handler.NewFlagAPIHandler(mockDao, &handler.FlagAPIHandlerOptions{})
	expectedHealthHandler := handler.NewHealthHandler(mockDao)

	tests := []struct {
		name        string
		dao         dao.FlagStorage
		want        handler.Handlers
		wantErr     assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name:        "should error if a nil dao is passed",
			dao:         nil,
			want:        handler.Handlers{},
			wantErr:     assert.Error,
			expectedErr: handler.ErrMissingDao,
		},
		{
			name: "should return a handler with a dao",
			dao:  mockDao,
			want: handler.Handlers{
				FlagAPIHandler: &expectedFlagAPIHandler,
				HealthHandler:  &expectedHealthHandler,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler.InitHandlers(tt.dao)
			if !tt.wantErr(t, err, fmt.Sprintf("InitHandlers(%v)", tt.dao)) {
				assert.Equal(t, tt.expectedErr, err)
				return
			}
			assert.Equalf(t, tt.want, got, "InitHandlers(%v)", tt.dao)
		})
	}
}
