package response

// MerchantListResponse 商户列表返回参数
type MerchantListResponse struct {
	Id           uint   `json:"id"`            // 商户ID
	MerchantName string `json:"merchant_name"` // 商户名称
	MerchantCode string `json:"merchant_code"` // 商户号
	ParentId     string `json:"parent_id"`     // 上级商户的username
	Username     string `json:"username"`      // 账号
	Nickname     string `json:"nickname"`      // 昵称
	Status       uint   `json:"status"`        // 是否允许登录:取值说明在模型里

	CommissionRate    uint `json:"commission_rate"`     // 成数
	MaxCommissionRate uint `json:"max_commission_rate"` // 上级成数
	RollingRate       uint `json:"rolling_rate"`        // 洗码比
	MaxRollingRate    uint `json:"max_rolling_rate"`    // 上级洗码比

	WalletType  uint                   `json:"wallet_type"`  // 钱包类型:取值说明在模型里
	LoginTime   string                 `json:"login_time"`   // 上次登录时间
	LoginIp     string                 `json:"login_ip"`     // 登录IP
	Currency    string                 `json:"currency"`     // 商户币种
	PlayerCount int64                  `json:"player_count"` // 玩家数量
	MaxTip      uint                   `json:"max_tip"`      // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip      uint                   `json:"min_tip"`      // 荷官可接收的最小打赏金额，存储为最小单位
	Digits      uint                   `json:"digits"`       // 币种的最小单位倍数
	Children    []MerchantListResponse `json:"children"`     // 子账号信息
}

// MerchantDetailResponse 商户详情返回参数
type MerchantDetailResponse struct {
	Id           uint   `json:"id"`            // 商户ID
	MerchantCode string `json:"merchant_code"` // 商户号
	MerchantName string `json:"merchant_name"` // 商户名称
	ParentId     string `json:"parent_id"`     // 上级商户的username
	Username     string `json:"username"`      // 账号
	Nickname     string `json:"nickname"`      // 昵称

	CommissionRate    uint `json:"commission_rate"`     // 成数
	MaxCommissionRate uint `json:"max_commission_rate"` // 上级成数
	RollingRate       uint `json:"rolling_rate"`        // 洗码比
	MaxRollingRate    uint `json:"max_rolling_rate"`    // 上级洗码比

	Status     uint `json:"status"`      // 是否允许登录:取值说明在模型里
	WalletType uint `json:"wallet_type"` // 钱包类型:取值说明在模型里

	ChipIds         []uint `json:"chip_ids"`          // 筹码
	ChipTemplateIds []uint `json:"chip_template_ids"` // 筹码模板
	RoleId          uint   `json:"role_id"`           // 角色ID
	TemplateId      string `json:"template_id"`       // 商户限红模板ID

	AesKeyBase64  string `json:"aes_key_base64"`  // aesKey的base64字符串
	Md5SaltBase64 string `json:"md5_salt_base64"` // md5盐的base64字符串
	OrToken       string `json:"or_token"`        // 商户api调用验证or_token
	OrTokenKey    string `json:"or_token_key"`    // 生成token用的商户唯一key，永久有效
	WalletUrl     string `json:"wallet_url"`      // 单钱包URL
	WhiteList     string `json:"white_list"`      // 白名单IP,逗号分隔
	Currency      string `json:"currency"`        // 币种

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

	RoleName  string `json:"role_name"`  // 角色名称
	MaxTip    uint   `json:"max_tip"`    // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip    uint   `json:"min_tip"`    // 荷官可接收的最小打赏金额，存储为最小单位
	Digits    uint   `json:"digits"`     // 币种的最小单位倍数
	TipConfig []uint `json:"tip_config"` // 商户打赏配置
}

// CheckUsernameResponse 判断账号是否可用返回参数
type CheckUsernameResponse struct {
	Available bool `json:"available"` // 是否可用
}

// LoginMerchantResponse 商户登录返回参数
type LoginMerchantResponse struct {
	Id          uint   `json:"id"`           // 商户账号ID
	ParentId    string `json:"parent_id"`    // 上级商户的username
	Username    string `json:"username"`     // 账号
	Nickname    string `json:"nickname"`     // 昵称
	Status      uint   `json:"status"`       // 是否允许登录:取值说明在模型里
	Token       string `json:"token"`        // 登录token
	ExpiresAt   string `json:"expires_at"`   // token过期时间
	DefaultPath string `json:"default_path"` // 默认路劲
	Currency    string `json:"currency"`     // 币种
}

// ResetMerchantPasswordResponse 重置商户密码返回结构体
type ResetMerchantPasswordResponse struct {
	Id       uint   `json:"id"`       // 商户ID
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
}

// MerchantBoxResponse 获取商户下拉框数据返回参数
type MerchantBoxResponse struct {
	Id         uint   `json:"id"`          // ID
	Username   string `json:"username"`    // 账号
	TemplateId string `json:"template_id"` // 商户限红模板ID
}

// MerchantApiKeyResponse 获取api相关密钥返回参数
type MerchantApiKeyResponse struct {
	Id            uint   `json:"id"`              // 商户ID
	MerchantCode  string `json:"merchant_code"`   // 商户号
	Username      string `json:"username"`        // 账号
	AesKeyBase64  string `json:"aes_key_base64"`  // aesKey的base64字符串
	Md5SaltBase64 string `json:"md5_salt_base64"` // md5盐的base64字符串
	OrToken       string `json:"or_token"`        // 商户api调用验证or_token
	OrTokenKey    string `json:"or_token_key"`    // 生成token用的商户唯一key，永久有效
}

// RegenerateAesKeySaltResponse 重新生成商户AesKeySalt返回参数
type RegenerateAesKeySaltResponse struct {
	Id           uint   `json:"id"`             // 商户ID
	AesKeyBase64 string `json:"aes_key_base64"` // aesKey的base64字符串
}

// RegenerateMd5SaltResponse 重新生成商户md5盐返回参数
type RegenerateMd5SaltResponse struct {
	Id            uint   `json:"id"`              // 商户ID
	Md5SaltBase64 string `json:"md5_salt_base64"` // md5盐的base64字符串
}

// RegenerateTokenKeyResponse 重新生成商户tokenKey返回参数
type RegenerateTokenKeyResponse struct {
	Id         uint   `json:"id"`           // 商户ID
	OrToken    string `json:"or_token"`     // 商户api调用验证or_token
	OrTokenKey string `json:"or_token_key"` // 生成token用的商户唯一key，永久有效
}
