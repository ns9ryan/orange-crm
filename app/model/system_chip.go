// ==========================================================================
// 模型-系统筹码表
// ==========================================================================

package model

import (
	"orange-crm/global"
)

// systemChipColumns defines and stores column names for the table or_system_chip.
type systemChipColumns struct {
	Id     string // ID
	Name   string // 筹码名称
	Value  string // 面额
	Digits string // 筹码的最小单位倍数
	Status string // 是否启用:0否,1是
}

// SystemChipColumns holds the columns for the table or_system_chip.
var SystemChipColumns = systemChipColumns{
	Id:     "id",
	Name:   "name",
	Value:  "value",
	Digits: "digits",
	Status: "status",
}

// SystemChip 系统筹码表结构体
type SystemChip struct {
	ID     uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Name   string `gorm:"size:32;not null;comment:筹码名称"`
	Value  uint   `gorm:"type:int unsigned;not null;comment:面额"`
	Digits uint   `gorm:"type:int unsigned;not null;default:1;comment:筹码的最小单位倍数"`
	Status uint   `gorm:"type:tinyint unsigned;not null;default:1;comment:是否启用:0否,1是"`
}

// RegisterTable 数据迁移
func (m *SystemChip) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统筹码表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		data := []SystemChip{
			{Name: "面值：0.1", Value: 1, Digits: 10},
			{Name: "面值：0.2", Value: 2, Digits: 10},
			{Name: "面值：0.5", Value: 5, Digits: 10},
			{Name: "面值：1", Value: 1, Digits: 1},
			{Name: "面值：2", Value: 2, Digits: 1},
			{Name: "面值：5", Value: 5, Digits: 1},
			{Name: "面值：10", Value: 10, Digits: 1},
			{Name: "面值：20", Value: 20, Digits: 1},
			{Name: "面值：50", Value: 50, Digits: 1},
			{Name: "面值：100", Value: 100, Digits: 1},
			{Name: "面值：200", Value: 200, Digits: 1},
			{Name: "面值：500", Value: 500, Digits: 1},
			{Name: "面值：1000", Value: 1000, Digits: 1},
			{Name: "面值：2000", Value: 2000, Digits: 1},
			{Name: "面值：5000", Value: 5000, Digits: 1},
			{Name: "面值：1W", Value: 10000, Digits: 1},
			{Name: "面值：2W", Value: 20000, Digits: 1},
			{Name: "面值：5W", Value: 50000, Digits: 1},
			{Name: "面值：10W", Value: 100000, Digits: 1},
			{Name: "面值：20W", Value: 200000, Digits: 1},
			{Name: "面值：50W", Value: 500000, Digits: 1},
			{Name: "面值：100W", Value: 1000000, Digits: 1},
			{Name: "面值：200W", Value: 2000000, Digits: 1},
			{Name: "面值：500W", Value: 5000000, Digits: 1},
		}
		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
