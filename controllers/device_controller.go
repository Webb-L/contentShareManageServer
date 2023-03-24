package controllers

import (
	model "contentShareManage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeviceIndex 获取所有设备
//
//	@Summary		显示设备信息
//	@Description	根据数据开始索引和结束索引查询数据.
//	@Tags			device
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string	true	"连接服务器密码"
//	@Param			start		query		int		false	"设备数据开始索引"
//	@Param			end			query		int		false	"设备数据结束索引"
//	@Success		200			{object}	models.Device
//	@Failure		401			string		string	"密码错误！"
//	@Failure		403			string		string	"您没有权限操作！"
//	@Failure		404			string		string	"找不到设备！"
//	@Router			/device/ [get]
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
		context.JSON(http.StatusOK, make([]model.Device, 0))
		return
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

// DeviceStore 新增设备
//
//	@Summary		新增设备
//	@Description	新增设备,默认只有发送权限.
//	@Tags			device
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string			true	"连接服务器密码"
//	@Param			device		body		models.Device	true	"新设备信息"
//	@Success		200			{object}	models.Device
//	@Success		202			string		string	"设备已存在！"
//	@Failure		401			string		string	"密码错误！"
//	@Router			/device/ [post]
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

// DeviceShow 查询某台设备
//
//	@Summary		显示设备信息
//	@Description	返回某台设备信息
//	@Tags			device
//	@Accept			json
//	@Produce		json
//	@Param			password	query		string	true	"连接服务器密码"
//	@Param			name		path		string	true	"设备名"
//	@Success		200			{object}	models.Device
//	@Failure		401			string		string	"密码错误！"
//	@Failure		403			string		string	"您没有权限操作！"
//	@Failure		404			string		string	"找不到设备！"
//	@Router			/device/:name/ [get]
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

// DeviceUpdate 更新设备
//
//	@Summary		更新设备
//	@Description	更新设备名称更新设备
//	@Tags			device
//	@Accept			json
//	@Produce		json
//	@Param			password	query	string			true	"连接服务器密码"
//	@Param			name		path	string			true	"设备名"
//	@Param			device		body	models.Device	true	"更新设备信息"
//	@Success		200			string	string			"更新成功！"
//	@Failure		401			string	string			"密码错误！"
//	@Failure		403			string	string			"您没有权限操作！"
//	@Failure		404			string	string			"找不到设备！"
//	@Router			/device/:name/ [put]
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

// DeviceDestroy 删除设备
//
//	@Summary		删除设备
//	@Description	根据设备名称删除设备
//	@Tags			device
//	@Accept			json
//	@Produce		json
//	@Param			password	query	string	true	"连接服务器密码"
//	@Param			name		path	string	true	"设备名"
//	@Success		200			string	string	"更新成功！"
//	@Success		202			string	string	"管理员账户不能删除！"
//	@Failure		401			string	string	"密码错误！"
//	@Failure		403			string	string	"您没有权限操作！"
//	@Failure		404			string	string	"找不到设备！"
//	@Router			/device/:name/ [delete]
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
