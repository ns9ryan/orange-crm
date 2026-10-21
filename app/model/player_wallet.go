package model

import (
	"orange-crm/global"
)

// playerWalletColumns defines and stores column names for the table or_player_wallet.
type playerWalletColumns struct {
	Id                 string // ID
	RootKey            string // 顶级Key
	MerchantId         string // 对应商户表的username
	Currency           string // 币种
	MerchantCurrencyId string // 商户币种唯一标识
	Username           string // 玩家账号,每个玩家的唯一值
	Digits             string // 币种的最小单位倍数
	CreditSum          string // 累计充值额度，存储为最小单位
	DebitSum           string // 累计提现额度，存储为最小单位
	Balance            string // 账户余额，存储为最小单位
	Tip                string // 总打赏金额，存储为最小单位
	TotalBet           string // 总投注金额，存储为最小单位
	TotalBetEffective  string // 总有效投注金额，存储为最小单位
	TotalWater         string // 总洗码金额，存储为最小单位
	TotalWinLost       string // 总输赢金额，存储为最小单位
	LastBet            string // 上次投注金额，存储为最小单位
	LastBetEffective   string // 上次有效投注金额，存储为最小单位
	LastWater          string // 上次洗码金额，存储为最小单位
	LastWinLost        string // 上次输赢金额，存储为最小单位
	LastTip            string // 上次打赏金额，存储为最小单位
	CreatedTime        string // 创建时间
	UpdatedTime        string // 更新时间
}

// PlayerWalletColumns holds the columns for the table or_player_wallet.
var PlayerWalletColumns = playerWalletColumns{
	Id:                 "id",
	RootKey:            "root_key",
	MerchantId:         "merchant_id",
	Currency:           "currency",
	MerchantCurrencyId: "merchant_currency_id",
	Username:           "username",
	Digits:             "digits",
	CreditSum:          "credit_sum",
	DebitSum:           "debit_sum",
	Balance:            "balance",
	Tip:                "tip",
	TotalBet:           "total_bet",
	TotalBetEffective:  "total_bet_effective",
	TotalWater:         "total_water",
	TotalWinLost:       "total_win_lost",
	LastBet:            "last_bet",
	LastBetEffective:   "last_bet_effective",
	LastWater:          "last_water",
	LastWinLost:        "last_win_lost",
	LastTip:            "last_tip",
	CreatedTime:        "created_time",
	UpdatedTime:        "updated_time",
}

// PlayerWallet 玩家钱包表结构体
type PlayerWallet struct {
	ID                 uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	RootKey            uint   `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`
	MerchantId         string `gorm:"size:32;not null;comment:对应商户表的username"`
	Currency           string `gorm:"type:varchar(10);not null;comment:币种"`
	MerchantCurrencyId string `gorm:"size:50;not null;comment:商户币种唯一标识"`
	Username           string `gorm:"uniqueindex;size:30;not null;comment:玩家账号,每个玩家的唯一值"`
	Digits             uint   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	CreditSum          uint   `gorm:"not null;default:0;comment:累计充值额度，存储为最小单位"`
	DebitSum           uint   `gorm:"not null;default:0;comment:累计提现额度，存储为最小单位"`
	Balance            uint   `gorm:"not null;default:0;comment:账户余额，存储为最小单位"`
	Tip                uint   `gorm:"not null;default:0;comment:总打赏金额，存储为最小单位"`
	TotalBet           uint   `gorm:"not null;default:0;comment:总投注金额，存储为最小单位"`
	TotalBetEffective  uint   `gorm:"not null;default:0;comment:总有效投注金额，存储为最小单位"`
	TotalWater         uint   `gorm:"not null;default:0;comment:总洗码金额，存储为最小单位"`
	TotalWinLost       int    `gorm:"not null;default:0;comment:总输赢金额，存储为最小单位"`
	LastBet            uint   `gorm:"not null;default:0;comment:上次投注金额，存储为最小单位"`
	LastBetEffective   uint   `gorm:"not null;default:0;comment:上次有效投注金额，存储为最小单位"`
	LastWater          uint   `gorm:"not null;default:0;comment:上次洗码金额，存储为最小单位"`
	LastWinLost        int    `gorm:"not null;default:0;comment:上次输赢金额，存储为最小单位"`
	LastTip            uint   `gorm:"not null;default:0;comment:上次打赏金额，存储为最小单位"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *PlayerWallet) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='玩家钱包表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		//// 填充数据
		//data := []PlayerWallet{
		//	{ID: 1, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player001", Digits: 100, Balance: 10000},
		//	{ID: 2, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player002", Digits: 100, Balance: 10000},
		//	{ID: 3, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player003", Digits: 100, Balance: 10000},
		//	{ID: 4, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player004", Digits: 100, Balance: 10000},
		//	{ID: 5, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player005", Digits: 100, Balance: 10000},
		//	{ID: 6, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player006", Digits: 100, Balance: 10000},
		//	{ID: 7, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player007", Digits: 100, Balance: 10000},
		//	{ID: 8, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player008", Digits: 100, Balance: 10000},
		//	{ID: 9, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player009", Digits: 100, Balance: 10000},
		//	{ID: 10, MerchantId: "merchant001", Currency: "CNY", MerchantCurrencyId: "merchant001_CNY", Username: "player010", Digits: 100, Balance: 10000},
		//}
		//
		//if err := global.DB.Create(&data).Error; err != nil {
		//	return err
		//}
	}
	return nil
}
