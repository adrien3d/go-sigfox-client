package server

import (
	"github.com/joho/godotenv"
	"os"
)

func (a *API) SetupViper() error {

	filename := ".env"
	switch os.Getenv("SIGFOX_CLIENT_ENV") {
	case "testing":
		filename = "../.env.testing"
	case "prod":
		filename = ".env.prod"
	}

	err := godotenv.Overload(filename)
	if err != nil {
		return err
	}

	a.Config.SetEnvPrefix("sigfox_client")
	a.Config.AutomaticEnv()

	a.SetupViperDefaults()

	return nil
}

func (a *API) SetupViperDefaults() {
	// a.Config.SetDefault("rate_limit_requests_per_second", 5)
	a.Config.SetDefault("rate_limit_activated", false)
}
