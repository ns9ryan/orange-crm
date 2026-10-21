package model

import (
	"orange-crm/global"
)

// merchantOperationLogColumns defines and stores column names for the table or_merchant_operation_log.
type merchantOperationLogColumns struct {
	Id                         string // ID
	ParentId                   string // 上级账号
	Username                   string // 账号
	Module                     string // 模块名（如：用户管理）
	Action                     string // 操作动作（如：修改、删除）
	Content                    string // 操作内容描述
	IsSuccess                  string // 是否成功:0否,1是
	ErrorMessage               string // 错误信息（失败时记录）
	MerchantOperationEventType string // 事件类型
	Latency                    string // 延迟(毫秒)
	RequestUrl                 string // 请求地址
	RequestMethod              string // 请求方法（GET/POST/...）
	RequestParams              string // 请求参数
	ResponseData               string // 响应数据
	Ip                         string // 请求 IP
	UserAgent                  string // 请求 User-Agent
	StatusCode                 string // HTTP 状态码
	CreatedTime                string // 创建时间
	UpdatedTime                string // 更新时间
}

// MerchantOperationLogColumns holds the columns for the table or_merchant_operation_log.
var MerchantOperationLogColumns = merchantOperationLogColumns{
	Id:                         "id",
	ParentId:                   "parent_id",
	Username:                   "username",
	Module:                     "module",
	Action:                     "action",
	Content:                    "content",
	IsSuccess:                  "is_success",
	ErrorMessage:               "error_message",
	MerchantOperationEventType: "merchant_operation_event_type",
	Latency:                    "latency",
	RequestUrl:                 "request_url",
	RequestMethod:              "request_method",
	RequestParams:              "request_params",
	ResponseData:               "response_data",
	Ip:                         "ip",
	UserAgent:                  "user_agent",
	StatusCode:                 "status_code",
	CreatedTime:                "created_time",
	UpdatedTime:                "updated_time",
}

// MerchantOperationLog 商户操作记录表结构体
type MerchantOperationLog struct {
	ID                         uint                       `gorm:"primarykey;type:int unsigned;comment:ID"`
	ParentId                   string                     `gorm:"size:32;not null;default:'';comment:上级账号"`
	Username                   string                     `gorm:"size:32;not null;default:'';comment:账号"`
	Module                     string                     `gorm:"size:64;not null;default:'';comment:模块名（如：用户管理）"`
	Action                     string                     `gorm:"size:64;not null;default:'';comment:操作动作（如：修改、删除）"`
	Content                    string                     `gorm:"type:text not null;comment:操作内容描述"`
	IsSuccess                  uint                       `gorm:"type:tinyint unsigned;not null;comment:是否成功:0否,1是"`
	ErrorMessage               string                     `gorm:"type:text not null;comment:错误信息（失败时记录）"`
	MerchantOperationEventType MerchantOperationEventType `gorm:"type:int unsigned;not null;comment:事件类型"`

	Latency       uint   `gorm:"type:int unsigned;not null;comment:延迟(毫秒)"`
	RequestUrl    string `gorm:"not null;default:'';comment:请求地址"`
	RequestMethod string `gorm:"size:10;not null;default:'';comment:请求方法（GET/POST/...）"`
	RequestParams string `gorm:"type:text not null;comment:请求参数"`
	ResponseData  string `gorm:"type:text not null;comment:响应数据"`
	Ip            string `gorm:"size:64;not null;default:'';comment:请求 IP"`
	UserAgent     string `gorm:"type:text not null;comment:请求 User-Agent"`
	StatusCode    uint   `gorm:"type:int unsigned;not null;comment:HTTP 状态码"`
	BaseModel
}

// MerchantOperationEventType 商户操作日志事件类型
type MerchantOperationEventType uint

const (
	MerchantOperationMerchant                 MerchantOperationEventType = 1 // 商户模块
	MerchantOperationMerchantTemplateBetlimit MerchantOperationEventType = 2 // 商户限红模板模块
	MerchantOperationSystemNotice             MerchantOperationEventType = 3 // 系统消息
	MerchantOperationMerchantNotice           MerchantOperationEventType = 4 // 商户公告
)

var MerchantOperationLogEventMap = map[MerchantOperationEventType]string{
	MerchantOperationMerchant:                 "商户模块",
	MerchantOperationMerchantTemplateBetlimit: "商户限红模板模块",
	MerchantOperationSystemNotice:             "系统消息模块",
	MerchantOperationMerchantNotice:           "商户公告模块",
}

// RegisterTable 数据迁移
func (m *MerchantOperationLog) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='商户操作记录表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
