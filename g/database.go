package g

import (
	"log"
	//引入 mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	//引入 sqlite 驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbp *gorm.DB

//Conn 给其他模块调用的连接池获取方法
func Conn() *gorm.DB {
	return dbp
}

//InitDB 初始化数据库连接池
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
}

//CloseDB 关闭数据库连接池
func CloseDB() (err error) {
	err = dbp.Close()
	return
}
