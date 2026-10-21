package model

import (
	"orange-crm/global"
)

// RoleMenu 角色菜单表结构体
type RoleMenu struct {
	RoleId uint `gorm:"uniqueindex:roleId_menuId;type:int unsigned;not null;comment:角色ID"`
	MenuId uint `gorm:"uniqueindex:roleId_menuId;type:int unsigned;not null;comment:菜单ID"`
}

// RegisterTable 数据迁移
func (m *RoleMenu) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='角色菜单表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 获取管理员的菜单
		var menuAdmins []Menu
		global.DB.Where("use_type = 0").Find(&menuAdmins)

		var data []RoleMenu
		// 总公司（超管）角色数据
		for _, menu := range menuAdmins {
			data = append(data, RoleMenu{RoleId: 1, MenuId: menu.ID})
		}

		// 商户菜单
		var menuMerchants []Menu
		global.DB.Where("use_type = 1").Find(&menuMerchants)
		// or（商户1）
		for _, menu := range menuMerchants {
			data = append(data, RoleMenu{RoleId: 2, MenuId: menu.ID})
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
