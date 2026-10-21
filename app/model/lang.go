package model

import (
	"orange-crm/global"
)

// Lang 系统语言表结构体
type Lang struct {
	ID          uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Lang        string `gorm:"unique;size:10;not null;comment:语言"`
	Description string `gorm:"size:20;not null;comment:描述"`
}

// RegisterTable 数据迁移
func (m *Lang) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统语言表'").Migrator().CreateTable(m); err != nil {
			return err
		}
		data := []Lang{
			{Lang: "zh", Description: "汉语"},
			{Lang: "en", Description: "英语"},
			{Lang: "vi", Description: "越南语"},
			{Lang: "pt", Description: "葡萄牙语"},
			{Lang: "ru", Description: "俄语"},
			{Lang: "hi", Description: "印地语"},
			{Lang: "id", Description: "印尼语"},
			{Lang: "th", Description: "泰语"},
			{Lang: "es", Description: "西班牙语"},
		}
		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
