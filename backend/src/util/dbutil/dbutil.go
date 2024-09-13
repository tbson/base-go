package dbutil

import (
	"fmt"

	"app/common/setting"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var e error

func InitDb() {
	host := setting.DB_HOST
	user := setting.DB_USER
	password := setting.DB_PASSWORD
	dbName := setting.DB_NAME
	port := setting.DB_PORT
	timeZone := setting.TIME_ZONE
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbName, port, timeZone)
	db, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func Db() *gorm.DB {
	return db
}
