package model

import (
	"orange-crm/global"
)

// systemNoticeMerchantColumns defines and stores column names for the table or_system_notice_merchant.
type systemNoticeMerchantColumns struct {
	Id          string // ID
	NoticeId    string // 公告ID
	MerchantId  string // 商户ID
	OperatorId  string // 操作的账号
	Read        string // 是否已读:0未读,1已读
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// SystemNoticeMerchantColumns holds the columns for the table or_system_notice_merchant.
var SystemNoticeMerchantColumns = systemNoticeMerchantColumns{
	Id:          "id",
	NoticeId:    "notice_id",
	MerchantId:  "merchant_id",
	OperatorId:  "operator_id",
	Read:        "read",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// SystemNoticeMerchant  系统公告与商户关联表结构体
type SystemNoticeMerchant struct {
	ID         uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	NoticeId   string `gorm:"size:32;not null;comment:公告ID"`
	MerchantId string `gorm:"size:32;not null;comment: 商户ID"`
	OperatorId string `gorm:"size:32;not null;comment:操作的账号"`
	Read       uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否已读:0未读,1已读"`
	BaseModel
	SystemNotice SystemNotice `gorm:"foreignKey:NoticeId;references:NoticeId"`
}

// RegisterTable 数据迁移
func (m *SystemNoticeMerchant) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统公告与商户关联表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
