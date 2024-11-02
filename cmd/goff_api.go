package cmd

import (
	"fmt"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/dao/pgimpl"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/log"
	"go.uber.org/zap"
)

func NewGOFeatureFlagManagementAPICommand() GOFeatureFlagManagementAPICommand {
	return GOFeatureFlagManagementAPICommand{}
}

type GOFeatureFlagManagementAPICommand struct {
	databaseDao dao.FlagStorage
	apiHandlers handler.Handlers
	logger      zap.Logger
}

func (g *GOFeatureFlagManagementAPICommand) Run() error {
	if err := g.initDependencies(); err != nil {
		return fmt.Errorf("error while initialising dependencies: %w", err)
	}
	apiServer, err := api.New(":3001", g.apiHandlers)
	if err != nil {
		return fmt.Errorf("impossible to initialize API server: %w", err)
	}
	apiServer.Start()
	defer func() { apiServer.Stop() }()
	return nil
}

// initDependencies initializes the dependencies of the API server
func (g *GOFeatureFlagManagementAPICommand) initDependencies() error {
	// Initialize uber-go/zap logger
	log.InitLogger()

	// init database connection
	var err error
	if g.databaseDao, err = pgimpl.NewPostgresDao("localhost", 5432, "gofeatureflag", "goff-user", "my-secret-pw"); err != nil {
		return fmt.Errorf("impossible to initialize database connection: %w", err)
	}

	// init API handlers
	if g.apiHandlers, err = handler.InitHandlers(g.databaseDao); err != nil {
		return fmt.Errorf("impossible to initialize API handlers: %w", err)
	}
	return nil
}
