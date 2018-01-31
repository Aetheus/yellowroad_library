package commands

import (
	"log"
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes"
)


func ServerCommand(configuration config.Configuration){
	container := containers.NewAppContainer(configuration)

	err := routes.Init(container)
	if (err != nil) {
		log.Fatal(err)
	}
}