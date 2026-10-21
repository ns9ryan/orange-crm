package model

import (
	"encoding/json"
	"orange-crm/global"
	"time"
)

// betRecordColumns defines and stores column names for the table or_bet_record.
type betRecordColumns struct {
	Id           string // ID
	RecordId     string // 投注明细表的ID
	UserId       string // 帐号的ID
	MerchantId   string // 商户ID
	MerchantCode string // 商户码
	ParentId     string // 上级的ID
	RootKey      string // 顶级Key
	GameId       string // 游戏id
	DeskNo       string // 桌号
	Type         string // 游戏类型
	SubType      string // 游戏类型二级
	Set          string // 桌号
	Round        string // 局号
	RoundId      string // 游戏局数全局ID
	Currency     string // 货币种类
	BetTime      string // 投注时间
	BetMode      string // 投注的模式
	BetAreaId    string // 投注区域ID
	BetAmount    string // 投注额
	BetEffective string // 有效投注额
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

// BetRecordColumns holds the columns for the table or_bet_record.
var BetRecordColumns = betRecordColumns{
	Id:           "id",
	RecordId:     "record_id",
	UserId:       "user_id",
	MerchantId:   "merchant_id",
	MerchantCode: "merchant_code",
	ParentId:     "parent_id",
	RootKey:      "root_key",
	GameId:       "game_id",
	DeskNo:       "desk_no",
	Type:         "type",
	SubType:      "sub_type",
	Set:          "set",
	Round:        "round",
	RoundId:      "round_id",
	Currency:     "currency",
	BetTime:      "bet_time",
	BetMode:      "bet_mode",
	BetAreaId:    "bet_area_id",
	BetAmount:    "bet_amount",
	BetEffective: "bet_effective",
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

// BetRecord 投注明细表结构体
type BetRecord struct {
	ID           uint            `gorm:"primarykey;type:int unsigned;comment:ID"`
	RecordId     string          `gorm:"size:32;not null;default:'';comment:投注明细表的ID"`
	UserId       string          `gorm:"size:32;not null;default:'';comment:帐号的ID"`
	MerchantId   string          `gorm:"size:32;not null;default:'';comment:商户ID"`
	MerchantCode string          `gorm:"size:32;not null;default:'';comment:商户码"`
	ParentId     string          `gorm:"size:32;not null;default:'';comment:上级的ID"`
	RootKey      uint            `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`
	GameId       string          `gorm:"size:30;not null;default:'';comment:游戏id"`
	DeskNo       int             `gorm:"type:tinyint unsigned;not null;default:0;comment:桌号"`
	Type         string          `gorm:"size:20;not null;default:'';comment:游戏类型"`
	SubType      string          `gorm:"size:20;not null;default:'';comment:游戏类型二级"`
	Set          int             `gorm:"type:smallint unsigned;not null;default:0;comment:桌号"`
	Round        int             `gorm:"type:smallint unsigned;not null;default:0;comment:局号"`
	RoundId      string          `gorm:"size:24;not null;default:'';comment:游戏局数全局ID"`
	Currency     string          `gorm:"size:20;not null;default:'';comment:货币种类"`
	BetTime      *time.Time      `gorm:"default:NULL;comment:投注时间"`
	BetMode      int             `gorm:"type:smallint unsigned;not null;default:0;comment:投注的模式"`
	BetAreaId    int             `gorm:"type:mediumint  unsigned;not null;default:0;comment:投注区域ID"`
	BetAmount    int             `gorm:"type:int unsigned;not null;default:0;comment:投注额"`
	BetEffective int             `gorm:"type:int unsigned;not null;default:0;comment:有效投注额"`
	WinLost      int             `gorm:"type:int ;not null;default:0;comment:胜负金额"`
	Digits       int             `gorm:"type:mediumint unsigned;not null;default:0;comment:投注额的小数位"`
	Winner       int             `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖结果,游戏种类不同"`
	Value1       int             `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数1"`
	Value2       int             `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数2"`
	Value3       int             `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数3"`
	AnswerTime   *time.Time      `gorm:"default:NULL;comment:开奖时间"`
	Answer       json.RawMessage `gorm:"type:json;comment:开奖的详细内容"`
	CreatedTime  time.Time       `gorm:"autoCreateTime;type:datetime(3);comment:创建时间"`
	UpdatedTime  time.Time       `gorm:"autoUpdateTime;type:datetime(3);comment:更新时间"`
	BetArea      BetArea         `gorm:"foreignKey:BetAreaId;references:AreaId"`
	GameCategory GameCategory    `gorm:"foreignKey:SubCode;references:SubType"`
}

// RegisterTable 数据迁移
func (m *BetRecord) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='投注明细表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
