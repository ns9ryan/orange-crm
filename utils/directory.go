package utils

import "os"

// CreateDirectory 创建文件夹
func CreateDirectory(path string) error {
	// 使用 os.MkdirAll 来递归创建目录
	return os.MkdirAll(path, os.ModePerm)
}
