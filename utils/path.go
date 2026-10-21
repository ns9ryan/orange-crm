package utils

import (
	"go.uber.org/zap"
	"orange-crm/global"
	"os"
	"path/filepath"
)

// ProjectRoot 获取当前项目根目录（支持独立运行和被托管运行）
func ProjectRoot() string {
	root, err := os.Getwd()
	if err != nil {
		global.LOG.Error("os.Getwd failed:", zap.Any("err", err))
	}
	// 如果是独立运行，返回根目录/name
	if _, ok := os.LookupEnv("RUN_AS_STANDALONE"); ok {
		return root
	}

	elem := []string{
		root,
	}

	for _, v := range global.CONFIG.Server.OriginPathPrefix {
		elem = append(elem, v)
	}

	return filepath.Join(elem...)
}

// GetDirByName 通过name获取对应的文件夹路径
func GetDirByName(name string) string {
	return filepath.Join(ProjectRoot(), name)
}
