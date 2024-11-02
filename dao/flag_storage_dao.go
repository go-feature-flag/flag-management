package dao

import (
	"context"

	daoErr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/go-feature-flag/app-api/model"
)

type FlagStorage interface {
	// GetFlags return all the flags
	GetFlags(ctx context.Context) ([]model.FeatureFlag, daoErr.DaoError)

	// GetFlagByID return a flag by its ID
	GetFlagByID(ctx context.Context, id string) (model.FeatureFlag, daoErr.DaoError)

	// GetFlagByName return a flag by its name
	GetFlagByName(ctx context.Context, name string) (model.FeatureFlag, daoErr.DaoError)

	// CreateFlag create a new flag, return the id of the flag
	CreateFlag(ctx context.Context, flag model.FeatureFlag) (string, daoErr.DaoError)

	// UpdateFlag update a flag
	UpdateFlag(ctx context.Context, flag model.FeatureFlag) daoErr.DaoError

	// DeleteFlagByID delete a flag
	DeleteFlagByID(ctx context.Context, id string) daoErr.DaoError

	// Ping check that the data layer is available
	Ping() daoErr.DaoError
}
