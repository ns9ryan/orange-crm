package model

import (
	"orange-crm/global"
)

// AnswerRot 开答表-轮盘
type AnswerRot struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerRot) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-轮盘'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
