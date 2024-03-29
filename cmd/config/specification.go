package config

import "github.com/kelseyhightower/envconfig"

// Specification holds all environment variables of the Application.
type Specification struct {
	XApiKeys struct {
		Scooter string `envconfig:"SCOOTER_API_KEY"`
		Mobile  string `envconfig:"MOBILE_API_KEY"`
	}
	Database struct {
		Driver string `envconfig:"DATABASE_DRIVER" default:"mysql"`
		DSN    string `envconfig:"DATABASE_DSN"`
	}
	Server struct {
		Port string `envconfig:"APP_PORT" default:":80"`
	}
	Log struct {
		Level string `envconfig:"LOG_LEVEL" default:"trace"`
	}
}

// LoadEnvironmentVariables process the given spec to load all environment variables.
func LoadEnvironmentVariables(spec *Specification) error {
	err := envconfig.Process("", spec)
	if err != nil {
		return err
	}

	return nil
}
