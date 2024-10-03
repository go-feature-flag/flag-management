package dbmodel

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
)

type JSONB json.RawMessage

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &j)
}

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

	variations, err := json.Marshal(mff.Variations)
	if err != nil {
		return FeatureFlag{}, err
	}

	var metadata JSONB
	if mff.Metadata != nil {
		metadataBytes, err := json.Marshal(mff.Metadata)
		if err != nil {
			return FeatureFlag{}, err
		}
		metadata = JSONB(metadataBytes)
	}

	return FeatureFlag{
		ID:              id,
		Name:            mff.Name,
		Description:     mff.Description,
		Variations:      JSONB(variations),
		Type:            mff.VariationType,
		BucketingKey:    mff.BucketingKey,
		Metadata:        metadata,
		TrackEvents:     mff.TrackEvents,
		Disable:         mff.Disable,
		Version:         mff.Version,
		CreatedDate:     mff.CreatedDate,
		LastUpdatedDate: mff.LastUpdatedDate,
		LastModifiedBy:  mff.LastModifiedBy,
	}, nil
}

func (ff *FeatureFlag) ToModelFeatureFlag(rules []Rule) (model.FeatureFlag, error) {
	var apiRules = make([]model.Rule, 0)
	var defaultRule *model.Rule
	for _, rule := range rules {
		convertedRule, err := rule.ToModelRule()
		if err != nil {
			return model.FeatureFlag{}, err
		}
		if rule.IsDefault {
			defaultRule = &convertedRule
			continue
		}
		apiRules = append(apiRules, convertedRule)
	}

	var variations map[string]*interface{}
	err := json.Unmarshal(ff.Variations, &variations)
	if err != nil {
		return model.FeatureFlag{}, err
	}

	var metadata map[string]interface{}
	err = json.Unmarshal(ff.Metadata, &metadata)
	if err != nil {
		return model.FeatureFlag{}, err
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
	}, nil
}
