package main

import (
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/database/migrations"
	"yellowroad_library/http/routes"
)
var configuration = config.Load("./config.json")
var container = containers.NewAppContainer(configuration)

func main() {
	err := migrations.Migrate(configuration)
	if (err != nil){
		panic(err)
	}

	routes.Init(container)
}
