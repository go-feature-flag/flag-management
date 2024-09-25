package dbmodel

import (
	"encoding/json"
	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
	"time"
)

type Rule struct {
	ID              uuid.UUID `db:"id"`
	FeatureFlagID   uuid.UUID `db:"feature_flag_id"`
	IsDefault       bool      `db:"is_default"`
	Name            string    `db:"name"`
	Query           string    `db:"query"`
	VariationResult *string   `db:"variation_result"`
	Percentages     *string   `db:"percentages"` // JSONB is stored as string
	Disable         bool      `db:"disable"`
	CreatedDate     time.Time `db:"created_date"`
	LastUpdatedDate time.Time `db:"last_updated_date"`

	// De normalized and added fields in FeatureFlag
	ProgressiveRolloutInitialVariation  *string    `db:"progressive_rollout_initial_variation"`
	ProgressiveRolloutEndVariation      *string    `db:"progressive_rollout_end_variation"`
	ProgressiveRolloutInitialPercentage *float64   `db:"progressive_rollout_initial_percentage"`
	ProgressiveRolloutEndPercentage     *float64   `db:"progressive_rollout_end_percentage"`
	ProgressiveRolloutStartDate         *time.Time `db:"progressive_rollout_start_date"`
	ProgressiveRolloutEndDate           *time.Time `db:"progressive_rollout_end_date"`
}

func (f *Rule) ToAPI() (model.Rule, error) {
	var percentages *map[string]float64 = nil
	if f.Percentages != nil {
		err := json.Unmarshal([]byte(*f.Percentages), &percentages)
		if err != nil {
			return model.Rule{}, err
		}
	}

	res := model.Rule{
		Name:            f.Name,
		Query:           f.Query,
		VariationResult: f.VariationResult,
		Percentages:     percentages,
		Disable:         f.Disable,
	}

	if f.ProgressiveRolloutInitialVariation != nil &&
		f.ProgressiveRolloutEndVariation != nil &&
		f.ProgressiveRolloutInitialPercentage != nil &&
		f.ProgressiveRolloutEndPercentage != nil {

		res.ProgressiveRollout = &model.ProgressiveRollout{
			Initial: &model.ProgressiveRolloutStep{
				Variation:  f.ProgressiveRolloutInitialVariation,
				Percentage: f.ProgressiveRolloutInitialPercentage,
				Date:       f.ProgressiveRolloutStartDate,
			},
			End: &model.ProgressiveRolloutStep{
				Variation:  f.ProgressiveRolloutEndVariation,
				Percentage: f.ProgressiveRolloutEndPercentage,
				Date:       f.ProgressiveRolloutEndDate,
			},
		}
	}

	return res, nil
}

//func extractProgressiveRollout(progressiveRolloutId string, pr *model.ProgressiveRollout) (ProgressiveRollout, error) {
//	if pr == nil {
//		return ProgressiveRollout{}, nil
//	}
//	if pr.Initial == nil {
//		return ProgressiveRollout{}, errors.New("initial is required")
//	}
//	if pr.End == nil {
//		return ProgressiveRollout{}, errors.New("end is required")
//	}
//	if pr.Initial.Variation == nil {
//		return ProgressiveRollout{}, errors.New("initial variation is required")
//	}
//	if pr.End.Variation == nil {
//		return ProgressiveRollout{}, errors.New("end variation is required")
//	}
//	if pr.Initial.Date == nil {
//		return ProgressiveRollout{}, errors.New("initial date is required")
//	}
//	if pr.End.Date == nil {
//		return ProgressiveRollout{}, errors.New("end date is required")
//	}
//	if pr.Initial.Percentage == nil {
//		initial := float64(0)
//		pr.Initial.Percentage = &initial
//	}
//	if pr.End.Percentage == nil {
//		end := float64(100)
//		pr.End.Percentage = &end
//	}
//
//	id := progressiveRolloutId
//	if id == "" {
//		id = uuid.New()
//	}
//
//	return ProgressiveRollout{
//		ID:                id,
//		InitialVariation:  *pr.Initial.Variation,
//		EndVariation:      *pr.End.Variation,
//		InitialPercentage: *pr.Initial.Percentage,
//		EndPercentage:     *pr.End.Percentage,
//		StartDate:         *pr.Initial.Date,
//		EndDate:           *pr.End.Date,
//	}, nil
//
//}
