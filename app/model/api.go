package model

import (
	"orange-crm/global"
)

// apiColumns defines and stores column names for the table or_api.
type apiColumns struct {
	Id                string // ID
	Path              string // 路劲
	Method            string // 请求方法
	Group             string // 分组
	HasDbWrite        string // 是否对数据库有新增或编辑操作:0否,1是
	Description       string // api简介
	CanAssignMerchant string // 是否可分配给商户:0否,1是
	IsEditable        string // 是否可编辑:0不可编辑,1可编辑
	IsDeletable       string // 是否可删除:0不可删除,1可删除
	CreatedTime       string // 创建时间
	UpdatedTime       string // 更新时间
}

// ApiColumns holds the columns for the table or_api.
var ApiColumns = apiColumns{
	Id:                "id",
	Path:              "path",
	Method:            "method",
	Group:             "group",
	HasDbWrite:        "has_db_write",
	Description:       "description",
	CanAssignMerchant: "can_assign_merchant",
	IsEditable:        "is_editable",
	IsDeletable:       "is_deletable",
	CreatedTime:       "created_time",
	UpdatedTime:       "updated_time",
}

// Api Api表结构体
type Api struct {
	ID                uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	Path              string `gorm:"not null;comment:路劲"`
	Method            string `gorm:"size:20;not null;comment:请求方法"`
	Group             string `gorm:"not null;comment:分组"`
	HasDbWrite        uint   `gorm:"type:tinyint unsigned;not null;comment:是否对数据库有新增或编辑操作:0否,1是"`
	Description       string `gorm:"not null;comment:api简介"`
	CanAssignMerchant uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否可分配给商户:0否,1是"`
	IsEditable        uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否可编辑:0不可编辑,1可编辑"`
	IsDeletable       uint   `gorm:"type:tinyint unsigned;not null;default:0;comment:是否可删除:0不可删除,1可删除"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *Api) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='Api表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		adminPrefix := "/admin"
		adminData := []Api{
			// 权限管理============================================================================================================================================
			// 菜单
			{Path: adminPrefix + "/menu/list", Method: "POST", Group: "菜单模块", HasDbWrite: 0, Description: "菜单列表"},
			{Path: adminPrefix + "/menu/add", Method: "POST", Group: "菜单模块", HasDbWrite: 1, Description: "创建菜单"},
			{Path: adminPrefix + "/menu/edit", Method: "POST", Group: "菜单模块", HasDbWrite: 1, Description: "编辑菜单"},
			{Path: adminPrefix + "/menu/detail/*", Method: "GET", Group: "菜单模块", HasDbWrite: 0, Description: "菜单详情"},
			{Path: adminPrefix + "/menu/del/*", Method: "POST", Group: "菜单模块", HasDbWrite: 1, Description: "删除菜单"},

			// api
			{Path: adminPrefix + "/api/list", Method: "POST", Group: "api模块", HasDbWrite: 0, Description: "api列表"},
			{Path: adminPrefix + "/api/add", Method: "POST", Group: "api模块", HasDbWrite: 1, Description: "创建api"},
			{Path: adminPrefix + "/api/edit", Method: "POST", Group: "api模块", HasDbWrite: 1, Description: "编辑api"},
			{Path: adminPrefix + "/api/detail/*", Method: "GET", Group: "api模块", HasDbWrite: 0, Description: "api详情"},
			{Path: adminPrefix + "/api/del/*", Method: "POST", Group: "api模块", HasDbWrite: 1, Description: "删除api"},
			{Path: adminPrefix + "/api/getApiGroupBox", Method: "GET", Group: "api模块", HasDbWrite: 0, Description: "获取api分组选择框"},

			// 角色
			{Path: adminPrefix + "/role/list", Method: "POST", Group: "角色模块", HasDbWrite: 0, Description: "角色列表"},
			{Path: adminPrefix + "/role/add", Method: "POST", Group: "角色模块", HasDbWrite: 1, Description: "创建角色"},
			{Path: adminPrefix + "/role/edit", Method: "POST", Group: "角色模块", HasDbWrite: 1, Description: "编辑角色"},
			{Path: adminPrefix + "/role/detail/*", Method: "GET", Group: "角色模块", HasDbWrite: 0, Description: "角色详情"},
			{Path: adminPrefix + "/role/del/*", Method: "POST", Group: "角色模块", HasDbWrite: 1, Description: "删除角色"},
			{Path: adminPrefix + "/role/setRoleMenus", Method: "POST", Group: "角色模块", HasDbWrite: 1, Description: "设置角色菜单"},
			{Path: adminPrefix + "/role/setRoleApis", Method: "POST", Group: "角色模块", HasDbWrite: 1, Description: "设置角色Apis"},
			{Path: adminPrefix + "/role/getMenuListByRole", Method: "POST", Group: "角色模块", HasDbWrite: 0, Description: "根据角色获取菜单列表"},

			// 管理员
			{Path: adminPrefix + "/admin/list", Method: "POST", Group: "管理员模块", HasDbWrite: 0, Description: "管理员列表"},
			{Path: adminPrefix + "/admin/add", Method: "POST", Group: "管理员模块", HasDbWrite: 1, Description: "创建管理员"},
			{Path: adminPrefix + "/admin/edit", Method: "POST", Group: "管理员模块", HasDbWrite: 1, Description: "编辑管理员"},
			{Path: adminPrefix + "/admin/detail/*", Method: "GET", Group: "管理员模块", HasDbWrite: 0, Description: "管理员详情"},
			{Path: adminPrefix + "/admin/del/*", Method: "POST", Group: "管理员模块", HasDbWrite: 1, Description: "删除管理员"},
			{Path: adminPrefix + "/admin/editCurrentPassword", Method: "POST", Group: "管理员模块", HasDbWrite: 1, Description: "编辑当前登录的管理员密码"},
			{Path: adminPrefix + "/admin/resetPassword", Method: "POST", Group: "管理员模块", HasDbWrite: 1, Description: "重置密码"},
			{Path: adminPrefix + "/admin/getAdminInfo", Method: "POST", Group: "管理员模块", HasDbWrite: 0, Description: "获取管理员信息"},
			{Path: adminPrefix + "/admin/logoutAdmin", Method: "POST", Group: "管理员模块", HasDbWrite: 0, Description: "退出登录"},

			// 商户管理============================================================================================================================================
			// 商户
			{Path: adminPrefix + "/merchant/list", Method: "POST", Group: "商户模块", HasDbWrite: 0, Description: "商户列表"},
			{Path: adminPrefix + "/merchant/add", Method: "POST", Group: "商户模块", HasDbWrite: 1, Description: "创建商户"},
			{Path: adminPrefix + "/merchant/addChildren", Method: "POST", Group: "商户模块", HasDbWrite: 1, Description: "创建子商户"},
			{Path: adminPrefix + "/merchant/edit", Method: "POST", Group: "商户模块", HasDbWrite: 1, Description: "编辑商户"},
			{Path: adminPrefix + "/merchant/del/*", Method: "POST", Group: "商户模块", HasDbWrite: 1, Description: "删除商户"},
			{Path: adminPrefix + "/merchant/checkUsername", Method: "POST", Group: "商户模块", HasDbWrite: 0, Description: "检查账号是否可用"},
			{Path: adminPrefix + "/merchant/resetPassword", Method: "POST", Group: "商户模块", HasDbWrite: 1, Description: "重置密码"},
			{Path: adminPrefix + "/merchant/getApiKey/*", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "获取api相关密钥"},
			{Path: adminPrefix + "/merchant/detail/*", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "商户详情"},
			{Path: adminPrefix + "/merchant/detailByMerchantKey/*", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "通过key获取详情"},
			{Path: adminPrefix + "/merchant/regenerateAesKeySalt/*", Method: "GET", Group: "商户模块", HasDbWrite: 1, Description: "重新生成AesKeySalt"},
			{Path: adminPrefix + "/merchant/regenerateMd5Salt/*", Method: "GET", Group: "商户模块", HasDbWrite: 1, Description: "重新生成Md5Salt"},
			{Path: adminPrefix + "/merchant/regenerateTokenKey/*", Method: "GET", Group: "商户模块", HasDbWrite: 1, Description: "重新生成商户tokenKey"},
			{Path: adminPrefix + "/merchant/getBox", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "商户下拉框数据"},
			{Path: adminPrefix + "/merchant/getParentBox", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "主商户下拉框数据"},
			{Path: adminPrefix + "/merchant/getSubBox/*", Method: "GET", Group: "商户模块", HasDbWrite: 0, Description: "子商户下拉框数据"},
			// 商户限红模板
			{Path: adminPrefix + "/merchant_template_betlimit/list", Method: "POST", Group: "商户限红模板", HasDbWrite: 0, Description: "商户限红模板列表"},
			{Path: adminPrefix + "/merchant_template_betlimit/add", Method: "POST", Group: "商户限红模板", HasDbWrite: 1, Description: "创建商户限红模板"},
			{Path: adminPrefix + "/merchant_template_betlimit/edit", Method: "POST", Group: "商户限红模板", HasDbWrite: 1, Description: "编辑商户限红模板"},
			{Path: adminPrefix + "/merchant_template_betlimit/del/*", Method: "POST", Group: "商户限红模板", HasDbWrite: 1, Description: "删除商户限红模板"},
			{Path: adminPrefix + "/merchant_template_betlimit/getTemplateBoxByMerchant", Method: "POST", Group: "商户限红模板", HasDbWrite: 0, Description: "获取商户限红模板下拉框数据"},
			{Path: adminPrefix + "/merchant_template_betlimit/detail/*", Method: "GET", Group: "商户限红模板", HasDbWrite: 0, Description: "商户限红模板详情"},
			// 商户输赢表
			{Path: adminPrefix + "/merchant_dayaccount/list", Method: "POST", Group: "商户限红模板", HasDbWrite: 0, Description: "获取主商户输赢报表"},
			{Path: adminPrefix + "/merchant_dayaccount/subList", Method: "POST", Group: "商户限红模板", HasDbWrite: 0, Description: "获取主商户下子商户的输赢报表"},

			// 玩家管理============================================================================================================================================
			// 玩家
			{Path: adminPrefix + "/player/list", Method: "POST", Group: "玩家模块", HasDbWrite: 0, Description: "玩家列表"},
			{Path: adminPrefix + "/player/add", Method: "POST", Group: "玩家模块", HasDbWrite: 1, Description: "创建玩家"},
			{Path: adminPrefix + "/player/edit", Method: "POST", Group: "玩家模块", HasDbWrite: 1, Description: "编辑玩家"},
			{Path: adminPrefix + "/player/detail/*", Method: "GET", Group: "玩家模块", HasDbWrite: 0, Description: "玩家详情"},
			{Path: adminPrefix + "/player/del/*", Method: "POST", Group: "玩家模块", HasDbWrite: 1, Description: "删除玩家"},
			{Path: adminPrefix + "/player/resetPassword", Method: "POST", Group: "玩家模块", HasDbWrite: 1, Description: "重置密码"},
			// 投注记录
			{Path: adminPrefix + "/bet_record/list", Method: "POST", Group: "投注记录", HasDbWrite: 0, Description: "投注记录"},
			// 玩家输赢表
			{Path: adminPrefix + "/player_dayaccount/list", Method: "POST", Group: "玩家输赢表", HasDbWrite: 0, Description: "玩家输赢表列表"},
			// 上分/下分记录列表
			{Path: adminPrefix + "/player_transaction/creditAndDebitList", Method: "POST", Group: "上分/下分记录列表", HasDbWrite: 0, Description: "上分/下分记录列表"},

			// 游戏管理============================================================================================================================================
			// 游戏桌台
			{Path: adminPrefix + "/game_table/list", Method: "POST", Group: "游戏桌台模块", HasDbWrite: 0, Description: "游戏桌台列表"},
			{Path: adminPrefix + "/game_table/add", Method: "POST", Group: "游戏桌台模块", HasDbWrite: 1, Description: "创建游戏桌台"},
			{Path: adminPrefix + "/game_table/edit", Method: "POST", Group: "游戏桌台模块", HasDbWrite: 1, Description: "编辑游戏桌台"},
			{Path: adminPrefix + "/game_table/del/*", Method: "POST", Group: "游戏桌台模块", HasDbWrite: 1, Description: "删除游戏桌台"},
			{Path: adminPrefix + "/game_table/detail/*", Method: "GET", Group: "游戏桌台模块", HasDbWrite: 0, Description: "游戏桌台详情"},
			{Path: adminPrefix + "/game_table/getBoxWithGroup", Method: "GET", Group: "游戏桌台模块", HasDbWrite: 0, Description: "获取游戏桌台,分组下拉框"},
			{Path: adminPrefix + "/game_table/getBox", Method: "GET", Group: "游戏桌台模块", HasDbWrite: 0, Description: "游戏桌台下拉框"},
			// 开答记录
			{Path: adminPrefix + "/game_answer/list", Method: "POST", Group: "开答记录", HasDbWrite: 0, Description: "开答记录"},
			// 限红方案
			{Path: adminPrefix + "/template_plan/list", Method: "POST", Group: "限红方案", HasDbWrite: 0, Description: "限红方案列表"},
			{Path: adminPrefix + "/template_plan/add", Method: "POST", Group: "限红方案", HasDbWrite: 1, Description: "创建限红方案"},
			{Path: adminPrefix + "/template_plan/edit", Method: "POST", Group: "限红方案", HasDbWrite: 1, Description: "编辑限红方案"},
			{Path: adminPrefix + "/template_plan/del/*", Method: "POST", Group: "限红方案", HasDbWrite: 1, Description: "删除限红方案"},
			{Path: adminPrefix + "/template_plan/getBoxByCategorySub", Method: "POST", Group: "限红方案", HasDbWrite: 0, Description: "根据子分类获取限红方案下拉框数据"},
			{Path: adminPrefix + "/template_plan/detail/*", Method: "GET", Group: "限红方案", HasDbWrite: 0, Description: "限红方案详情"},
			// 限红模板
			{Path: adminPrefix + "/template_plan_form/list", Method: "POST", Group: "限红模板", HasDbWrite: 0, Description: "限红模板列表"},
			{Path: adminPrefix + "/template_plan_form/add", Method: "POST", Group: "限红模板", HasDbWrite: 1, Description: "创建限红模板"},
			{Path: adminPrefix + "/template_plan_form/edit", Method: "POST", Group: "限红模板", HasDbWrite: 1, Description: "编辑限红模板"},
			{Path: adminPrefix + "/template_plan_form/del/*", Method: "POST", Group: "限红模板", HasDbWrite: 1, Description: "删除限红模板"},
			{Path: adminPrefix + "/template_plan_form/getBox", Method: "GET", Group: "限红模板", HasDbWrite: 0, Description: "根据子分类获取限红模板下拉框数据"},
			{Path: adminPrefix + "/template_plan_form/detail/*", Method: "GET", Group: "限红模板", HasDbWrite: 0, Description: "限红模板详情"},

			// 平台管理============================================================================================================================================
			// 币种
			{Path: adminPrefix + "/currency/list", Method: "POST", Group: "币种模块", HasDbWrite: 0, Description: "币种列表"},
			{Path: adminPrefix + "/currency/add", Method: "POST", Group: "币种模块", HasDbWrite: 1, Description: "创建币种"},
			{Path: adminPrefix + "/currency/detail/*", Method: "GET", Group: "币种模块", HasDbWrite: 0, Description: "币种详情"},
			{Path: adminPrefix + "/currency/del/*", Method: "POST", Group: "币种模块", HasDbWrite: 1, Description: "删除币种"},
			{Path: adminPrefix + "/currency/setExchangeRate", Method: "POST", Group: "币种模块", HasDbWrite: 1, Description: "设置币种汇率"},
			{Path: adminPrefix + "/currency/exchangeRateHistoryList", Method: "POST", Group: "币种模块", HasDbWrite: 0, Description: "币种汇率历史记录"},
			// 控端账号
			{Path: adminPrefix + "/control_account/list", Method: "POST", Group: "控端账号模块", HasDbWrite: 0, Description: "控端账号列表"},
			{Path: adminPrefix + "/control_account/add", Method: "POST", Group: "控端账号模块", HasDbWrite: 1, Description: "创建控端账号"},
			{Path: adminPrefix + "/control_account/edit", Method: "POST", Group: "控端账号模块", HasDbWrite: 1, Description: "编辑控端账号"},
			{Path: adminPrefix + "/control_account/detail/*", Method: "GET", Group: "控端账号模块", HasDbWrite: 0, Description: "控端账号详情"},
			{Path: adminPrefix + "/control_account/del/*", Method: "POST", Group: "控端账号模块", HasDbWrite: 1, Description: "删除控端账号"},
			{Path: adminPrefix + "/control_account/resetPassword", Method: "POST", Group: "控端账号模块", HasDbWrite: 1, Description: "重置密码"},
			{Path: adminPrefix + "/control_account/getLevelBox", Method: "GET", Group: "控端账号模块", HasDbWrite: 0, Description: "获取控端账号等级选择框"},
			{Path: adminPrefix + "/control_account/getBox", Method: "GET", Group: "控端账号模块", HasDbWrite: 0, Description: "获取控端账号下拉框数据"},
			// 荷官
			{Path: adminPrefix + "/dealer/list", Method: "POST", Group: "荷官模块", HasDbWrite: 0, Description: "荷官列表"},
			{Path: adminPrefix + "/dealer/add", Method: "POST", Group: "荷官模块", HasDbWrite: 1, Description: "创建荷官"},
			{Path: adminPrefix + "/dealer/edit", Method: "POST", Group: "荷官模块", HasDbWrite: 1, Description: "编辑荷官"},
			{Path: adminPrefix + "/dealer/detail/*", Method: "GET", Group: "荷官模块", HasDbWrite: 0, Description: "荷官详情"},
			{Path: adminPrefix + "/dealer/del/*", Method: "POST", Group: "荷官模块", HasDbWrite: 1, Description: "删除荷官"},
			{Path: adminPrefix + "/dealer/uploadAvatar", Method: "POST", Group: "荷官模块", HasDbWrite: 0, Description: "上传头像"},
			{Path: adminPrefix + "/dealer/generateEmployeeNumber", Method: "GET", Group: "荷官模块", HasDbWrite: 0, Description: "生成工号"},
			// 系统公告
			{Path: adminPrefix + "/system_notice/list", Method: "POST", Group: "系统公告模块", HasDbWrite: 0, Description: "系统公告列表"},
			{Path: adminPrefix + "/system_notice/add", Method: "POST", Group: "系统公告模块", HasDbWrite: 1, Description: "创建系统公告"},
			{Path: adminPrefix + "/system_notice/edit", Method: "POST", Group: "系统公告模块", HasDbWrite: 1, Description: "编辑系统公告"},
			{Path: adminPrefix + "/system_notice/del/*", Method: "POST", Group: "系统公告模块", HasDbWrite: 1, Description: "删除系统公告"},
			{Path: adminPrefix + "/system_notice/detail/*", Method: "GET", Group: "系统公告模块", HasDbWrite: 0, Description: "系统公告详情"},

			// 日志管理============================================================================================================================================
			{Path: adminPrefix + "/log_module/adminLoginLogList", Method: "POST", Group: "日志管理模块", HasDbWrite: 0, Description: "管理员登录日志列表"},
			{Path: adminPrefix + "/log_module/merchantLoginLogList", Method: "POST", Group: "日志管理模块", HasDbWrite: 0, Description: "商户登录日志列表"},
			{Path: adminPrefix + "/log_module/adminOperationLogList", Method: "POST", Group: "日志管理模块", HasDbWrite: 0, Description: "管理员操作日志列表"},
			{Path: adminPrefix + "/log_module/getAdminOperationLogEventBox", Method: "GET", Group: "日志管理模块", HasDbWrite: 0, Description: "管理员操作日志事件下拉框"},
			{Path: adminPrefix + "/log_module/playerLoginLogList", Method: "POST", Group: "日志管理模块", HasDbWrite: 0, Description: "会员登录日志列表"},
			{Path: adminPrefix + "/log_module/controlLoginLogList", Method: "POST", Group: "日志管理模块", HasDbWrite: 0, Description: "控端登录日志列表"},

			// 系统数据============================================================================================================================================
			// 游戏分类
			{Path: adminPrefix + "/game_category/list", Method: "POST", Group: "系统数据", HasDbWrite: 0, Description: "游戏分类列表"},
			{Path: adminPrefix + "/game_category/getBox", Method: "GET", Group: "系统数据", HasDbWrite: 0, Description: "游戏分类选择框"},
			// 下注区域
			{Path: adminPrefix + "/bet_area/list", Method: "POST", Group: "系统数据", HasDbWrite: 0, Description: "下注区域列表"},
			{Path: adminPrefix + "/bet_area/getBoxByCategorySub", Method: "POST", Group: "系统数据", HasDbWrite: 0, Description: "根据子分类获取下注区域下拉框数据"},
			// 系统币种
			{Path: adminPrefix + "/system_currency/list", Method: "POST", Group: "系统数据", HasDbWrite: 0, Description: "系统币种列表"},
			{Path: adminPrefix + "/system_currency/getSystemCurrencyNoUse", Method: "POST", Group: "系统数据", HasDbWrite: 0, Description: "获取未使用的系统币种"},
			// 筹码
			{Path: adminPrefix + "/system_chip/list", Method: "POST", Group: "筹码模块", HasDbWrite: 0, Description: "筹码列表"},
			{Path: adminPrefix + "/system_chip/getBox", Method: "GET", Group: "筹码模块", HasDbWrite: 0, Description: "筹码选择框"},

			// 限红
			{Path: adminPrefix + "/betting_limit/list", Method: "POST", Group: "限红模块", HasDbWrite: 0, Description: "限红列表"},
			{Path: adminPrefix + "/betting_limit/add", Method: "POST", Group: "限红模块", HasDbWrite: 1, Description: "创建限红"},
			{Path: adminPrefix + "/betting_limit/edit", Method: "POST", Group: "限红模块", HasDbWrite: 1, Description: "编辑限红"},
			{Path: adminPrefix + "/betting_limit/detail/*", Method: "GET", Group: "限红模块", HasDbWrite: 0, Description: "限红详情"},
			{Path: adminPrefix + "/betting_limit/del/*", Method: "POST", Group: "限红模块", HasDbWrite: 1, Description: "删除限红"},
		}
		if err := global.DB.Create(&adminData).Error; err != nil {
			return err
		}

		merchantPrefix := "/merchant"
		merchantData := []Api{
			// 商户菜单
			{Path: merchantPrefix + "/menu/getMenuList", Method: "GET", Group: "商户菜单-[商户]", HasDbWrite: 0, Description: "商户菜单", CanAssignMerchant: global.CanAssignMerchant},

			// 商户管理============================================================================================================================================
			// 商户
			{Path: merchantPrefix + "/merchant/list", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "商户列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/addChildren", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "创建子商户", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/edit", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "编辑商户", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/del/*", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "删除子商户", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/checkUsername", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "检查账号是否可用", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/resetPassword", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "重置密码", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/editCurrentPassword", Method: "POST", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "编辑当前登录的商户密码", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/getApiKey/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "获取api相关密钥", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/detail/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "商户详情", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/detailByMerchantKey/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "通过key获取详情", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/regenerateAesKeySalt/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "重新生成AesKeySalt", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/regenerateMd5Salt/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "重新生成Md5Salt", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/regenerateTokenKey/*", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 1, Description: "重新生成商户tokenKey", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant/getSubBox", Method: "GET", Group: "商户模块-[商户]", HasDbWrite: 0, Description: "子商户下拉框数据", CanAssignMerchant: global.CanAssignMerchant},
			// 商户限红模板
			{Path: merchantPrefix + "/merchant_template_betlimit/list", Method: "POST", Group: "商户限红模板-[商户]", HasDbWrite: 0, Description: "商户限红模板列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_template_betlimit/add", Method: "POST", Group: "商户限红模板-[商户]", HasDbWrite: 1, Description: "创建商户限红模板", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_template_betlimit/edit", Method: "POST", Group: "商户限红模板-[商户]", HasDbWrite: 1, Description: "编辑商户限红模板", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_template_betlimit/del/*", Method: "POST", Group: "商户限红模板-[商户]", HasDbWrite: 1, Description: "删除商户限红模板", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_template_betlimit/getTemplateBoxByMerchant", Method: "POST", Group: "商户限红模板-[商户]", HasDbWrite: 0, Description: "获取商户限红模板下拉框数据", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_template_betlimit/detail/*", Method: "GET", Group: "商户限红模板-[商户]", HasDbWrite: 0, Description: "商户限红模板详情", CanAssignMerchant: global.CanAssignMerchant},
			// 商户输赢表
			{Path: merchantPrefix + "/merchant_dayaccount/list", Method: "POST", Group: "商户输赢表-[商户]", HasDbWrite: 0, Description: "获取商户自身的输赢报表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_dayaccount/subList", Method: "POST", Group: "商户输赢表-[商户]", HasDbWrite: 0, Description: "获取主商户下子商户的输赢报表（仅主商户可用）", CanAssignMerchant: global.CanAssignMerchant},

			// 玩家管理============================================================================================================================================
			// 玩家
			{Path: merchantPrefix + "/player/list", Method: "POST", Group: "玩家模块-[商户]", HasDbWrite: 0, Description: "玩家列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/player/add", Method: "POST", Group: "玩家模块-[商户]", HasDbWrite: 1, Description: "创建玩家", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/player/edit", Method: "POST", Group: "玩家模块-[商户]", HasDbWrite: 1, Description: "编辑玩家", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/player/detail/*", Method: "GET", Group: "玩家模块-[商户]", HasDbWrite: 0, Description: "玩家详情", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/player/resetPassword", Method: "POST", Group: "玩家模块-[商户]", HasDbWrite: 1, Description: "重置密码", CanAssignMerchant: global.CanAssignMerchant},
			// 投注记录
			{Path: merchantPrefix + "/bet_record/list", Method: "POST", Group: "投注记录-[商户]", HasDbWrite: 0, Description: "投注记录", CanAssignMerchant: global.CanAssignMerchant},
			// 玩家输赢表
			{Path: merchantPrefix + "/player_dayaccount/list", Method: "POST", Group: "玩家输赢表-[商户]", HasDbWrite: 0, Description: "玩家输赢表列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/player_transaction/creditAndDebitList", Method: "POST", Group: "上分/下分记录列表-[商户]", HasDbWrite: 0, Description: "上分/下分记录列表", CanAssignMerchant: global.CanAssignMerchant},

			// 游戏管理============================================================================================================================================
			// 游戏桌台
			{Path: merchantPrefix + "/game_table/list", Method: "POST", Group: "游戏桌台模块-[商户]", HasDbWrite: 0, Description: "游戏桌台列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/game_table/detail/*", Method: "GET", Group: "游戏桌台模块-[商户]", HasDbWrite: 0, Description: "游戏桌台详情", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/game_table/getBoxWithGroup", Method: "GET", Group: "游戏桌台模块-[商户]", HasDbWrite: 0, Description: "获取游戏桌台,分组下拉框", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/game_table/getBox", Method: "GET", Group: "游戏桌台模块-[商户]", HasDbWrite: 0, Description: "游戏桌台下拉框", CanAssignMerchant: global.CanAssignMerchant},
			// 开答记录
			{Path: merchantPrefix + "/game_answer/list", Method: "POST", Group: "开答记录-[商户]", HasDbWrite: 0, Description: "开答记录", CanAssignMerchant: global.CanAssignMerchant},
			// 限红方案
			{Path: merchantPrefix + "/template_plan/list", Method: "POST", Group: "限红方案-[商户]", HasDbWrite: 0, Description: "限红方案列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/template_plan/getBoxByCategorySub", Method: "POST", Group: "限红方案-[商户]", HasDbWrite: 0, Description: "根据子分类获取限红方案下拉框数据", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/template_plan/detail/*", Method: "GET", Group: "限红方案-[商户]", HasDbWrite: 0, Description: "限红方案详情", CanAssignMerchant: global.CanAssignMerchant},
			// 限红模板
			{Path: merchantPrefix + "/template_plan_form/list", Method: "POST", Group: "限红模板-[商户]", HasDbWrite: 0, Description: "限红模板列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/template_plan_form/getBox", Method: "GET", Group: "限红模板-[商户]", HasDbWrite: 0, Description: "根据子分类获取限红模板下拉框数据", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/template_plan_form/detail/*", Method: "GET", Group: "限红模板-[商户]", HasDbWrite: 0, Description: "限红模板详情", CanAssignMerchant: global.CanAssignMerchant},

			// 日志管理============================================================================================================================================
			{Path: merchantPrefix + "/log_module/merchantLoginLogList", Method: "POST", Group: "日志管理模块-[商户]", HasDbWrite: 0, Description: "商户登录日志列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/log_module/merchantOperationLogList", Method: "POST", Group: "日志管理模块-[商户]", HasDbWrite: 0, Description: "商户操作日志列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/log_module/getMerchantOperationLogEventBox", Method: "GET", Group: "日志管理模块-[商户]", HasDbWrite: 0, Description: "商户操作日志事件下拉框", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/log_module/playerLoginLogList", Method: "POST", Group: "日志管理模块-[商户]", HasDbWrite: 0, Description: "会员登录日志列表", CanAssignMerchant: global.CanAssignMerchant},

			// 系统数据============================================================================================================================================
			// 游戏分类
			{Path: merchantPrefix + "/game_category/list", Method: "POST", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "游戏分类列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/game_category/getBox", Method: "GET", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "游戏分类选择框", CanAssignMerchant: global.CanAssignMerchant},
			// 下注区域
			{Path: merchantPrefix + "/bet_area/list", Method: "POST", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "下注区域列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/bet_area/getBoxByCategorySub", Method: "POST", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "根据子分类获取下注区域下拉框数据", CanAssignMerchant: global.CanAssignMerchant},
			// 系统币种
			{Path: merchantPrefix + "/system_currency/list", Method: "POST", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "系统币种列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/system_currency/getSystemCurrencyNoUse", Method: "POST", Group: "系统数据-[商户]", HasDbWrite: 0, Description: "获取未使用的系统币种", CanAssignMerchant: global.CanAssignMerchant},
			// 筹码
			{Path: merchantPrefix + "/system_chip/list", Method: "POST", Group: "筹码模块-[商户]", HasDbWrite: 0, Description: "筹码列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/system_chip/getBox", Method: "GET", Group: "筹码模块-[商户]", HasDbWrite: 0, Description: "筹码选择框", CanAssignMerchant: global.CanAssignMerchant},

			// 平台管理============================================================================================================================================
			// 系统公告
			{Path: merchantPrefix + "/system_notice/listForNew", Method: "POST", Group: "系统公告模块-[商户]", HasDbWrite: 0, Description: "商户新消息", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/system_notice/listForSelf", Method: "POST", Group: "系统公告模块-[商户]", HasDbWrite: 0, Description: "商户收到的系统公共列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/system_notice/markRead", Method: "POST", Group: "系统公告模块-[商户]", HasDbWrite: 1, Description: "标记已读", CanAssignMerchant: global.CanAssignMerchant},
			// 商户公告
			{Path: merchantPrefix + "/merchant_notice/list", Method: "POST", Group: "商户公告模块-[商户]", HasDbWrite: 0, Description: "商户公告列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_notice/add", Method: "POST", Group: "商户公告模块-[商户]", HasDbWrite: 1, Description: "创建商户公告", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_notice/edit", Method: "POST", Group: "商户公告模块-[商户]", HasDbWrite: 1, Description: "编辑商户公告", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_notice/del/*", Method: "POST", Group: "商户公告模块-[商户]", HasDbWrite: 1, Description: "删除商户公告", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/merchant_notice/detail/*", Method: "GET", Group: "商户公告模块-[商户]", HasDbWrite: 0, Description: "商户公告详情", CanAssignMerchant: global.CanAssignMerchant},
			// 币种
			{Path: merchantPrefix + "/currency/list", Method: "POST", Group: "币种模块-[商户]", HasDbWrite: 0, Description: "币种列表", CanAssignMerchant: global.CanAssignMerchant},
			{Path: merchantPrefix + "/currency/detail/*", Method: "GET", Group: "币种模块-[商户]", HasDbWrite: 0, Description: "币种详情", CanAssignMerchant: global.CanAssignMerchant},
		}
		if err := global.DB.Create(&merchantData).Error; err != nil {
			return err
		}
	}
	return nil
}
