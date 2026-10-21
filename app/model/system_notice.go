package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
	"time"
)

// systemNoticeColumns defines and stores column names for the table or_system_notice.
type systemNoticeColumns struct {
	Id          string // ID
	NoticeId    string // 公告ID
	Title       string // 公告标题
	Content     string // 公告内容
	StartTime   string // 开始时间
	EndTime     string // 结束时间
	IsAll       string // 是否发给所有商户:0是,1否
	OperatorId  string // 发布公告的账号
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// SystemNoticeColumns holds the columns for the table or_system_notice.
var SystemNoticeColumns = systemNoticeColumns{
	Id:          "id",
	NoticeId:    "notice_id",
	Title:       "title",
	Content:     "content",
	StartTime:   "start_time",
	EndTime:     "end_time",
	IsAll:       "is_all",
	OperatorId:  "operator_id",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// SystemNotice  系统公告表结构体
type SystemNotice struct {
	ID         uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	NoticeId   string     `gorm:"uniqueindex;size:36;not null;default:'';comment:公告ID"`
	Title      string     `gorm:"size:100;not null;comment:公告标题"`
	Content    string     `gorm:"type:text not null;comment:公告内容"`
	StartTime  *time.Time `gorm:"type:datetime;not null;comment:开始时间"`
	EndTime    *time.Time `gorm:"default:NULL;comment:结束时间"`
	IsAll      uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否发给所有商户:0是,1否"`
	OperatorId string     `gorm:"size:32;not null;comment:发布公告的账号"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *SystemNotice) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='系统公告表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *SystemNotice) AfterDelete(tx *gorm.DB) (err error) {
	// 删除系统公告与商户关联表记录
	tx.Where(SystemNoticeMerchantColumns.NoticeId, m.NoticeId).Delete(&SystemNoticeMerchant{})
	return
}
