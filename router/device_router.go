package router

import (
	controller "contentShareManage/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(r *gin.Engine) {
	deviceGroup := r.Group("/device", func(context *gin.Context) {
		fmt.Println("/device 中间件")
	})
	{
		deviceGroup.GET("/", controller.Index)
		deviceGroup.POST("/", controller.Store)
		deviceGroup.GET("/:name", controller.Show)
		deviceGroup.DELETE("/:name", controller.Destroy)
	}
}
