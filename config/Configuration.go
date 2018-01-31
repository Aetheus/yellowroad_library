package config

import (
	"encoding/json"
	"io/ioutil"
	"yellowroad_library/utils/app_error"
	"errors"
)

/*Configuration :
This is the Configuration struct - for the actual config files, check the root directory
for "config.json". If that doesn't exist yet, use the "_sample_config.json" file as a reference
to create your own config.json
*/
type Configuration struct {
	Database struct {
		Host     string
		Username string
		Password string
		Database string //the database name
		SSLMode  string
		Driver   string

		MigrationsDir string
	}

	Web struct {
		Port int
		AllowOrigins []string
	}

	JWT struct {
		ExpiryDurationInDays int
		SecretKey string
	}
}

//TODO : write logic to do validity checking for all the config fields
func (config Configuration) isValid() (isValid bool, reason string) {
	return true, ""
}

func Load(pathToConfigFile string) (Configuration, app_error.AppError) {
	var config Configuration

	raw, err := ioutil.ReadFile(pathToConfigFile)
	if err != nil {
		return config, app_error.Wrap(err)
	}

	if err := json.Unmarshal(raw, &config); err != nil {
		return config, app_error.Wrap(err)
	}

	if isValidConfig, reason := config.isValid(); !isValidConfig {
		err := errors.New(reason)
		return config, app_error.Wrap(err)
	}

	return config, nil
}
