// =================================================================================
// 模型层-游戏房间
// =================================================================================

package model

import (
	"orange-crm/global"
)

// gameTableColumns defines and stores column names for the table or_game_table.
type gameTableColumns struct {
	Id               string // ID
	GameTableCode    string // 游戏桌台标识符（如 bacx001）
	GameCategoryCode string // 游戏分类
	CategorySub      string // 游戏子分类
	DeskNo           string // 桌号
	BetInterval      string // 投注间隔时间,秒,0表示不限制
	UnbetRounds      string // 允许未下注局数,局,0表示不限制
	Status           string // 是否启用:0否,1是
	GameServiceName  string // 游戏服务名
	StreamUrl        string // 视频流地址
	Description      string // 游戏说明
	CreatedTime      string // 创建时间
	UpdatedTime      string // 更新时间
}

// GameTableColumns holds the columns for the table or_game_table.
var GameTableColumns = gameTableColumns{
	Id:               "id",
	GameTableCode:    "game_table_code",
	GameCategoryCode: "game_category_code",
	CategorySub:      "category_sub",
	DeskNo:           "desk_no",
	BetInterval:      "bet_interval",
	UnbetRounds:      "unbet_rounds",
	Status:           "status",
	GameServiceName:  "game_service_name",
	StreamUrl:        "stream_url",
	Description:      "description",
	CreatedTime:      "created_time",
	UpdatedTime:      "updated_time",
}

// GameTable  游戏桌台表结构体
type GameTable struct {
	ID               uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	GameTableCode    string `gorm:"uniqueindex;size:32;not null;comment:游戏桌台标识符（如 bacx001）"`
	GameCategoryCode string `gorm:"size:32;not null;comment:游戏分类"`
	CategorySub      string `gorm:"size:32;not null;comment:游戏子分类"`
	DeskNo           uint   `gorm:"type:int unsigned;not null;comment:桌号"`
	BetInterval      uint   `gorm:"type:int unsigned;not null;comment:投注间隔时间,秒"`
	UnbetRounds      uint   `gorm:"type:int unsigned;not null;comment:允许未下注局数,局,0表示不限制"`
	Status           uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否启用:0否,1是"`
	GameServiceName  string `gorm:"size:32;not null;default:'';comment:游戏服务名"`
	StreamUrl        string `gorm:"not null;default:'';comment:视频流地址"`
	Description      string `gorm:"size:100;not null;default:'';comment:游戏说明"`
	BaseModel
	GameCategory GameCategory `gorm:"foreignKey:SubCode;references:CategorySub"`
}

// RegisterTable 数据迁移
func (m *GameTable) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='游戏桌台表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 填充数据
		data := []GameTable{
			{GameTableCode: "BACX001", GameCategoryCode: "BACX", CategorySub: "CLASSIC_BACX", DeskNo: 1, BetInterval: 10, Status: global.CanUse, GameServiceName: "BacXService", StreamUrl: "", Description: "Baccarat 01"},
			{GameTableCode: "BACX002", GameCategoryCode: "BACX", CategorySub: "CLASSIC_BACX", DeskNo: 2, BetInterval: 20, Status: global.CanUse, GameServiceName: "BacXService", StreamUrl: "", Description: "Baccarat 02"},
			{GameTableCode: "BACX003", GameCategoryCode: "BACX", CategorySub: "CLASSIC_BACX", DeskNo: 3, BetInterval: 30, Status: global.CanUse, GameServiceName: "BacXService", StreamUrl: "", Description: "Baccarat 03"},
			{GameTableCode: "DTX001", GameCategoryCode: "DTX", CategorySub: "DTX", DeskNo: 4, BetInterval: 10, Status: global.CanUse, GameServiceName: "LongHuService", StreamUrl: "", Description: "Dragon Tiger 01"},
			{GameTableCode: "ROT", GameCategoryCode: "ROT", CategorySub: "ROT", DeskNo: 5, BetInterval: 30, Status: global.CanUse, GameServiceName: "RotService", StreamUrl: "", Description: "Roulette 01"},
			{GameTableCode: "SICBO", GameCategoryCode: "SICBO", CategorySub: "SICBO", DeskNo: 6, BetInterval: 10, Status: global.CanUse, GameServiceName: "SicboService", StreamUrl: "rtsp://admin:a1234567@192.168.3.112", Description: "Sic Bo 01"},
			{GameTableCode: "FTAN", GameCategoryCode: "FTAN", CategorySub: "FTAN", DeskNo: 7, BetInterval: 10, Status: global.CanUse, GameServiceName: "FtanService", StreamUrl: "", Description: "Fan Tan 01"},
			{GameTableCode: "SDIE", GameCategoryCode: "SDIE", CategorySub: "SDIE", DeskNo: 8, BetInterval: 10, Status: global.CanUse, GameServiceName: "SdieService", StreamUrl: "", Description: "Se Diep 01"},
			{GameTableCode: "FSC", GameCategoryCode: "FSC", CategorySub: "FSC", DeskNo: 9, BetInterval: 10, Status: global.CanUse, GameServiceName: "FscService", StreamUrl: "", Description: "Fish Shrimp Crab 01"},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
