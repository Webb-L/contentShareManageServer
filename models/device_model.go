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

// GetDevices @title 获取所有设备。
// @description 获取所有的设备。
// @return []Device 返回所有设备。
func GetDevices() []Device {
	return devices
}

// Size @title 设备数量。
// @description 获取有多少设备连接。
// @return int 返回设备数量。
func Size() int {
	return len(devices)
}

// AddDevice @title 添加设备。
// @description 如果有新设备加入就添加到devices中。
// @param device Device 新增的设备。
// @return int 返回设备数量。
func AddDevice(device Device) int {
	devices = append(devices, device)
	return len(devices)
}

// QueryDevice @title 查询设备。
// @description 根据deviceName查询设备是否存在。如果存在就返回设备信息否则就返回空。
// @param deviceName string 设备名。
// @return Device|nil 设备信息或空。
func QueryDevice(deviceName string) interface{} {
	for _, device := range devices {
		fmt.Println(device.Name, deviceName)
		if device.Name == deviceName {
			return device
		}
	}
	return nil
}

// DeleteDevice @title 删除设备。
// @description 根据设备名称和设备密码删除设备。删除成功就返回true否则就返回false。
// @param name string 设备名称。
// @param password string 设备密码。
// @return bool 返回删除状态。
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

// UpdateDevice @title 更新设备
// @description 更新设备。
// @param device Device 更新的设备信息。
// @return bool 如果更新成功就返回true否则返回false
func UpdateDevice(device Device) bool {
	for index, oldDevice := range devices {
		if oldDevice.Name == device.Name {
			devices[index] = device
			return true
		}
	}
	return false
}

// ExistDevice @title 判断设备是否存在。
// @description 根据deviceName参数判断这个设备是否存在如果存在就返回true否则false。
// @param deviceName string 查询的设备名称。
// @return bool 是否存在如果存在就返回true否则false。
func ExistDevice(deviceName string) bool {
	for _, device := range devices {
		if device.Name == deviceName {
			return true
		}
	}
	return false
}

// BuildToken @title 生成token。
// @description 根据随机数进行md5编码返回token。
// @return string md5编码后的字符串。
func BuildToken() string {
	rand.Seed(time.Now().UnixMilli())
	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int()))))
}

// QueryDeviceByToken @title 根据token查询设备。
// @description 根据token查询设备是否存在。如果存在就返回设备信息否则就返回空。
// @param token string 设备token。
// @return Device|nil 设备信息或空。
func QueryDeviceByToken(token string) interface{} {
	for _, device := range devices {
		if device.Token == token {
			return device
		}
	}
	return nil
}
