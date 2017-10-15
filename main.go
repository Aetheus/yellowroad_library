package main

import (
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/database/migrations"
	"yellowroad_library/http/routes"
)


func main() {
	var err error

	configuration, err:= config.Load("./config.json")
	if (err != nil){
		panic(err)
	}

	container := containers.NewAppContainer(configuration)
	err = migrations.Migrate(configuration)
	if (err != nil){
		panic(err)
	}

	routes.Init(container)
}
