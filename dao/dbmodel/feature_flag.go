package dbmodel

import (
	"encoding/json"
	"errors"
	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
	"time"
)

type FeatureFlag struct {
	ID              string         `db:"id"`
	Name            string         `db:"name"`
	Description     *string        `db:"description"`
	Variations      string         `db:"variations"` // JSONB is stored as string
	Type            model.FlagType `db:"type"`       // variationType is stored as string
	BucketingKey    *string        `db:"bucketing_key"`
	Metadata        *string        `db:"metadata"` // JSONB is stored as string
	TrackEvents     *bool          `db:"track_events"`
	Disable         *bool          `db:"disable"`
	Version         *string        `db:"version"`
	CreatedDate     time.Time      `db:"created_date"`
	LastUpdatedDate time.Time      `db:"last_updated_date"`
}

func (f *FeatureFlag) ToAPI(rules []Rule) (model.FeatureFlag, error) {
	var variations *map[string]*interface{}
	err := json.Unmarshal([]byte(f.Variations), &variations)
	if err != nil {
		return model.FeatureFlag{}, err
	}

	var metadata *map[string]interface{} = nil
	if f.Metadata != nil {
		err = json.Unmarshal([]byte(*f.Metadata), &metadata)
		if err != nil {
			return model.FeatureFlag{}, err
		}
	}

	var convertedRules []model.Rule
	var defaultRule *model.Rule = nil

	if rules != nil {
		convertedRules = make([]model.Rule, 0, len(rules))
		for _, rule := range rules {
			convertedRule, err := rule.ToAPI()
			if err != nil {
				return model.FeatureFlag{}, err
			}
			if rule.IsDefault && defaultRule == nil {
				defaultRule = &convertedRule
			} else {
				convertedRules = append(convertedRules, convertedRule)
			}
		}
	}

	return model.FeatureFlag{
		ID:              f.ID,
		Name:            f.Name,
		CreatedDate:     f.CreatedDate,
		LastUpdatedDate: f.LastUpdatedDate,
		Description:     f.Description,
		VariationType:   f.Type,
		Rules:           &convertedRules,
		DefaultRule:     defaultRule,
		Variations:      variations,
		BucketingKey:    f.BucketingKey,
		Disable:         f.Disable,
		Version:         f.Version,
		Metadata:        metadata,
	}, nil

}

func NewFeatureFlag(f model.FeatureFlag) (FeatureFlag, error) {
	id := f.ID
	if id == "" {
		id = uuid.New().String()
	}

	// Variations
	if f.Variations == nil {
		return FeatureFlag{}, errors.New("variations is required")
	}
	variations, err := json.Marshal(f.Variations)
	if err != nil {
		return FeatureFlag{}, err
	}

	// Variations
	var metadata *string = nil
	if f.Metadata != nil {
		metadataBytes, err := json.Marshal(f.Variations)
		if err != nil {
			return FeatureFlag{}, err
		}
		metadataStr := string(metadataBytes)
		metadata = &metadataStr
	}

	return FeatureFlag{
		ID:              id,
		Name:            f.Name,
		Description:     f.Description,
		Variations:      string(variations),
		Type:            f.VariationType,
		BucketingKey:    f.BucketingKey,
		Metadata:        metadata,
		TrackEvents:     f.TrackEvents,
		Disable:         f.Disable,
		Version:         f.Version,
		CreatedDate:     f.CreatedDate,
		LastUpdatedDate: f.LastUpdatedDate,
	}, nil

}

//func FeatureFlagFromAPI(f model.FeatureFlag) error {
//	res := FeatureFlag{}
//	rollouts := make([]ProgressiveRollout, 0)
//
//	// DefaultRule
//	if f.DefaultRule == nil {
//		return errors.New("defaultRule is required")
//	}
//	if f.DefaultRule.ProgressiveRollout != nil {
//		//p, err := extractProgressiveRollout(f.DefaultRule.ProgressiveRollout)
//		//if err != nil {
//		//	return err
//		//}
//		//rollouts = append(rollouts, p)
//		//res.DefaultProgressiveRolloutID = &p.ID
//	}
//	if f.DefaultRule.Percentages != nil {
//		percentages, err := json.Marshal(f.DefaultRule.Percentages)
//		if err != nil {
//			return err
//		}
//		percentagesStr := string(percentages)
//		res.DefaultPercentages = &percentagesStr
//	}
//	if f.DefaultRule.VariationResult != nil {
//		res.DefaultVariationResult = f.DefaultRule.VariationResult
//	}
//

//	res.Variations = string(variations)
//	res.Name = f.Name
//	res.Description = f.Description
//
//}
