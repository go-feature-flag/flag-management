package model

// FeatureFlagStatusUpdate represents the input for updating the status of a feature flag.
type FeatureFlagStatusUpdate struct {
	Disable bool `json:"disable"`
}
