package models

var devices = make([]Device, 0)

// Device 设备
type Device struct {
	Name     string // 设备名称
	PassWord string // 设备密码
	Type     int    // 设备类型（Android、IOS、PC）
	IsAdmin  bool   // 是否是管理员
	IsSend   bool   // 是否可以发送
}

// GetDevices @title 获取所有设备。
// @description 获取所有的设备。
// @return []Device 返回所有设备。
func GetDevices() []Device {
	return devices
}

// AddDevice @title 添加设备。
// @description 如果有新设备加入就添加到devices中。
// @param device Device 新增的设备。
// @return int 返回设备数量。
func AddDevice(device Device) int {
	devices = append(devices, device)
	return len(devices)
}
