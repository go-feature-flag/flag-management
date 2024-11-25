package dao

import (
	"context"
	"fmt"
	daoErr "github.com/go-feature-flag/flag-management/server/dao/err"

	"github.com/go-feature-flag/flag-management/server/model"
	_ "github.com/lib/pq" // we import the driver used by sqlx
)

func NewInMemoryMockDao() (*InMemoryMockDao, error) {
	return &InMemoryMockDao{
		flags: []model.FeatureFlag{},
	}, nil
}

type InMemoryMockDao struct {
	flags []model.FeatureFlag

	errorOnPing bool
}

// GetFlags return all the flags
func (m *InMemoryMockDao) GetFlags(ctx context.Context) ([]model.FeatureFlag, daoErr.DaoError) {
	if ctx.Value("error") != nil {
		if err, ok := ctx.Value("error").(daoErr.DaoErrorCode); ok {
			return nil, daoErr.NewDaoError(err, fmt.Errorf("error on get flags"))
		}
		return nil, daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error on get flags"))
	}
	return m.flags, nil
}

// GetFlagByID return a flag by its ID
func (m *InMemoryMockDao) GetFlagByID(ctx context.Context, id string) (model.FeatureFlag, daoErr.DaoError) {
	if ctx.Value("error") != nil {
		if err, ok := ctx.Value("error").(daoErr.DaoErrorCode); ok {
			return model.FeatureFlag{}, daoErr.NewDaoError(err, fmt.Errorf("error on get flag by id"))
		}
		return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error on get flag by id"))
	}
	for _, flag := range m.flags {
		if flag.ID == id {
			return flag, nil
		}
	}
	return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with id %s not found", id))
}

// GetFlagByName return a flag by its name
func (m *InMemoryMockDao) GetFlagByName(ctx context.Context, name string) (model.FeatureFlag, daoErr.DaoError) {
	if ctx.Value("error") != nil {
		if err, ok := ctx.Value("error").(daoErr.DaoErrorCode); ok {
			return model.FeatureFlag{}, daoErr.NewDaoError(err, fmt.Errorf("error on get flag by name"))
		}
		return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error on get flag by name"))
	}
	for _, flag := range m.flags {
		if flag.Name == name {
			return flag, nil
		}
	}
	return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with name %s not found", name))
}

// CreateFlag create a new flag, return the id of the flag
func (m *InMemoryMockDao) CreateFlag(ctx context.Context, flag model.FeatureFlag) (string, daoErr.DaoError) {
	if ctx.Value("error_create") != nil {
		if err, ok := ctx.Value("error_create").(daoErr.DaoErrorCode); ok {
			return "", daoErr.NewDaoError(err, fmt.Errorf("error creating flag"))
		}
		return "", daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error creating flag"))
	}

	m.flags = append(m.flags, flag)
	return flag.ID, nil
}

func (m *InMemoryMockDao) UpdateFlag(ctx context.Context, flag model.FeatureFlag) daoErr.DaoError {
	if ctx.Value("error_update") != nil {
		if err, ok := ctx.Value("error_update").(daoErr.DaoErrorCode); ok {
			return daoErr.NewDaoError(err, fmt.Errorf("error on update flags"))
		}
		return daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error on update flags"))
	}
	for index, f := range m.flags {
		if f.ID == flag.ID {
			m.flags[index] = flag
			return nil
		}
	}
	return daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with id %s not found", flag.ID))
}

func (m *InMemoryMockDao) DeleteFlagByID(ctx context.Context, id string) daoErr.DaoError {
	if ctx.Value("error_delete") != nil {
		if err, ok := ctx.Value("error_delete").(daoErr.DaoErrorCode); ok {
			return daoErr.NewDaoError(err, fmt.Errorf("error on get flags"))
		}
		return daoErr.NewDaoError(daoErr.UnknownError, fmt.Errorf("error on get flags"))
	}

	newInmemoryFlagList := []model.FeatureFlag{}
	for _, f := range m.flags {
		if f.ID != id {
			newInmemoryFlagList = append(newInmemoryFlagList, f)
		}
	}
	m.flags = newInmemoryFlagList
	return nil
}

func (m *InMemoryMockDao) Ping() daoErr.DaoError {
	if m.errorOnPing {
		return daoErr.NewDaoError(daoErr.DatabaseNotInitialized, fmt.Errorf("error on ping"))
	}
	return nil
}

func (m *InMemoryMockDao) OnPingReturnError(v bool) {
	m.errorOnPing = v
}

func (m *InMemoryMockDao) SetFlags(flags []model.FeatureFlag) {
	m.flags = flags
}
