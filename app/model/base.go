package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"orange-crm/global"
	"time"
)

// BaseModel 公共model
type BaseModel struct {
	CreatedTime time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedTime time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// BaseDeleteModel 公共model-带软删除
type BaseDeleteModel struct {
	CreatedTime time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedTime time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeletedTime gorm.DeletedAt `gorm:"index;comment:删除时间"`
}

type SGAnswerTable struct {
	ID      uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	GameId  string `gorm:"size:30;not null;default:'';comment:游戏id"`
	DeskNo  int    `gorm:"type:tinyint unsigned;not null;default:0;comment:桌号"`
	Type    string `gorm:"size:20;not null;default:'';comment:游戏类型"`
	Set     int    `gorm:"type:smallint unsigned;not null;default:0;comment:桌号"`
	Round   int    `gorm:"type:smallint unsigned;not null;default:0;comment:桌号"`
	RoundId string `gorm:"size:32;not null;default:'';comment:游戏的局数ID"`
	Win     int    `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖结果"`

	Value1 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数1"`
	Value2 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数2"`
	Value3 int `gorm:"type:tinyint unsigned;not null;default:0;comment:开奖参数3"`

	AnswerTime *time.Time `gorm:"default:NULL;comment:开奖时间"`

	Dealer string          `gorm:"size:30;not null;default:'';comment:开奖的荷管"`
	Cards  json.RawMessage `gorm:"type:json;comment:牌的内容"`

	BaseModel
}

// sGAnswerColumns defines and stores column names for the table or_answer_bacx.
type sGAnswerColumns struct {
	Id          string // ID
	GameId      string // 游戏id
	DeskNo      string // 桌号
	Type        string // 游戏类型
	Set         string // 桌号
	Round       string // 桌号
	RoundId     string // 游戏的局数ID
	Win         string // 开奖结果
	Value1      string // 开奖参数1
	Value2      string // 开奖参数2
	Value3      string // 开奖参数3
	AnswerTime  string // 开奖时间
	Dealer      string // 开奖的荷管
	Cards       string // 牌的内容
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// SGAnswerColumns holds the columns for the table or_answer_bacx.
var SGAnswerColumns = sGAnswerColumns{
	Id:          "id",
	GameId:      "game_id",
	DeskNo:      "desk_no",
	Type:        "type",
	Set:         "set",
	Round:       "round",
	RoundId:     "round_id",
	Win:         "win",
	Value1:      "value1",
	Value2:      "value2",
	Value3:      "value3",
	AnswerTime:  "answer_time",
	Dealer:      "dealer",
	Cards:       "cards",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

type GameAnswerDTO struct {
	ID          uint   // ID
	GameId      string // 游戏id
	DeskNo      int    // 桌号
	Type        string // 游戏类型
	Set         int    // 桌号
	Round       int    // 桌号
	RoundId     string // 游戏的局数ID
	Win         int    // 开奖结果
	Value1      int    // 开奖参数1
	Value2      int    // 开奖参数2
	Value3      int    // 开奖参数3
	Dealer      string // 开奖的荷管
	AnswerTime  string // 开奖时间
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
	Cards       []byte // 牌的内容
}

func (m *SGAnswerTable) ToDTO() GameAnswerDTO {
	return GameAnswerDTO{
		ID:          m.ID,                                          // ID
		GameId:      m.GameId,                                      // 游戏id
		DeskNo:      m.DeskNo,                                      // 桌号
		Type:        m.Type,                                        // 游戏类型
		Set:         m.Set,                                         // 桌号
		Round:       m.Round,                                       // 桌号
		RoundId:     m.RoundId,                                     // 游戏的局数ID
		Win:         m.Win,                                         // 开奖结果
		Value1:      m.Value1,                                      // 开奖参数1
		Value2:      m.Value2,                                      // 开奖参数2
		Value3:      m.Value3,                                      // 开奖参数3
		Dealer:      m.Dealer,                                      // 开奖的荷管
		AnswerTime:  m.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
		CreatedTime: m.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
		UpdatedTime: m.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
		Cards:       m.Cards,                                       // 牌的内容
	}
}

// GetCards 实现接口方法：返回原始 Cards 字节
func (m *SGAnswerTable) GetCards() []byte {
	return m.Cards
}

type IGameAnswerModel interface {
	ToDTO() GameAnswerDTO // 返回基础字段
	GetCards() []byte     // 返回原始 Cards 字节
}
