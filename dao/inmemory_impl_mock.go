package dao

import (
	"context"
	"fmt"
	daoErr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/go-feature-flag/app-api/model"
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
	for _, flag := range m.flags {
		if flag.ID == id {
			return flag, nil
		}
	}
	return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with id %s not found", id))
}

// GetFlagByName return a flag by its name
func (m *InMemoryMockDao) GetFlagByName(ctx context.Context, name string) (model.FeatureFlag, daoErr.DaoError) {
	for _, flag := range m.flags {
		if flag.Name == name {
			return flag, nil
		}
	}
	return model.FeatureFlag{}, daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with name %s not found", name))
}

// CreateFlag create a new flag, return the id of the flag
func (m *InMemoryMockDao) CreateFlag(ctx context.Context, flag model.FeatureFlag) (string, daoErr.DaoError) {
	m.flags = append(m.flags, flag)
	return flag.ID, nil
}

func (m *InMemoryMockDao) UpdateFlag(ctx context.Context, flag model.FeatureFlag) daoErr.DaoError {
	for index, f := range m.flags {
		if f.ID == flag.ID {
			m.flags[index] = flag
			return nil
		}
	}
	return daoErr.NewDaoError(daoErr.NotFound, fmt.Errorf("flag with id %s not found", flag.ID))
}

func (m *InMemoryMockDao) DeleteFlagByID(ctx context.Context, id string) daoErr.DaoError {
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
