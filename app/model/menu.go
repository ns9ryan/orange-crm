package model

import (
	"encoding/json"
	"orange-crm/global"
)

// MenuType 菜单类型
type MenuType uint

var (
	PageMenu MenuType = 1 // 页面
	BtnMenu  MenuType = 2 // 按钮
)

// menuColumns defines and stores column names for the table or_menu.
type menuColumns struct {
	Id          string // ID
	Pid         string // 父级ID
	MenuName    string // 菜单名称
	Status      string // 状态:0禁用,1正常
	Sort        string // 排序,值越小越前面
	MenuType    string // 类型:1页面,2按钮
	Description string // 描述
	Path        string // 跳转路径
	Icon        string // 图标
	LangData    string // 语言
	UseType     string // 用途:0管理员使用,1商户使用
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// MenuColumns holds the columns for the table or_menu.
var MenuColumns = menuColumns{
	Id:          "id",
	Pid:         "pid",
	MenuName:    "menu_name",
	Status:      "status",
	Sort:        "sort",
	MenuType:    "menu_type",
	Description: "description",
	Path:        "path",
	Icon:        "icon",
	LangData:    "lang_data",
	UseType:     "use_type",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// Menu 菜单表结构体
type Menu struct {
	ID          uint            `gorm:"primarykey;type:int unsigned;comment:ID"`
	Pid         uint            `gorm:"type:int unsigned;not null;default:0;comment:父级ID"`
	MenuName    string          `gorm:"size:50;not null;comment:菜单名称"`
	Status      uint            `gorm:"type:tinyint unsigned;not null;default:0;comment:状态:0禁用,1正常"`
	Sort        uint            `gorm:"type:int unsigned;not null;default:1;comment:排序,值越小越前面"`
	MenuType    MenuType        `gorm:"type:tinyint unsigned;not null;comment:类型:1页面,2按钮"`
	Description string          `gorm:"not null;default:'';comment:描述"`
	Path        string          `gorm:"not null;default:'';comment:跳转路径"`
	Icon        string          `gorm:"not null;default:'';comment:图标"`
	LangData    json.RawMessage `gorm:"type:json;not null;comment:语言"`
	UseType     uint            `gorm:"type:tinyint unsigned;not null;default:0;comment:用途:0管理员使用,1商户使用"`
	BaseModel
	Children []Menu `gorm:"foreignKey:Pid;references:ID"`
}

// RegisterTable 数据迁移
func (m *Menu) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='菜单表'").Migrator().CreateTable(m); err != nil {
			return err
		}
		data := []Menu{
			// 仪表盘==========================================================================
			{ID: 1, Pid: 0, MenuName: "仪表盘", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/home", Icon: "DataBoard", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"仪表盘"},{"lang":"en","menu_name":"Dashboard"}]`)},

			// 权限管理==========================================================================
			{ID: 11, Pid: 0, MenuName: "权限管理", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "", Icon: "Collection", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"权限管理"},{"lang":"en","menu_name":"Permission Management"}]`)},
			// 子集
			{ID: 12, Pid: 11, MenuName: "管理员管理", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/superadmin/user", Icon: "User", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"管理员管理"},{"lang":"en","menu_name":"Administrator Management"}]`)},
			{ID: 13, Pid: 11, MenuName: "角色管理", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/superadmin/authority", Icon: "Collection", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"角色管理"},{"lang":"en","menu_name":"Role Management"}]`)},
			{ID: 14, Pid: 11, MenuName: "菜单管理", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/superadmin/menus", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"菜单管理"},{"lang":"en","menu_name":"Menu Management"}]`)},
			{ID: 15, Pid: 11, MenuName: "api管理", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/superadmin/apis", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"api管理"},{"lang":"en","menu_name":"API Management"}]`)},

			// 商户管理==========================================================================
			{ID: 21, Pid: 0, MenuName: "商户管理", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "", Icon: "OfficeBuilding", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户管理"},{"lang":"en","menu_name":"Merchant Management"}]`)},
			// 子集
			{ID: 22, Pid: 21, MenuName: "商户列表", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/merchant/management", Icon: "User", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户列表"},{"lang":"en","menu_name":"Merchant List"}]`)},
			{ID: 23, Pid: 21, MenuName: "商户流水", Status: global.CannotUse, Sort: 2, MenuType: PageMenu, Path: "/merchant/transaction", Icon: "Wallet", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户流水"},{"lang":"en","menu_name":"Merchant turnover"}]`)},
			{ID: 24, Pid: 21, MenuName: "商户限红模板", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/merchant/betLimitTemp", Icon: "SuitcaseLine", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户限红模板"},{"lang":"en","menu_name":"Merchant limited red template"}]`)},
			{ID: 25, Pid: 21, MenuName: "商户输赢报表", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/merchant/dayaccount", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户输赢报表"},{"lang":"en","menu_name":"Merchant Winnings Report"}]`)},

			// 会员管理==========================================================================
			{ID: 31, Pid: 0, MenuName: "会员管理", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "", Icon: "User", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员管理"},{"lang":"en","menu_name":"Member Management"}]`)},
			// 子集
			{ID: 32, Pid: 31, MenuName: "会员列表", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/player/management", Icon: "Monitor", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员列表"},{"lang":"en","menu_name":"Member List"}]`)},
			{ID: 33, Pid: 31, MenuName: "投注记录", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/player/betRecord", Icon: "TrophyBase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"投注记录"},{"lang":"en","menu_name":"Betting records"}]`)},
			{ID: 34, Pid: 31, MenuName: "会员流水", Status: global.CannotUse, Sort: 3, MenuType: PageMenu, Path: "/player/transaction", Icon: "Wallet", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员流水"},{"lang":"en","menu_name":"Member Transactions"}]`)},
			{ID: 35, Pid: 31, MenuName: "会员报表", Status: global.CannotUse, Sort: 4, MenuType: PageMenu, Path: "/player/record", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员报表"},{"lang":"en","menu_name":"Member Report"}]`)},
			{ID: 36, Pid: 31, MenuName: "会员输赢报表", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "/player/dayaccount", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员输赢报表"},{"lang":"en","menu_name":"Member Win/Loss Report"}]`)},
			{ID: 37, Pid: 31, MenuName: "上下分记录", Status: global.CanUse, Sort: 6, MenuType: PageMenu, Path: "/player/creditAndDebit", Icon: "ShoppingCart", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"上分/下分记录"},{"lang":"en","menu_name":"Credit/Debit Records"}]`)},

			// 游戏管理==========================================================================
			{ID: 41, Pid: 0, MenuName: "游戏管理", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "", Icon: "Orange", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"游戏管理"},{"lang":"en","menu_name":"Game Management"}]`)},
			// 子集
			{ID: 42, Pid: 41, MenuName: "桌台列表", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/game/gameTable", Icon: "Dish", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"桌台列表"},{"lang":"en","menu_name":"Table List"}]`)},
			{ID: 43, Pid: 41, MenuName: "开答记录", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/game/gameAnswerRecord", Icon: "TrophyBase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"开答记录"},{"lang":"en","menu_name":"Answer record"}]`)},
			{ID: 44, Pid: 41, MenuName: "游戏限红方案", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/game/betLimitCase", Icon: "Suitcase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"游戏限红方案"},{"lang":"en","menu_name":"Game red limit plan"}]`)},
			{ID: 45, Pid: 41, MenuName: "系统限红模板", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/game/betLimitTemplate", Icon: "SuitcaseLine", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统限红模板"},{"lang":"en","menu_name":"System limited red template"}]`)},

			// 平台管理==========================================================================
			{ID: 51, Pid: 0, MenuName: "平台管理", Status: global.CanUse, Sort: 6, MenuType: PageMenu, Path: "", Icon: "Setting", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"平台管理"},{"lang":"en","menu_name":"Platform Management"}]`)},
			// 子集
			{ID: 52, Pid: 51, MenuName: "币种管理", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/platform/currency", Icon: "Coin", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"币种管理"},{"lang":"en","menu_name":"Currency Management"}]`)},
			{ID: 53, Pid: 51, MenuName: "控端账号管理", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/platform/controlend", Icon: "SetUp", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"控端账号管理"},{"lang":"en","menu_name":"Controller Account Management"}]`)},
			{ID: 54, Pid: 51, MenuName: "荷官管理", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/platform/dealers", Icon: "Microphone", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"荷官管理"},{"lang":"en","menu_name":"Dealer Management"}]`)},
			{ID: 55, Pid: 51, MenuName: "公告管理(系统)", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/platform/merchantNotice", Icon: "Notification", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"公告管理"},{"lang":"en","menu_name":"Announcement Management"}]`)},
			{ID: 56, Pid: 51, MenuName: "服务地址管理", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "/platform/serviceNode", Icon: "Connection", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"服务地址管理"},{"lang":"en","menu_name":"Service address management"}]`)},

			// 日志管理==========================================================================
			{ID: 61, Pid: 0, MenuName: "日志管理", Status: global.CanUse, Sort: 7, MenuType: PageMenu, Path: "", Icon: "Bell", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"日志管理"},{"lang":"en","menu_name":"Log Management"}]`)},
			// 子集
			{ID: 62, Pid: 61, MenuName: "管理员登录日志", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/log/adminLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"管理员登录日志"},{"lang":"en","menu_name":"Administrator login log"}]`)},
			{ID: 63, Pid: 61, MenuName: "控端账号登录日志", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/log/ctrlLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"控端账号登录日志"},{"lang":"en","menu_name":"Control terminal account login log"}]`)},
			{ID: 64, Pid: 61, MenuName: "商户登录日志", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/log/merchantLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户登录日志"},{"lang":"en","menu_name":"Merchant Login Log"}]`)},
			{ID: 65, Pid: 61, MenuName: "会员登录日志", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/log/playerLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员登录日志"},{"lang":"en","menu_name":"Player Login Log"}]`)},
			{ID: 66, Pid: 61, MenuName: "系统操作日志", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "/log/sysOperate", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统操作日志"},{"lang":"en","menu_name":"System operation log"}]`)},

			// 系统数据==========================================================================
			{ID: 71, Pid: 0, MenuName: "系统数据", Status: global.CanUse, Sort: 8, MenuType: PageMenu, Path: "", Icon: "Lock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统数据"},{"lang":"en","menu_name":"System data"}]`)},
			// 子集
			{ID: 72, Pid: 71, MenuName: "游戏分类数据", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/system/gameCategory", Icon: "Warning", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"游戏分类数据"},{"lang":"en","menu_name":"Game classification data"}]`)},
			{ID: 73, Pid: 71, MenuName: "下注区域数据", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/system/betArea", Icon: "TurnOff", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"下注区域数据"},{"lang":"en","menu_name":"Betting area data"}]`)},
			{ID: 74, Pid: 71, MenuName: "全币种数据", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/system/currency", Icon: "Operation", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"全币种数据"},{"lang":"en","menu_name":"All currency data"}]`)},
			{ID: 75, Pid: 71, MenuName: "筹码数据", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/system/chip", Icon: "Notebook", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"筹码数据"},{"lang":"en","menu_name":"Chip data"}]`)},
			{ID: 76, Pid: 71, MenuName: "系统语言", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "/system/sysLanguage", Icon: "Reading", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统语言"},{"lang":"en","menu_name":"System language"}]`)},
			//{ID: 16, Pid: 10, MenuName: "限红管理", Status: global.CanUse, Sort: 6, MenuType: PageMenu, Path: "/system/bet-limit", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"限红管理"},{"lang":"en","menu_name":"Red limit management"}]`)},

			// 测试
			{ID: 100, Pid: 0, MenuName: "演示菜单", Status: global.CanUse, Sort: 9, MenuType: PageMenu, Path: "", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"演示菜单"},{"lang":"en","menu_name":"Demo Menu"}]`)},
			{ID: 101, Pid: 100, MenuName: "演示子菜单", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"演示子菜单"},{"lang":"en","menu_name":"Demo Submenu"}]`)},

			// 以下是商户端菜单
			// 仪表盘==========================================================================
			{ID: 201, Pid: 0, MenuName: "仪表盘（商户端）", Status: global.CanUse, Sort: 10, MenuType: PageMenu, Path: "/home", Icon: "DataBoard", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"仪表盘"},{"lang":"en","menu_name":"Dashboard"}]`), UseType: 1},

			// 商户管理==========================================================================
			{ID: 211, Pid: 0, MenuName: "商户管理（商户端）", Status: global.CanUse, Sort: 11, MenuType: PageMenu, Path: "", Icon: "OfficeBuilding", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户管理"},{"lang":"en","menu_name":"Merchant Management"}]`), UseType: 1},
			{ID: 212, Pid: 211, MenuName: "商户列表", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/merchant/management", Icon: "User", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户列表"},{"lang":"en","menu_name":"Merchant List"}]`), UseType: 1},
			// 子集
			{ID: 213, Pid: 211, MenuName: "商户流水（商户端）", Status: global.CannotUse, Sort: 2, MenuType: PageMenu, Path: "/merchant/transaction", Icon: "Wallet", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户流水"},{"lang":"en","menu_name":"Merchant turnover"}]`), UseType: 1},
			{ID: 214, Pid: 211, MenuName: "商户限红模板（商户端）", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/merchant/betLimitTemp", Icon: "SuitcaseLine", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户限红模板"},{"lang":"en","menu_name":"Merchant limited red template"}]`), UseType: 1},
			{ID: 215, Pid: 211, MenuName: "商户输赢报表", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/merchant/dayaccount", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户输赢报表"},{"lang":"en","menu_name":"Merchant Winnings Report"}]`), UseType: 1},

			// 会员管理==========================================================================
			{ID: 221, Pid: 0, MenuName: "会员管理（商户端）", Status: global.CanUse, Sort: 12, MenuType: PageMenu, Path: "", Icon: "User", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员管理"},{"lang":"en","menu_name":"Member Management"}]`), UseType: 1},
			// 子集
			{ID: 222, Pid: 221, MenuName: "会员列表（商户端）", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/player/management", Icon: "Monitor", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员列表"},{"lang":"en","menu_name":"Member List"}]`), UseType: 1},
			{ID: 223, Pid: 221, MenuName: "投注记录（商户端）", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/player/betRecord", Icon: "TrophyBase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"投注记录"},{"lang":"en","menu_name":"Betting records"}]`), UseType: 1},
			{ID: 224, Pid: 221, MenuName: "会员流水（商户端）", Status: global.CannotUse, Sort: 3, MenuType: PageMenu, Path: "/player/transaction", Icon: "Wallet", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员流水"},{"lang":"en","menu_name":"Member Transactions"}]`), UseType: 1},
			{ID: 225, Pid: 221, MenuName: "会员报表（商户端）", Status: global.CannotUse, Sort: 4, MenuType: PageMenu, Path: "/player/record", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员报表"},{"lang":"en","menu_name":"Member Report"}]`), UseType: 1},
			{ID: 226, Pid: 221, MenuName: "会员输赢报表（商户端）", Status: global.CanUse, Sort: 5, MenuType: PageMenu, Path: "/player/dayaccount", Icon: "Tickets", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员输赢报表"},{"lang":"en","menu_name":"Member Win/Loss Report"}]`), UseType: 1},
			{ID: 227, Pid: 221, MenuName: "上下分记录（商户端）", Status: global.CanUse, Sort: 6, MenuType: PageMenu, Path: "/player/creditAndDebit", Icon: "ShoppingCart", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"上分/下分记录"},{"lang":"en","menu_name":"Credit/Debit Records"}]`), UseType: 1},

			// 游戏管理==========================================================================
			{ID: 231, Pid: 0, MenuName: "游戏管理（商户端）", Status: global.CanUse, Sort: 13, MenuType: PageMenu, Path: "", Icon: "Orange", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"游戏管理"},{"lang":"en","menu_name":"Game Management"}]`), UseType: 1},
			// 子集
			{ID: 232, Pid: 231, MenuName: "桌台列表（商户端）", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/game/gameTable", Icon: "Dish", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"桌台列表"},{"lang":"en","menu_name":"Table List"}]`), UseType: 1},
			{ID: 233, Pid: 231, MenuName: "开答记录（商户端）", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/game/gameAnswerRecord", Icon: "TrophyBase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"开答记录"},{"lang":"en","menu_name":"Answer record"}]`), UseType: 1},
			{ID: 234, Pid: 231, MenuName: "游戏限红方案（商户端）", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/game/betLimitCase", Icon: "Suitcase", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"游戏限红方案"},{"lang":"en","menu_name":"Game red limit plan"}]`), UseType: 1},
			{ID: 235, Pid: 231, MenuName: "系统限红模板（商户端）", Status: global.CanUse, Sort: 4, MenuType: PageMenu, Path: "/game/betLimitTemplate", Icon: "SuitcaseLine", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统限红模板"},{"lang":"en","menu_name":"System limited red template"}]`), UseType: 1},

			// 日志管理==========================================================================
			{ID: 241, Pid: 0, MenuName: "日志管理（商户端）", Status: global.CanUse, Sort: 14, MenuType: PageMenu, Path: "", Icon: "Bell", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"日志管理"},{"lang":"en","menu_name":"Log Management"}]`), UseType: 1},
			// 子集
			{ID: 242, Pid: 241, MenuName: "商户登录日志（商户端）", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/log/merchantLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"商户登录日志"},{"lang":"en","menu_name":"Merchant Login Log"}]`), UseType: 1},
			{ID: 243, Pid: 241, MenuName: "会员登录日志（商户端）", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/log/playerLogin", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"会员登录日志"},{"lang":"en","menu_name":"Player Login Log"}]`), UseType: 1},
			{ID: 244, Pid: 241, MenuName: "系统操作日志（商户端）", Status: global.CanUse, Sort: 3, MenuType: PageMenu, Path: "/log/sysOperate", Icon: "AlarmClock", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统操作日志"},{"lang":"en","menu_name":"System operation log"}]`), UseType: 1},

			// 平台管理==========================================================================
			{ID: 251, Pid: 0, MenuName: "平台管理（商户端）", Status: global.CanUse, Sort: 15, MenuType: PageMenu, Path: "", Icon: "Setting", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"平台管理"},{"lang":"en","menu_name":"Platform Management"}]`), UseType: 1},
			// 子集
			{ID: 252, Pid: 251, MenuName: "公告管理(商户)", Status: global.CanUse, Sort: 1, MenuType: PageMenu, Path: "/platform/gameNotice", Icon: "Notification", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"公告管理"},{"lang":"en","menu_name":"Announcement Management"}]`), UseType: 1},
			{ID: 253, Pid: 251, MenuName: "系统消息", Status: global.CanUse, Sort: 2, MenuType: PageMenu, Path: "/platform/systemNotice", Icon: "Notification", LangData: json.RawMessage(`[{"lang":"zh","menu_name":"系统消息"},{"lang":"en","menu_name":"System Notifications"}]`), UseType: 1},
		}
		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
