package request

// CreditAndDebitListRequest 上分/下分记录列表请求参数
type CreditAndDebitListRequest struct {
	PageRequest        // 分页参数
	MerchantId  string `json:"merchant_id"`                            // 对应商户表的username
	Username    string `json:"username"`                               // 账号
	StartTime   string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime     string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}
