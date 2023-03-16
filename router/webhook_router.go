package router

import (
	"contentShareManage/controllers"
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebHookRouter(r *gin.Engine) {
	webHookGroup := r.Group("/webhook", func(context *gin.Context) {
		token := context.Param("token")
		if token == "" {
			context.String(http.StatusUnauthorized, "您没有权限操作！")
			context.Abort()
			context.Next()
			return
		}

		temp := model.QueryDeviceByTokenType(token, 6)
		if temp == nil {
			context.String(http.StatusUnauthorized, "您没有权限操作！")
			context.Abort()
			context.Next()
			return
		}
		device := temp.(model.Device)
		context.Set("deviceName", device.Name)
	})
	{
		webHookGroup.GET("/:token/", controllers.SendContentGet)
		webHookGroup.POST("/:token/", controllers.SendContentPost)
	}
}
