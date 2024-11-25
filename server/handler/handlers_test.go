package handler_test

import (
	"fmt"
	dao2 "github.com/go-feature-flag/flag-management/server/dao"
	handler2 "github.com/go-feature-flag/flag-management/server/handler"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitHandlers(t *testing.T) {
	mockDao, err := dao2.NewInMemoryMockDao()
	require.NoError(t, err)
	expectedFlagAPIHandler := handler2.NewFlagAPIHandler(mockDao, &handler2.FlagAPIHandlerOptions{})
	expectedHealthHandler := handler2.NewHealthHandler(mockDao)

	tests := []struct {
		name        string
		dao         dao2.FlagStorage
		want        handler2.Handlers
		wantErr     assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name:        "should error if a nil dao is passed",
			dao:         nil,
			want:        handler2.Handlers{},
			wantErr:     assert.Error,
			expectedErr: handler2.ErrMissingDao,
		},
		{
			name: "should return a handler with a dao",
			dao:  mockDao,
			want: handler2.Handlers{
				FlagAPIHandler: &expectedFlagAPIHandler,
				HealthHandler:  &expectedHealthHandler,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler2.InitHandlers(tt.dao)
			if !tt.wantErr(t, err, fmt.Sprintf("InitHandlers(%v)", tt.dao)) {
				assert.Equal(t, tt.expectedErr, err)
				return
			}
			assert.Equalf(t, tt.want, got, "InitHandlers(%v)", tt.dao)
		})
	}
}
