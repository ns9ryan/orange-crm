// =================================================================================
// 模型层-控端账号
// =================================================================================

package model

import (
	"golang.org/x/crypto/bcrypt"
	"orange-crm/global"
	"time"
)

// controlAccountColumns defines and stores column names for the table or_control_account.
type controlAccountColumns struct {
	Id          string // ID
	Username    string // 账号
	Nickname    string // 昵称
	Level       string // 级别
	Status      string // 是否允许登录:0否,1是
	LoginTime   string // 上次登录时间
	LoginIp     string // 登录IP
	Password    string // 密码
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// ControlAccountColumns holds the columns for the table or_control_account.
var ControlAccountColumns = controlAccountColumns{
	Id:          "id",
	Username:    "username",
	Nickname:    "nickname",
	Level:       "level",
	Status:      "status",
	LoginTime:   "login_time",
	LoginIp:     "login_ip",
	Password:    "password",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// ControlAccount 控端账号表结构体
type ControlAccount struct {
	ID        uint                `gorm:"primarykey;type:int unsigned;comment:ID"`
	Username  string              `gorm:"uniqueindex;size:30;not null;default:'';comment:账号"`
	Nickname  string              `gorm:"size:50;not null;default:'';comment:昵称"`
	Level     ControlAccountLevel `gorm:"type:int unsigned;not null;default:1;comment:级别"`
	Status    uint                `gorm:"type:tinyint unsigned;not null;default:0;comment:是否允许登录:0否,1是"`
	LoginTime *time.Time          `gorm:"default:NULL;comment:上次登录时间"`
	LoginIp   string              `gorm:"size:50;comment:登录IP"`
	Password  string              `gorm:"not null;default:'';comment:密码"`
	BaseModel
}

// ControlAccountLevel 控端账号级别
type ControlAccountLevel uint

type ControlAccountLevelStruct struct {
	Name  string              `json:"name"`  // 等级名称
	Value ControlAccountLevel `json:"value"` // 等级对应的值
}

const (
	ControlAccountLevel1 ControlAccountLevel = iota + 1
	ControlAccountLeve2
)

var ControlAccountLevelMap = map[ControlAccountLevel]ControlAccountLevelStruct{
	ControlAccountLevel1: {
		Name:  "等级1",
		Value: ControlAccountLevel1,
	},
	ControlAccountLeve2: {
		Name:  "等级2",
		Value: ControlAccountLeve2,
	},
}

// RegisterTable 数据迁移
func (m *ControlAccount) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='控端账号表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 添加测试控端账号
		password := "123456"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		data := []ControlAccount{
			{Username: "ctl001", Nickname: "测试控端账号001", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl002", Nickname: "测试控端账号002", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl003", Nickname: "测试控端账号003", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl004", Nickname: "测试控端账号004", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl005", Nickname: "测试控端账号005", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl006", Nickname: "测试控端账号006", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl007", Nickname: "测试控端账号007", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
			{Username: "ctl008", Nickname: "测试控端账号008", Status: global.CanLogin, Level: ControlAccountLevel1, Password: string(hash)},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
