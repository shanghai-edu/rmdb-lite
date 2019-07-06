package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanghai-edu/rmdb-lite/controller/utils"
)

func readOnlyCheckMidd(c *gin.Context) {
	xAPIKey := c.Request.Header.Get("X-API-KEY")
	user := utils.GetUserFromKey(xAPIKey)
	if user.Role == 0 {
		c.JSON(http.StatusUnauthorized, utils.ErrorRes(utils.InvalidAPIKEY))
		c.Abort()
		return
	}
}

func adminCheckMidd(c *gin.Context) {
	xAPIKey := c.Request.Header.Get("X-API-KEY")
	user := utils.GetUserFromKey(xAPIKey)
	if user.Role != 3 {
		c.JSON(http.StatusUnauthorized, utils.ErrorRes(utils.InvalidAPIKEY))
		c.Abort()
		return
	}
}
