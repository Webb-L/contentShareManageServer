package models

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var devices = make([]Device, 0)

// Device 设备
type Device struct {
	Name    string // 设备名称
	Token   string // Token
	Type    int    // 设备类型（Android、IOS、PC）
	IsAdmin bool   // 是否是管理员
	IsRead  bool   // 是否可以查看
	IsSend  bool   // 是否可以发送
}

// GetDevices 获取所有设备。
func GetDevices() []Device {
	return devices
}

// Size 获取有多少设备连接。
func Size() int {
	return len(devices)
}

// AddDevice 如果有新设备加入就添加到devices中。
func AddDevice(device Device) int {
	devices = append(devices, device)
	return len(devices)
}

// QueryDevice 根据deviceName查询设备是否存在。如果存在就返回设备信息否则就返回空。
func QueryDevice(deviceName string) interface{} {
	for _, device := range devices {
		fmt.Println(device.Name, deviceName)
		if device.Name == deviceName {
			return device
		}
	}
	return nil
}

// DeleteDevice 根据设备名称和设备密码删除设备。删除成功就返回true否则就返回false。
func DeleteDevice(name string) bool {
	deleteIndex := -1
	for index, device := range devices {
		if device.Name == name {
			deleteIndex = index
		}
	}

	temp := make([]Device, 0)
	if deleteIndex != -1 {
		temp = append(temp, devices[:deleteIndex]...)
		temp = append(temp, devices[deleteIndex+1:]...)
		devices = temp
		return true
	}
	return false
}

// UpdateDevice 更新设备
func UpdateDevice(device Device) bool {
	for index, oldDevice := range devices {
		if oldDevice.Name == device.Name {
			devices[index] = device
			return true
		}
	}
	return false
}

// ExistDevice 根据deviceName参数判断这个设备是否存在如果存在就返回true否则false。
func ExistDevice(deviceName string) bool {
	for _, device := range devices {
		if device.Name == deviceName {
			return true
		}
	}
	return false
}

// BuildToken 根据随机数进行md5编码返回token。
func BuildToken() string {
	rand.Seed(time.Now().UnixMilli())
	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int()))))
}

// QueryDeviceByToken 根据token查询设备是否存在。如果存在就返回设备信息否则就返回空。
func QueryDeviceByToken(token string) interface{} {
	for _, device := range devices {
		if device.Token == token {
			return device
		}
	}
	return nil
}
