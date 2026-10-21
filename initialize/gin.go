package initialize

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/middleware"
	"orange-crm/global"
	"orange-crm/router"
	"orange-crm/router/admin"
	"orange-crm/router/merchant"
)

// Gin 初始化gin
func Gin() *gin.Engine {
	Router := gin.New()

	if global.CONFIG.Server.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
		Router = gin.New()
		Router.Use(middleware.GinRecovery(true), middleware.GinLogger(), middleware.GormLogger())
	} else {
		Router.Use(gin.Recovery(), gin.Logger())
	}

	Router.Use(middleware.UseI18n(), middleware.Cors(), middleware.StaticServerWithCORS("/upload", "./upload"))

	// 将项目根目录下的 upload 目录映射为 /upload 路径
	Router.Static("/upload", "./upload")

	// 基础路由分组不需要鉴权
	{
		PublicGroup := Router.Group("")
		test := router.Test{}
		test.Group(PublicGroup) // 测试路由	/

		common := router.Common{}
		common.Group(PublicGroup) // 公共路由-不需要登录和鉴权	/
		//PublicGroup.Use(middleware.OperationRecord())

		dataInit := router.DataInit{}
		dataInit.Group(PublicGroup) // 公共路由-不需要登录和鉴权	/data_init
	}

	// 管理端路由分组-需要带上/admin前缀
	{
		// 需要登录的路由分组
		adminLoginGroup := Router.Group("admin")
		adminLoginGroup.Use(middleware.CheckLogin(), middleware.AdminLogMiddleware())

		// 需要登录并鉴权的路由分组
		adminPrivateGroup := Router.Group("admin")
		adminPrivateGroup.Use(middleware.CheckLogin(), middleware.CasbinMiddleware(), middleware.AdminLogMiddleware())

		admin.Entry.SystemCurrency.Group(adminPrivateGroup)           // 系统币种路由	/admin/system_currency
		admin.Entry.SystemChip.Group(adminPrivateGroup)               // 系统筹码路由	/admin/system_chip
		admin.Entry.BetArea.Group(adminPrivateGroup)                  // 下注区域路由	/admin/bet_area
		admin.Entry.GameCategory.Group(adminPrivateGroup)             // 游戏分组路由	/admin/game_category
		admin.Entry.Dealer.Group(adminPrivateGroup)                   // 荷官路由	/admin/dealer
		admin.Entry.Currency.Group(adminPrivateGroup)                 // 币种路由	/admin/currency
		admin.Entry.GameTable.Group(adminPrivateGroup)                // 游戏桌台路由	/admin/game_table
		admin.Entry.ControlAccount.Group(adminPrivateGroup)           // 控端路由	/admin/control_account
		admin.Entry.GameAnswer.Group(adminPrivateGroup)               // 开答明细路由	/admin/game_answer
		admin.Entry.TemplatePlan.Group(adminPrivateGroup)             // 限红方案路由	/admin/template_plan
		admin.Entry.TemplatePlanForm.Group(adminPrivateGroup)         // 限红方案路由	/admin/template_plan_form
		admin.Entry.MerchantTemplateBetlimit.Group(adminPrivateGroup) // 限红方案路由	/admin/merchant_template_betlimit
		admin.Entry.LogModule.Group(adminPrivateGroup)                // 日志模块路由	/admin/log_module
		admin.Entry.Admin.Group(adminLoginGroup, adminPrivateGroup)   // 管理员路由	/admin/admin
		admin.Entry.Role.Group(adminPrivateGroup)                     // 角色路由	/admin/role
		admin.Entry.Menu.Group(adminPrivateGroup)                     // 菜单路由	/admin/menu
		admin.Entry.Api.Group(adminPrivateGroup)                      // api路由	/admin/api
		admin.Entry.Merchant.Group(adminPrivateGroup)                 // 商户路由	/admin/merchant
		admin.Entry.Player.Group(adminPrivateGroup)                   // 玩家路由	/admin/player
		admin.Entry.BetRecord.Group(adminPrivateGroup)                // 投注明细路由	/admin/bet_record
		admin.Entry.PlayerDayaccount.Group(adminPrivateGroup)         // 玩家输赢报表路由	/admin/player_dayaccount
		admin.Entry.MerchantDayaccount.Group(adminPrivateGroup)       // 商户输赢报表路由	/admin/merchant_dayaccount
		admin.Entry.SystemNotice.Group(adminPrivateGroup)             // 系统公告路由	/admin/system_notice
		admin.Entry.ServiceNode.Group(adminPrivateGroup)              // 服务节点路由	/admin/service_node
		admin.Entry.PlayerTransaction.Group(adminPrivateGroup)        // 会员交易流水路由	/admin/player_transaction
	}

	// 商户路由分组-需要带上/merchant前缀
	{
		// 需要登录的路由分组
		merchantLoginGroup := Router.Group("merchant")
		merchantLoginGroup.Use(middleware.CheckLogin(), middleware.MerchantLogMiddleware())

		// 需要登录并鉴权的路由分组
		merchantPrivateGroup := Router.Group("merchant")
		merchantPrivateGroup.Use(middleware.CheckLogin(), middleware.CasbinMiddleware(), middleware.MerchantLogMiddleware())

		merchant.Entry.SystemCurrency.Group(merchantPrivateGroup)               // 系统币种路由	/merchant/system_currency
		merchant.Entry.SystemChip.Group(merchantPrivateGroup)                   // 系统筹码路由	/merchant/system_chip
		merchant.Entry.BetArea.Group(merchantPrivateGroup)                      // 下注区域路由	/merchant/bet_area
		merchant.Entry.GameCategory.Group(merchantPrivateGroup)                 // 游戏分组路由	/merchant/game_category
		merchant.Entry.Menu.Group(merchantPrivateGroup)                         // 菜单路由	/merchant/menu
		merchant.Entry.Merchant.Group(merchantLoginGroup, merchantPrivateGroup) // 商户路由	/merchant/merchant
		merchant.Entry.MerchantDayaccount.Group(merchantPrivateGroup)           // 商户输赢报表路由	/merchant/merchant_dayaccount
		merchant.Entry.MerchantTemplateBetlimit.Group(merchantPrivateGroup)     // 商户限红模板商户路由	/merchant/merchant_template_betlimit
		merchant.Entry.GameTable.Group(merchantPrivateGroup)                    // 游戏桌台路由	/merchant/game_table
		merchant.Entry.GameAnswer.Group(merchantPrivateGroup)                   // 开答明细路由	/merchant/game_answer
		merchant.Entry.TemplatePlan.Group(merchantPrivateGroup)                 // 限红方案路由	/merchant/template_plan
		merchant.Entry.TemplatePlanForm.Group(merchantPrivateGroup)             // 限红方案路由	/merchant/template_plan_form
		merchant.Entry.LogModule.Group(merchantPrivateGroup)                    // 日志模块路由	/merchant/log_module
		merchant.Entry.Player.Group(merchantPrivateGroup)                       // 玩家路由	/merchant/player
		merchant.Entry.BetRecord.Group(merchantPrivateGroup)                    // 投注明细路由	/merchant/bet_record
		merchant.Entry.PlayerDayaccount.Group(merchantPrivateGroup)             // 玩家输赢报表路由	/merchant/player_dayaccount
		merchant.Entry.SystemNotice.Group(merchantPrivateGroup)                 // 系统消息路由	/merchant/system_notice
		merchant.Entry.MerchantNotice.Group(merchantPrivateGroup)               // 商户公告路由	/merchant/system_notice
		merchant.Entry.PlayerTransaction.Group(merchantPrivateGroup)            // 会员交易流水路由	/merchant/player_transaction
		merchant.Entry.Currency.Group(merchantPrivateGroup)                     // 币种路由	/merchant/currency
	}

	return Router
}
