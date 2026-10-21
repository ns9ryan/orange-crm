package model

type Group struct {
	// 系统数据
	Lang           Lang           // 语言表
	SystemChip     SystemChip     // 系统筹码表
	SystemCurrency SystemCurrency // 系统币种表
	BetArea        BetArea        // 下注区域表
	GameCategory   GameCategory   // 游戏分类表
	SystemTip      SystemTip      // 系统打赏配置表

	// 平台管理
	Currency                    Currency                    // 币种表
	CurrencyExchangeRate        CurrencyExchangeRate        // 币种汇率表
	CurrencyExchangeRateHistory CurrencyExchangeRateHistory // 币种汇率历史表
	Dealer                      Dealer                      // 荷官表
	ControlAccount              ControlAccount              // 控端账号表
	ControlLoginLog             ControlLoginLog             // 控端登录日志表
	SystemNotice                SystemNotice                // 系统公告表
	SystemNoticeMerchant        SystemNoticeMerchant        // 系统公告与商户关联表
	MerchantNotice              MerchantNotice              // 商户公告表
	MerchantNoticePlayer        MerchantNoticePlayer        // 商户公告与玩家关联表
	ServiceNode                 ServiceNode                 // 服务节点池表

	// 日志管理
	AdminLoginLog        AdminLoginLog        // 管理员登录日志表
	AdminOperationLog    AdminOperationLog    // 管理员操作记录表
	MerchantOperationLog MerchantOperationLog // 商户操作记录表
	LoginLog             LoginLog             // 商户+会员登录日志表

	// 游戏管理
	GameTable            GameTable            // 游戏桌台表
	BetRecord            BetRecord            // 投注明细表
	AnswerBacx           AnswerBacx           // 开答表-百家乐
	AnswerDtx            AnswerDtx            // 开答表-龙虎
	AnswerFtan           AnswerFtan           // 开答表-番摊
	AnswerRot            AnswerRot            // 开答表-轮盘
	AnswerSicbo          AnswerSicbo          // 开答表-骰宝
	AnswerSdie           AnswerSdie           // 开答表-色碟
	TemplatePlan         TemplatePlan         // 限红方案表
	TemplatePlanArea     TemplatePlanArea     // 限红方案投注区域详情表
	TemplatePlanForm     TemplatePlanForm     // 限红模板表
	TemplatePlanFormArea TemplatePlanFormArea // 限红模板投注区域表

	// 商户管理
	Merchant                     Merchant                     // 商户表
	MerchantCurrency             MerchantCurrency             // 商户币种表
	MerchantTemplateBetlimit     MerchantTemplateBetlimit     // 商户限红模板表
	MerchantTemplateBetlimitArea MerchantTemplateBetlimitArea // 商户限红模板投注区域表
	MerchantDayaccount           MerchantDayaccount           // 商户输赢表

	// 玩家管理
	Player            Player            // 玩家表
	PlayerWallet      PlayerWallet      // 玩家钱包表
	PlayerTransaction PlayerTransaction // 玩家余额变动日志表
	PlayerDayaccount  PlayerDayaccount  // 玩家输赢表

	// 权限管理
	Admin      Admin      // 管理员表
	Api        Api        // api表
	Menu       Menu       // 菜单表
	Role       Role       // 角色表
	RoleApi    RoleApi    // 角色api表
	RoleMenu   RoleMenu   // 角色菜单表
	CasbinRule CasbinRule // casbin表

	GameNotifyLog GameNotifyLog // 游戏通知接口请求日志表

	// =========================================================================================todo
	BetAccountRecord BetAccountRecord // 投注帐目表
	TipRecord        TipRecord        // 打赏明细表
}

// Model 模型入口
var Model = new(Group)
