package cmd_test

import (
	"github.com/go-feature-flag/flag-management/server/cmd"
	"github.com/go-feature-flag/flag-management/server/dao"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGOFeatureFlagManagementAPICommand(t *testing.T) {
	mockDao, _ := dao.NewInMemoryMockDao()
	defaultOptions := cmd.APICommandOptions{
		OverrideDefaultDao: mockDao,
	}
	tests := []struct {
		name           string
		options        cmd.APICommandOptions
		wantErr        assert.ErrorAssertionFunc
		expectedErrMsg string
	}{
		{
			name:    "should not have any error when creating a new GOFeatureFlagManagementAPICommand",
			options: defaultOptions,
			wantErr: assert.NoError,
		},
		{
			name:           "should have an error when creating a new GOFeatureFlagManagementAPICommand with no access to DB",
			options:        cmd.APICommandOptions{},
			wantErr:        assert.Error,
			expectedErrMsg: "error while initializing dependencies: impossible to initialize database connection: connection string is empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cmd.NewGOFeatureFlagManagementAPICommand(tt.options)
			tt.wantErr(t, err)
			if tt.expectedErrMsg != "" {
				assert.Equal(t, tt.expectedErrMsg, err.Error())
				return
			}
		})
	}
}

func TestNewGOFeatureFlagManagementAPICommandShouldStartAPI(t *testing.T) {
	mockDao, _ := dao.NewInMemoryMockDao()
	defaultOptions := cmd.APICommandOptions{
		OverrideDefaultDao: mockDao,
	}
	cli, err := cmd.NewGOFeatureFlagManagementAPICommand(defaultOptions)
	require.NoError(t, err)
	go cli.Run()
	time.Sleep(200 * time.Millisecond)
	resp, err := http.Get("http://localhost:3001/health")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
