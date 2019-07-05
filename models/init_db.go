package models

import (
	"encoding/csv"
	"errors"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shanghai-edu/rmdb-lite/g"
	"github.com/toolkits/file"
)

type Router struct {
	gorm.Model
	IP         string `gorm:"unique_index" json:"ip" binding:"required"`
	Node       string `json:"node" binding:"required"`
	NodeDetail string `json:"node_detail" binding:"required"`
}

func loadCsvData(csvFile string) (routers []Router, err error) {
	var router Router
	cf, err := os.Open(csvFile)
	if err != nil {
		return
	}
	defer cf.Close()
	csvReader := csv.NewReader(cf)
	_, err = csvReader.Read()
	if err != nil {
		return
	}
	rows, err := csvReader.ReadAll()
	if err != nil {
		return
	}
	for _, row := range rows {
		router.IP = row[0]
		router.Node = row[1]
		router.NodeDetail = row[2]
		routers = append(routers, router)
	}
	return
}

func removeExistFile(f string) (err error) {
	if file.IsExist(f) {
		if file.IsFile(f) {
			err = file.Remove(f)
		} else {
			err = errors.New(f + "is not directory, not file")
		}
	}
	return
}

func InitData(csvFile string) (err error) {
	if g.Config().Sqlite != "" {
		err = InitSqliteData(csvFile)
	} else {
		err = InitMysqlData(csvFile)
	}
	return
}

func InitSqliteData(csvFile string) (err error) {
	var routers []Router
	dbFile := g.Config().Sqlite
	err = removeExistFile(dbFile)
	if err != nil {
		return
	}
	routers, err = loadCsvData(csvFile)
	if err != nil {
		return
	}

	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		return
	}
	defer db.Close()

	db.CreateTable(&routers)
	for _, router := range routers {
		err = db.Create(&router).Error
	}
	if err != nil {
		db.Close()
		file.Remove(dbFile)
		return
	}
	return
}

func InitMysqlData(csvFile string) (err error) {
	var routers []Router

	routers, err = loadCsvData(csvFile)
	if err != nil {
		return
	}

	db, err := gorm.Open("mysql", g.Config().Mysql)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.DropTableIfExists(&routers).Error
	if err != nil {
		return
	}

	db.CreateTable(&routers)
	for _, router := range routers {
		err = db.Create(&router).Error
	}
	if err != nil {
		db.Close()
		return
	}
	return
}
