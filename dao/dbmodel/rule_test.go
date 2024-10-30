package dbmodel_test

import (
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/dao/dbmodel"
	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromModelRule(t *testing.T) {
	ruleID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	flagID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174111")
	type args struct {
		r             model.Rule
		featureFlagID uuid.UUID
		isDefault     bool
		orderInder    int
	}
	tests := []struct {
		name    string
		args    args
		want    dbmodel.Rule
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should convert default rule and ignore query, orderIndex and disable",
			args: args{
				r: model.Rule{
					ID:              ruleID.String(),
					Name:            "defaultRule",
					VariationResult: testutils.String("A"),
					Query:           `targetingKey eq "foo"`,
					Disable:         true,
				},
				featureFlagID: flagID,
				isDefault:     true,
				orderInder:    10,
			},

			wantErr: assert.NoError,
			want: dbmodel.Rule{
				ID:              ruleID,
				Name:            "defaultRule",
				FeatureFlagID:   flagID,
				Query:           testutils.String(""),
				Disable:         false,
				OrderIndex:      -1,
				VariationResult: testutils.String("A"),
				IsDefault:       true,
			},
		},
		{
			name: "should convert rule and keep query, orderIndex and disable if not default",
			args: args{
				r: model.Rule{
					ID:              ruleID.String(),
					Name:            "rule 1",
					VariationResult: testutils.String("A"),
					Query:           `targetingKey eq "foo"`,
					Disable:         true,
				},
				featureFlagID: flagID,
				isDefault:     false,
				orderInder:    10,
			},

			wantErr: assert.NoError,
			want: dbmodel.Rule{
				ID:              ruleID,
				Name:            "rule 1",
				FeatureFlagID:   flagID,
				Query:           testutils.String(`targetingKey eq "foo"`),
				Disable:         true,
				OrderIndex:      10,
				VariationResult: testutils.String("A"),
				IsDefault:       false,
			},
		},
		{
			name: "should convert rule with percentage",
			args: args{
				r: model.Rule{
					ID:   ruleID.String(),
					Name: "rule 1",
					Percentages: &map[string]float64{
						"A": 50,
						"B": 50,
					},
					Query:   `targetingKey eq "foo"`,
					Disable: true,
				},
				featureFlagID: flagID,
				isDefault:     false,
				orderInder:    10,
			},

			wantErr: assert.NoError,
			want: dbmodel.Rule{
				ID:            ruleID,
				Name:          "rule 1",
				FeatureFlagID: flagID,
				Query:         testutils.String(`targetingKey eq "foo"`),
				Disable:       true,
				OrderIndex:    10,
				IsDefault:     false,
				Percentages: testutils.JSONB(map[string]interface{}{
					"A": float64(50),
					"B": float64(50),
				}),
			},
		},
		{
			name: "should convert rule with progressive rollout",
			args: args{
				r: model.Rule{
					ID:   ruleID.String(),
					Name: "rule 1",
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
					Query:   `targetingKey eq "foo"`,
					Disable: true,
				},
				featureFlagID: flagID,
				isDefault:     false,
				orderInder:    10,
			},

			wantErr: assert.NoError,
			want: dbmodel.Rule{
				ID:                                  ruleID,
				Name:                                "rule 1",
				FeatureFlagID:                       flagID,
				Query:                               testutils.String(`targetingKey eq "foo"`),
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbmodel.FromModelRule(tt.args.r, tt.args.featureFlagID, tt.args.isDefault, tt.args.orderInder)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFromModelRuleCreateUUID(t *testing.T) {
	rule := model.Rule{
		Name: "rule 1",
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
		Query:   `targetingKey eq "foo"`,
		Disable: true,
	}

	got, err := dbmodel.FromModelRule(rule, uuid.New(), false, 1)
	require.NoError(t, err)
	assert.NotNil(t, got.ID)
	assert.NotEqual(t, uuid.Nil, got.ID)
}

func TestFromModelRuleErrorParsingUUID(t *testing.T) {
	rule := model.Rule{
		Name: "rule 1",
		ID:   "invalid-uuid",
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
		Query:   `targetingKey eq "foo"`,
		Disable: true,
	}

	_, err := dbmodel.FromModelRule(rule, uuid.New(), false, 1)
	require.Error(t, err)
	assert.Equal(t, "invalid UUID length: 12", err.Error())
}
