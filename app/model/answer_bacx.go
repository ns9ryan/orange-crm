package model

import (
	"orange-crm/global"
)

// AnswerBacx 开答表-百家乐
type AnswerBacx struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerBacx) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-百家乐'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
