package model_test

import (
	"testing"

	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFeatureFlag_GetRules(t *testing.T) {

	tests := []struct {
		name string
		flag *model.FeatureFlag
		want []model.Rule
	}{
		{
			name: "should return empty rules if not set",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: &model.Rule{
					ID:              uuid.New().String(),
					VariationResult: testutils.String("A"),
				},
			},
			want: []model.Rule{},
		},
		{
			name: "should return empty rules if set to nil",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: &model.Rule{
					ID:              uuid.New().String(),
					VariationResult: testutils.String("A"),
				},
				Rules: nil,
			},
			want: []model.Rule{},
		},
		{
			name: "should return empty rules if set to empty array",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: &model.Rule{
					ID:              uuid.New().String(),
					VariationResult: testutils.String("A"),
				},
				Rules: &[]model.Rule{},
			},
			want: []model.Rule{},
		},
		{
			name: "should return rules in the same order if rules are set",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: &model.Rule{
					ID:              uuid.New().String(),
					VariationResult: testutils.String("A"),
				},
				Rules: &[]model.Rule{
					{
						ID:              "875f9582-35ad-4912-a980-a66887e5de8f",
						Name:            "rule1",
						Query:           "query1 eq 'value1'",
						VariationResult: testutils.String("A"),
					},
					{
						ID:              "6ad0a484-3063-4d90-9536-ca0b8aaf43e6",
						Name:            "rule1",
						Query:           "query2 eq 'value2'",
						VariationResult: testutils.String("B"),
					},
				},
			},
			want: []model.Rule{
				{
					ID:              "875f9582-35ad-4912-a980-a66887e5de8f",
					Name:            "rule1",
					Query:           "query1 eq 'value1'",
					VariationResult: testutils.String("A"),
				},
				{
					ID:              "6ad0a484-3063-4d90-9536-ca0b8aaf43e6",
					Name:            "rule1",
					Query:           "query2 eq 'value2'",
					VariationResult: testutils.String("B"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.flag.GetRules())
		})
	}
}

func TestFeatureFlag_GetDefaultRule(t *testing.T) {
	tests := []struct {
		name string
		flag *model.FeatureFlag
		want model.Rule
	}{
		{
			name: "should return empty rules if not set",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
			},
			want: model.Rule{},
		},
		{
			name: "should return empty rules if set to nil",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: nil,
			},
			want: model.Rule{},
		},
		{
			name: "should return default rule if set",
			flag: &model.FeatureFlag{
				Variations: &map[string]*interface{}{
					"A": testutils.Interface("a"),
					"B": testutils.Interface("b"),
				},
				DefaultRule: &model.Rule{
					ID:              "875f9582-35ad-4912-a980-a66887e5de8f",
					VariationResult: testutils.String("A"),
				},
			},
			want: model.Rule{
				ID:              "875f9582-35ad-4912-a980-a66887e5de8f",
				VariationResult: testutils.String("A"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.flag.GetDefaultRule(), "GetDefaultRule()")
		})
	}
}
