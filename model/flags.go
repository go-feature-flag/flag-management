package model

import (
	"time"
)

type FlagType string

const (
	FlagTypeBoolean FlagType = "boolean"
	FlagTypeString  FlagType = "string"
	FlagTypeInteger FlagType = "integer"
	FlagTypeDouble  FlagType = "double"
	FlagTypeJSON    FlagType = "json"
)

type FeatureFlag struct {
	ID              string    `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	CreatedDate     time.Time `json:"createdDate" db:"created_date"`
	LastUpdatedDate time.Time `json:"lastUpdatedDate" db:"last_updated_date"`
	LastModifiedBy  string    `json:"LastModifiedBy" db:"last_modified_by"`
	Description     *string   `json:"description" db:"description"`
	VariationType   FlagType  `json:"type" db:"type"`
	// Variations are all the variations available for this flag. The minimum is 2 variations and, we don't have any max
	// limit except if the variationValue is a bool, the max is 2.
	Variations *map[string]interface{} `json:"variations,omitempty"`

	// Rules is the list of Rule for this flag.
	// This an optional field.
	Rules *[]Rule `json:"targeting,omitempty"` // nolint: lll

	// BucketingKey defines a source for a dynamic targeting key
	BucketingKey *string `json:"bucketingKey,omitempty"`

	// DefaultRule is the rule applied after checking that any other rules
	// matched the user.
	DefaultRule *Rule `json:"defaultRule,omitempty"` // nolint: lll

	// Experimentation is your struct to configure an experimentation.
	// It will allow you to configure a start date and an end date for your flag.
	// When the experimentation is not running, the flag will serve the default value.
	// Experimentation *ExperimentationDto `json:"experimentation,omitempty""`

	// Metadata is a field containing information about your flag such as an issue tracker link, a description, etc ...
	Metadata *map[string]interface{} `json:"metadata,omitempty"` // nolint: lll

	// Disable is true if the flag is disabled.
	Disable *bool `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`

	// Version (optional) This field contains the version of the flag.
	// The version is manually managed when you configure your flags and, it is used to display the information
	// in the notifications and data collection.
	Version *string `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty"`

	// TrackEvents is false if you don't want to export the data in your data exporter.
	// Default value is true
	TrackEvents *bool `json:"trackEvents,omitempty" yaml:"trackEvents,omitempty" toml:"trackEvents,omitempty"`
}

type Rule struct {
	// Id of the rule
	ID string `json:"id" db:"id"`
	// Name is the name of the rule, this field is mandatory if you want
	// to update the rule during scheduled rollout
	Name string `json:"name,omitempty"`
	// Query represents an antlr query in the nikunjy/rules format
	Query string `json:"query,omitempty"`

	// VariationResult represents the variation name to use if the rule apply for the user.
	// In case we have a percentage field in the config VariationResult is ignored
	VariationResult *string `json:"variation,omitempty"` // nolint: lll

	// Percentages represents the percentage we should give to each variation.
	// example: variationA = 10%, variationB = 80%, variationC = 10%
	Percentages *map[string]float64 `json:"percentage,omitempty" ` // nolint: lll

	// ProgressiveRollout is your struct to configure a progressive rollout deployment of your flag.
	// It will allow you to ramp up the percentage of your flag over time.
	// You can decide at which percentage you starts with and at what percentage you ends with in your release ramp.
	// Before the start date we will serve the initial percentage and, after we will serve the end percentage.
	ProgressiveRollout *ProgressiveRollout `json:"progressiveRollout,omitempty" yaml:"progressiveRollout,omitempty" toml:"progressiveRollout,omitempty" jsonschema:"title=progressiveRollout,description=Configure a progressive rollout deployment of your flag."` // nolint: lll

	// Disable indicates that this rule is disabled.
	Disable bool `json:"disable,omitempty" ` // nolint: lll
}

type ProgressiveRollout struct {
	// Initial contains a description of the initial state of the rollout.
	Initial *ProgressiveRolloutStep `json:"initial,omitempty" yaml:"initial,omitempty" toml:"initial,omitempty" jsonschema:"title=initial,description=A description of the initial state of the rollout."` // nolint: lll

	// End contains what describes the end status of the rollout.
	End *ProgressiveRolloutStep `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty" jsonschema:"title=initial,description=A description of the end state of the rollout."` // nolint: lll
}

// ProgressiveRolloutStep define a progressive rollout step (initial and end)
type ProgressiveRolloutStep struct {
	// Variation - name of the variation for this step
	Variation *string `json:"variation,omitempty" yaml:"variation,omitempty" toml:"variation,omitempty" jsonschema:"required,title=variation,description=Name of the variation to apply."` // nolint: lll

	// Percentage is the percentage (initial or end) for the progressive rollout
	Percentage *float64 `json:"percentage,omitempty" yaml:"percentage,omitempty" toml:"percentage,omitempty" jsonschema:"required,title=percentage,description=The percentage (initial or end) for the progressive rollout."` // nolint: lll

	// Date is the time it starts or ends.
	Date *time.Time `json:"date,omitempty" yaml:"date,omitempty" toml:"date,omitempty" jsonschema:"required,title=date,description=Date is the time it starts or ends."` // nolint: lll
}

func (ff *FeatureFlag) GetRules() []Rule {
	if ff.Rules == nil {
		return []Rule{}
	}
	return *ff.Rules
}

func (ff *FeatureFlag) GetDefaultRule() Rule {
	if ff.DefaultRule == nil {
		return Rule{}
	}
	return *ff.DefaultRule
}
