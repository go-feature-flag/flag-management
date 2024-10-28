package dbmodel_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/dao/dbmodel"
	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFromModelFeatureFlag(t *testing.T) {
	flagId, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	tests := []struct {
		name    string
		mff     model.FeatureFlag
		want    dbmodel.FeatureFlag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should convert model.FeatureFlag to dbmodel.FeatureFlag",
			mff: model.FeatureFlag{
				ID:          flagId.String(),
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: &map[string]interface{}{
					"A": "a",
					"B": "b",
				},
				VariationType: "boolean",
				BucketingKey:  testutils.String("teamID"),
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DefaultRule: &model.Rule{
					ID:              "123e4567-e89b-12d3-a456-426614174111",
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
			},
			wantErr: assert.NoError,
			want: dbmodel.FeatureFlag{
				ID:          flagId,
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: dbmodel.JSONB(map[string]interface{}{
					"A": "a",
					"B": "b",
				}),
				Type:         model.FlagTypeBoolean,
				BucketingKey: testutils.String("teamID"),
				Metadata: dbmodel.JSONB(map[string]interface{}{
					"key": "value",
				}),
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}, {
			name: "should convert model.FeatureFlag to dbmodel.FeatureFlag without metadata",
			mff: model.FeatureFlag{
				ID:          flagId.String(),
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: &map[string]interface{}{
					"A": "a",
					"B": "b",
				},
				VariationType:   "boolean",
				BucketingKey:    testutils.String("teamID"),
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DefaultRule: &model.Rule{
					ID:              "123e4567-e89b-12d3-a456-426614174111",
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
			},
			wantErr: assert.NoError,
			want: dbmodel.FeatureFlag{
				ID:          flagId,
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: dbmodel.JSONB(map[string]interface{}{
					"A": "a",
					"B": "b",
				}),
				Type:            model.FlagTypeBoolean,
				BucketingKey:    testutils.String("teamID"),
				Metadata:        nil,
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}, {
			name: "should convert model.FeatureFlag to dbmodel.FeatureFlag without variation",
			mff: model.FeatureFlag{
				ID:              flagId.String(),
				Name:            "my-flag",
				Description:     testutils.String("my flag description"),
				VariationType:   "boolean",
				BucketingKey:    testutils.String("teamID"),
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DefaultRule: &model.Rule{
					ID:              "123e4567-e89b-12d3-a456-426614174111",
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
			},
			wantErr: assert.NoError,
			want: dbmodel.FeatureFlag{
				ID:              flagId,
				Name:            "my-flag",
				Description:     testutils.String("my flag description"),
				Variations:      nil,
				Type:            model.FlagTypeBoolean,
				BucketingKey:    testutils.String("teamID"),
				Metadata:        nil,
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "should error if invalid UUID",
			mff: model.FeatureFlag{
				ID:              "invalid-uuid",
				Name:            "my-flag",
				Description:     testutils.String("my flag description"),
				VariationType:   "boolean",
				BucketingKey:    testutils.String("teamID"),
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DefaultRule: &model.Rule{
					ID:              "123e4567-e89b-12d3-a456-426614174111",
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should error if empty ID",
			mff: model.FeatureFlag{
				ID:              "",
				Name:            "my-flag",
				Description:     testutils.String("my flag description"),
				VariationType:   "boolean",
				BucketingKey:    testutils.String("teamID"),
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdatedDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DefaultRule: &model.Rule{
					ID:              "123e4567-e89b-12d3-a456-426614174111",
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbmodel.FromModelFeatureFlag(tt.mff)
			if !tt.wantErr(t, err, fmt.Sprintf("FromModelFeatureFlag(%v)", tt.mff)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FromModelFeatureFlag(%v)", tt.mff)
		})
	}
}
