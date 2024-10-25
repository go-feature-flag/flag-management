package handler

import (
	"context"
	"fmt"
	"github.com/go-feature-flag/app-api/dao"
	daoErr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
			flags: []model.FeatureFlag{
				{
					ID:          "926214f3-80c1-46e6-a913-b2d40b92a932",
					Name:        "flag1",
					Description: testutils.String("description1"),
					Variations: &map[string]*interface{}{
						"variation1": testutils.Interface("A"),
						"variation2": testutils.Interface("B"),
					},
					LastModifiedBy:  "foo",
					LastUpdatedDate: time.Unix(1729849827, 0),
					CreatedDate:     time.Unix(1729849827, 0),
					DefaultRule: &model.Rule{
						VariationResult: testutils.String("variation1"),
					},
				},
			},
			expectedBody: "[{\"id\":\"926214f3-80c1-46e6-a913-b2d40b92a932\",\"name\":\"flag1\",\"createdDate\":\"2024-10-25T11:50:27+02:00\",\"lastUpdatedDate\":\"2024-10-25T11:50:27+02:00\",\"LastModifiedBy\":\"foo\",\"description\":\"description1\",\"type\":\"\",\"variations\":{\"variation1\":\"A\",\"variation2\":\"B\"},\"defaultRule\":{\"id\":\"\",\"variation\":\"variation1\"}}]\n",
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

			h := NewFlagAPIHandler(mockDao)
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
