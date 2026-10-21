package model

import (
	"orange-crm/global"
)

// AnswerFsc 开答表-鱼虾蟹
type AnswerFsc struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerFsc) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-鱼虾蟹'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
