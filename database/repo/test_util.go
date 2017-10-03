package repo

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/config"
	"github.com/jinzhu/gorm"
)


func WithGormDBConnection (configuration config.Configuration,onward func(*gorm.DB)) func(){
	return func(){
		dbSettings := configuration.Database

		var connectionString = fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=%s password=%s",
			dbSettings.Host,
			dbSettings.Username,
			dbSettings.Database,
			dbSettings.SSLMode,
			dbSettings.Password,
		)

		db, err := gorm.Open(dbSettings.Driver, connectionString)
		db.LogMode(true)
		So(err,ShouldBeNil)

		Reset(func() {
			//add code for reset here later
		})

		onward(db)
	}
}
