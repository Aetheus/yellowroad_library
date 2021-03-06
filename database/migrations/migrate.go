package migrations

import (
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"fmt"
	"yellowroad_library/config"
	"errors"
)

func Migrate(configuration config.Configuration, migrationsDirectoryPath string) (error) {
	fmt.Println("Attempting migrations ... ")
	//"postgres://mattes:secret@localhost:5432/database?sslmode=disable"
	connString, err := createConnectionString(configuration)
	if(err != nil) {
		return err
	}

	migrater, err := migrate.New(migrationsDirectoryPath, connString)
	if (err != nil ){
		return err
	}


	if err = migrater.Up(); err != nil {
		if (err.Error() == "no change"){
			//reminder: if you added a new SQL script and it still showed "no change", make sure the filename ends with ".up.sql"
			fmt.Println("No migrations to run")
			return nil
		}else {
			fmt.Println(err.Error())
			return err
		}
	}

	fmt.Println("Migrations complete")
	return nil

}

func createConnectionString(configuration config.Configuration) (string, error) {
	var dbSettings =  configuration.Database

	if (dbSettings.Driver == "postgres"){
		connectionString := fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			dbSettings.Username,
			dbSettings.Password,
			dbSettings.Host,
			dbSettings.Database,
			dbSettings.SSLMode,

		)
		return connectionString, nil
	} else {
		return "", errors.New("No support for migrating " + dbSettings.Driver + " databases yet!")
	}

}


