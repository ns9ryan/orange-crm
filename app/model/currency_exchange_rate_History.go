package model

import (
	"orange-crm/global"
	"time"
)

// currencyExchangeRateHistoryColumns defines and stores column names for the table or_currency_exchange_rate_history.
type currencyExchangeRateHistoryColumns struct {
	Id                 string // ID
	FromCode           string // 原币种编码
	ToCode             string // 目标币种编码
	ExchangeRateInt    string // 汇率整数表示,例如 1374
	Digits             string // 汇率倍数,例如 10000
	EffectiveStartTime string // 汇率的生效开始时间
	EffectiveEndTime   string // 汇率生效结束时间
	RateSource         string // 来源:1系统
	OperatorId         string // 操作者ID,当前登录账号的username
	CreatedTime        string // 创建时间
	UpdatedTime        string // 更新时间
}

// CurrencyExchangeRateHistoryColumns holds the columns for the table or_currency_exchange_rate_history.
var CurrencyExchangeRateHistoryColumns = currencyExchangeRateHistoryColumns{
	Id:                 "id",
	FromCode:           "from_code",
	ToCode:             "to_code",
	ExchangeRateInt:    "exchange_rate_int",
	Digits:             "digits",
	EffectiveStartTime: "effective_start_time",
	EffectiveEndTime:   "effective_end_time",
	RateSource:         "rate_source",
	OperatorId:         "operator_id",
	CreatedTime:        "created_time",
	UpdatedTime:        "updated_time",
}

// CurrencyExchangeRateHistory 币种汇率历史表结构体
type CurrencyExchangeRateHistory struct {
	ID                 uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	FromCode           string     `gorm:"type:varchar(10);not null;index:idx_currency_base_currency;comment:原币种编码"`
	ToCode             string     `gorm:"type:varchar(10);not null;index:idx_currency_base_currency;comment:目标币种编码"`
	ExchangeRateInt    uint       `gorm:"type:int;unsigned;not null;comment:汇率整数表示,例如 1374"`
	Digits             uint       `gorm:"type:int;unsigned;not null;comment:汇率倍数,例如 10000"`
	EffectiveStartTime time.Time  `gorm:"comment:汇率的生效开始时间"`
	EffectiveEndTime   *time.Time `gorm:"default:NULL;comment:汇率生效结束时间"`
	RateSource         uint       `gorm:"type:tinyint unsigned;not null;default:1;comment:来源:1系统"`
	OperatorId         string     `gorm:"size:32;not null;comment:操作者ID,当前登录账号的username"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *CurrencyExchangeRateHistory) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='币种汇率历史表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		nowTime := time.Now()
		// 填充数据
		data := []CurrencyExchangeRateHistory{
			{FromCode: "CNY", ToCode: "USD", ExchangeRateInt: 1374, Digits: 10000, EffectiveStartTime: nowTime, OperatorId: ""},     // 人民币对美元 1 CNY = 0.1374 USD
			{FromCode: "VND", ToCode: "USD", ExchangeRateInt: 3958, Digits: 100000000, EffectiveStartTime: nowTime, OperatorId: ""}, // 越南盾对美元 1 VND = 0.00003958 USD
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
