package dbmodel

import (
	"errors"
	"time"

	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
)

type FeatureFlag struct {
	ID              uuid.UUID      `db:"id"`
	Name            string         `db:"name"`
	Description     *string        `db:"description"`
	Variations      JSONB          `db:"variations"`
	Type            model.FlagType `db:"type"`
	BucketingKey    *string        `db:"bucketing_key"`
	Metadata        JSONB          `db:"metadata"`
	TrackEvents     *bool          `db:"track_events"`
	Disable         *bool          `db:"disable"`
	Version         *string        `db:"version"`
	CreatedDate     time.Time      `db:"created_date"`
	LastUpdatedDate time.Time      `db:"last_updated_date"`
	LastModifiedBy  string         `db:"last_modified_by"`
}

func FromModelFeatureFlag(mff model.FeatureFlag) (FeatureFlag, error) {
	id, err := uuid.Parse(mff.ID)
	if err != nil {
		return FeatureFlag{}, err
	}

	ff := FeatureFlag{
		ID:              id,
		Name:            mff.Name,
		Description:     mff.Description,
		Type:            mff.VariationType,
		BucketingKey:    mff.BucketingKey,
		TrackEvents:     mff.TrackEvents,
		Disable:         mff.Disable,
		Version:         mff.Version,
		CreatedDate:     mff.CreatedDate,
		LastUpdatedDate: mff.LastUpdatedDate,
		LastModifiedBy:  mff.LastModifiedBy,
	}
	if mff.Variations != nil {
		ff.Variations = JSONB(*mff.Variations)
	}
	if mff.Metadata != nil {
		ff.Metadata = JSONB(*mff.Metadata)
	}
	return ff, nil
}

func (ff *FeatureFlag) ToModelFeatureFlag(rules []Rule) (model.FeatureFlag, error) {
	var apiRules = make([]model.Rule, 0)
	var defaultRule *model.Rule
	hasDefaultRule := false
	for _, rule := range rules {
		convertedRule := rule.ToModelRule()
		if rule.IsDefault {
			hasDefaultRule = true
			defaultRule = &convertedRule
			continue
		}
		apiRules = append(apiRules, convertedRule)
	}
	if !hasDefaultRule {
		return model.FeatureFlag{}, errors.New("default rule is required")
	}

	variations := make(map[string]interface{})
	if ff.Variations != nil {
		variations = ff.Variations
	}
	metadata := make(map[string]interface{})
	if ff.Metadata != nil {
		metadata = ff.Metadata
	}
	return model.FeatureFlag{
		ID:              ff.ID.String(),
		Name:            ff.Name,
		Description:     ff.Description,
		Variations:      &variations,
		VariationType:   ff.Type,
		BucketingKey:    ff.BucketingKey,
		Metadata:        &metadata,
		TrackEvents:     ff.TrackEvents,
		Disable:         ff.Disable,
		Version:         ff.Version,
		CreatedDate:     ff.CreatedDate,
		LastUpdatedDate: ff.LastUpdatedDate,
		Rules:           &apiRules,
		DefaultRule:     defaultRule,
		LastModifiedBy:  ff.LastModifiedBy,
	}, nil
}
