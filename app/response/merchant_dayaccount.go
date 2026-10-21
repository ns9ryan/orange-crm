package response

// MerchantDayaccountListResponse 商户输赢表返回结构体
type MerchantDayaccountListResponse struct {
	MerchantId       string `json:"merchant_id"`        // 对应商户表的username
	ParentId         string `json:"parent_id"`          // 上级商户的userName
	Currency         string `json:"currency"`           // 币种
	Digits           uint   `json:"digits"`             // 币种的最小单位倍数
	LastBet          string `json:"last_bet"`           // 累计投注金额，存储为最小单位
	LastBetEffective string `json:"last_bet_effective"` // 累计有效投注金额，存储为最小单位
	LastWater        string `json:"last_water"`         // 累计洗码金额，存储为最小单位
	LastWinLost      string `json:"last_win_lost"`      // 累计输赢金额，存储为最小单位
	LastTip          string `json:"last_tip"`           // 累计打赏金额，存储为最小单位
	StartTime        string `json:"start_time"`         // 开始时间
	EndTime          string `json:"end_time"`           // 结束时间
	HasChildren      uint   `json:"has_children"`       // 是否存在子商户数据:0否,1是
	CommissionRate   uint   `json:"commission_rate"`    // 成数
	RollingRate      uint   `json:"rolling_rate"`       // 洗码比
	AccountDate      string `json:"account_date"`       // 结算归属日期
}

// GetSubMerchantDayaccountResponse 获取主商户下子商户的输赢报表返回结构体
type GetSubMerchantDayaccountResponse struct {
	MerchantId       string `json:"merchant_id"`        // 子商户 ID
	ParentId         string `json:"parent_id"`          // 主商户的 ID
	Currency         string `json:"currency"`           // 币种
	Digits           uint   `json:"digits"`             // 币种的最小单位倍数
	LastBet          string `json:"last_bet"`           // 累计投注金额，存储为最小单位
	LastBetEffective string `json:"last_bet_effective"` // 累计有效投注金额，存储为最小单位
	LastWater        string `json:"last_water"`         // 累计洗码金额，存储为最小单位
	LastWinLost      string `json:"last_win_lost"`      // 累计输赢金额，存储为最小单位
	LastTip          string `json:"last_tip"`           // 累计打赏金额，存储为最小单位
	StartTime        string `json:"start_time"`         // 开始时间
	EndTime          string `json:"end_time"`           // 结束时间
	HasChildren      uint   `json:"has_children"`       // 子商户没有下级，所以值为 0
	CommissionRate   uint   `json:"commission_rate"`    // 成数
	RollingRate      uint   `json:"rolling_rate"`       // 洗码比
	AccountDate      string `json:"account_date"`       // 结算归属日期
}
