package controller

import (
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {

	readOnly := r.Group("/api/v1")
	readOnly.Use(readOnlyCheckMidd)
	readOnly.GET("/router", getRouter)
	readOnly.POST("/router/multi", getMultiRouters)

	admin := r.Group("/api/v1")
	admin.Use(adminCheckMidd)
	admin.GET("/router/all", getAllRouters)
	admin.DELETE("/router", deleteRouter)
	admin.POST("/router", createRouter)
	admin.PUT("/router", updateRouter)

}
