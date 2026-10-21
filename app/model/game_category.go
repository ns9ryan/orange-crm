// =================================================================================
// 模型层-游戏分类
// =================================================================================

package model

import (
	"orange-crm/global"
)

// gameCategoryColumns defines and stores column names for the table or_game_category.
type gameCategoryColumns struct {
	Id               string // ID
	GameCategoryCode string // 分类
	SubCode          string // 子分类
	Description      string // 描述
	Status           string // 是否启用:0否,1是
	IsHide           string // 是否隐藏:0否,1是
}

// GameCategoryColumns holds the columns for the table or_game_category.
var GameCategoryColumns = gameCategoryColumns{
	Id:               "id",
	GameCategoryCode: "game_category_code",
	SubCode:          "sub_code",
	Description:      "description",
	Status:           "status",
	IsHide:           "is_hide",
}

// GameCategory  游戏分类表结构体
type GameCategory struct {
	ID               uint           `gorm:"primarykey;type:int unsigned;comment:ID"`
	GameCategoryCode string         `gorm:"size:32;not null;comment:分类"`
	SubCode          string         `gorm:"uniqueindex;size:32;not null;comment:子分类"`
	Description      string         `gorm:"size:100;not null;default:'';comment:描述"`
	Status           uint           `gorm:"type:tinyint unsigned;not null;default:1;comment:是否启用:0否,1是"`
	IsHide           uint           `gorm:"type:tinyint unsigned;not null;default:0;comment:是否隐藏:0否,1是"`
	GameTables       []GameTable    `gorm:"foreignKey:CategorySub;references:SubCode"`
	TemplatePlans    []TemplatePlan `gorm:"foreignKey:CategorySub;references:SubCode"`
}

// RegisterTable 数据迁移
func (m *GameCategory) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='游戏分类表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 填充数据
		data := []GameCategory{
			{ID: 1, GameCategoryCode: "BACX", SubCode: "CLASSIC_BACX", Description: "经典百家乐"},
			{ID: 2, GameCategoryCode: "BACX", SubCode: "STREAMER_BACX", Description: "主播百家乐", IsHide: 1},
			{ID: 3, GameCategoryCode: "BACX", SubCode: "SPEED_BACX", Description: "极速百家乐", IsHide: 1},
			{ID: 4, GameCategoryCode: "BACX", SubCode: "COMMUNAL_BACX", Description: "共咪百家乐", IsHide: 1},
			{ID: 5, GameCategoryCode: "BACX", SubCode: "PRIVATE_BACX", Description: "包桌百家乐", IsHide: 1},
			{ID: 6, GameCategoryCode: "BACX", SubCode: "ESPORTS_BACX", Description: "电竞百家乐", IsHide: 1},
			{ID: 7, GameCategoryCode: "DTX", SubCode: "DTX", Description: "龙虎"},
			{ID: 8, GameCategoryCode: "ROT", SubCode: "ROT", Description: "轮盘"},
			{ID: 9, GameCategoryCode: "SICBO", SubCode: "SICBO", Description: "筛宝"},
			{ID: 10, GameCategoryCode: "FTAN", SubCode: "FTAN", Description: "翻滩"},
			{ID: 11, GameCategoryCode: "SDIE", SubCode: "SDIE", Description: "色碟"},
			{ID: 12, GameCategoryCode: "BacXM", SubCode: "BacXM", Description: "百家乐多投"},
			{ID: 13, GameCategoryCode: "FSC", SubCode: "FSC", Description: "鱼虾蟹"},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
