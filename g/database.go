package g

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbp *gorm.DB

func Conn() *gorm.DB {
	return dbp
}

func InitDB() {
	if Config().Sqlite != "" {
		db, err := gorm.Open("sqlite3", Config().Sqlite)
		if err != nil {
			log.Fatalln("Init DB Connect Failed: ", err)
		}
		if Config().LogLevel == "debug" {
			db.LogMode(true)
		}
		dbp = db
		return
	}
	db, err := gorm.Open("mysql", Config().Mysql)
	if err != nil {
		log.Fatalln("Init DB Connect Failed: ", err)
	}
	if Config().LogLevel == "debug" {
		db.LogMode(true)
	}
	dbp = db
	return
}

func CloseDB() (err error) {
	err = dbp.Close()
	return
}
