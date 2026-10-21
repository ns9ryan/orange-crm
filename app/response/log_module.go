package response

// AdminLoginLogResponse 管理员登录日志返回参数
type AdminLoginLogResponse struct {
	Id        uint   `json:"id"`         // ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	RoleName  string `json:"role_name"`  // 角色名称
	LoginIp   string `json:"login_ip"`   // 登录IP
	LoginTime string `json:"login_time"` // 登录时间
}

// MerchantLoginLogResponse 商户登录日志返回参数
type MerchantLoginLogResponse struct {
	Id        uint   `json:"id"`         // ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	RoleName  string `json:"role_name"`  // 角色名称
	LoginIp   string `json:"login_ip"`   // 登录IP
	LoginTime string `json:"login_time"` // 登录时间
	IsOk      uint   `json:"is_ok"`      // 是否成功:取值说明在模型里
	Reason    string `json:"reason"`     // 失败原因
}

// PlayerLoginLogResponse 会员登录日志返回参数
type PlayerLoginLogResponse struct {
	Id        uint   `json:"id"`         // ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	LoginIp   string `json:"login_ip"`   // 登录IP
	LoginTime string `json:"login_time"` // 登录时间
	IsOk      uint   `json:"is_ok"`      // 是否成功:取值说明在模型里
	Reason    string `json:"reason"`     // 失败原因
}

// AdminOperationLogListResponse 管理员操作日志列表返回参数
type AdminOperationLogListResponse struct {
	Id                      uint   `json:"id"`           // ID
	Username                string `json:"username"`     // 账号
	Module                  string `json:"module"`       // 模块名
	Action                  string `json:"action"`       // 操作动作
	Content                 string `json:"content"`      // 操作内容描述
	AdminOperationEventType uint   `json:"event_type"`   // 事件类型
	CreatedTime             string `json:"created_time"` // 操作时间
}

// MerchantOperationLogListResponse 商户操作日志列表返回参数
type MerchantOperationLogListResponse struct {
	Id                         uint   `json:"id"`           // ID
	Username                   string `json:"username"`     // 账号
	Module                     string `json:"module"`       // 模块名
	Action                     string `json:"action"`       // 操作动作
	Content                    string `json:"content"`      // 操作内容描述
	MerchantOperationEventType uint   `json:"event_type"`   // 事件类型
	CreatedTime                string `json:"created_time"` // 操作时间
}

// ControlLoginLogResponse 控端登录日志返回参数
type ControlLoginLogResponse struct {
	Id        uint   `json:"id"`         // ID
	Username  string `json:"username"`   // 账号
	LoginIp   string `json:"login_ip"`   // 登录IP
	LoginTime string `json:"login_time"` // 登录时间
	IsOk      uint   `json:"is_ok"`      // 是否成功:取值说明在模型里
	Reason    string `json:"reason"`     // 失败原因
}
