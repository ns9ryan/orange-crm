package response

// LoginAdminResponse 管理员登录返回结构体
type LoginAdminResponse struct {
	Id          uint   `json:"id"`           // 管理员ID
	RoleId      uint   `json:"role_id"`      // 角色ID
	Username    string `json:"username"`     // 账号
	Nickname    string `json:"nickname"`     // 昵称
	Avatar      string `json:"avatar"`       // 头像
	Status      uint   `json:"status"`       // 是否允许登录:取值说明在模型里
	Token       string `json:"token"`        // 登录token
	ExpiresAt   string `json:"expires_at"`   // token过期时间
	DefaultPath string `json:"default_path"` // 默认路劲
}

// AdminListResponse 管理员列表返回结构体
type AdminListResponse struct {
	Id        uint   `json:"id"`         // 管理员ID
	Username  string `json:"username"`   // 账号
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	Status    uint   `json:"status"`     // 是否允许登录:取值说明在模型里
	RoleId    uint   `json:"role_id"`    // 角色ID
	RoleName  string `json:"role_name"`  // 角色名称
	LoginTime string `json:"login_time"` // 登录时间
	LoginIp   string `json:"login_ip"`   // 登录IP
}

// AdminDetailResponse 管理员详情返回结构体
type AdminDetailResponse struct {
	Id        uint   `json:"id"`         // 管理员ID
	Username  string `json:"username"`   // 账号f
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	Status    uint   `json:"status"`     // 是否允许登录:取值说明在模型里
	RoleId    uint   `json:"role_id"`    // 角色ID
	RoleName  string `json:"role_name"`  // 角色名称
	LoginTime string `json:"login_time"` // 登录时间
	LoginIp   string `json:"login_ip"`   // 登录IP
}

// ResetAdminPasswordResponse 重置管理员密码返回结构体
type ResetAdminPasswordResponse struct {
	Id       uint   `json:"id"`       // 管理员ID
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
}

// GetAdminInfoResponse 获取管理员信息返回结构体
type GetAdminInfoResponse struct {
	Id        uint   `json:"id"`         // 管理员ID
	Username  string `json:"username"`   // 账号f
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	Status    uint   `json:"status"`     // 是否允许登录:取值说明在模型里
	RoleId    uint   `json:"role_id"`    // 角色ID
	RoleName  string `json:"role_name"`  // 角色名称
	LoginTime string `json:"login_time"` // 登录时间
	LoginIp   string `json:"login_ip"`   // 登录IP
}
