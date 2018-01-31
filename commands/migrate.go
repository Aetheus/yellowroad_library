package commands

import (
	"fmt"
	"log"

	"yellowroad_library/database/migrations"
	"yellowroad_library/config"
)

func MigrateCommand(configuration config.Configuration){

	fmt.Println("Migration Tool: Initializing ... ")

	fmt.Println("Attempting to run migrations")
	appErr := migrations.Migrate(configuration)
	if appErr != nil {
		log.Fatal(appErr)
	}

	fmt.Println("Migration process completed. Exiting now ...")
}