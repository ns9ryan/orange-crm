package model

import (
	"orange-crm/global"
	"time"
)

// tipRecordColumns defines and stores column names for the table or_tip_record.
type tipRecordColumns struct {
	Id           string // ID
	RecordId     string // 打赏表的ID
	UserName     string // 帐号的ID
	ParentId     string // 上级的ID
	MerchantId   string // 商户ID
	MerchantCode string // 商户码
	RootKey      string // 顶级Key
	GameId       string // 游戏id
	DeskNo       string // 桌号
	Type         string // 游戏类型
	Currency     string // 货币种类
	TipAmount    string // 打赏金额
	Digits       string // 投注额的小数位
	DealerId     string // 打赏给哪个荷官
	TipTime      string // 打赏时间
}

// TipRecordColumns holds the columns for the table or_tip_record.
var TipRecordColumns = tipRecordColumns{
	Id:           "id",
	RecordId:     "record_id",
	UserName:     "user_name",
	ParentId:     "parent_id",
	MerchantId:   "merchant_id",
	MerchantCode: "merchant_code",
	RootKey:      "root_key",
	GameId:       "game_id",
	DeskNo:       "desk_no",
	Type:         "type",
	Currency:     "currency",
	TipAmount:    "tip_amount",
	Digits:       "digits",
	DealerId:     "dealer_id",
	TipTime:      "tip_time",
}

// TipRecord 投注帐目表结构体
type TipRecord struct {
	ID       uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	RecordId string `gorm:"size:32;not null;default:'';comment:打赏表的ID"`

	UserName     string `gorm:"size:32;not null;default:'';comment:帐号的ID"`
	ParentId     string `gorm:"size:32;not null;default:'';comment:上级的ID"`
	MerchantId   string `gorm:"size:32;not null;default:'';comment:商户ID"`
	MerchantCode string `gorm:"size:32;not null;default:'';comment:商户码"`

	RootKey uint `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`

	GameId string `gorm:"size:30;not null;default:'';comment:游戏id"`
	DeskNo int    `gorm:"type:tinyint unsigned;not null;default:0;comment:桌号"`

	Type string `gorm:"size:10;not null;default:'';comment:游戏类型"`

	Currency string `gorm:"size:20;not null;default:'';comment:货币种类"`

	TipAmount int `gorm:"type:int unsigned;not null;default:0;comment:打赏金额"`
	Digits    int `gorm:"type:mediumint unsigned;not null;default:0;comment:投注额的小数位"`

	DealerId string `gorm:"size:30;not null;default:'';comment:打赏给哪个荷官"`

	TipTime *time.Time `gorm:"default:NULL;comment:打赏时间"`
}

// RegisterTable 数据迁移
func (m *TipRecord) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='打赏明细表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
