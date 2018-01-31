package main

import (
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes"
	"log"
)


func main() {
	var err error

	configuration, err:= config.Load("./config.json")
	if (err != nil){
		log.Fatal(err)
	}

	container := containers.NewAppContainer(configuration)

	routes.Init(container)
}
