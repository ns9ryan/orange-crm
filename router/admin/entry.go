package admin

// Group 路由分组结构体
type Group struct {
	SystemCurrency           SystemCurrency           // 系统币种路由
	SystemChip               SystemChip               // 系统筹码路由
	BetArea                  BetArea                  // 下注区域路由
	GameCategory             GameCategory             // 游戏分类路由
	Dealer                   Dealer                   // 荷官路由
	ControlAccount           ControlAccount           // 控端账号路由
	Currency                 Currency                 // 币种路由
	GameTable                GameTable                // 游戏桌台路由
	GameAnswer               GameAnswer               // 游戏开答路由
	TemplatePlan             TemplatePlan             // 限红方案路由
	TemplatePlanForm         TemplatePlanForm         // 限红模板路由
	MerchantTemplateBetlimit MerchantTemplateBetlimit // 商户限红模板路由
	LogModule                LogModule                // 日志模块路由
	BetRecord                BetRecord                // 投注明细路由
	Merchant                 Merchant                 // 商户路由
	PlayerDayaccount         PlayerDayaccount         // 玩家输赢报表路由
	MerchantDayaccount       MerchantDayaccount       // 商户输赢报表路由
	Admin                    Admin                    // 管理员账号路由
	Role                     Role                     // 角色管理路由
	Menu                     Menu                     // 菜单管理路由
	Api                      Api                      // api路由
	Player                   Player                   // 玩家路由
	SystemNotice             SystemNotice             // 系统公告路由
	ServiceNode              ServiceNode              // 服务节点路由
	PlayerTransaction        PlayerTransaction        // 会员交易流水路由
}

// Entry 路由分组路口
var Entry = new(Group)
