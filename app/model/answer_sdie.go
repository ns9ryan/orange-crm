package model

import (
	"orange-crm/global"
)

// AnswerSdie 开答表-色碟
type AnswerSdie struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerSdie) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-色碟'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
