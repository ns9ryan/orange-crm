package model

import (
	"orange-crm/global"
	"time"
)

// controlLoginLogColumns defines and stores column names for the table or_control_login_log.
type controlLoginLogColumns struct {
	Id          string // ID
	Username    string // 账号
	LoginTime   string // 登录时间
	LoginIp     string // 登录IP
	IsOk        string // 是否成功:0否,1是
	Reason      string // 失败原因
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// ControlLoginLogColumns holds the columns for the table or_control_login_log.
var ControlLoginLogColumns = controlLoginLogColumns{
	Id:          "id",
	Username:    "username",
	LoginTime:   "login_time",
	LoginIp:     "login_ip",
	IsOk:        "is_ok",
	Reason:      "reason",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// ControlLoginLog 控端登录日志表结构体
type ControlLoginLog struct {
	ID        uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	Username  string     `gorm:"size:30;not null;default:'';comment:账号"`
	LoginTime *time.Time `gorm:"default:NULL;comment:登录时间"`
	LoginIp   string     `gorm:"size:50;comment:登录IP"`
	IsOk      uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否成功:0否,1是"`
	Reason    string     `gorm:"not null;default:'';comment:失败原因"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *ControlLoginLog) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='控端登录日志表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
