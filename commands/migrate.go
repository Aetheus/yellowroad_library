package commands

import (
	"fmt"
	"log"

	"yellowroad_library/database/migrations"
	"yellowroad_library/config"
	"path"
	"yellowroad_library/utils/app_error"
)

func MigrateCommand(configuration config.Configuration, workingDirectory string){

	fmt.Println("Migration Tool: Initializing ... ")

	fmt.Println("Attempting to run migrations")

	migrationDirectoryPath := "file:///" + path.Join(workingDirectory,"database","migrations")
	err := migrations.Migrate(configuration, migrationDirectoryPath)
	if err != nil {
		err := app_error.Wrap(err)
		err.Stacktrace()
		log.Fatalf("Error:\n\t%s\nStackTrace:\n%s",err.Error(),err.Stacktrace())
	}

	fmt.Println("Migration process completed. Exiting now ...")
}