package request

// PlayerDayaccountListRequest 玩家输赢表请求参数
type PlayerDayaccountListRequest struct {
	PageRequest          // 分页参数
	Username    string   `json:"username"`                                                            // 玩家账号
	MerchantId  string   `json:"merchant_id"`                                                         // 商户账号
	StartTime   string   `json:"start_time" v:"required|date-format:Y-m-d H:i:s"`                     // 开始时间
	EndTime     string   `json:"end_time" v:"required|date-format:Y-m-d H:i:s|after-equal:StartTime"` // 结束时间
	PlayerIds   []string `json:"player_ids"`                                                          // 玩家账号-用来判断总网和商户查询
}
