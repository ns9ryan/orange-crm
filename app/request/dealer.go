package request

// SaveDealerRequest 创建、编辑荷官请求参数
type SaveDealerRequest struct {
	Id             uint   `json:"id" v:"checkId"`                                                       // ID
	Nickname       string `json:"nickname" v:"required|length:2,50"`                                    // 昵称
	EmployeeNumber string `json:"employee_number" v:"required-if:Id,0|length:2,50|checkEmployeeNumber"` // 工号
	DealerType     uint   `json:"dealer_type" v:"required|in:1,2"`                                      // 类型:取值说明在模型里
	Level          uint   `json:"level" v:"required|min:1"`                                             // 等级
	LeadName       string `json:"lead_name" v:"length:2,50"`                                            // 上级名称
	Status         uint   `json:"status" v:"required|in:0,1"`                                           // 状态:取值说明在模型里
	Avatar         string `json:"avatar" v:"max-length:255"`                                            // 头像
}

// DealerListRequest 荷官列表请求参数
type DealerListRequest struct {
	PageRequest // 分页参数
}
