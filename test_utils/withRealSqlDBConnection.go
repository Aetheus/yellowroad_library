package test_utils

import (
	"database/sql"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/config"
)


func WithRealSqlDBConnection(onward func(*sql.DB)) func(){

	return func(){
		pathToConfigFile := APP_ROOT+"/test_utils/config_for_mocks.json"
		configuration, configErr := config.Load(pathToConfigFile)
		So(configErr, ShouldBeNil)


		dbSettings := configuration.Database
		connectionString := fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=%s password=%s",
			dbSettings.Host,
			dbSettings.Username,
			dbSettings.Database,
			dbSettings.SSLMode,
			dbSettings.Password,
		)

		db, err := sql.Open("postgres", connectionString)
		So(err,ShouldBeNil)

		Reset(func() {
			//add code for reset here later
		})

		onward(db)
	}
}

func WithRealSqlDBTransaction(onward func(tx *sql.Tx)) func(){

	return func(){
		pathToConfigFile := APP_ROOT+"/test_utils/config_for_mocks.json"
		configuration, configErr := config.Load(pathToConfigFile)
		So(configErr, ShouldBeNil)


		dbSettings := configuration.Database
		connectionString := fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=%s password=%s",
			dbSettings.Host,
			dbSettings.Username,
			dbSettings.Database,
			dbSettings.SSLMode,
			dbSettings.Password,
		)

		db, err := sql.Open("postgres", connectionString)
		So(err,ShouldBeNil)

		tx, err := db.Begin()
		So(err,ShouldBeNil)

		Reset(func() {
			tx.Rollback()
		})

		onward(tx)
	}
}
