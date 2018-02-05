package commands

import (
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes"
	"yellowroad_library/utils/app_error"
)


func ServerCommand(configuration config.Configuration){
	container, err := containers.NewAppContainer(configuration)
	if (err != nil) {
		LogErrorAndExit(app_error.Wrap(err))
	}

	err = routes.Init(container)
	if (err != nil) {
		LogErrorAndExit(app_error.Wrap(err))
	}
}