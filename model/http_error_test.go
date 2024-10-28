package model_test

import (
	"errors"
	"testing"

	"github.com/go-feature-flag/app-api/model"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPError(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		err      error
		wantCode int
		wantErr  model.HTTPError
	}{
		{
			name:     "should return HTTPError with given code and error",
			code:     404,
			err:      errors.New("not found"),
			wantCode: 404,
			wantErr:  model.HTTPError{ErrorDetails: "not found", Code: 404},
		},
		{
			name:     "should return HTTPError with empty error message if error is nil",
			code:     500,
			err:      nil,
			wantCode: 500,
			wantErr:  model.HTTPError{ErrorDetails: "No error passed, please report the issue", Code: 500},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotErr := model.NewHTTPError(tt.code, tt.err)
			assert.Equal(t, tt.wantCode, gotCode)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
