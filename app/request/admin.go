package request

// LoginAdminRequest 管理员登录请求参数
type LoginAdminRequest struct {
	Username string `json:"username" v:"required|length:2,30"` // 用户名
	Password string `json:"password" v:"required|length:6,30"` // 密码
}

// SaveAdminRequest 创建、编辑管理员请求参数
type SaveAdminRequest struct {
	Id       uint   `json:"id" v:"checkId"`                                          // ID
	Username string `json:"username" v:"required-if:Id,0|length:2,30|checkUsername"` // 用户名,创建的时候必传
	Password string `json:"password" v:"required-if:Id,0|length:6,30"`               // 密码,创建的时候必传
	Nickname string `json:"nickname" v:"required|length:2,50"`                       // 昵称
	Avatar   string `json:"avatar" v:"max-length:255"`                               // 头像
	Status   uint   `json:"status" v:"required|in:0,1"`                              // 是否允许登录:取值说明在模型里
	RoleId   uint   `json:"role_id" v:"required|min:1|checkRoleId"`                  // 角色ID
}

// AdminListRequest 管理员列表请求参数
type AdminListRequest struct {
	PageRequest // 分页参数
}

// EditCurrentPasswordRequest 编辑当前登录的管理员密码请求参数
type EditCurrentPasswordRequest struct {
	OldPassword string `json:"old_password" v:"required|length:6,30"` // 原密码
	NewPassword string `json:"new_password" v:"required|length:6,30"` // 新密码
}
