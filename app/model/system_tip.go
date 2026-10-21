package model

import (
	"orange-crm/global"
)

// systemTipColumns defines and stores column names for the table or_system_tip.
type systemTipColumns struct {
	Id       string // ID
	Currency string // 币种
	Digits   string // 币种的最小单位倍数
	Value    string // 面额
}

// SystemTipColumns holds the columns for the table or_system_tip.
var SystemTipColumns = systemTipColumns{
	Id:       "id",
	Currency: "currency",
	Digits:   "digits",
	Value:    "value",
}

// SystemTip 系统打赏配置表结构体
type SystemTip struct {
	ID       uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Currency string `gorm:"type:varchar(10);not null;comment:币种"`
	Digits   uint   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	Value    uint   `gorm:"not null;comment:面额"`
}

// RegisterTable 数据迁移
func (m *SystemTip) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统打赏配置表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		data := []SystemTip{
			// 美元1-100
			{Currency: "USD", Digits: 100, Value: 100},   // 1
			{Currency: "USD", Digits: 100, Value: 200},   // 2
			{Currency: "USD", Digits: 100, Value: 500},   // 5
			{Currency: "USD", Digits: 100, Value: 1000},  // 10
			{Currency: "USD", Digits: 100, Value: 5000},  // 50
			{Currency: "USD", Digits: 100, Value: 10000}, // 100

			// 印度卢比80-8500
			{Currency: "INR", Digits: 100, Value: 10000},  // 100
			{Currency: "INR", Digits: 100, Value: 20000},  // 200
			{Currency: "INR", Digits: 100, Value: 50000},  // 500
			{Currency: "INR", Digits: 100, Value: 100000}, // 1000
			{Currency: "INR", Digits: 100, Value: 200000}, // 2000
			{Currency: "INR", Digits: 100, Value: 300000}, // 3000

			// 越南盾20000-1000000
			{Currency: "VND", Digits: 1, Value: 20000},   // 20000
			{Currency: "VND", Digits: 1, Value: 50000},   // 50000
			{Currency: "VND", Digits: 1, Value: 100000},  // 100000
			{Currency: "VND", Digits: 1, Value: 200000},  // 200000
			{Currency: "VND", Digits: 1, Value: 500000},  // 500000
			{Currency: "VND", Digits: 1, Value: 1000000}, // 1000000

			// 欧元1-100
			{Currency: "EUR", Digits: 100, Value: 100},   // 1
			{Currency: "EUR", Digits: 100, Value: 200},   // 2
			{Currency: "EUR", Digits: 100, Value: 500},   // 5
			{Currency: "EUR", Digits: 100, Value: 1000},  // 10
			{Currency: "EUR", Digits: 100, Value: 5000},  // 50
			{Currency: "EUR", Digits: 100, Value: 10000}, // 100

			// 印尼盾20000-1000000IDR
			{Currency: "IDR", Digits: 1, Value: 20000},   // 20000
			{Currency: "IDR", Digits: 1, Value: 50000},   // 50000
			{Currency: "IDR", Digits: 1, Value: 100000},  // 100000
			{Currency: "IDR", Digits: 1, Value: 200000},  // 200000
			{Currency: "IDR", Digits: 1, Value: 500000},  // 500000
			{Currency: "IDR", Digits: 1, Value: 1000000}, // 1000000

			// 卢布80-8500
			{Currency: "RUB", Digits: 100, Value: 10000},  // 100
			{Currency: "RUB", Digits: 100, Value: 20000},  // 200
			{Currency: "RUB", Digits: 100, Value: 50000},  // 500
			{Currency: "RUB", Digits: 100, Value: 100000}, // 1000
			{Currency: "RUB", Digits: 100, Value: 200000}, // 2000
			{Currency: "RUB", Digits: 100, Value: 300000}, // 3000

			// 泰达币1-100
			{Currency: "USDT", Digits: 100, Value: 100},   // 1
			{Currency: "USDT", Digits: 100, Value: 200},   // 2
			{Currency: "USDT", Digits: 100, Value: 500},   // 5
			{Currency: "USDT", Digits: 100, Value: 1000},  // 10
			{Currency: "USDT", Digits: 100, Value: 5000},  // 50
			{Currency: "USDT", Digits: 100, Value: 10000}, // 100
		}
		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
