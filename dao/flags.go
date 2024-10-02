package dao

import (
	"context"

	"github.com/go-feature-flag/app-api/model"
)

type Flags interface {
	// GetFlags return all the flags
	GetFlags(ctx context.Context) ([]model.FeatureFlag, error)

	// GetFlagById return a flag by its ID
	GetFlagByID(ctx context.Context, id string) (model.FeatureFlag, error)

	// GetFlagByName return a flag by its name
	GetFlagByName(ctx context.Context, name string) (model.FeatureFlag, error)

	// CreateFlag create a new flag, return the id of the flag
	CreateFlag(ctx context.Context, flag model.FeatureFlag) (string, error)

	// UpdateFlag update a flag
	UpdateFlag(ctx context.Context, flag model.FeatureFlag) error

	// DeleteFlagByID delete a flag
	DeleteFlagByID(ctx context.Context, id string) error
}
