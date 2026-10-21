package model

import (
	"orange-crm/global"
	"time"
)

// loginLogColumns defines and stores column names for the table or_login_log.
type loginLogColumns struct {
	Id          string // ID
	UserType    string // 用户类型:1玩家,2商户
	Username    string // 账号
	LoginTime   string // 登录时间
	LoginIp     string // 登录IP
	IsOk        string // 是否成功:0否,1是
	Reason      string // 失败原因
	RootKey     string // 顶级Key
	MerchantId  string // 对应商户表的username
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// LoginLogColumns holds the columns for the table or_login_log.
var LoginLogColumns = loginLogColumns{
	Id:          "id",
	UserType:    "user_type",
	Username:    "username",
	LoginTime:   "login_time",
	LoginIp:     "login_ip",
	IsOk:        "is_ok",
	Reason:      "reason",
	RootKey:     "root_key",
	MerchantId:  "merchant_id",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// LoginLog 登录日志表结构体
type LoginLog struct {
	ID         uint             `gorm:"primarykey;type:int unsigned;comment:ID"`
	UserType   LoginLogUserType `gorm:"type:tinyint unsigned;not null;comment:用户类型:1玩家,2商户"`
	Username   string           `gorm:"size:30;not null;default:'';comment:账号"`
	LoginTime  *time.Time       `gorm:"default:NULL;comment:登录时间"`
	LoginIp    string           `gorm:"size:50;comment:登录IP"`
	IsOk       uint             `gorm:"type:tinyint unsigned;not null;default:0;comment:是否成功:0否,1是"`
	Reason     string           `gorm:"not null;default:'';comment:失败原因"`
	RootKey    uint             `gorm:"type:int unsigned;default:0;comment:顶级Key"`
	MerchantId string           `gorm:"size:32;comment:对应商户表的username"`
	BaseModel
	Merchant Merchant `gorm:"foreignKey:Username;references:Username"`
	Player   Player   `gorm:"foreignKey:Username;references:Username"`
}

// LoginLogUserType 登录日志表用户类型
type LoginLogUserType uint

const (
	LoginLogPlayer   LoginLogUserType = 1 // 玩家
	LoginLogMerchant LoginLogUserType = 2 // 商户
)

// RegisterTable 数据迁移
func (m *LoginLog) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='登录日志表(商户+会员)'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
