package request

// SaveSystemNoticeRequest 创建、编辑系统公告请求参数
type SaveSystemNoticeRequest struct {
	Id          uint     `json:"id" v:"checkId"`                                                      // ID
	Title       string   `json:"title" v:"required|length:2,100|checkTitle"`                          // 公告标题
	Content     string   `json:"content" v:"required"`                                                // 公告内容
	StartTime   string   `json:"start_time" v:"required|date-format:Y-m-d H:i:s"`                     // 开始时间
	EndTime     string   `json:"end_time" v:"required|date-format:Y-m-d H:i:s|after-equal:StartTime"` // 结束时间
	IsAll       uint     `json:"is_all" v:"required|in:0,1"`                                          // 是否发给所有商户:取值说明在模型里
	MerchantIds []string `json:"merchant_ids" v:"required-if:IsAll,1"`                                // 指定发给的商户
}

// SystemNoticeListRequest 系统公告列表请求参数
type SystemNoticeListRequest struct {
	PageRequest // 分页参数
}

// SystemNoticeListForSelfRequest 商户收到的系统公共列表请求参数
type SystemNoticeListForSelfRequest struct {
	PageRequest       // 分页参数
	Read        *uint `json:"read" v:"in:0,1"` // 是否已读:取值说明在模型里
}

// SystemNoticeMarkReadRequest 商户系统公告标记已读请求参数
type SystemNoticeMarkReadRequest struct {
	NoticeId string `json:"notice_id" v:"required|length:2,50"` // 公告ID
}
