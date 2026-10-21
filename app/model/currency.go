package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
)

// currencyColumns defines and stores column names for the table or_currency.
type currencyColumns struct {
	Id         string // ID
	Code       string // 币种编码
	Name       string // 币种名称
	Symbol     string // 币种符号
	Digits     string // 币种的最小单位倍数
	MaxTip     string // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip     string // 荷官可接收的最小打赏金额，存储为最小单位
	OperatorId string // 操作者ID,当前登录账号的username
}

// CurrencyColumns holds the columns for the table or_currency.
var CurrencyColumns = currencyColumns{
	Id:         "id",
	Code:       "code",
	Name:       "name",
	Symbol:     "symbol",
	Digits:     "digits",
	MaxTip:     "max_tip",
	MinTip:     "min_tip",
	OperatorId: "operator_id",
}

// Currency 币种表结构体
type Currency struct {
	ID                       uint                   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Code                     string                 `gorm:"type:varchar(10);not null;unique;comment:币种编码"`
	Name                     string                 `gorm:"type:varchar(32);not null;comment:币种名称"`
	Symbol                   string                 `gorm:"type:varchar(10);not null;comment:币种符号"`
	Digits                   uint                   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	MaxTip                   uint                   `gorm:"not null;default:0;comment:荷官可接收的最大打赏金额，存储为最小单位"`
	MinTip                   uint                   `gorm:"not null;default:0;comment:荷官可接收的最小打赏金额，存储为最小单位"`
	OperatorId               string                 `gorm:"size:32;not null;comment:操作者ID,当前登录账号的username"`
	DefaultChipIds           string                 `gorm:"not null;comment:默认筹码ID列表"`
	CurrencyExchangeRateList []CurrencyExchangeRate `gorm:"foreignKey:FromCode;references:Code"`
	SystemTips               []SystemTip            `gorm:"foreignKey:Currency;references:Code"`
}

// RegisterTable 数据迁移
func (m *Currency) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统币种表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		data := []Currency{
			{ID: 1, Code: "USD", Name: "美元", Symbol: "$", Digits: 100, MinTip: 100, MaxTip: 10000, OperatorId: "admin", DefaultChipIds: "1,3,4,5,6,7"},
			{ID: 2, Code: "VND", Name: "越南盾", Symbol: "₫", Digits: 1, MinTip: 20000, MaxTip: 1000000, OperatorId: "admin", DefaultChipIds: "13,14,15,16,17,18"},
			{ID: 3, Code: "RUB", Name: "卢布", Symbol: "₽", Digits: 100, MinTip: 8000, MaxTip: 850000, OperatorId: "admin", DefaultChipIds: "6,7,9,10,11,12"},
			{ID: 4, Code: "IDR", Name: "印尼盾", Symbol: "Rp", Digits: 1, MinTip: 20000, MaxTip: 1000000, OperatorId: "admin", DefaultChipIds: "13,14,15,16,17,18"},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *Currency) AfterDelete(tx *gorm.DB) (err error) {
	// 删除币种汇率
	tx.Where("from_code = ? || to_code = ?", m.Code, m.Code).Delete(&CurrencyExchangeRate{})

	// 删除币种汇率历史记录
	tx.Where("from_code = ? || to_code = ?", m.Code, m.Code).Delete(&CurrencyExchangeRateHistory{})

	return
}
