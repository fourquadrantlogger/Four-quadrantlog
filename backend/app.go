package main

import (
	_ "fourquadrantlog/assist/xlog"
	"fourquadrantlog/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/log", controller.CreateLog)
	r.POST("/api/blob", controller.CreateBlob)
	r.PUT("/api/log", controller.UpdateLog)
	r.DELETE("/api/log:logid", controller.DeleteLog)
	r.GET("/api/blob/:blobid", controller.GetBlob)
	r.GET("/api/log/:logid", controller.GetLog)
	r.GET("/api/log", controller.GetLogs)

	r.Run(":10008")
}
