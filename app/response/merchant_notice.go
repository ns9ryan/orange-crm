package response

// MerchantNoticeListResponse 商户公告列表返回结构体
type MerchantNoticeListResponse struct {
	Id          uint     `json:"id"`           // ID
	NoticeId    string   `json:"notice_id"`    // 公告ID
	Title       string   `json:"title"`        // 公告标题
	Content     string   `json:"content"`      // 公告内容
	StartTime   string   `json:"start_time"`   // 开始时间
	EndTime     string   `json:"end_time"`     // 结束时间
	IsAll       uint     `json:"is_all"`       // 是否发给所有会员:取值说明在模型里
	PlayerIds   []string `json:"player_ids"`   // 指定发给的会员
	OperatorId  string   `json:"operator_id"`  // 发布公告的账号
	CreatedTime string   `json:"created_time"` // 创建时间
}

// MerchantNoticeDetailResponse 商户公告详情返回结构体
type MerchantNoticeDetailResponse struct {
	Id         uint     `json:"id"`          // ID
	NoticeId   string   `json:"notice_id"`   // 公告ID
	Title      string   `json:"title"`       // 公告标题
	Content    string   `json:"content"`     // 公告内容
	StartTime  string   `json:"start_time"`  // 开始时间
	EndTime    string   `json:"end_time"`    // 结束时间
	IsAll      uint     `json:"is_all"`      // 是否发给所有会员:取值说明在模型里
	PlayerIds  []string `json:"player_ids"`  // 指定发给的会员
	OperatorId string   `json:"operator_id"` // 发布公告的账号
}
