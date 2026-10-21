package validate

type Group struct {
	Test                     Test                     // 测试相关验证器
	Common                   Common                   // 公共相关验证器
	SystemChip               SystemChip               // 筹码相关验证器
	SystemCurrency           SystemCurrency           // 系统币种相关验证器
	BetArea                  BetArea                  // 下注区域相关验证器
	GameCategory             GameCategory             // 游戏分类相关验证器
	Currency                 Currency                 // 币种相关验证器
	Dealer                   Dealer                   // 荷官相关验证器
	ControlAccount           ControlAccount           // 控端账号相关验证器
	GameTable                GameTable                // 游戏桌台相关验证器
	GameAnswer               GameAnswer               // 开答相关验证器
	TemplatePlan             TemplatePlan             // 限红方案相关验证器
	TemplatePlanForm         TemplatePlanForm         // 限红模板相关验证器
	MerchantTemplateBetlimit MerchantTemplateBetlimit // 商户限红模板相关验证器
	LogModule                LogModule                // 日志模块相关验证器
	BetRecord                BetRecord                // 投注明细相关验证器
	Role                     Role                     // 角色相关验证器
	Admin                    Admin                    // 管理员相关验证器
	Menu                     Menu                     // 菜单相关验证器
	Api                      Api                      // api相关验证器
	Merchant                 Merchant                 // 商户相关验证器
	Player                   Player                   // 玩家相关验证器
	PlayerDayaccount         PlayerDayaccount         // 玩家输赢相关验证器
	MerchantDayaccount       MerchantDayaccount       // 商户输赢相关验证器
	SystemNotice             SystemNotice             // 系统公告相关验证器
	DataInit                 DataInit                 // 数据初始化相关验证器
	ServiceNode              ServiceNode              // 服务节点相关验证器
	MerchantNotice           MerchantNotice           // 商户公告相关验证器
	PlayerTransaction        PlayerTransaction        // 会员交易流水相关证器
}

var Validate = new(Group)
