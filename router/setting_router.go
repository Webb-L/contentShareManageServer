package router

import (
	controller "contentShareManage/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitSettingRouter(r *gin.Engine) {
	deviceGroup := r.Group("/settings", func(context *gin.Context) {
		fmt.Println(context.Request.URL.Path)
	})
	{
		deviceGroup.GET("/password/", controller.GetPassword)
	}
}
