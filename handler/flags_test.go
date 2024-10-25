package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/dao"
	daoErr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlagsHandler_GetAllFeatureFlags(t *testing.T) {
	type test struct {
		name             string
		ctx              context.Context
		flags            []model.FeatureFlag
		expectedHTTPCode int
		expectedBody     string
	}
	tests := []test{
		{
			name:             "should return an empty array of flags if there are no flags",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            make([]model.FeatureFlag, 0),
			expectedBody:     "[]\n",
		},
		{
			name:             "should return a flag with a default rule",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `[{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a111","name":"flagr6w8","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}},{"id":"926214f3-80c1-46e6-a913-b2d40b92a222","name":"flagr576987209","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}]`,
		},
		{
			name:             "should return a 500 if an error occured ",
			ctx:              context.WithValue(context.Background(), "error", daoErr.UnknownError),
			expectedHTTPCode: http.StatusInternalServerError,
			flags:            make([]model.FeatureFlag, 0),
			expectedBody:     "{\"errorDetails\":\"error on get flags\",\"code\":500}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)

			h := handler.NewFlagAPIHandler(mockDao, nil)
			req := httptest.NewRequestWithContext(tt.ctx, http.MethodGet, "/v1/flags", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			require.NoError(t, h.GetAllFeatureFlags(c))
			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
			fmt.Println(rec.Body.String())
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestFlagsHandler_GetFeatureFlagByID(t *testing.T) {
	type test struct {
		name             string
		ctx              context.Context
		flags            []model.FeatureFlag
		ID               string
		expectedHTTPCode int
		expectedBody     string
	}
	tests := []test{
		{
			name:             "should return a 404 if the flag does not exist",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusNotFound,
			flags:            defaultInMemoryFlags(),
			ID:               "926214f3-80c1-46e6-a913-b2d40b92a965",
			expectedBody:     "{\"errorDetails\":\"flag not found\",\"code\":404}\n",
		},
		{
			name:             "should return a flag if the id exists",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			ID:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			expectedBody:     "{\"id\":\"926214f3-80c1-46e6-a913-b2d40b92a932\",\"name\":\"flag1\",\"createdDate\":\"2024-10-25T11:50:27Z\",\"lastUpdatedDate\":\"2024-10-25T11:50:27Z\",\"LastModifiedBy\":\"foo\",\"description\":\"description1\",\"type\":\"string\",\"variations\":{\"variation1\":\"A\",\"variation2\":\"B\"},\"defaultRule\":{\"id\":\"\",\"variation\":\"variation1\"}}\n",
		},
		{
			name:             "should return a 400 if the id is not a valid UUID",
			ctx:              context.WithValue(context.Background(), "error", daoErr.InvalidUUID),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			ID:               "invalidUUID",
			expectedBody:     "{\"errorDetails\":\"invalid UUID format\",\"code\":400}\n",
		},
		{
			name:             "should return a 500 if unknown error",
			ctx:              context.WithValue(context.Background(), "error", daoErr.UnknownError),
			expectedHTTPCode: http.StatusInternalServerError,
			flags:            defaultInMemoryFlags(),
			ID:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			expectedBody:     "{\"errorDetails\":\"error on get flag by id\",\"code\":500}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)

			h := handler.NewFlagAPIHandler(mockDao, nil)
			req := httptest.NewRequestWithContext(
				tt.ctx, http.MethodGet, "/v1/flags/:id", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.ID)

			require.NoError(t, h.GetFeatureFlagByID(c))
			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestFlagsHandler_CreateNewFlag(t *testing.T) {
	nilRules := make([]model.Rule, 1)
	type test struct {
		name             string
		ctx              context.Context
		flags            []model.FeatureFlag
		newFlag          model.FeatureFlag
		newFlagAsString  string // newFlag as string will take over newFlag if not empty
		expectedHTTPCode int
		expectedBody     string
	}
	tests := []test{
		{
			name:             "should return an error if the name is empty",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"flag name is required\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				Name:        "",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return an error if you start inserting a flag with the same name",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusConflict,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"flag with name flag1 already exists\",\"code\":409}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a932",
				Name:        "flag1",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return an error if malformed JSON",
			ctx:              context.WithValue(context.Background(), "error", daoErr.UnknownError),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"code=400, message=Unmarshal type error: expected=model.FeatureFlag, got=string, field=, offset=4, internal=json: cannot unmarshal string into Go value of type model.FeatureFlag\",\"code\":400}\n",
			newFlagAsString:  `"id":"926214f3-80c1-46e6-a913-b2d40b92a93","name":"flag2","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2024-10-25T11:50:27Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"}}`,
		},
		{
			name:             "should return an error if error when finding flag by name",
			ctx:              context.WithValue(context.Background(), "error", daoErr.UnknownError),
			expectedHTTPCode: http.StatusInternalServerError,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"error on get flag by name\",\"code\":500}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 400 if error when converting the body in db format",
			ctx:              context.WithValue(context.Background(), "error_create", daoErr.ConversionError),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"error creating flag\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 500 if error when calling the DB",
			ctx:              context.WithValue(context.Background(), "error_create", daoErr.DatabaseNotInitialized),
			expectedHTTPCode: http.StatusInternalServerError,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"error creating flag\",\"code\":500}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 400 if no default rule",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"flag default rule is required\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
			},
		},
		{
			name:             "should return a 400 if default rule has no variation result",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"flag default rule is invalid\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name: "defaultRule",
				},
			},
		},
		{
			name:             "should return a 400 if a targeting rule is empty",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"invalid rule rule1\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{{Name: "rule1"}},
			},
		},
		{
			name:             "should return a 400 if a targeting rule is nil",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"targeting rule is nil\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &nilRules,
			},
		},
		{
			name:             "should return a 400 if a targeting rule has no query",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"errorDetails\":\"query is required for targeting rules\",\"code\":400}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{
					{
						Name:            "rule1",
						VariationResult: testutils.String("variation1"),
					},
				},
			},
		},
		{
			name:             "should return a 201 and uuid if flag created",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusCreated,
			flags:            defaultInMemoryFlags(),
			expectedBody:     "{\"id\":\"926214f3-80c1-46e6-a913-b2d40b92a93\",\"name\":\"flag2\",\"createdDate\":\"2020-01-01T00:00:00Z\",\"lastUpdatedDate\":\"2020-01-01T00:00:00Z\",\"LastModifiedBy\":\"toto\",\"description\":\"description1\",\"type\":\"string\",\"variations\":{\"variation1\":\"A\",\"variation2\":\"B\"},\"targeting\":[{\"id\":\"\",\"name\":\"rule1\",\"query\":\"targetingKey eq \\\"value\\\"\",\"variation\":\"variation1\"}],\"defaultRule\":{\"id\":\"\",\"name\":\"defaultRule\",\"variation\":\"variation1\"}}\n",
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{
					{
						Name:            "rule1",
						Query:           "targetingKey eq \"value\"",
						VariationResult: testutils.String("variation1"),
					},
				},
			},
		},
		{
			name:             "should return a 400 if no variation type specified",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"flag type is required","code":400}`,
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{
					{
						Name:            "rule1",
						Query:           "targetingKey eq \"value\"",
						VariationResult: testutils.String("variation1"),
					},
				},
			},
		},
		{
			name:             "should return a 400 if empty variation type",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"flag type is required","code":400}`,
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{
					{
						Name:            "rule1",
						Query:           "targetingKey eq \"value\"",
						VariationResult: testutils.String("variation1"),
					},
				},
			},
		},
		{
			name:             "should return a 400 if not supported variation type",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"flag type notsupported not supported","code":400}`,
			newFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a93",
				Name:        "flag2",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "notsupported",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Unix(1729849827, 0),
				CreatedDate:     time.Unix(1729849827, 0),
				DefaultRule: &model.Rule{
					Name:            "defaultRule",
					VariationResult: testutils.String("variation1"),
				},
				Rules: &[]model.Rule{
					{
						Name:            "rule1",
						Query:           "targetingKey eq \"value\"",
						VariationResult: testutils.String("variation1"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)
			h := handler.NewFlagAPIHandler(mockDao, &handler.FlagAPIHandlerOptions{Clock: &testutils.ClockMock{}})

			var body io.Reader
			if tt.newFlagAsString != "" {
				body = bytes.NewReader([]byte(tt.newFlagAsString))
			} else {
				b, err := json.Marshal(tt.newFlag)
				require.NoError(t, err)
				body = bytes.NewReader(b)
			}

			req := httptest.NewRequestWithContext(
				tt.ctx, http.MethodPost, "/v1/flags", body)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			require.NoError(t, h.CreateNewFlag(c))
			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestFlagsHandler_DeleteFlagByID(t *testing.T) {
	type test struct {
		name                  string
		ctx                   context.Context
		flags                 []model.FeatureFlag
		id                    string
		expectedHTTPCode      int
		expectedBody          string
		expectedNumberOfFlags int
	}
	tests := []test{
		{
			name:             "should return an error if the id is not a valid UUID",
			ctx:              context.WithValue(context.Background(), "error_delete", daoErr.InvalidUUID),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"invalid UUID format","code":400}`,
			id:               "invalidUUID",
		},
		{
			name:                  "should not return an error if the id does not exist",
			ctx:                   context.Background(),
			expectedHTTPCode:      http.StatusNoContent,
			flags:                 defaultInMemoryFlags(),
			id:                    "926214f3-80c1-46e6-a913-b2d40b92a000",
			expectedNumberOfFlags: 3,
		},
		{
			name:                  "should not return an error if a real id is provided",
			ctx:                   context.Background(),
			expectedHTTPCode:      http.StatusNoContent,
			flags:                 defaultInMemoryFlags(),
			id:                    "926214f3-80c1-46e6-a913-b2d40b92a932",
			expectedNumberOfFlags: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)
			h := handler.NewFlagAPIHandler(mockDao, nil)
			req := httptest.NewRequestWithContext(
				tt.ctx, http.MethodDelete, "/v1/flags/:id", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)
			require.NoError(t, h.DeleteFlagByID(c))

			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
			if rec.Code == http.StatusNoContent {
				flagAfterDelete, err := mockDao.GetFlags(tt.ctx)
				require.NoError(t, err)
				assert.Equal(t, tt.expectedNumberOfFlags, len(flagAfterDelete))
			} else {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestFlagsHandler_UpdateFlagByID(t *testing.T) {
	type test struct {
		name                string
		ctx                 context.Context
		flags               []model.FeatureFlag
		updatedFlag         model.FeatureFlag
		updatedFlagAsString string
		id                  string
		expectedHTTPCode    int
		expectedBody        string
	}
	tests := []test{
		{
			name:             "should return a 400 if id is not a valid UUID",
			ctx:              context.WithValue(context.Background(), "error_update", daoErr.InvalidUUID),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"invalid UUID format","code":400}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			updatedFlag: model.FeatureFlag{
				Name:        "flag1",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
				CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 404 if id does not exist",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusNotFound,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"flag not found","code":404}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92aaaa",
			updatedFlag: model.FeatureFlag{
				Name:        "flag1",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("A"),
					"variation2": testutils.Interface("B"),
				},
				VariationType:   "string",
				LastModifiedBy:  "foo",
				LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
				CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:                "should return a 400 if json is malformed",
			ctx:                 context.Background(),
			expectedHTTPCode:    http.StatusBadRequest,
			flags:               defaultInMemoryFlags(),
			expectedBody:        `{"errorDetails":"code=400, message=unexpected EOF, internal=unexpected EOF","code":400}`,
			id:                  "926214f3-80c1-46e6-a913-b2d40b92a932",
			updatedFlagAsString: "{",
		},
		{
			name:             "should return a 200 if flag is modified",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"C","variation2":"D"},"defaultRule":{"id":"","variation":"variation1"}}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			updatedFlag: model.FeatureFlag{
				ID:          "926214f3-80c1-46e6-a913-b2d40b92a932",
				Name:        "flag1",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("C"),
					"variation2": testutils.Interface("D"),
				},
				VariationType:  "string",
				LastModifiedBy: "foo",
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 200 if flag is modified with no ID in body",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"C","variation2":"D"},"defaultRule":{"id":"","variation":"variation1"}}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			updatedFlag: model.FeatureFlag{
				Name:        "flag1",
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("C"),
					"variation2": testutils.Interface("D"),
				},
				VariationType:  "string",
				LastModifiedBy: "foo",
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
		{
			name:             "should return a 400 if modified flag is not valid (test without name)",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			expectedBody:     `{"errorDetails":"flag name is required","code":400}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
			updatedFlag: model.FeatureFlag{
				Description: testutils.String("description1"),
				Variations: &map[string]*interface{}{
					"variation1": testutils.Interface("C"),
					"variation2": testutils.Interface("D"),
				},
				VariationType:  "string",
				LastModifiedBy: "foo",
				DefaultRule: &model.Rule{
					VariationResult: testutils.String("variation1"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)
			options := &handler.FlagAPIHandlerOptions{
				Clock: testutils.ClockMock{},
			}
			h := handler.NewFlagAPIHandler(mockDao, options)

			var body io.Reader
			if tt.updatedFlagAsString != "" {
				body = bytes.NewReader([]byte(tt.updatedFlagAsString))
			} else {
				b, err := json.Marshal(tt.updatedFlag)
				require.NoError(t, err)
				body = bytes.NewReader(b)
			}

			req := httptest.NewRequestWithContext(
				tt.ctx, http.MethodPut, "/v1/flags/:id", body)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)
			require.NoError(t, h.UpdateFlagByID(c))
			fmt.Println(rec.Body.String())
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
		})
	}
}

func TestFlagsHandler_UpdateFeatureFlagStatus(t *testing.T) {
	type test struct {
		name             string
		ctx              context.Context
		flags            []model.FeatureFlag
		body             string
		id               string
		expectedHTTPCode int
		expectedBody     string
	}
	tests := []test{
		{
			name:             "should return a 200 if status updated (disabled)",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			body:             `{"disable": true}`,
			expectedBody:     `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":true}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
		},
		{
			name:             "should return a 200 if status updated (enabled)",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusOK,
			flags:            defaultInMemoryFlags(),
			body:             `{"disable": false}`,
			expectedBody:     `{"id":"926214f3-80c1-46e6-a913-b2d40b92a932","name":"flag1","createdDate":"2024-10-25T11:50:27Z","lastUpdatedDate":"2020-01-01T00:00:00Z","LastModifiedBy":"foo","description":"description1","type":"string","variations":{"variation1":"A","variation2":"B"},"defaultRule":{"id":"","variation":"variation1"},"disable":false}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
		},
		{
			name:             "should return a 404 if flag does not exist",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusNotFound,
			flags:            defaultInMemoryFlags(),
			body:             `{"disable": false}`,
			expectedBody:     `{"errorDetails":"flag not found","code":404}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92aaaa",
		},
		{
			name:             "should return a 400 if body is not a valid json",
			ctx:              context.Background(),
			expectedHTTPCode: http.StatusBadRequest,
			flags:            defaultInMemoryFlags(),
			body:             `{"disable": false`,
			expectedBody:     `{"errorDetails":"code=400, message=unexpected EOF, internal=unexpected EOF","code":400}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
		},
		{
			name:             "should return a 500 if db not available",
			ctx:              context.WithValue(context.Background(), "error_update", daoErr.DatabaseNotInitialized),
			expectedHTTPCode: http.StatusInternalServerError,
			flags:            defaultInMemoryFlags(),
			body:             `{"disable": false}`,
			expectedBody:     `{"errorDetails":"error on update flags","code":500}`,
			id:               "926214f3-80c1-46e6-a913-b2d40b92a932",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			mockDao, err := dao.NewInMemoryMockDao()
			require.NoError(t, err)
			mockDao.SetFlags(tt.flags)
			options := &handler.FlagAPIHandlerOptions{
				Clock: testutils.ClockMock{},
			}
			h := handler.NewFlagAPIHandler(mockDao, options)
			req := httptest.NewRequestWithContext(
				tt.ctx, http.MethodPut, "/v1/flags/:id/status", bytes.NewReader([]byte(tt.body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)
			require.NoError(t, h.UpdateFeatureFlagStatus(c))
			// TODO: delete 👇
			fmt.Println(rec.Body.String())
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			assert.Equal(t, tt.expectedHTTPCode, rec.Code)
		})
	}
}

func defaultInMemoryFlags() []model.FeatureFlag {
	return []model.FeatureFlag{
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a932",
			Name:        "flag1",
			Description: testutils.String("description1"),
			Variations: &map[string]*interface{}{
				"variation1": testutils.Interface("A"),
				"variation2": testutils.Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: testutils.String("variation1"),
			},
		},
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a111",
			Name:        "flagr6w8",
			Description: testutils.String("description1"),
			Variations: &map[string]*interface{}{
				"variation1": testutils.Interface("A"),
				"variation2": testutils.Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: testutils.String("variation1"),
			},
		},
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a222",
			Name:        "flagr576987209",
			Description: testutils.String("description1"),
			Variations: &map[string]*interface{}{
				"variation1": testutils.Interface("A"),
				"variation2": testutils.Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: testutils.String("variation1"),
			},
		},
	}
}
