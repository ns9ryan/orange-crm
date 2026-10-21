package model

import (
	"orange-crm/global"
)

// roleApiColumns defines and stores column names for the table or_role_api.
type roleApiColumns struct {
	RoleId string // 角色ID
	ApiId  string // apiID
}

// RoleApiColumns holds the columns for the table or_role_api.
var RoleApiColumns = roleApiColumns{
	RoleId: "role_id",
	ApiId:  "api_id",
}

// RoleApi 角色api表结构体
type RoleApi struct {
	RoleId uint `gorm:"uniqueindex:roleId_apiId;type:int unsigned;not null;comment:角色ID"`
	ApiId  uint `gorm:"uniqueindex:roleId_apiId;type:int unsigned;not null;comment:apiID"`
}

// RegisterTable 数据迁移
func (m *RoleApi) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='角色api表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 获取所有的api
		var apiModels []Api
		global.DB.Find(&apiModels)

		var data []RoleApi
		// 总公司（超管）角色数据
		for _, apiModel := range apiModels {
			data = append(data, RoleApi{RoleId: 1, ApiId: apiModel.ID})
		}

		// 获取商户的api
		var apiMerchantModels []Api
		global.DB.Where("can_assign_merchant = ?", 1).Find(&apiMerchantModels)
		for _, apiModel := range apiMerchantModels {
			data = append(data, RoleApi{RoleId: 2, ApiId: apiModel.ID})
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
