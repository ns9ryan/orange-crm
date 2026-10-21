package model

import (
	"orange-crm/global"
)

// AnswerSicbo 开答表-骰宝
type AnswerSicbo struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerSicbo) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-骰宝'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
