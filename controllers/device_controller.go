package controllers

import (
	model "contentShareManage/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	devices := model.GetDevices()
	context.JSON(http.StatusOK, devices)
}

func Store(context *gin.Context) {
	device := model.Device{}
	context.BindJSON(&device)
	context.JSON(http.StatusOK, device)
}

func Show(context *gin.Context) {
	fmt.Println(context.Param("name"))
}

// Destroy @title 删除设备
// @description 删除设备（管理员权限）。
func Destroy(context *gin.Context) {

}
