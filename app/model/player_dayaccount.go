package model

import (
	"orange-crm/global"
	"time"
)

// playerDayaccountColumns defines and stores column names for the table or_player_dayaccount.
type playerDayaccountColumns struct {
	Id               string // ID
	AccountDate      string // 结算的归属日
	Username         string // 玩家账号,每个玩家的唯一值
	ParentId         string // 它上级的商户Username
	MerchantId       string // 最上级商户表的username
	RootKey          string // 顶级Key
	Currency         string // 货币的代码
	Digits           string // 币种的最小单位倍数
	LastBet          string // 上次投注金额，存储为最小单位
	LastBetEffective string // 上次有效投注金额，存储为最小单位
	LastWater        string // 上次洗码金额，存储为最小单位
	LastWinLost      string // 上次输赢金额，存储为最小单位
	LastTip          string // 最近一天的打赏金额，存储为最小单位
	StarTime         string // 结算区间起始时间
	EndTime          string // 结算区间终止时间
}

// PlayerDayaccountColumns holds the columns for the table or_player_dayaccount.
var PlayerDayaccountColumns = playerDayaccountColumns{
	Id:               "id",
	AccountDate:      "account_date",
	Username:         "username",
	ParentId:         "parent_id",
	MerchantId:       "merchant_id",
	RootKey:          "root_key",
	Currency:         "currency",
	Digits:           "digits",
	LastBet:          "last_bet",
	LastBetEffective: "last_bet_effective",
	LastWater:        "last_water",
	LastWinLost:      "last_win_lost",
	LastTip:          "last_tip",
	StarTime:         "star_time",
	EndTime:          "end_time",
}

// PlayerDayaccount 玩家输赢表结构体
type PlayerDayaccount struct {
	ID uint `gorm:"primarykey;type:int unsigned;comment:ID"`

	AccountDate *time.Time `gorm:"default:NULL;comment:结算的归属日 "`
	Username    string     `gorm:"size:30;not null;comment:玩家账号,每个玩家的唯一值"`

	ParentId   string `gorm:"size:32;not null;comment:它上级的商户Username"`
	MerchantId string `gorm:"size:32;not null;comment:最上级商户表的username"`
	RootKey    uint   `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`

	Currency string `gorm:"size:20;not null;comment:货币的代码"`

	Digits           uint `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	LastBet          uint `gorm:"type:int unsigned;not null;default:0;comment:上次投注金额，存储为最小单位"`
	LastBetEffective uint `gorm:"type:int unsigned;not null;default:0;comment:上次有效投注金额，存储为最小单位"`
	LastWater        uint `gorm:"type:int unsigned;not null;default:0;comment:上次洗码金额，存储为最小单位"`
	LastWinLost      int  `gorm:"type:int(10) not null;default:0;comment:上次输赢金额，存储为最小单位"`
	LastTip          uint `gorm:"type:int unsigned;not null;default:0;comment:最近一天的打赏金额，存储为最小单位"`

	StarTime *time.Time `gorm:"default:NULL;comment:结算区间起始时间"`
	EndTime  *time.Time `gorm:"default:NULL;comment:结算区间终止时间"`
}

// RegisterTable 数据迁移
func (m *PlayerDayaccount) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='玩家输赢表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
