package model_test

import (
	"testing"

	"github.com/go-feature-flag/app-api/model"
	"github.com/stretchr/testify/assert"
)

func TestFlagTypeFromValue(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      model.FlagType
		expectErr bool
	}{
		{"valid boolean type", "boolean", model.FlagTypeBoolean, false},
		{"valid string type", "string", model.FlagTypeString, false},
		{"valid integer type", "integer", model.FlagTypeInteger, false},
		{"valid double type", "double", model.FlagTypeDouble, false},
		{"valid json type", "json", model.FlagTypeJSON, false},
		{"empty type", "", "", true},
		{"unsupported type", "unsupported", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.FlagTypeFromValue(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
