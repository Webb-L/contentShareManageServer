package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// ContentIndex 获取所有内容
//
//	@Summary		显示内容信息
//	@Description	根据数据开始索引和结束索引查询数据.
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string	true	"连接服务器密码"
//	@Param			deviceName	query		string	false	"设备名"
//	@Param			start		query		int		false	"内容数据开始索引"
//	@Param			end			query		int		false	"内容数据结束索引"
//	@Success		200			{object}	models.ContentWithDevice
//	@Failure		401			string		string	"密码错误！"
//	@Failure		403			string		string	"您没有权限操作！"
//	@Failure		404			string		string	"找不到内容！"
//	@Router			/content/ [get]
func ContentIndex(context *gin.Context) {
	var contents []model.ContentWithDevice = nil
	deviceName, exists := context.GetQuery("deviceName")
	if exists {
		contents = model.QueryContentsByDeviceName(deviceName)
	} else {
		contents = model.GetContents()
	}

	sort.Slice(contents, func(i, j int) bool {
		return contents[i].UpdateDate > contents[j].UpdateDate
	})

	startString, exists := context.GetQuery("start")
	endString, exists := context.GetQuery("end")

	start, err := strconv.Atoi(startString)
	if err != nil {
		start = 0
	}

	end, err := strconv.Atoi(endString)
	if err != nil {
		end = 10
	}

	if start > len(contents) {
		context.JSON(http.StatusOK, make([]model.ContentWithDevice, 0))
		return
	}

	if end > len(contents) {
		end = len(contents)
	}

	context.JSON(http.StatusOK, contents[start:end])
}

// ContentStore 新增内容
//
//	@Summary		新增内容
//	@Description	新增内容
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string			true	"连接服务器密码"
//	@Param			content		body		models.Content	true	"新设备信息"
//	@Success		200			{object}	models.ContentWithDevice
//	@Failure		401			string		string	"密码错误！"
//	@Failure		403			string		string	"您没有权限操作！"
//	@Router			/content/ [post]
func ContentStore(context *gin.Context) {
	content := model.Content{}
	context.BindJSON(&content)
	content.Id = model.BuildId()
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusBadRequest, "找不到设备！")
		return
	}
	content.DeviceName = deviceName.(string)
	content.CreateDate = time.Now().UnixMilli()
	content.UpdateDate = time.Now().UnixMilli()
	model.AddContent(content)
	context.JSON(http.StatusOK, content.ToContentWithDevice())
}

// ContentShow 查询某条内容
//
//	@Summary		显示内容信息
//	@Description	返回某条内容信息
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string	true	"连接服务器密码"
//	@Param			id			path		string	true	"内容id"
//	@Success		200			{object}	models.ContentWithDevice
//	@Failure		400			string		string	"找不到设备！"
//	@Failure		400			string		string	"找不到内容！"
//	@Failure		401			string		string	"密码错误！"
//	@Failure		403			string		string	"您没有权限操作！"
//	@Router			/content/:id/ [get]
func ContentShow(context *gin.Context) {
	_content := model.QueryContentById(context.Param("id"))
	if _content == nil {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}
	content := _content.(model.Content)
	context.JSON(http.StatusOK, content.ToContentWithDevice())
}

// ContentUpdate 更新内容
//
//	@Summary		更新内容
//	@Description	更新某条内容
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			password	query	string			true	"连接服务器密码"
//	@Param			id			path	string			true	"内容id"
//	@Param			content		body	models.Content	true	"更新内容信息"
//	@Success		200			string	string			"更新成功！"
//	@Success		202			string	string			"更新失败！"
//	@Failure		400			string	string			"找不到设备！"
//	@Failure		400			string	string			"找不到内容！"
//	@Failure		401			string	string			"密码错误！"
//	@Failure		403			string	string			"您没有权限操作！"
//	@Router			/device/:id/ [put]
func ContentUpdate(context *gin.Context) {
	tempContent, exists := context.Get("content")
	if !exists {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
		return
	}
	content := tempContent.(model.Content)
	newContent := model.Content{}
	context.BindJSON(&newContent)

	content.Type = newContent.Type
	content.Text = newContent.Text
	content.UpdateDate = time.Now().UnixMilli()
	if model.UpdateContent(content, deviceName.(string)) {
		context.String(http.StatusOK, "更新成功！")
	} else {
		context.String(http.StatusAccepted, "更新失败！")
	}
}

// ContentDestroy 删除内容
//
//	@Summary		删除内容
//	@Description	根据内容名称删除内容
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			password	query	string	true	"连接服务器密码"
//	@Param			id			path	string	true	"内容id"
//	@Success		200			string	string	"删除成功！"
//	@Success		202			string	string	"删除失败！"
//	@Failure		401			string	string	"密码错误！"
//	@Failure		403			string	string	"您没有权限操作！"
//	@Failure		404			string	string	"找不到设备！"
//	@Router			/device/:id/ [delete]
func ContentDestroy(context *gin.Context) {
	id := context.Param("id")
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
		return
	}

	if model.DeleteContent(id, deviceName.(string)) {
		context.String(http.StatusOK, "删除成功！")
	} else {
		context.String(http.StatusAccepted, "删除失败！")
	}
}
