package model

import (
	"orange-crm/global"
	"time"
)

// merchantDayaccountColumns defines and stores column names for the table or_merchant_dayaccount.
type merchantDayaccountColumns struct {
	Id             string // ID
	AccountDate    string // 结算归属日期
	MerchantId     string // 商户账号
	ParentId       string // 上级商户账号
	RootKey        string // 顶级Key
	Currency       string // 货币的代码
	Digits         string // 币种的最小单位倍数
	Bet            string // 当天投注金额，存储为最小单位
	BetEffective   string // 当天有效投注金额，存储为最小单位
	Water          string // 当天洗码金额，存储为最小单位
	WinLost        string // 当天输赢金额，存储为最小单位
	Tip            string // 当天打赏金额，存储为最小单位
	CommissionRate string // 成数
	RollingRate    string // 洗码比
	StarTime       string // 结算区间起始时间
	EndTime        string // 结算区间起始时间
}

// MerchantDayaccountColumns holds the columns for the table or_merchant_dayaccount.
var MerchantDayaccountColumns = merchantDayaccountColumns{
	Id:             "id",
	AccountDate:    "account_date",
	MerchantId:     "merchant_id",
	ParentId:       "parent_id",
	RootKey:        "root_key",
	Currency:       "currency",
	Digits:         "digits",
	Bet:            "bet",
	BetEffective:   "bet_effective",
	Water:          "water",
	WinLost:        "win_lost",
	Tip:            "tip",
	CommissionRate: "commission_rate",
	RollingRate:    "rolling_rate",
	StarTime:       "star_time",
	EndTime:        "end_time",
}

// MerchantDayaccount 商户输赢表结构体
type MerchantDayaccount struct {
	ID uint `gorm:"primarykey;type:int unsigned;comment:ID"`

	AccountDate *time.Time `gorm:"default:NULL;comment:结算归属日期"`

	MerchantId string `gorm:"size:32;not null;comment:商户账号"`
	ParentId   string `gorm:"size:32;not null;comment:上级商户账号"`
	RootKey    uint   `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`

	Currency string `gorm:"size:20;not null;comment:货币的代码"`
	Digits   uint   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`

	Bet          uint `gorm:"type:int unsigned;not null;default:0;comment:当天投注金额，存储为最小单位"`
	BetEffective uint `gorm:"type:int unsigned;not null;default:0;comment:当天有效投注金额，存储为最小单位"`
	Water        uint `gorm:"type:int unsigned;not null;default:0;comment:当天洗码金额，存储为最小单位"`
	WinLost      int  `gorm:"type:int(10) not null;default:0;comment:当天输赢金额，存储为最小单位"`
	Tip          uint `gorm:"type:int unsigned;not null;default:0;comment:当天打赏金额，存储为最小单位"`

	CommissionRate uint `gorm:"type:int unsigned;not null;default:0;comment:成数"`
	RollingRate    uint `gorm:"type:int unsigned;not null;default:0;comment:洗码比"`

	StarTime *time.Time `gorm:"default:NULL;comment:结算区间起始时间"`
	EndTime  *time.Time `gorm:"default:NULL;comment:结算区间起始时间"`
}

// RegisterTable 数据迁移
func (m *MerchantDayaccount) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='商户输赢表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
