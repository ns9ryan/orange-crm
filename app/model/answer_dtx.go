package model

import (
	"orange-crm/global"
)

// AnswerDtx 开答表-龙虎
type AnswerDtx struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerDtx) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-龙虎'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
