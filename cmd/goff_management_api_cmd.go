package cmd

import (
	"fmt"
	"os"

	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/config"
	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/dao/pgimpl"
	"github.com/go-feature-flag/app-api/handler"
	"github.com/go-feature-flag/app-api/log"
	"github.com/spf13/pflag"
)

type APICommandOptions struct {
	// OverrideDefaultDao is used to override the default dao used by the API server
	// This is useful for testing purpose to provide a mock dao
	OverrideDefaultDao dao.FlagStorage
}

// NewGOFeatureFlagManagementAPICommand creates a new instance of the GOFeatureFlagManagementAPICommand
// It initializes the dependencies of the API server and returns the command
// After you can call Run to start the API server.
// options - is used to override default services for testing purpose.
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
	apiServer     *api.Server
	options       APICommandOptions
	configuration *config.Configuration
}

func (g *GOFeatureFlagManagementAPICommand) Run() {
	g.apiServer.Start()
	defer func() { _ = g.apiServer.Stop() }()
}

// initDependencies initializes the dependencies of the API server
func (g *GOFeatureFlagManagementAPICommand) initDependencies() error {
	// Initialize uber-go/zap logger
	log.InitLogger()

	//init config
	f := pflag.NewFlagSet("config", pflag.ContinueOnError)
	f.String("postgresConnectionString", "", "Connection string to connect to the postgres database")
	f.String("serverAddress", ":3001", "Address where the API server will listen")
	f.String("mode", "production", "Application mode (development or production)")
	_ = f.Parse(os.Args[1:])

	c, err := config.LoadConfiguration(f)
	if err != nil {
		return fmt.Errorf("impossible to load configuration: %w", err)
	}
	g.configuration = c

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
	if g.apiServer, err = api.New(g.configuration.ServerAddress, apiHandlers); err != nil {
		return fmt.Errorf("impossible to initialize API server: %w", err)
	}
	return nil
}

func (g *GOFeatureFlagManagementAPICommand) initDatabaseAccess() (dao.FlagStorage, error) {
	if g.options.OverrideDefaultDao != nil {
		return g.options.OverrideDefaultDao, nil
	}
	databaseDao, err := pgimpl.NewPostgresDao(g.configuration.PostgresConnectionString)
	if err != nil {
		return nil, err
	}
	return databaseDao, nil
}
