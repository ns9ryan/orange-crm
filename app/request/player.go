package request

// SavePlayerRequest 创建、编辑玩家请求参数
type SavePlayerRequest struct {
	Id         uint   `json:"id" v:"checkId"`                                          // ID
	MerchantId string `json:"merchant_id" v:"required|checkMerchantId"`                // 商户账号
	Username   string `json:"username" v:"required-if:Id,0|length:2,30|checkUsername"` // 用户名,创建的时候必传
	Nickname   string `json:"nickname" v:"required|length:2,50"`                       // 昵称
	ParentId   string `json:"parent_id" v:"length:2,30|checkParentId"`                 // 上级玩家账号

	Status     uint `json:"status" v:"required|in:0,1"`        // 是否允许登录:取值说明在模型里
	BetPaused  uint `json:"bet_paused" v:"required|in:0,1"`    // 是否暂停投注:取值说明在模型里
	MaxWin     uint `json:"max_win" v:"required|max:99999999"` // 最高赢额,0表示不限制
	PlayerType uint `json:"player_type" v:"required|in:1,2"`   // 玩家类型:取值说明在模型里
	MoneyType  uint `json:"money_type" v:"required|in:1,2"`    // 玩家模式类型:取值说明在模型里
	Level      uint `json:"level" v:"required|min:1"`          // 等级

	Email      string `json:"email" v:"max-length:100"`                  // 邮箱
	Phone      string `json:"phone" v:"max-length:20"`                   // 手机号
	Password   string `json:"password" v:"required-if:Id,0|length:6,30"` // 密码,创建的时候必传
	ChipIds    []uint `json:"chip_ids" v:"required|checkChipIds"`        // 筹码ID数组
	TemplateId string `json:"template_id" v:"required|checkTemplateId"`  // 商户限红模板ID
}

// PlayerListRequest 玩家列表请求参数
type PlayerListRequest struct {
	PageRequest        // 分页参数
	SortRequest        // 排序参数
	MerchantId  string `json:"merchant_id"`           // 商户账号
	Username    string `json:"username"`              // 用户名
	Nickname    string `json:"nickname"`              // 昵称
	Currency    string `json:"currency"`              // 币种
	Status      *uint  `json:"status" v:"in:0,1"`     // 是否冻结:取值说明在模型里
	BetPaused   *uint  `json:"bet_paused" v:"in:0,1"` // 是否暂停投注:取值说明在模型里
	IsOnline    *uint  `json:"is_online" v:"in:0,1"`  // 是否在线:取值说明在模型里
}

// LoginPlayerRequest 会员登录请求参数
type LoginPlayerRequest struct {
	Username string `json:"username" v:"required|length:2,30"` // 账号
	Password string `json:"password" v:"required|length:6,30"` // 密码
}
