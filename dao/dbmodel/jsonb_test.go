package dbmodel_test

import (
	"encoding/json"
	"testing"

	"github.com/go-feature-flag/app-api/dao/dbmodel"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var theTestMap = map[string]interface{}{
	"key":         "value",
	"another_key": "another_value",
	"nested_key": map[string]interface{}{
		"child_key": "child_value",
	},
}

var theTestString = `{"field":{"another_key":"another_value","key":"value","nested_key":{"child_key":"child_value"}}}`

type TestStruct struct {
	Field dbmodel.JSONB `json:"field"`
}

func TestJSONBMarshalling(t *testing.T) {
	theTest := TestStruct{
		Field: theTestMap,
	}
	b, err := json.Marshal(theTest)
	assert.Nil(t, err)
	assert.Equal(t, theTestString, string(b))
}

func TestJSONBUnmarshalling(t *testing.T) {
	theTest := TestStruct{}
	err := json.Unmarshal([]byte(theTestString), &theTest)
	assert.Nil(t, err)
	assert.Equal(t, dbmodel.JSONB(theTestMap), theTest.Field)
}
