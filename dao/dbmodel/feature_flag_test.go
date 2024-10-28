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
	flagID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	tests := []struct {
		name    string
		mff     model.FeatureFlag
		want    dbmodel.FeatureFlag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should convert model.FeatureFlag to dbmodel.FeatureFlag",
			mff: model.FeatureFlag{
				ID:          flagID.String(),
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
				ID:          flagID,
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
				ID:          flagID.String(),
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
				ID:          flagID,
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
				ID:              flagID.String(),
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
				ID:              flagID,
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

func TestToModelFeatureFlag(t *testing.T) {
	flagID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	ruleIdDefault, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174111")
	ruleId1, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174333")
	ruleId2, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174222")

	tests := []struct {
		name   string
		dbFF   dbmodel.FeatureFlag
		dbRule []dbmodel.Rule
		want   model.FeatureFlag
	}{
		{
			name: "should convert model.FeatureFlag to dbmodel.FeatureFlag",
			dbFF: dbmodel.FeatureFlag{
				ID:          flagID,
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: dbmodel.JSONB(map[string]interface{}{
					"A": "a",
					"B": "b",
				}),
				Type:         model.FlagTypeString,
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
			dbRule: []dbmodel.Rule{
				{
					ID:            ruleId1,
					Name:          "Rule 1",
					FeatureFlagID: flagID,
					Disable:       false,
					Percentages:   dbmodel.JSONB(map[string]interface{}{"A": float64(50), "B": float64(50)}),
					IsDefault:     false,
					Query:         `targetingKey eq "foo"`,
					OrderIndex:    6,
				},
				{
					ID:                                  ruleId2,
					Name:                                "rule 2",
					FeatureFlagID:                       flagID,
					Query:                               `targetingKey eq "bar"`,
					Disable:                             true,
					OrderIndex:                          10,
					IsDefault:                           false,
					ProgressiveRolloutInitialVariation:  testutils.String("A"),
					ProgressiveRolloutEndVariation:      testutils.String("B"),
					ProgressiveRolloutInitialPercentage: testutils.Float64(0),
					ProgressiveRolloutEndPercentage:     testutils.Float64(100),
					ProgressiveRolloutStartDate:         testutils.Time(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
					ProgressiveRolloutEndDate:           testutils.Time(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:              ruleIdDefault,
					Name:            "defaultRule",
					FeatureFlagID:   flagID,
					Disable:         false,
					VariationResult: testutils.String("A"),
					IsDefault:       true,
				},
			},
			want: model.FeatureFlag{
				ID:          flagID.String(),
				Name:        "my-flag",
				Description: testutils.String("my flag description"),
				Variations: &map[string]interface{}{
					"A": "a",
					"B": "b",
				},
				VariationType: model.FlagTypeString,
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
					ID:              ruleIdDefault.String(),
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
				},
				Rules: &[]model.Rule{
					{
						ID:          ruleId1.String(),
						Name:        "Rule 1",
						Query:       `targetingKey eq "foo"`,
						Disable:     false,
						Percentages: &map[string]float64{"A": 50, "B": 50},
					},
					{
						ID:      ruleId2.String(),
						Name:    "rule 2",
						Query:   `targetingKey eq "bar"`,
						Disable: true,
						ProgressiveRollout: &model.ProgressiveRollout{
							Initial: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("A"),
								Percentage: testutils.Float64(0),
								Date:       testutils.Time(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
							},
							End: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("B"),
								Percentage: testutils.Float64(100),
								Date:       testutils.Time(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dbFF.ToModelFeatureFlag(tt.dbRule)
			assert.Equalf(t, tt.want, got, "FromModelFeatureFlag")
		})
	}
}
