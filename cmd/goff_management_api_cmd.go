package cmd

import (
	"fmt"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/dao/pgimpl"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/log"
)

type APICommandOptions struct {
	// OverrideDefaultDao is used to override the default dao used by the API server
	// This is useful for testing purpose to provide a mock dao
	OverrideDefaultDao dao.FlagStorage
}

// NewGOFeatureFlagManagementAPICommand creates a new instance of the GOFeatureFlagManagementAPICommand
// It initializes the dependencies of the API server and returns the command
// After you can call Run to start the API server.
func NewGOFeatureFlagManagementAPICommand(options APICommandOptions) (GOFeatureFlagManagementAPICommand, error) {
	g := GOFeatureFlagManagementAPICommand{
		options: options,
	}
	if err := g.initDependencies(); err != nil {
		return g, fmt.Errorf("error while initializing dependencies: %w", err)
	}
	return g, nil
}

type GOFeatureFlagManagementAPICommand struct {
	apiServer *api.Server
	options   APICommandOptions
}

func (g *GOFeatureFlagManagementAPICommand) Run() {
	g.apiServer.Start()
	defer func() { _ = g.apiServer.Stop() }()
}

// initDependencies initializes the dependencies of the API server
func (g *GOFeatureFlagManagementAPICommand) initDependencies() error {
	// Initialize uber-go/zap logger
	log.InitLogger()

	// init database connection
	databaseDao, err := g.initDatabaseAccess()
	if err != nil {
		return fmt.Errorf("impossible to initialize database connection: %w", err)
	}

	// init API handlers
	apiHandlers, err := handler.InitHandlers(databaseDao)
	if err != nil {
		return fmt.Errorf("impossible to initialize API handlers: %w", err)
	}

	// init API server
	if g.apiServer, err = api.New(":3001", apiHandlers); err != nil {
		return fmt.Errorf("impossible to initialize API server: %w", err)
	}
	return nil
}

func (g *GOFeatureFlagManagementAPICommand) initDatabaseAccess() (dao.FlagStorage, error) {
	if g.options.OverrideDefaultDao != nil {
		return g.options.OverrideDefaultDao, nil
	}
	databaseDao, err := pgimpl.NewPostgresDao("localhost", 5432, "gofeatureflag", "goff-user", "my-secret-pw")
	if err != nil {
		return nil, err
	}
	return databaseDao, nil
}
