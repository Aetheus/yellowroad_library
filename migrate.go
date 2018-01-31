package main

import (
	"yellowroad_library/database/migrations"
	"yellowroad_library/config"
	"fmt"
	"log"
	"os"
	"path"
)

func main(){

	fmt.Println("Migration Tool: Initializing ... ")

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrationPath := os.Getenv("config_path")
	if migrationPath == "" {
		migrationPath = path.Join(workingDir,"config.json")
		fmt.Println("No 'config_path' environment was found. Defaulting to:", migrationPath)
	}

	fmt.Println("Attempting to read configuration file from:",migrationPath)
	configuration, err := config.Load(migrationPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Attempting to run migrations")
	appErr := migrations.Migrate(configuration)
	if appErr != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration process completed. Exiting now ...")
}