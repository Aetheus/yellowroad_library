package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*Configuration :
This is the Configuratoin struct - for the actual config files, check the root directory
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
	}

	Web struct {
		Port int
	}
}

//TODO : write logic to do validity checking for all the config fields
func (config Configuration) isValid() (isValid bool, reason string) {
	return true, ""
}

func Load(pathToConfigFile string) Configuration {
	raw, err := ioutil.ReadFile(pathToConfigFile)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	var config Configuration
	if err := json.Unmarshal(raw, &config); err != nil {
		fmt.Println(err.Error())
		panic(err.Error)
	}

	if isValidConfig, reason := config.isValid(); !isValidConfig {
		panic(reason)
	}

	return config
}
