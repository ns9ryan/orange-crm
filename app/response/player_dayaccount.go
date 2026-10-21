package response

// PlayerDayaccountListResponse 玩家输赢表返回结构体
type PlayerDayaccountListResponse struct {
	Username         string `json:"username"`           // 玩家账号,每个玩家的唯一值
	MerchantId       string `json:"merchant_id"`        // 对应商户表的username
	Digits           uint   `json:"digits"`             // 币种的最小单位倍数
	LastBet          string `json:"last_bet"`           // 累计投注金额，存储为最小单位
	LastBetEffective string `json:"last_bet_effective"` // 累计有效投注金额，存储为最小单位
	LastWater        string `json:"last_water"`         // 累计洗码金额，存储为最小单位
	LastWinLost      string `json:"last_win_lost"`      // 累计输赢金额，存储为最小单位
	LastTip          string `json:"last_tip"`           // 累计打赏金额，存储为最小单位
	StartTime        string `json:"start_time"`         // 开始时间
	EndTime          string `json:"end_time"`           // 结束时间
	Currency         string `json:"currency"`           // 币种
}
