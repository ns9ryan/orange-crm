package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
	"time"
)

// merchantNoticeColumns defines and stores column names for the table or_merchant_notice.
type merchantNoticeColumns struct {
	Id          string // ID
	NoticeId    string // 公告ID
	MerchantId  string // 商户账号
	Title       string // 公告标题
	Content     string // 公告内容
	StartTime   string // 公告开始时间
	EndTime     string // 结束时间
	IsAll       string // 是否发给所有商户:0是,1否
	OperatorId  string // 发布公告的账号
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// MerchantNoticeColumns holds the columns for the table or_merchant_notice.
var MerchantNoticeColumns = merchantNoticeColumns{
	Id:          "id",
	NoticeId:    "notice_id",
	MerchantId:  "merchant_id",
	Title:       "title",
	Content:     "content",
	StartTime:   "start_time",
	EndTime:     "end_time",
	IsAll:       "is_all",
	OperatorId:  "operator_id",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// MerchantNotice  商户公告表结构体
type MerchantNotice struct {
	ID         uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	NoticeId   string     `gorm:"uniqueindex;size:36;not null;default:'';comment:公告ID"`
	MerchantId string     `gorm:"size:32;not null;comment:商户账号"`
	Title      string     `gorm:"size:100;not null;comment:公告标题"`
	Content    string     `gorm:"type:text not null;comment:公告内容"`
	StartTime  *time.Time `gorm:"type:datetime;not null;comment:公告开始时间"`
	EndTime    *time.Time `gorm:"default:NULL;comment:结束时间"`
	IsAll      uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否发给所有商户:0是,1否"`
	OperatorId string     `gorm:"size:32;not null;comment:发布公告的账号"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *MerchantNotice) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='商户公告表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *MerchantNotice) AfterDelete(tx *gorm.DB) (err error) {
	// 删除系统公告与商户关联表记录
	tx.Where(MerchantNoticePlayerColumns.NoticeId, m.NoticeId).Delete(&MerchantNoticePlayer{})
	return
}
