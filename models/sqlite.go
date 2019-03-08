package models

import (
	"encoding/csv"
	"errors"
	"log"

	"os"

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

func InitData(dbFile string, csvFile string) (err error) {
	var routers []Router
	err = removeExistFile(dbFile)
	if err != nil {
		return
	}
	db, err := gorm.Open("sqlite3", dbFile)

	if err != nil {
		return
	}
	routers, err = loadCsvData(csvFile)
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

func ReadRouter(ip string) (router Router) {
	db := g.Conn()
	db.Where("ip = ?", ip).First(&router)
	return
}

func ReadAllRouters() (routers []Router) {
	db := g.Conn()
	db.Find(&routers)
	return
}

func ReadMultiRouters(ips []string) (routers []Router, failedList []string) {
	db := g.Conn()
	failedList = []string{}
	for _, ip := range ips {
		router := Router{}
		db.Where("ip = ?", ip).First(&router)
		if router.ID == 0 {
			failedList = append(failedList, ip)
		} else {
			routers = append(routers, router)
		}
	}
	return
}

func UpdateRouter(newRouter Router) (err error) {
	db := g.Conn()
	router := Router{}
	db.Where("ip = ?", newRouter.IP).First(&router)
	if router.ID == 0 {
		err = errors.New("record not found")
		return
	}
	router.IP = newRouter.IP
	router.Node = newRouter.Node
	router.NodeDetail = newRouter.NodeDetail
	err = db.Save(&router).Error
	if err != nil {
		log.Printf("UpdateRouter Failed: %s", err)
	}
	return
}

func DeleteRouter(ip string) (err error) {
	db := g.Conn()
	router := Router{}
	db.Where("ip = ?", ip).First(&router)
	if router.ID == 0 {
		err = errors.New("record not found")
		return
	}
	err = db.Delete(&router).Error
	if err != nil {
		log.Printf("DeleteRouter Failed: %s", err)
	}
	return
}

func CreateRouter(newRouter Router) (err error) {
	db := g.Conn()
	router := Router{}
	db.Where("ip = ?", newRouter.IP).First(&router)
	if router.ID != 0 {
		err = errors.New("router already exists")
		return
	}
	db.Unscoped().Where("ip = ?", newRouter.IP).Find(&router)
	if router.ID != 0 {
		router.IP = newRouter.IP
		router.Node = newRouter.Node
		router.NodeDetail = newRouter.NodeDetail
		router.DeletedAt = nil
		err = db.Unscoped().Save(&router).Error
	} else {
		err = db.Create(&newRouter).Error
	}
	if err != nil {
		log.Printf("CreateRouter Failed: %s", err)
	}
	return
}
