package request

// MerchantDayaccountListRequest 商户输赢表请求参数
type MerchantDayaccountListRequest struct {
	PageRequest        // 分页参数
	MerchantId  string `json:"merchant_id"`                                                         // 商户账号
	StartTime   string `json:"start_time" v:"required|date-format:Y-m-d H:i:s"`                     // 开始时间
	EndTime     string `json:"end_time" v:"required|date-format:Y-m-d H:i:s|after-equal:StartTime"` // 结束时间
}

// GetSubMerchantDayaccountRequest 获取主商户下子商户的输赢报表请求参数
type GetSubMerchantDayaccountRequest struct {
	ParentId  string `json:"parent_id" v:"required"`                                              // 主商户账号
	StartTime string `json:"start_time" v:"required|date-format:Y-m-d H:i:s"`                     // 开始时间
	EndTime   string `json:"end_time" v:"required|date-format:Y-m-d H:i:s|after-equal:StartTime"` // 结束时间
}
