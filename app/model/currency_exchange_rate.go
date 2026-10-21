package model

import (
	"orange-crm/global"
)

// currencyExchangeRateColumns defines and stores column names for the table or_currency_exchange_rate.
type currencyExchangeRateColumns struct {
	Id              string // ID
	FromCode        string // 原币种编码
	ToCode          string // 目标币种编码
	ExchangeRateInt string // 汇率整数表示,例如 1374
	Digits          string // 汇率倍数,例如 10000
	OperatorId      string // 操作者ID,当前登录账号的username
	CreatedTime     string // 创建时间
	UpdatedTime     string // 更新时间
}

// CurrencyExchangeRateColumns holds the columns for the table or_currency_exchange_rate.
var CurrencyExchangeRateColumns = currencyExchangeRateColumns{
	Id:              "id",
	FromCode:        "from_code",
	ToCode:          "to_code",
	ExchangeRateInt: "exchange_rate_int",
	Digits:          "digits",
	OperatorId:      "operator_id",
	CreatedTime:     "created_time",
	UpdatedTime:     "updated_time",
}

// CurrencyExchangeRate 币种汇率表结构体
type CurrencyExchangeRate struct {
	ID              uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	FromCode        string `gorm:"type:varchar(10);not null;index:idx_currency_base_currency;comment:原币种编码"`
	ToCode          string `gorm:"type:varchar(10);not null;index:idx_currency_base_currency;comment:目标币种编码"`
	ExchangeRateInt uint   `gorm:"type:int;unsigned;not null;comment:汇率整数表示,例如 1374"`
	Digits          uint   `gorm:"type:int;unsigned;not null;comment:汇率倍数,例如 10000"`
	OperatorId      string `gorm:"size:32;not null;comment:操作者ID,当前登录账号的username"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *CurrencyExchangeRate) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='币种汇率表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 填充数据
		data := []CurrencyExchangeRate{
			{ID: 1, FromCode: "CNY", ToCode: "USD", ExchangeRateInt: 1374, Digits: 10000, OperatorId: ""},     // 人民币对美元 1 CNY = 0.1374 USD
			{ID: 2, FromCode: "VND", ToCode: "USD", ExchangeRateInt: 3958, Digits: 100000000, OperatorId: ""}, // 越南盾对美元 1 VND = 0.00003958 USD
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

//{BaseModel: BaseModel{ID: 1}, FromCurrencyCode: "CNY", ToCurrencyCode: "USD", ExchangeRateInt: 1374, Digits: 10000},     // 人民币对美元 1 CNY = 0.1374 USD
//{BaseModel: BaseModel{ID: 2}, FromCurrencyCode: "KRW", ToCurrencyCode: "USD", ExchangeRateInt: 7, Digits: 10000},        // 韩元对美元 1 KRW = 0.0007 USD
//{BaseModel: BaseModel{ID: 3}, FromCurrencyCode: "MYR", ToCurrencyCode: "USD", ExchangeRateInt: 2234, Digits: 10000},     // 马来西亚币对美元 1 MYR = 0.2234 USD
//{BaseModel: BaseModel{ID: 4}, FromCurrencyCode: "USD", ToCurrencyCode: "CNY", ExchangeRateInt: 7278, Digits: 1000},      // 美元对人民币 1 USD = 7.278 CNY
//{BaseModel: BaseModel{ID: 5}, FromCurrencyCode: "JPY", ToCurrencyCode: "USD", ExchangeRateInt: 64, Digits: 10000},       // 日元对美元 1 JPY = 0.0064 USD
//{BaseModel: BaseModel{ID: 6}, FromCurrencyCode: "THB", ToCurrencyCode: "USD", ExchangeRateInt: 294, Digits: 10000},      // 泰铢对美元 1 THB = 0.0294 USD
//{BaseModel: BaseModel{ID: 7}, FromCurrencyCode: "BTC", ToCurrencyCode: "USD", ExchangeRateInt: 10219764, Digits: 100},   // 比特币对美元 1 BTC = 102197.64 USD
//{BaseModel: BaseModel{ID: 8}, FromCurrencyCode: "VND", ToCurrencyCode: "USD", ExchangeRateInt: 3958, Digits: 100000000}, // 越南盾对美元 1 VND = 0.00003958 USD
//{BaseModel: BaseModel{ID: 9}, FromCurrencyCode: "EUR", ToCurrencyCode: "USD", ExchangeRateInt: 1037, Digits: 1000},      // 欧元对美元 1 EUR = 1.037 USD
//{BaseModel: BaseModel{ID: 10}, FromCurrencyCode: "HKD", ToCurrencyCode: "USD", ExchangeRateInt: 1285, Digits: 10000},    // 港币对美元 1 HKD = 0.1285 USD
//{BaseModel: BaseModel{ID: 11}, FromCurrencyCode: "AUD", ToCurrencyCode: "USD", ExchangeRateInt: 6235, Digits: 10000},    // 澳元对美元 1 AUD = 0.6235 USD
//{BaseModel: BaseModel{ID: 12}, FromCurrencyCode: "TWD", ToCurrencyCode: "USD", ExchangeRateInt: 305, Digits: 10000},     // 台币对美元 1 TWD = 0.0305 USD
//{BaseModel: BaseModel{ID: 13}, FromCurrencyCode: "IDR", ToCurrencyCode: "USD", ExchangeRateInt: 1, Digits: 10000},       // 印尼盾对美元 1 IDR = 0.0001 USD
//{BaseModel: BaseModel{ID: 14}, FromCurrencyCode: "INR", ToCurrencyCode: "USD", ExchangeRateInt: 116, Digits: 10000},     // 印度卢比对美元 1 INR = 0.0116 USD
//{BaseModel: BaseModel{ID: 15}, FromCurrencyCode: "BND", ToCurrencyCode: "USD", ExchangeRateInt: 7361, Digits: 10000},    // 文莱币对美元 1 BND = 0.7361 USD
