package main

import (
	"github.com/go-feature-flag/flag-management/server/cmd"
	"log"
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
	command, err := cmd.NewGOFeatureFlagManagementAPICommand(cmd.APICommandOptions{})
	if err != nil {
		log.Fatal(err)
	}
	command.Run()
}
