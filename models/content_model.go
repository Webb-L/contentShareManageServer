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
}

// GetContents @title 获取所有设备。
// @description 获取所有的设备。
// @return []Content 返回所有设备。
func GetContents() []Content {
	return contents
}

// AddContent @title 添加内容。
// @description 如果有新内容加入就添加到contents中。
// @param content Content 新增的内容。
// @return int 返回内容数量。
func AddContent(content Content) int {
	contents = append(contents, content)
	return len(contents)
}

// BuildId @title 生成内容Id。
// @description 根据随机数进行md5编码返回id。
// @return string md5编码后的字符串。
func BuildId() string {
	rand.Seed(time.Now().UnixMilli())
	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int()+rand.Int()))))
}

// QueryContentById @title 根据id查询内容。
// @description 根据id查询内容，如果查询成功就返回内容否则就返回空。
// @param id 内容id。
// @return Content|null 如果查询成功就返回内容否则就返回空。
func QueryContentById(id string) any {
	for _, content := range contents {
		fmt.Println(content.Id, id, content.Id == id)
		if content.Id == id {
			return content
		}
	}
	return nil
}

// QueryContentsByDeviceName @title 根据设备名查询内容。
// @description 根据设备名查询内容返回内容数组。
// @param deviceName string 设备名。
// @return []Content 内容数组。
func QueryContentsByDeviceName(deviceName string) []Content {
	temp := make([]Content, 0)
	for _, content := range contents {
		if content.DeviceName == deviceName {
			temp = append(temp, content)
		}
	}
	return temp
}

// DeleteContent @title 删除内容。
// @description 根据id和设备名删除设备。删除成功就返回true否则就返回false。
// @param id string 内容id。
// @param deviceName string 设备名。
// @return bool 返回删除状态。
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
		return true
	}

	return false
}

// UpdateContent @title 更新内容
// @description 更新内容。
// @param content Content 更新的设备信息。
// @return bool 如果更新成功就返回true否则返回false
func UpdateContent(content Content) bool {
	for index, oldContent := range contents {
		if oldContent.Id == content.Id {
			contents[index] = content
			return true
		}
	}
	return false
}
