package utils

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"os"
)

// ImportAndProcessJSON 从 JSON 文件导入数据到数据库，并支持对每条数据执行额外处理逻辑
func ImportAndProcessJSON[T any](db *gorm.DB, filePath string, tableName string, processFunc func(*T) error) error {
	// 打开 JSON 文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	// 解析 JSON 内容
	var data []T
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	// 可选：对每项执行处理逻辑
	if processFunc != nil {
		for i := range data {
			if err = processFunc(&data[i]); err != nil {
				return fmt.Errorf("processing item %d failed: %v", i, err)
			}
		}
	}

	// 插入数据库
	if err = db.Create(&data).Error; err != nil {
		return fmt.Errorf("failed to insert into table %s: %v", tableName, err)
	}

	return nil
}
