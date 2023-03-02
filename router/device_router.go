package router

import (
	controller "contentShareManage/controllers"
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDeviceRouter(r *gin.Engine) {
	deviceGroup := r.Group("/device", func(context *gin.Context) {
		queryPassword, exists := context.GetQuery("password")
		if password == "" || !exists || queryPassword != password {
			context.String(http.StatusUnauthorized, "密码错误！")
			context.Abort()
			context.Next()
			return
		}

		if context.Request.Method != "GET" && context.Request.Method != "POST" {
			// 检测Token 是否有效。
			token := context.GetHeader("Token")
			if token == "" {
				context.String(http.StatusUnauthorized, "您没有权限操作！")
				context.Abort()
				context.Next()
				return
			}
			temp := model.QueryDeviceByToken(token)
			if temp == nil {
				context.String(http.StatusUnauthorized, "您没有权限操作！")
				context.Abort()
				context.Next()
				return
			}

			// 判断Token的用户是否有权限操作。
			device := temp.(model.Device)
			if !device.IsAdmin {
				context.String(http.StatusForbidden, "您没有权限操作！")
				context.Abort()
				context.Next()
				return
			}

			// 查询被操作的设备
			name := context.Param("name")
			if name == "" {
				context.String(http.StatusNotFound, "找不到设备！")
				context.Abort()
				context.Next()
				return
			}

			queryDevice := model.QueryDevice(name)
			if queryDevice == nil {
				context.String(http.StatusNotFound, "找不到设备！")
				context.Abort()
				context.Next()
				return
			}

			context.Set("device", queryDevice.(model.Device))
		}
	})
	{
		deviceGroup.GET("/", controller.DeviceIndex)
		deviceGroup.POST("/", controller.DeviceStore)
		deviceGroup.GET("/:name", controller.DeviceShow)
		deviceGroup.PUT("/:name", controller.DeviceUpdate)
		deviceGroup.DELETE("/:name", controller.DeviceDestroy)
	}
}
