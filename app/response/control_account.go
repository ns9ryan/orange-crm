package response

// ControlAccountLevelBox 控端账号等级选择框返回值
type ControlAccountLevelBox struct {
	Name  string `json:"name"`  // 等级名称
	Value uint   `json:"value"` // 等级对应的值
}

// ControlAccountListResponse 控端账号列表返回结构体
type ControlAccountListResponse struct {
	Id        uint   `json:"id"`         // 控端账号ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	Level     uint   `json:"level"`      // 等级
	Status    uint   `json:"status"`     // 是否允许登录:取值说明在模型里
	LoginTime string `json:"login_time"` // 上次登录时间
	LoginIp   string `json:"login_ip"`   // 登录IP
}

// ControlAccountDetailResponse 控端账号详情返回结构体
type ControlAccountDetailResponse struct {
	Id        uint   `json:"id"`         // 控端账号ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	Level     uint   `json:"level"`      // 等级
	Status    uint   `json:"status"`     // 是否允许登录:取值说明在模型里
	LoginTime string `json:"login_time"` // 上次登录时间
	LoginIp   string `json:"login_ip"`   // 登录IP
}

// ResetControlAccountPasswordResponse 重置控端账号密码返回结构体
type ResetControlAccountPasswordResponse struct {
	Id       uint   `json:"id"`       // 控端账号ID
	Password string `json:"password"` // 密码
}

// ControlAccountBoxResponse 获取控端下拉框数据返回参数
type ControlAccountBoxResponse struct {
	Id       uint   `json:"id"`       // ID
	Username string `json:"username"` // 账号
}
