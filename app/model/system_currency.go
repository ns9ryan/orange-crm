package model

import (
	"orange-crm/global"
)

// systemCurrencyColumns defines and stores column names for the table or_system_currency.
type systemCurrencyColumns struct {
	Id          string // ID
	Code        string // 字母编码
	Name        string // 名称
	Symbol      string // 符号
	Digits      string // 币种的最小单位倍数
	Number      string // 数字编码
	Region      string // 国家/地区/组织
	MaxTip      string // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip      string // 荷官可接收的最小打赏金额，存储为最小单位
	Description string // 备注
}

// SystemCurrencyColumns holds the columns for the table or_system_currency.
var SystemCurrencyColumns = systemCurrencyColumns{
	Id:          "id",
	Code:        "code",
	Name:        "name",
	Symbol:      "symbol",
	Digits:      "digits",
	Number:      "number",
	Region:      "region",
	MaxTip:      "max_tip",
	MinTip:      "min_tip",
	Description: "description",
}

// SystemCurrency 系统币种表结构体
type SystemCurrency struct {
	ID             uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Code           string `gorm:"type:varchar(10);not null;unique;comment:字母编码"`
	Name           string `gorm:"type:varchar(32);not null;comment:名称"`
	Symbol         string `gorm:"type:varchar(10);not null;comment:符号"`
	Digits         uint   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	Number         string `gorm:"type:varchar(10);not null;comment:数字编码"`
	Region         string `gorm:"size:20;not null;comment:国家/地区/组织"`
	MaxTip         uint   `gorm:"not null;default:0;comment:荷官可接收的最大打赏金额，存储为最小单位"`
	MinTip         uint   `gorm:"not null;default:0;comment:荷官可接收的最小打赏金额，存储为最小单位"`
	Description    string `gorm:"size:20;not null;comment:备注"`
	DefaultChipIds string `gorm:"not null;comment:默认筹码ID列表"`
}

// RegisterTable 数据迁移
func (m *SystemCurrency) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统币种表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		data := []SystemCurrency{
			{ID: 1, Code: "USD", Name: "美元", Symbol: "$", Digits: 100, Number: "840", Region: "美国", MinTip: 100, MaxTip: 10000, Description: "US Dollar", DefaultChipIds: "1,3,4,5,6,7"},
			{ID: 2, Code: "INR", Name: "印度卢比", Symbol: "₹", Digits: 100, Number: "356", Region: "印度", MinTip: 8000, MaxTip: 850000, Description: "Indian Rupee", DefaultChipIds: "5,6,7,8,9,10"},
			{ID: 3, Code: "VND", Name: "越南盾", Symbol: "₫", Digits: 1, Number: "704", Region: "越南", MinTip: 20000, MaxTip: 1000000, Description: "Dong", DefaultChipIds: "13,14,15,16,17,18"},
			{ID: 4, Code: "EUR", Name: "欧元", Symbol: "€", Digits: 100, Number: "978", Region: "欧盟", MinTip: 100, MaxTip: 10000, Description: "Euro", DefaultChipIds: "1,3,4,5,6,7"},
			{ID: 5, Code: "IDR", Name: "印尼盾", Symbol: "Rp", Digits: 1, Number: "360", Region: "印度尼西亚", MinTip: 20000, MaxTip: 1000000, Description: "Rupiah", DefaultChipIds: "13,14,15,16,17,18"},
			{ID: 6, Code: "RUB", Name: "卢布", Symbol: "₽", Digits: 100, Number: "643", Region: "俄罗斯联邦", MinTip: 8000, MaxTip: 850000, Description: "Russian Ruble", DefaultChipIds: "6,7,9,10,11,12"},
			{ID: 7, Code: "USDT", Name: "泰达币", Symbol: "USDT", Digits: 100, Number: "10000", Region: "虚拟货币", MinTip: 100, MaxTip: 10000, Description: "USDT", DefaultChipIds: "1,3,4,5,6,7"},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
