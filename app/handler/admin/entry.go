package admin

type Group struct {
	SystemCurrency           SystemCurrency           // 系统币种
	SystemChip               SystemChip               // 系统筹码
	BetArea                  BetArea                  // 下注区域
	GameCategory             GameCategory             // 游戏分类管理
	Dealer                   Dealer                   // 荷官管理
	GameAnswer               GameAnswer               // 游戏开答
	GameTable                GameTable                // 游戏桌台管理
	TemplatePlan             TemplatePlan             // 限红方案管理
	TemplatePlanForm         TemplatePlanForm         // 限红模板管理
	MerchantTemplateBetlimit MerchantTemplateBetlimit // 商户限红模板管理
	LogModule                LogModule                // 日志模块管理
	BetRecord                BetRecord                // 投注明细
	ControlAccount           ControlAccount           // 控端账号管理
	Admin                    Admin                    // 管理员管理
	Role                     Role                     // 角色管理
	Menu                     Menu                     // 菜单管理
	Api                      Api                      // api管理
	Currency                 Currency                 // 币种管理
	Merchant                 Merchant                 // 商户管理
	Player                   Player                   // 玩家管理
	PlayerDayaccount         PlayerDayaccount         // 玩家输赢报表
	MerchantDayaccount       MerchantDayaccount       // 商户输赢报表
	SystemNotice             SystemNotice             // 系统公告
	ServiceNode              ServiceNode              // 服务节点
	PlayerTransaction        PlayerTransaction        // 会员交易流水
}

// Entry http请求处理入口
var Entry = new(Group)
