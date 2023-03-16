package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	InitDeviceRouter(r)
	InitContentRouter(r)
	InitWebHookRouter(r)
	InitSettingRouter(r)
}
