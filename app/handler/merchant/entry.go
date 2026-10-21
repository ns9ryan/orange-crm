package merchant

type Group struct {
	SystemCurrency           SystemCurrency           // 系统币种
	SystemChip               SystemChip               // 系统筹码
	BetArea                  BetArea                  // 下注区域
	GameCategory             GameCategory             // 游戏分类管理
	Menu                     Menu                     // 菜单
	Merchant                 Merchant                 // 商户
	MerchantDayaccount       MerchantDayaccount       // 商户输赢报表
	MerchantTemplateBetlimit MerchantTemplateBetlimit // 商户限红模板
	GameTable                GameTable                // 游戏桌台
	GameAnswer               GameAnswer               // 游戏开答
	TemplatePlan             TemplatePlan             // 限红方案
	TemplatePlanForm         TemplatePlanForm         // 限红模板
	Player                   Player                   // 玩家管理
	PlayerDayaccount         PlayerDayaccount         // 玩家输赢报表
	BetRecord                BetRecord                // 投注明细列表
	LogModule                LogModule                // 日志模块管理
	SystemNotice             SystemNotice             // 系统公告管理
	MerchantNotice           MerchantNotice           // 商户公告管理
	PlayerTransaction        PlayerTransaction        // 会员交易流水
	Currency                 Currency                 // 币种管理
}

// Entry http请求处理入口
var Entry = new(Group)
