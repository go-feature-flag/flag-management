package testutils

import (
	"time"

	"github.com/go-feature-flag/flag-management/server/model"
)

func DefaultInMemoryFlags() []model.FeatureFlag {
	return []model.FeatureFlag{
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a932",
			Name:        "flag1",
			Description: String("description1"),
			Variations: &map[string]interface{}{
				"variation1": Interface("A"),
				"variation2": Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: String("variation1"),
			},
		},
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a111",
			Name:        "flagr6w8",
			Description: String("description1"),
			Variations: &map[string]interface{}{
				"variation1": Interface("A"),
				"variation2": Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: String("variation1"),
			},
		},
		{
			ID:          "926214f3-80c1-46e6-a913-b2d40b92a222",
			Name:        "flagr576987209",
			Description: String("description1"),
			Variations: &map[string]interface{}{
				"variation1": Interface("A"),
				"variation2": Interface("B"),
			},
			VariationType:   "string",
			LastModifiedBy:  "foo",
			LastUpdatedDate: time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			CreatedDate:     time.Date(2024, 10, 25, 11, 50, 27, 0, time.UTC),
			DefaultRule: &model.Rule{
				VariationResult: String("variation1"),
			},
		},
	}
}
