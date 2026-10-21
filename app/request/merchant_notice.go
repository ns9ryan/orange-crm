package request

// SaveMerchantNoticeRequest 创建、编辑系统公告请求参数
type SaveMerchantNoticeRequest struct {
	Id        uint     `json:"id"`                                                                  // ID
	Title     string   `json:"title" v:"required|length:2,100|checkTitle"`                          // 公告标题
	Content   string   `json:"content" v:"required"`                                                // 公告内容
	StartTime string   `json:"start_time" v:"required|date-format:Y-m-d H:i:s"`                     // 开始时间
	EndTime   string   `json:"end_time" v:"required|date-format:Y-m-d H:i:s|after-equal:StartTime"` // 结束时间
	IsAll     uint     `json:"is_all" v:"required|in:0,1"`                                          // 是否发给所有会员:取值说明在模型里
	PlayerIds []string `json:"player_ids" v:"required-if:IsAll,1"`                                  // 指定发给的会员
}

// MerchantNoticeListRequest 系统公告列表请求参数
type MerchantNoticeListRequest struct {
	PageRequest // 分页参数
}
