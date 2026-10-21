package request

// AdminLoginLogListRequest 管理员登录日志列表请求参数
type AdminLoginLogListRequest struct {
	PageRequest        // 分页参数
	Username    string `json:"username"`                               // 账号
	StartTime   string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime     string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}

// MerchantLoginLogListRequest 商户登录日志列表请求参数
type MerchantLoginLogListRequest struct {
	PageRequest        // 分页参数
	Username    string `json:"username"`                               // 账号
	StartTime   string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime     string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}

// PlayerLoginLogListRequest 会员登录日志列表请求参数
type PlayerLoginLogListRequest struct {
	PageRequest        // 分页参数
	MerchantId  string `json:"merchant_id"`                            // 商户账号
	Username    string `json:"username"`                               // 账号
	StartTime   string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime     string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}

// AdminOperationLogListRequest 管理员操作日志列表请求参数
type AdminOperationLogListRequest struct {
	PageRequest                    // 分页参数
	Username                string `json:"username"`                               // 账号
	AdminOperationEventType string `json:"event_type"`                             // 事件类型
	StartTime               string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime                 string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}

// MerchantOperationLogListRequest 商户操作日志列表请求参数
type MerchantOperationLogListRequest struct {
	PageRequest                       // 分页参数
	Username                   string `json:"username"`                               // 账号
	MerchantOperationEventType string `json:"event_type"`                             // 事件类型
	StartTime                  string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime                    string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}

// ControlLoginLogListRequest 控端登录日志列表请求参数
type ControlLoginLogListRequest struct {
	PageRequest        // 分页参数
	Username    string `json:"username"`                               // 账号
	StartTime   string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime     string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}
