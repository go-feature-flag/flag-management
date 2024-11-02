package handler

import (
	"errors"

	"github.com/go-feature-flag/app-api/dao"
)

type Handlers struct {
	FlagAPIHandler *FlagAPIHandler
	HealthHandler  *HealthHandler
}

func InitHandlers(dao dao.FlagStorage) (Handlers, error) {
	if dao == nil {
		return Handlers{}, ErrMissingDao
	}
	flagAPIHandler := NewFlagAPIHandler(dao, &FlagAPIHandlerOptions{})
	healthHandler := NewHealthHandler(dao)
	return Handlers{
		FlagAPIHandler: &flagAPIHandler,
		HealthHandler:  &healthHandler,
	}, nil
}

var ErrMissingFlagAPIHandler = errors.New("flagAPIHandler cannot be nil")
var ErrMissingHealthHandler = errors.New("healthHandler cannot be nil")
var ErrMissingDao = errors.New("dao cannot be nil")
