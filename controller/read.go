package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanghai-edu/rmdb-lite/controller/utils"
	"github.com/shanghai-edu/rmdb-lite/models"
)

type ResData struct {
	IP         string `json:"ip"`
	Node       string `json:"node"`
	NodeDetail string `json:"node_detail"`
}

type MultiRouterRes struct {
	Routers    []ResData `json:"routers"`
	FailedList []string  `json:"failed_list"`
}

func getRouter(c *gin.Context) {
	ip := c.Query("ip")
	router := models.ReadRouter(ip)

	if router.ID == 0 {
		c.JSON(http.StatusNotFound, utils.ErrorRes(utils.RecordNotFound))
		return
	}
	x_api_key := c.Request.Header.Get("X-API-KEY")
	user := utils.GetUserFromKey(x_api_key)
	resData := &ResData{
		IP:   router.IP,
		Node: router.Node,
	}
	if user.Role == 1 {
		c.JSON(http.StatusOK, utils.SuccessRes(resData))
		return
	}
	if user.Role > 1 {
		resData.NodeDetail = router.NodeDetail
		c.JSON(http.StatusOK, utils.SuccessRes(resData))
		return
	}
}

func getAllRouters(c *gin.Context) {
	routers := models.ReadAllRouters()
	if len(routers) == 0 {
		c.JSON(http.StatusNotFound, utils.ErrorRes(utils.RecordNotFound))
		return
	}
	x_api_key := c.Request.Header.Get("X-API-KEY")
	user := utils.GetUserFromKey(x_api_key)
	resDatas := []ResData{}
	if user.Role == 1 {
		for _, router := range routers {
			resData := ResData{
				IP:   router.IP,
				Node: router.Node,
			}
			resDatas = append(resDatas, resData)
		}
		c.JSON(http.StatusOK, utils.SuccessRes(resDatas))
		return
	}
	if user.Role > 1 {
		for _, router := range routers {
			resData := ResData{
				IP:         router.IP,
				Node:       router.Node,
				NodeDetail: router.NodeDetail,
			}
			resDatas = append(resDatas, resData)
		}
		c.JSON(http.StatusOK, utils.SuccessRes(resDatas))
		return
	}
}

type multiRouterInput struct {
	Ips []string `json:"ips" binding:"required"`
}

func getMultiRouters(c *gin.Context) {
	inputs := multiRouterInput{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorRes(utils.BodyJsonDecodeError))
		return
	}
	routers, failedList := models.ReadMultiRouters(inputs.Ips)

	x_api_key := c.Request.Header.Get("X-API-KEY")
	user := utils.GetUserFromKey(x_api_key)
	var multiRouterRes MultiRouterRes
	resDatas := []ResData{}
	if user.Role == 1 {
		for _, router := range routers {
			resData := ResData{
				IP:   router.IP,
				Node: router.Node,
			}
			resDatas = append(resDatas, resData)
		}
		multiRouterRes = MultiRouterRes{
			Routers:    resDatas,
			FailedList: failedList,
		}
		c.JSON(http.StatusOK, utils.SuccessRes(multiRouterRes))
		return
	}
	if user.Role > 1 {
		for _, router := range routers {
			resData := ResData{
				IP:         router.IP,
				Node:       router.Node,
				NodeDetail: router.NodeDetail,
			}
			resDatas = append(resDatas, resData)
		}
		multiRouterRes = MultiRouterRes{
			Routers:    resDatas,
			FailedList: failedList,
		}
		c.JSON(http.StatusOK, utils.SuccessRes(multiRouterRes))
		return
	}
}
