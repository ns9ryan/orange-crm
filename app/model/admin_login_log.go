package model

import (
	"orange-crm/global"
	"time"
)

// adminLoginLogColumns defines and stores column names for the table or_admin_login_log.
type adminLoginLogColumns struct {
	Id          string // ID
	Username    string // 账号
	LoginTime   string // 登录时间
	LoginIp     string // 登录IP
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// AdminLoginLogColumns holds the columns for the table or_admin_login_log.
var AdminLoginLogColumns = adminLoginLogColumns{
	Id:          "id",
	Username:    "username",
	LoginTime:   "login_time",
	LoginIp:     "login_ip",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// AdminLoginLog 管理员登录日志表结构体
type AdminLoginLog struct {
	ID        uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	Username  string     `gorm:"size:30;not null;default:'';comment:账号"`
	LoginTime *time.Time `gorm:"default:NULL;comment:登录时间"`
	LoginIp   string     `gorm:"size:50;comment:登录IP"`
	Admin     Admin      `gorm:"foreignKey:Username;references:Username"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *AdminLoginLog) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='管理员登录日志表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
