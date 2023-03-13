package main

import (
	router "contentShareManage/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRouter(r)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run(":8000")
}
