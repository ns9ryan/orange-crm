package response

// DealerAvatarResponse 荷官上传头像返回结构体
type DealerAvatarResponse struct {
	Src string `json:"src"`
}

// DealerListResponse 荷官列表返回结构体
type DealerListResponse struct {
	Id             uint   `json:"id"`              // 荷官ID
	Nickname       string `json:"nickname"`        // 昵称
	EmployeeNumber string `json:"employee_number"` // 工号
	BalanceAmount  uint   `json:"balance_amount"`  // 账户余额
	CurrencyCode   string `json:"currency_code"`   // 币种代码
	CurrencyName   string `json:"currency_name"`   // 币种名称
	CurrencySymbol string `json:"currency_symbol"` // 币种符号
	Digits         uint   `json:"digits"`          // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
	DealerType     uint   `json:"dealer_type"`     // 类型:取值说明在模型里
	Level          uint   `json:"level"`           // 等级
	LeadName       string `json:"lead_name"`       // 上级名称
	Status         uint   `json:"status"`          // 状态:取值说明在模型里
	Avatar         string `json:"avatar"`          // 头像
	Qrcode         string `json:"qrcode"`          // qrcode
}

// DealerDetailResponse 荷官详情返回结构体
type DealerDetailResponse struct {
	Id             uint   `json:"id"`              // 荷官ID
	Nickname       string `json:"nickname"`        // 昵称
	EmployeeNumber string `json:"employee_number"` // 工号
	BalanceAmount  uint   `json:"balance_amount"`  // 账户余额
	CurrencyCode   string `json:"currency_code"`   // 币种代码
	CurrencyName   string `json:"currency_name"`   // 币种名称
	CurrencySymbol string `json:"currency_symbol"` // 币种符号
	Digits         uint   `json:"digits"`          // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
	DealerType     uint   `json:"dealer_type"`     // 类型:取值说明在模型里
	Level          uint   `json:"level"`           // 等级
	LeadName       string `json:"lead_name"`       // 上级名称
	Status         uint   `json:"status"`          // 状态:取值说明在模型里
	Avatar         string `json:"avatar"`          // 头像
	Qrcode         string `json:"qrcode"`          // qrcode
}

// DealerEmployeeNumberResponse 荷官工号返回结构体
type DealerEmployeeNumberResponse struct {
	EmployeeNumber string `json:"employee_number"` // 工号
}
