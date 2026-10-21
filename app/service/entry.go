package service

type Group struct {
	Common                   Common                   // 公共服务
	SystemChip               SystemChip               // 筹码服务
	SystemCurrency           SystemCurrency           // 系统币种服务
	BetArea                  BetArea                  // 下注区域服务
	GameCategory             GameCategory             // 游戏分类服务
	Currency                 Currency                 // 币种服务
	Dealer                   Dealer                   // 荷官服务
	ControlAccount           ControlAccount           // 控端账号服务
	GameTable                GameTable                // 游戏桌台服务
	GameAnswer               GameAnswer               // 开答服务
	TemplatePlan             TemplatePlan             // 限红方案服务
	TemplatePlanForm         TemplatePlanForm         // 限红模板服务
	MerchantTemplateBetlimit MerchantTemplateBetlimit // 商户限红模板服务
	LogModule                LogModule                // 日志模块服务
	BetRecord                BetRecord                // 投注明细服务
	Admin                    Admin                    // 管理员服务
	Role                     Role                     // 角色服务
	Menu                     Menu                     // 菜单服务
	Api                      Api                      // api服务
	Merchant                 Merchant                 // 商户相关验证器
	Player                   Player                   // 玩家服务
	PlayerDayaccount         PlayerDayaccount         // 玩家输赢服务
	MerchantDayaccount       MerchantDayaccount       // 商户输赢服务
	SystemNotice             SystemNotice             // 系统公告服务
	AdminOperationLog        AdminOperationLog        // 管理员操作日志服务
	MerchantOperationLog     MerchantOperationLog     // 管理员操作日志服务
	DataInit                 DataInit                 // 数据初始化服务
	ServiceNode              ServiceNode              // 服务节点服务
	MerchantNotice           MerchantNotice           // 商户公告服务
	PlayerTransaction        PlayerTransaction        // 会员交易流水服务
	GameNotifyLog            GameNotifyLog            // 游戏通知接口请求日志服务
}

var Service = new(Group)
