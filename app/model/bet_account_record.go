package model

import (
	"encoding/json"
	"orange-crm/global"
	"time"
)

// betAccountRecordColumns defines and stores column names for the table or_bet_account_record.
type betAccountRecordColumns struct {
	Id           string // ID
	RecordId     string // 投注帐目表的ID
	UserId       string // 帐号的ID
	ParentId     string // 上级商户ID
	MerchantId   string // 最顶级商户ID
	MerchantCode string // 商户Code
	RootKey      string // 顶级Key
	GameId       string // 游戏id
	DeskNo       string // 桌号
	Type         string // 游戏类型
	Set          string // 桌号
	Round        string // 局号
	RoundId      string // 游戏局数全局ID
	Currency     string // 货币种类
	BetTime      string // 投注时间
	BetAmount    string // 投注额
	BetCount     string // 投注区域数
	WinLost      string // 胜负金额
	Digits       string // 投注额的小数位
	Winner       string // 开奖结果,游戏种类不同
	Value1       string // 开奖参数1
	Value2       string // 开奖参数2
	Value3       string // 开奖参数3
	AnswerTime   string // 开奖时间
	Answer       string // 开奖的详细内容
	CreatedTime  string // 创建时间
	UpdatedTime  string // 更新时间
}

// BetAccountRecordColumns holds the columns for the table or_bet_account_record.
var BetAccountRecordColumns = betAccountRecordColumns{
	Id:           "id",
	RecordId:     "record_id",
	UserId:       "user_id",
	ParentId:     "parent_id",
	MerchantId:   "merchant_id",
	MerchantCode: "merchant_code",
	RootKey:      "root_key",
	GameId:       "game_id",
	DeskNo:       "desk_no",
	Type:         "type",
	Set:          "set",
	Round:        "round",
	RoundId:      "round_id",
	Currency:     "currency",
	BetTime:      "bet_time",
	BetAmount:    "bet_amount",
	BetCount:     "bet_count",
	WinLost:      "win_lost",
	Digits:       "digits",
	Winner:       "winner",
	Value1:       "value1",
	Value2:       "value2",
	Value3:       "value3",
	AnswerTime:   "answer_time",
	Answer:       "answer",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}

// BetAccountRecord 投注帐目表结构体
type BetAccountRecord struct {
	ID       uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	RecordId string `gorm:"size:32;not null;default:'';comment:投注帐目表的ID"`

	UserId string `gorm:"size:32;not null;default:'';comment:帐号的ID"`

	ParentId     string `gorm:"size:32;not null;default:'';comment:上级商户ID"`
	MerchantId   string `gorm:"size:32;not null;default:'';comment:最顶级商户ID"`
	MerchantCode string `gorm:"size:32;not null;default:'';comment:商户Code"`

	RootKey uint `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`

	GameId string `gorm:"size:30;not null;default:'';comment:游戏id"`
	DeskNo int    `gorm:"type:tinyint unsigned;not null;default:0;comment:桌号"`
	Type   string `gorm:"size:20;not null;default:'';comment:游戏类型"`

	Set   int `gorm:"type:smallint unsigned;not null;default:0;comment:桌号"`
	Round int `gorm:"type:smallint unsigned;not null;default:0;comment:局号"`

	RoundId string `gorm:"size:24;not null;default:'';comment:游戏局数全局ID"`

	Currency string `gorm:"size:20;not null;default:'';comment:货币种类"`

	BetTime *time.Time `gorm:"default:NULL;comment:投注时间"`

	BetAmount int `gorm:"type:int unsigned;not null;default:0;comment:投注额"`
	BetCount  int `gorm:"type:int unsigned;not null;default:0;comment:投注区域数"`
	WinLost   int `gorm:"type:int ;not null;default:0;comment:胜负金额"`
	Digits    int `gorm:"type:mediumint unsigned;not null;default:0;comment:投注额的小数位"`

	Winner int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖结果,游戏种类不同"`

	Value1 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数1"`
	Value2 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数2"`
	Value3 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数3"`

	AnswerTime *time.Time      `gorm:"default:NULL;comment:开奖时间"`
	Answer     json.RawMessage `gorm:"type:json;comment:开奖的详细内容"`

	BaseModel
}

// RegisterTable 数据迁移
func (m *BetAccountRecord) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='投注帐目表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
