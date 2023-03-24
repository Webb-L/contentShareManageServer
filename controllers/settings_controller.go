package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Password = ""
var PassCount int

func init() {
	rand.Seed(time.Now().UnixMilli())
	Password = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int())+"password"+strconv.Itoa(rand.Int()))))
}

// GetPassword 获取服务器密码
//
//	@Summary		获取服务器密码
//	@Description	获取服务器密码，如果服务器密码已经获取就返回错误“密码已获取。如果忘记密码请重启服务端。”
//	@Tags			settings
//	@Accept			plain
//	@Produce		plain
//	@Success		200	string	string	"32位随机字符串"
//	@Failure		403	string	string	"密码已获取。如果忘记密码请重启服务端。"
//	@Router			/settings/password/ [get]
func GetPassword(context *gin.Context) {
	if PassCount > 0 {
		context.String(http.StatusOK, Password)
		PassCount--
	} else {
		context.String(http.StatusForbidden, "密码已获取。\n如果忘记密码请重启服务端。")
	}
}
