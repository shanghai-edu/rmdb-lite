package models

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	"github.com/shanghai-edu/rmdb-lite/g"
)

//ReadRouter 查询单个路由器
func ReadRouter(ip string) (router Router) {
	db := g.Conn()
	db.Where("ip = ?", ip).First(&router)
	return
}

//ReadAllRouters 查询所有路由器
func ReadAllRouters() (routers []Router) {
	db := g.Conn()
	db.Find(&routers)
	return
}

//ReadMultiRouters 查询多个路由器
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

//UpdateRouter 更新路由器
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
		log.Errorf("UpdateRouter Failed: %s", err)
	}
	return
}

//DeleteRouter 删除路由器
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
		log.Errorf("DeleteRouter Failed: %s", err)
	}
	return
}

//CreateRouter 创建路由器
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
		log.Errorf("CreateRouter Failed: %s", err)
	}
	return
}
