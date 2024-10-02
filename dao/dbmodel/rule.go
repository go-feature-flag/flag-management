package dbmodel

import (
	"encoding/json"
	"time"

	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
)

type Rule struct {
	ID                                  uuid.UUID  `db:"id"`
	FeatureFlagID                       uuid.UUID  `db:"feature_flag_id"`
	IsDefault                           bool       `db:"is_default"`
	Name                                string     `db:"name"`
	Query                               string     `db:"query"`
	VariationResult                     *string    `db:"variation_result"`
	Percentages                         *string    `db:"percentages"` // JSONB is stored as string
	Disable                             bool       `db:"disable"`
	ProgressiveRolloutInitialVariation  *string    `db:"progressive_rollout_initial_variation"`
	ProgressiveRolloutEndVariation      *string    `db:"progressive_rollout_end_variation"`
	ProgressiveRolloutInitialPercentage *float64   `db:"progressive_rollout_initial_percentage"`
	ProgressiveRolloutEndPercentage     *float64   `db:"progressive_rollout_end_percentage"`
	ProgressiveRolloutStartDate         *time.Time `db:"progressive_rollout_start_date"`
	ProgressiveRolloutEndDate           *time.Time `db:"progressive_rollout_end_date"`
	OrderIndex                          int        `db:"order_index"`
}

func FromModelRule(mr model.Rule, featureFlagID uuid.UUID, isDefault bool, orderIndex int) (Rule, error) {
	var id uuid.UUID
	if mr.ID != "" {
		var err error
		id, err = uuid.Parse(mr.ID)
		if err != nil {
			return Rule{}, err
		}
	} else {
		id = uuid.New()
	}

	var percentages *string
	if mr.Percentages != nil {
		percentagesJSON, err := json.Marshal(mr.Percentages)
		if err != nil {
			return Rule{}, err
		}
		percentagesStr := string(percentagesJSON)
		percentages = &percentagesStr
	}

	dbr := Rule{
		ID:            id,
		FeatureFlagID: featureFlagID,
		IsDefault:     isDefault,
		Name:          mr.Name,
		Query:         mr.Query,
		Disable:       mr.Disable,
		OrderIndex:    orderIndex,
	}

	if mr.VariationResult != nil {
		dbr.VariationResult = mr.VariationResult
	}

	if percentages != nil {
		dbr.Percentages = percentages
	}

	if mr.ProgressiveRollout != nil {
		dbr.ProgressiveRolloutInitialVariation = mr.ProgressiveRollout.Initial.Variation
		dbr.ProgressiveRolloutEndVariation = mr.ProgressiveRollout.End.Variation
		dbr.ProgressiveRolloutInitialPercentage = mr.ProgressiveRollout.Initial.Percentage
		dbr.ProgressiveRolloutEndPercentage = mr.ProgressiveRollout.End.Percentage
		dbr.ProgressiveRolloutStartDate = mr.ProgressiveRollout.Initial.Date
		dbr.ProgressiveRolloutEndDate = mr.ProgressiveRollout.End.Date
	}
	return dbr, nil
}

func (rule *Rule) ToAPI() (model.Rule, error) {
	apiRule := model.Rule{
		ID:      rule.ID.String(),
		Name:    rule.Name,
		Query:   rule.Query,
		Disable: rule.Disable,
	}

	if rule.VariationResult != nil {
		apiRule.VariationResult = rule.VariationResult
	}

	if rule.Percentages != nil {
		var percentages map[string]float64
		err := json.Unmarshal([]byte(*rule.Percentages), &percentages)
		if err != nil {
			return model.Rule{}, err
		}
		apiRule.Percentages = &percentages
	}

	if rule.ProgressiveRolloutInitialVariation != nil || rule.ProgressiveRolloutEndVariation != nil {
		apiRule.ProgressiveRollout = &model.ProgressiveRollout{
			Initial: &model.ProgressiveRolloutStep{
				Variation:  rule.ProgressiveRolloutInitialVariation,
				Percentage: rule.ProgressiveRolloutInitialPercentage,
				Date:       rule.ProgressiveRolloutStartDate,
			},
			End: &model.ProgressiveRolloutStep{
				Variation:  rule.ProgressiveRolloutEndVariation,
				Percentage: rule.ProgressiveRolloutEndPercentage,
				Date:       rule.ProgressiveRolloutEndDate,
			},
		}
	}
	return apiRule, nil
}
