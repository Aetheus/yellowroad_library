package database

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres" //for gorm

func Conn(dbType string, connectionString string) *gorm.DB {
	var dbConn *gorm.DB

	db, err := gorm.Open(dbType, connectionString)
	db.LogMode(true)
	if err != nil {
		panic("Failed to connect to the database!")
	}

	dbConn = db

	return dbConn
}
