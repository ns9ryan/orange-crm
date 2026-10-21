package model

import (
	"fmt"
	"orange-crm/global"
)

// roleColumns defines and stores column names for the table or_role.
type roleColumns struct {
	Id            string // ID
	Pid           string // 父级ID
	DefaultMenuId string // 默认菜单ID
	RoleName      string // 角色名称
	Status        string // 状态:0禁用,1正常
	IsEditable    string // 是否可编辑:0不可编辑,1可编辑
	IsDeletable   string // 是否可删除:0不可删除,1可删除
	CreatedTime   string // 创建时间
	UpdatedTime   string // 更新时间
}

// RoleColumns holds the columns for the table or_role.
var RoleColumns = roleColumns{
	Id:            "id",
	Pid:           "pid",
	DefaultMenuId: "default_menu_id",
	RoleName:      "role_name",
	Status:        "status",
	IsEditable:    "is_editable",
	IsDeletable:   "is_deletable",
	CreatedTime:   "created_time",
	UpdatedTime:   "updated_time",
}

// Role 角色表结构体
type Role struct {
	ID            uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Pid           uint   `gorm:"type:int unsigned;not null;default:0;comment:父级ID"`
	DefaultMenuId uint   `gorm:"type:int unsigned;not null;default:1;comment:默认菜单ID"`
	RoleName      string `gorm:"size:50;not null;comment:角色名称"`
	Status        uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:状态:0禁用,1正常"`
	IsEditable    uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否可编辑:0不可编辑,1可编辑"`
	IsDeletable   uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否可删除:0不可删除,1可删除"`
	BaseModel
	Children    []Role `gorm:"foreignKey:Pid;references:ID"`
	DefaultMenu Menu   `gorm:"foreignKey:DefaultMenuId"`
	RoleMenus   []RoleMenu
	RoleApis    []RoleApi
}

// RegisterTable 数据迁移
func (m *Role) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='角色表'").Migrator().CreateTable(m); err != nil {
			return err
		}
		data := []Role{
			{ID: 1, Pid: 0, RoleName: "总网", Status: global.CanUse},
			{ID: 2, Pid: 0, RoleName: "商户", Status: global.CanUse},
			{ID: 3, Pid: 2, RoleName: "普通商户", Status: global.CanUse},
		}
		fmt.Printf("data:%+v\n", data)
		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
