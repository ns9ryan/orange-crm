package model

import (
	"orange-crm/global"
)

// AnswerFtan 开答表-番摊
type AnswerFtan struct {
	SGAnswerTable
}

// RegisterTable 数据迁移
func (m *AnswerFtan) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='开答表-番摊'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
