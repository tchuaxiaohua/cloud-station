package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// fileCheck 检测文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	fmt.Println("无法确定文件是否存在:", err)
	return false
}

// FIleName 获取文件名称
func FileName(path string) string {
	ok := PathExists(path)
	if ok {
		return filepath.Base(path)
	}
	fmt.Println("文件不存在")
	return ""
}
