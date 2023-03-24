package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// SendContentGet 新增内容
//
//	@Summary		新增内容
//	@Description	使用GET方式 新增内容
//	@Tags			webhook
//	@Accept			json
//	@Produce		json
//	@Param			token	path		string	true	"WebHook凭证"
//	@Param			type	query		int		true	"内容类型"
//	@Param			text	query		string	true	"发送的内容"
//	@Success		200		{object}	models.Content
//	@Failure		400		string		string	"请设置类型！"
//	@Failure		400		string		string	"请输入正确的类型！"
//	@Failure		400		string		string	"请设置内容！"
//	@Failure		401		string		string	"您没有权限操作！"
//	@Failure		404		string		string	"找不到设备！"
//	@Router			/webhook/:token/ [get]
func SendContentGet(context *gin.Context) {
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
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

// SendContentPost 新增内容
//
//	@Summary		新增内容
//	@Description	使用POST方式 新增内容
//	@Tags			webhook
//	@Accept			json
//	@Produce		json
//	@Param			token	path		string			true	"WebHook凭证"
//	@Param			content	body		models.Content	true	"新内容信息"
//	@Success		200		{object}	models.Content
//	@Failure		404		string		string	"找不到设备！"
//	@Router			/webhook/:token/ [post]
func SendContentPost(context *gin.Context) {
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
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
