package daoerr_test

import (
	"fmt"
	daoerr "github.com/go-feature-flag/flag-management/server/dao/err"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDaoError(t *testing.T) {
	type args struct {
		code daoerr.DaoErrorCode
		err  error
	}
	tests := []struct {
		name     string
		args     args
		wantErr  error
		wantCode daoerr.DaoErrorCode
	}{
		{
			name: "Should be able to create a new DaoError",
			args: args{
				code: daoerr.NotFound,
				err:  fmt.Errorf("not found"),
			},
			wantErr:  fmt.Errorf("not found"),
			wantCode: daoerr.NotFound,
		},
		{
			name: "Should not fail with a nil error",
			args: args{
				code: daoerr.NotFound,
				err:  nil,
			},
			wantErr:  fmt.Errorf("unknown error"),
			wantCode: daoerr.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := daoerr.NewDaoError(tt.args.code, tt.args.err)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
			assert.Equal(t, tt.wantCode, err.Code())
		})
	}
}
