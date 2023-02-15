package controllers

import (
	model "contentShareManage/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ContentIndex @title 查询所有内容
func ContentIndex(context *gin.Context) {
	var contents []model.Content = nil
	deviceName, exists := context.GetQuery("deviceName")
	if exists {
		contents = model.QueryContentsByDeviceName(deviceName)
	} else {
		contents = model.GetContents()
	}

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
		start = 0
	}

	if end > len(contents) {
		end = len(contents)
	}

	context.JSON(http.StatusOK, contents[start:end])
}

// ContentStore @title 新增内容
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
	model.AddContent(content)
	context.JSON(http.StatusOK, content)
}

// ContentShow @title 查询某个内容
func ContentShow(context *gin.Context) {
	fmt.Println(context.Param("id"))
	content := model.QueryContentById(context.Param("id"))
	if content == nil {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}

	context.JSON(http.StatusOK, content)
}

// ContentUpdate @title 更新内容
func ContentUpdate(context *gin.Context) {
	tempContent, exists := context.Get("content")
	if !exists {
		context.String(http.StatusNotFound, "找不到内容！")
		return
	}

	content := tempContent.(model.Content)
	newContent := model.Content{}
	context.BindJSON(&newContent)
	newContent.Id = content.Id
	newContent.DeviceName = content.DeviceName
	model.UpdateContent(newContent)
	context.String(http.StatusOK, "更新成功！")
}

// ContentDestroy @title 删除内容
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
