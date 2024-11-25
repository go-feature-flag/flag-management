package config

import (
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/spf13/pflag"
)

var k = koanf.New(".")

type AppMode string

const (
	Development AppMode = "development"
	Production  AppMode = "production"
)

type Configuration struct {
	// ServerAddress is the address where the API server will listen
	ServerAddress string

	// database information
	PostgresConnectionString string

	// Mode is the mode in which the application is running (accepts "development" or "production")
	// If development, the application will run with verbose logging and no authentication will be required for the APIs
	// Default is "production"
	Mode AppMode
}

func LoadConfiguration(flagSet *pflag.FlagSet) (*Configuration, error) {
	if errBindFlag := k.Load(posflag.Provider(flagSet, ".", k), nil); errBindFlag != nil {
		return nil, errBindFlag
	}
	proxyConf := &Configuration{}
	errUnmarshal := k.Unmarshal("", &proxyConf)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return proxyConf, nil
}
