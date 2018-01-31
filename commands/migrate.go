package commands

import (
	"fmt"
	"path"

	"yellowroad_library/database/migrations"
	"yellowroad_library/config"
	"yellowroad_library/utils/app_error"
)

func MigrateCommand(configuration config.Configuration, workingDirectory string){

	fmt.Println("Migration Tool: Initializing ... ")

	fmt.Println("Attempting to run migrations")

	migrationDirectoryPath := "file:///" + path.Join(workingDirectory,"database","migrations")
	err := migrations.Migrate(configuration, migrationDirectoryPath)
	if err != nil {
		LogErrorAndExit(app_error.Wrap(err))
	}

	fmt.Println("Migration process completed. Exiting now ...")
}