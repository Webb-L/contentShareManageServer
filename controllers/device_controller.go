package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeviceIndex @title 查询设备
// @description 查询当前有多少设备
func DeviceIndex(context *gin.Context) {
	devices := model.GetDevices()
	startString, exists := context.GetQuery("start")
	if !exists {
		startString = "0"
	}
	endString, exists := context.GetQuery("end")
	if !exists {
		endString = "10"
	}

	start, err := strconv.Atoi(startString)
	if err != nil {
		start = 0
	}

	end, err := strconv.Atoi(endString)
	if err != nil {
		end = 10
	}

	if start > len(devices) {
		start = 0
	}

	if end > len(devices) {
		end = len(devices)
	}

	resultDevice := make([]model.Device, 0)
	for _, device := range devices[start:end] {
		device.Token = ""
		resultDevice = append(resultDevice, device)
	}

	context.JSON(http.StatusOK, resultDevice)
}

// DeviceStore @title 新增设备
// @description 保存新的设备。如果是第一台设备连接设置为管理员。
// @param context *gin.Context
func DeviceStore(context *gin.Context) {
	device := model.Device{}
	context.BindJSON(&device)
	firstDevice := model.Size() == 0
	device.IsAdmin = firstDevice
	device.IsRead = firstDevice
	device.IsSend = true
	device.Token = model.BuildToken()
	if model.ExistDevice(device.Name) {
		context.String(http.StatusAccepted, "设备已存在！")
		return
	}
	model.AddDevice(device)
	context.JSON(http.StatusOK, device)
}

// DeviceShow @title 查询设备信息
// @description 显示设备名称。
func DeviceShow(context *gin.Context) {
	tempDevice := model.QueryDevice(context.Param("name"))
	if tempDevice == nil {
		context.String(http.StatusNotFound, "找不到设备！")
		return
	}
	device := tempDevice.(model.Device)
	device.Token = ""
	context.JSON(http.StatusOK, device)
}

// DeviceUpdate @title 更新设备
// @description 更新设备信息。
func DeviceUpdate(context *gin.Context) {
	tempDevice, exists := context.Get("device")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
		return
	}

	// 找到设备
	device := tempDevice.(model.Device)
	newDevice := model.Device{}
	context.BindJSON(&newDevice)
	newDevice.Name = device.Name
	newDevice.Token = device.Token
	model.UpdateDevice(newDevice)
	context.String(http.StatusOK, "更新成功！")
}

// DeviceDestroy @title 删除设备
// @description 删除设备。
func DeviceDestroy(context *gin.Context) {
	tempDevice, exists := context.Get("device")
	if !exists {
		context.String(http.StatusNotFound, "找不到设备！")
		return
	}

	device := tempDevice.(model.Device)
	if device.IsAdmin {
		context.String(http.StatusAccepted, "管理员账户不能删除！")
	}
	if model.DeleteDevice(device.Name) {
		context.String(http.StatusOK, "删除成功！")
	} else {
		context.String(http.StatusNotFound, "找不到设备！")
	}
}
