package main

import (
	"github.com/go-feature-flag/app-api/api"
	"github.com/go-feature-flag/app-api/dao/pgimpl"
	"github.com/go-feature-flag/app-api/handler"
)

// version, releaseDate are override by the makefile during the build.
var version = "localdev"

// @title GO Feature Flag - configuration API
// @description.markdown
// @contact.name GO feature flag configuration API
// @contact.url https://gofeatureflag.org
// @contact.email contact@gofeatureflag.org
// @license.name MIT
// @license.url https://github.com/thomaspoignant/go-feature-flag/blob/main/LICENSE
// @x-logo {"url":"https://raw.githubusercontent.com/thomaspoignant/go-feature-flag/main/logo_128.png"}
// @BasePath /
// @in header
// @name Authorization
func main() {
	// TODO: add configuration management for the API server.
	data, err := pgimpl.NewPostgresDao("localhost", 5432, "gofeatureflag", "goff-user", "my-secret-pw")
	if err != nil {
		panic(err)
	}
	flagHandlers := handler.NewFlagAPIHandler(data, &handler.FlagAPIHandlerOptions{})
	healthHandlers := handler.NewHealth(data)

	apiServer := api.New(":3001", flagHandlers, healthHandlers)
	apiServer.Start()
	defer func() { apiServer.Stop() }()
}
