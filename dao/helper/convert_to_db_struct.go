package helper

import (
	"encoding/json"
	"github.com/go-feature-flag/app-api/dao/dbmodel"
	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
	"time"
)

func ConvertToDBStruct(flag model.FeatureFlag) (*dbmodel.DBConvertedFlag, error) {
	if flag.ID == "" {
		flag.ID = uuid.New().String()
	}

	variations, err := json.Marshal(flag.Variations)
	if err != nil {
		return nil, err
	}

	metadata, err := json.Marshal(flag.Metadata)
	if err != nil {
		return nil, err
	}

	percentages, err := json.Marshal(flag.Percentage)
	if err != nil {
		return nil, err
	}

	if flag.CreatedDate.IsZero() {
		flag.CreatedDate = time.Now()
	}

	ff := dbmodel.FeatureFlag{
		ID:                     flag.ID,
		Name:                   flag.Name,
		Description:            flag.Description,
		Variations:             string(variations),
		Type:                   string(flag.VariationType),
		BucketingKey:           &flag.BucketingKey,
		Metadata:               string(metadata),
		TrackEvents:            flag.TrackEvents,
		Disable:                flag.Disable,
		Version:                flag.Version,
		DefaultVariationResult: flag.DefaultRule.VariationResult,
		DefaultPercentages:     string(percentages),
		CreatedDate:            flag.CreatedDate,
		LastUpdatedDate:        time.Now(),
	}

	if flag.DefaultRule.ProgressiveRollout != nil {
		pr := dbmodel.ProgressiveRollout{}

		ff.DefaultProgressiveRolloutID = uuid.MustParse(flag.DefaultRule.ProgressiveRollout.ID)
	}

	// TODO: change return	type to dbmodel.FeatureFlag
	return nil, err
}
