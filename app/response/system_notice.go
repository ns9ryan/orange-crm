package response

// SystemNoticeListResponse 系统公告列表返回结构体
type SystemNoticeListResponse struct {
	Id          uint     `json:"id"`           // ID
	NoticeId    string   `json:"notice_id"`    // 公告ID
	Title       string   `json:"title"`        // 公告标题
	Content     string   `json:"content"`      // 公告内容
	StartTime   string   `json:"start_time"`   // 开始时间
	EndTime     string   `json:"end_time"`     // 结束时间
	IsAll       uint     `json:"is_all"`       // 是否发给所有商户:取值说明在模型里
	MerchantIds []string `json:"merchant_ids"` // 指定发给的商户
	OperatorId  string   `json:"operator_id"`  // 发布公告的账号
	CreatedTime string   `json:"created_time"` // 创建时间
}

// SystemNoticeDetailResponse 系统公告详情返回结构体
type SystemNoticeDetailResponse struct {
	Id          uint     `json:"id"`           // ID
	NoticeId    string   `json:"notice_id"`    // 公告ID
	Title       string   `json:"title"`        // 公告标题
	Content     string   `json:"content"`      // 公告内容
	StartTime   string   `json:"start_time"`   // 开始时间
	EndTime     string   `json:"end_time"`     // 结束时间
	IsAll       uint     `json:"is_all"`       // 是否发给所有商户:取值说明在模型里
	MerchantIds []string `json:"merchant_ids"` // 指定发给的商户
	OperatorId  string   `json:"operator_id"`  // 发布公告的账号
}

// SystemNoticeListForSelfResponse 商户收到的系统公共列表返回结构体
type SystemNoticeListForSelfResponse struct {
	Id        uint   `json:"id"`         // ID
	NoticeId  string `json:"notice_id"`  // 公告ID
	Title     string `json:"title"`      // 公告标题
	Content   string `json:"content"`    // 公告内容
	StartTime string `json:"start_time"` // 开始时间
	EndTime   string `json:"end_time"`   // 结束时间
	Read      uint   `json:"read"`       // 是否已读:取值说明在模型里
}

// LoginMerchantNoticeListResponse 商户登录的系统消息
type LoginMerchantNoticeListResponse struct {
	Id       uint   `json:"id"`        // ID
	NoticeId string `json:"notice_id"` // 公告ID
	Title    string `json:"title"`     // 公告标题
	Content  string `json:"content"`   // 公告内容
}
