package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"time"
)

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

func ContentShow(context *gin.Context) {
	_content := model.QueryContentById(context.Param("id"))
	if _content == nil {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}
	content := _content.(model.Content)
	context.JSON(http.StatusOK, content.ToContentWithDevice())
}

func ContentUpdate(context *gin.Context) {
	tempContent, exists := context.Get("content")
	if !exists {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusBadRequest, "找不到设备！")
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
		context.String(http.StatusOK, "更新失败！")
	}
}

func ContentDestroy(context *gin.Context) {
	id := context.Param("id")
	deviceName, exists := context.Get("deviceName")
	if !exists {
		context.String(http.StatusBadRequest, "找不到设备！")
		return
	}

	if model.DeleteContent(id, deviceName.(string)) {
		context.String(http.StatusOK, "删除成功！")
	} else {
		context.String(http.StatusForbidden, "删除失败！")
	}
}
