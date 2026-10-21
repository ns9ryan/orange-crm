package model

import (
	"orange-crm/global"
)

// merchantCurrencyColumns defines and stores column names for the table or_merchant_currency.
type merchantCurrencyColumns struct {
	Id                 string // ID
	Username           string // 商户账号
	Currency           string // 币种
	MerchantCurrencyId string // 唯一标识
	Digits             string // 币种的最小单位倍数
	Status             string // 是否启用:0否,1是
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
	BetPaused          string // 是否暂停投注（不开放）:0否,1是
	CommissionRate     string // 成数
	RollingRate        string // 洗码比
	LimitPlayer        string // 限制会员数量
	MaxTip             string // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip             string // 荷官可接收的最小打赏金额，存储为最小单位
	CreatedTime        string // 创建时间
	UpdatedTime        string // 更新时间
}

// MerchantCurrencyColumns holds the columns for the table or_merchant_currency.
var MerchantCurrencyColumns = merchantCurrencyColumns{
	Id:                 "id",
	Username:           "username",
	Currency:           "currency",
	MerchantCurrencyId: "merchant_currency_id",
	Digits:             "digits",
	Status:             "status",
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
	BetPaused:          "bet_paused",
	CommissionRate:     "commission_rate",
	RollingRate:        "rolling_rate",
	LimitPlayer:        "limit_player",
	MaxTip:             "max_tip",
	MinTip:             "min_tip",
	CreatedTime:        "created_time",
	UpdatedTime:        "updated_time",
}

// MerchantCurrency 商户币种表结构体
type MerchantCurrency struct {
	ID                 uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Username           string `gorm:"size:32;not null;comment:商户账号"`
	Currency           string `gorm:"type:varchar(10);not null;comment:币种"`
	MerchantCurrencyId string `gorm:"uniqueindex;size:50;not null;comment:唯一标识"`
	Digits             uint   `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	Status             uint   `gorm:"type:tinyint unsigned;not null;default:1;comment:是否启用:0否,1是"`
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
	BetPaused          uint   `gorm:"type:tinyint unsigned;not null;default:1;comment:是否暂停投注（不开放）:0否,1是"`
	CommissionRate     uint   `gorm:"type:int unsigned;not null;default:0;comment:成数"`
	RollingRate        uint   `gorm:"type:int unsigned;not null;default:0;comment:洗码比"`
	LimitPlayer        uint   `gorm:"type:int unsigned;not null;comment:限制会员数量"`
	MaxTip             uint   `gorm:"not null;default:0;comment:荷官可接收的最大打赏金额，存储为最小单位"`
	MinTip             uint   `gorm:"not null;default:0;comment:荷官可接收的最小打赏金额，存储为最小单位"`
	TipOptions         string `gorm:"not null;comment:打赏选项金额，以逗号隔开，存储为最小单位"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *MerchantCurrency) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='商户币种表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		//// 填充数据
		//data := []MerchantCurrency{
		//	{ID: 1, Username: "merchant001", Currency: "USD", MerchantCurrencyId: "merchant001_USD", Digits: 100, Status: 1, RollingRate: 1, CommissionRate: 2},
		//	{ID: 2, Username: "sub_merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", Digits: 100, Status: 1, RollingRate: 1, CommissionRate: 2},
		//	{ID: 3, Username: "merchant002", Currency: "USD", MerchantCurrencyId: "merchant002_USD", Digits: 1, Status: 1, RollingRate: 1, CommissionRate: 2},
		//	{ID: 4, Username: "sub_merchant002", Currency: "USD", MerchantCurrencyId: "sub_merchant002_USD", Digits: 100, RollingRate: 1, CommissionRate: 2},
		//}
		//
		//if err := global.DB.Create(&data).Error; err != nil {
		//	return err
		//}

		//// 导入 JSON 文件中的数据
		//if err := utils.ImportAndProcessJSON[MerchantCurrency](global.DB, "./data/init/merchant_currency.json", "MerchantCurrency", nil); err != nil {
		//	return err
		//}
	}
	return nil
}
