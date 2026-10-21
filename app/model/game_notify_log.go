package model

import (
	"orange-crm/global"
)

// GameNotifyLog 游戏通知接口请求日志表结构体
type GameNotifyLog struct {
	ID           uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Method       string `gorm:"size:10;not null;comment:HTTP方法，如GET/POST"`
	Url          string `gorm:"not null;comment:请求路径 /core/createmerchant"`
	Host         string `gorm:"size:100;comment:请求主机名，如 192.168.0.71"`
	Port         uint   `gorm:"type:int unsigned;comment:端口号，如 9002"`
	DurationMs   uint   `gorm:"type:int unsigned;not null;comment:响应耗时，单位毫秒"`
	RequestData  string `gorm:"type:text not null;comment:请求数据"`
	ResponseData string `gorm:"type:text not null;comment:响应数据"`
	Ret          uint   `gorm:"type:int unsigned;comment:返回编码:0失败,1成功"`
	ResponseId   uint   `gorm:"type:int unsigned;comment:哪个请求ID对应的回复"`
	MsgId        uint   `gorm:"type:int unsigned;comment:处理哪条消息的回复"`
	Success      uint   `gorm:"type:int unsigned;comment:处理的结果:0失败,1成功"`
	ErrCode      int    `gorm:"type:int not null;comment:错误码"`
	Remark       string `gorm:"type:text not null;comment:备注"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *GameNotifyLog) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='游戏通知接口请求日志表'").Migrator().CreateTable(m); err != nil {
			return err
		}

	}
	return nil
}
