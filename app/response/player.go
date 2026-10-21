package response

// PlayerListResponse 玩家列表返回结构体
type PlayerListResponse struct {
	Id                 uint   `json:"id"`                   // 玩家ID
	MerchantName       string `json:"merchant_name"`        // 商户名称
	MerchantId         string `json:"merchant_id"`          // 商户账号
	Currency           string `json:"currency"`             // 币种
	MerchantCurrencyId string `json:"merchant_currency_id"` // 商户币种唯一标识
	Username           string `json:"username"`             // 账号
	Nickname           string `json:"nickname"`             // 昵称
	ParentId           string `json:"parent_id"`            // 上级玩家的的username

	LoginTimeFirst string `json:"login_time_first"` // 首次次登录时间
	LoginIpFirst   string `json:"login_ip_first"`   // 首次次登录IP
	LoginTime      string `json:"login_time"`       // 上次登录时间
	LoginIp        string `json:"login_ip"`         // 上次登录IP

	Status     uint `json:"status"`      // 是否允许登录:取值说明在模型里
	BetPaused  uint `json:"bet_paused"`  // 是否暂停投注:取值说明在模型里
	IsOnline   uint `json:"is_online"`   // 是否在线:取值说明在模型里
	PlayerType uint `json:"player_type"` // 玩家类型:取值说明在模型里

	Digits            uint `json:"digits"`              // 币种的最小单位倍数
	CreditSum         uint `json:"credit_sum"`          // 累计充值额度
	Balance           uint `json:"balance"`             // 账户余额
	Tip               uint `json:"tip"`                 // 打赏金额
	TotalBet          uint `json:"total_bet"`           // 总投注金额
	TotalBetEffective uint `json:"total_bet_effective"` // 总有效投注金额
	TotalWater        uint `json:"total_water"`         // 总洗码金额
	TotalWinLost      int  `json:"total_win_lost"`      // 总输赢金额

	LastBetEffective uint `json:"last_bet_effective"` // 上次有效投注金额
}

// PlayerDetailResponse 玩家详情返回结构体
type PlayerDetailResponse struct {
	Id                 uint   `json:"id"`                   // 玩家ID
	MerchantName       string `json:"merchant_name"`        // 商户名称
	MerchantId         string `json:"merchant_id"`          // 商户账号
	Currency           string `json:"currency"`             // 币种
	MerchantCurrencyId string `json:"merchant_currency_id"` // 商户币种唯一标识
	ApiUser            string `json:"api_user"`             // 存储api会员原始账号
	Username           string `json:"username"`             // 账号
	Nickname           string `json:"nickname"`             // 昵称
	ParentId           string `json:"parent_id"`            // 上级玩家的的username

	LoginTimeFirst string `json:"login_time_first"` // 首次次登录时间
	LoginIpFirst   string `json:"login_ip_first"`   // 首次次登录IP
	LoginTime      string `json:"login_time"`       // 上次登录时间
	LoginIp        string `json:"login_ip"`         // 上次登录IP

	Status     uint `json:"status"`      // 是否允许登录:取值说明在模型里
	BetPaused  uint `json:"bet_paused"`  // 是否暂停投注:取值说明在模型里
	CreateType uint `json:"create_type"` // 玩家来源:取值说明在模型里
	MaxWin     uint `json:"max_win"`     // 最高赢额:取值说明在模型里
	DeviceType uint `json:"device_type"` // 设备类型:取值说明在模型里
	IsOnline   uint `json:"is_online"`   // 是否在线:取值说明在模型里
	PlayerType uint `json:"player_type"` // 玩家类型:取值说明在模型里
	MoneyType  uint `json:"money_type"`  // 玩家模式类型:取值说明在模型里
	Level      uint `json:"level"`       // 等级

	Email      string `json:"email"`       // 邮箱
	Phone      string `json:"phone"`       // 手机号
	ChipIds    []uint `json:"chip_ids"`    // 筹码
	TemplateId string `json:"template_id"` // 商户限红模板ID

	Digits            uint `json:"digits"`              // 币种的最小单位倍数
	CreditSum         uint `json:"credit_sum"`          // 累计充值额度
	Balance           uint `json:"balance"`             // 账户余额
	Tip               uint `json:"tip"`                 // 总打赏金额
	TotalBet          uint `json:"total_bet"`           // 总投注金额
	TotalBetEffective uint `json:"total_bet_effective"` // 总有效投注金额
	TotalWater        uint `json:"total_water"`         // 总洗码金额
	TotalWinLost      int  `json:"total_win_lost"`      // 总输赢金额

	LastBet          uint `json:"last_bet"`           // 上次投注金额
	LastBetEffective uint `json:"last_bet_effective"` // 上次有效投注金额
	LastWater        uint `json:"last_water"`         // 上次洗码金额
	LastWinLost      int  `json:"last_win_lost"`      // 上次输赢金额
	LastTip          uint `json:"last_tip"`           // 上次打赏金额
}

// ResetPlayerPasswordResponse 重置玩家密码返回结构体
type ResetPlayerPasswordResponse struct {
	Id       uint   `json:"id"`       // 玩家ID
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
}
