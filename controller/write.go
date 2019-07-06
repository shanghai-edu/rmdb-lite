package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanghai-edu/rmdb-lite/controller/utils"
	"github.com/shanghai-edu/rmdb-lite/models"
)

func deleteRouter(c *gin.Context) {
	ip := c.Query("ip")
	err := models.DeleteRouter(ip)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, utils.ErrorRes(utils.RecordNotFound))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorRes(utils.InternalAPIError))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessRes(nil))
}

func updateRouter(c *gin.Context) {
	inputs := models.Router{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorRes(utils.BodyJSONDecodeError))
		return
	}
	err := models.UpdateRouter(inputs)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, utils.ErrorRes(utils.RecordNotFound))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorRes(utils.InternalAPIError))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessRes(nil))
}

func createRouter(c *gin.Context) {
	inputs := models.Router{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorRes(utils.BodyJSONDecodeError))
		return
	}
	err := models.CreateRouter(inputs)
	if err != nil {
		if err.Error() == "router already exists" {
			c.JSON(http.StatusNotFound, utils.ErrorRes(utils.RecordAlreadyExists))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorRes(utils.InternalAPIError))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessRes(nil))
}
