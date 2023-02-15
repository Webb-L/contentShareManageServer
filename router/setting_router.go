package router

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var password = ""

func InitSettingRouter(r *gin.Engine) {
	deviceGroup := r.Group("/setting", func(context *gin.Context) {
		fmt.Println(context.Request.URL.Path)
	})
	{
		deviceGroup.GET("/password", func(context *gin.Context) {
			if password == "" {
				rand.Seed(time.Now().UnixMilli())
				password = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int())+"password"+strconv.Itoa(rand.Int()))))
				context.String(http.StatusOK, password)
			} else {
				context.String(http.StatusForbidden, "密码已获取。\n如果忘记密码请重启服务端。")
			}
		})
	}
}
