package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func SendContentGet(context *gin.Context) {
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusBadRequest, "找不到设备！")
		return
	}

	content := model.Content{}
	content.Id = model.BuildId()

	_type, typeExist := context.GetQuery("type")
	if !typeExist {
		context.String(http.StatusBadRequest, "请设置类型！")
		return
	}

	contentType, err := strconv.Atoi(_type)
	if err != nil {
		context.String(http.StatusBadRequest, "请输入正确的类型！")
		return
	}
	content.Type = contentType

	text, textExist := context.GetQuery("text")
	if !textExist {
		context.String(http.StatusBadRequest, "请设置内容！")
		return
	}
	content.Text = text

	content.DeviceName = deviceName.(string)
	content.CreateDate = time.Now().UnixMilli()
	content.UpdateDate = time.Now().UnixMilli()
	model.AddContent(content)
	context.JSON(http.StatusOK, content)
}
func SendContentPost(context *gin.Context) {
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusBadRequest, "找不到设备！")
		return
	}
	content := model.Content{}
	context.BindJSON(&content)
	content.Id = model.BuildId()
	content.DeviceName = deviceName.(string)
	content.CreateDate = time.Now().UnixMilli()
	content.UpdateDate = time.Now().UnixMilli()
	model.AddContent(content)
	context.JSON(http.StatusOK, content)
}
