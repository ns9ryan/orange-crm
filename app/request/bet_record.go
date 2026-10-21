package request

// BetRecordListRequest 投注明细列表请求参数
type BetRecordListRequest struct {
	PageRequest             // 分页参数
	SortRequest             // 排序参数
	GameTableCode    string `json:"game_table_code"`                        // 游戏id
	PlayerUsername   string `json:"player_username"`                        // 玩家账号
	MerchantUsername string `json:"merchant_username"`                      // 商户账号
	Currency         string `json:"currency"`                               // 币种
	StartTime        string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime          string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}
