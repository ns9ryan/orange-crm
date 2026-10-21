package request

// SaveRoleRequest 创建、编辑角色请求参数
type SaveRoleRequest struct {
	Id       uint   `json:"id" v:"checkId"`                                   // 角色ID
	Pid      uint   `json:"pid" v:"checkPid"`                                 // 父级ID
	RoleName string `json:"role_name" v:"required|length:2,50|checkRoleName"` // 角色名称
	Status   uint   `json:"status" v:"required|in:0,1"`                       // 状态:取值说明在模型里
}

// RoleListRequest 角色列表请求参数
type RoleListRequest struct {
	PageRequest       // 分页参数
	Status      *uint `json:"status" v:"in:0,1"` // 状态:取值说明在模型里
}

// SetRoleMenusRequest 设置角色菜单请求参数
type SetRoleMenusRequest struct {
	RoleId  uint   `json:"role_id" v:"required|min:1|checkId"` // 角色ID
	MenuIds []uint `json:"menu_ids" v:"required|checkMenuIds"` // 菜单ID数组
}

// SetRoleApisRequest 设置角色api权限请求参数
type SetRoleApisRequest struct {
	RoleId uint   `json:"role_id" v:"required|min:1|checkId"` // 角色ID
	ApiIds []uint `json:"api_ids" v:"checkApiIds"`            // apiID数组
}

// GetMenuListByRoleRequest 根据角色获取菜单列表
type GetMenuListByRoleRequest struct {
	RoleId uint `json:"role_id" v:"required|checkId"` // 角色ID
}
