package router

import (
	controller "contentShareManage/controllers"
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitContentRouter(r *gin.Engine) {
	contentGroup := r.Group("/content", func(context *gin.Context) {
		queryPassword, exists := context.GetQuery("password")
		if password == "" || !exists || queryPassword != password {
			context.String(http.StatusUnauthorized, "密码错误！")
			context.Abort()
			context.Next()
			return
		}

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
		device := temp.(model.Device)
		context.Set("deviceName", device.Name)
		// 权限判断
		if !device.IsAdmin && !device.IsRead && context.Request.Method == "GET" {
			context.String(http.StatusUnauthorized, "您没有权限操作！")
			context.Abort()
			context.Next()
			return
		}
		if !device.IsAdmin && !device.IsSend && context.Request.Method != "GET" {
			context.String(http.StatusUnauthorized, "您没有权限操作！")
			context.Abort()
			context.Next()
			return
		}

		if context.Request.Method != "GET" && context.Request.Method != "POST" {
			id := context.Param("id")
			if id == "" {
				context.String(http.StatusNotFound, "找不到内容！")
				context.Abort()
				context.Next()
				return
			}

			content := model.QueryContentById(id)
			if content == nil {
				context.String(http.StatusNotFound, "找不到内容！")
				context.Abort()
				context.Next()
				return
			}

			context.Set("content", content)
		}
	})
	{
		contentGroup.GET("/", controller.ContentIndex)
		contentGroup.POST("/", controller.ContentStore)
		contentGroup.GET("/:id", controller.ContentShow)
		contentGroup.PUT("/:id", controller.ContentUpdate)
		contentGroup.DELETE("/:id", controller.ContentDestroy)
	}
}
