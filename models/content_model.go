package models

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var contents = make([]Content, 0)

// Content 内容
type Content struct {
	Id         string // 内容id
	DeviceName string // 设备名称
	Type       int    // 内容类型（文本、HTML、JSON、XML...）
	Text       string // 内容
	CreateDate int64  // 创建时间
	UpdateDate int64  // 更新时间
}

func (receiver Content) toContentWithDevice() ContentWithDevice {
	return ContentWithDevice{
		receiver,
		QueryDeviceTypeByDeviceName(receiver.DeviceName),
	}
}

type ContentWithDevice struct {
	Content        // 内容
	DeviceType int // 设备类型
}

// GetContents 获取所有设备。
func GetContents() []ContentWithDevice {
	tempContent := make([]ContentWithDevice, 0)
	for _, content := range contents {
		contentWithDevice := content.toContentWithDevice()
		tempContent = append(tempContent, contentWithDevice)
	}
	return tempContent
}

// AddContent 如果有新内容加入就添加到contents中。
func AddContent(content Content) int {
	contents = append(contents, content)
	return len(contents)
}

// BuildId 根据随机数进行md5编码返回id
func BuildId() string {
	rand.Seed(time.Now().UnixMilli())
	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int()+rand.Int()))))
}

// QueryContentById 根据id查询内容，如果查询成功就返回内容否则就返回空。
func QueryContentById(id string) any {
	for _, content := range contents {
		if content.Id == id {
			return content
		}
	}
	return nil
}

// QueryContentsByDeviceName 根据设备名查询内容返回内容数组。
func QueryContentsByDeviceName(deviceName string) []ContentWithDevice {
	tempContents := make([]ContentWithDevice, 0)
	for _, content := range contents {
		if content.DeviceName == deviceName {
			tempContents = append(tempContents, content.toContentWithDevice())
		}
	}
	return tempContents
}

// DeleteContent 根据id和设备名删除设备。删除成功就返回true否则就返回false。
func DeleteContent(id string, deviceName string) bool {
	deleteIndex := -1
	for index, content := range contents {
		if content.Id == id && content.DeviceName == deviceName {
			deleteIndex = index
		}
	}

	temp := make([]Content, 0)
	if deleteIndex != -1 {
		temp = append(temp, contents[:deleteIndex]...)
		temp = append(temp, contents[deleteIndex+1:]...)
		contents = temp
		return true
	}

	return false
}

// UpdateContent 更新内容
func UpdateContent(content Content) bool {
	for index, oldContent := range contents {
		if oldContent.Id == content.Id {
			contents[index] = content
			return true
		}
	}
	return false
}
