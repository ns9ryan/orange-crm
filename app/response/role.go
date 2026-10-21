package response

// RoleListResponse 角色列表返回结构体
type RoleListResponse struct {
	Id          uint               `json:"id"`           // 角色ID
	Pid         uint               `json:"pid"`          // 角色父级ID
	RoleName    string             `json:"role_name"`    // 角色名称
	Status      uint               `json:"status"`       // 状态:取值说明在模型里
	IsEditable  uint               `json:"is_editable"`  // 是否可编辑:取值说明在模型里
	IsDeletable uint               `json:"is_deletable"` // 是否可删除:取值说明在模型里
	Children    []RoleListResponse `json:"children"`     // 子角色
}

// RoleDetailResponse 角色详情返回结构体
type RoleDetailResponse struct {
	Id          uint   `json:"id"`           // 角色ID
	Pid         uint   `json:"pid"`          // 角色父级ID
	RoleName    string `json:"role_name"`    // 角色名称
	Status      uint   `json:"status"`       // 状态:取值说明在模型里
	IsEditable  uint   `json:"is_editable"`  // 是否可编辑:取值说明在模型里
	IsDeletable uint   `json:"is_deletable"` // 是否可删除:取值说明在模型里
	MenuIds     []uint `json:"menu_ids"`     // 角色的菜单Id数组
	ApiIds      []uint `json:"api_ids"`      // 角色的apiId数组
}

// GetMenuListByRoleResponse 根据角色获取菜单列表返回结构体
type GetMenuListByRoleResponse struct {
	Id           uint                        `json:"id"`                 // 菜单ID
	Pid          uint                        `json:"pid"`                // 父级的菜单ID，0表示没有父级为顶级菜单
	MenuName     string                      `json:"menu_name"`          // 菜单名称
	Status       uint                        `json:"status"`             // 状态:取值说明在模型里
	Sort         uint                        `json:"sort"`               // 排序,值越小越前面，相同的pid才能进行排序
	MenuType     uint                        `json:"menu_type"`          // 类型:取值说明在模型里
	Description  string                      `json:"description"`        // 描述
	Path         string                      `json:"path"`               // 跳转路径
	Icon         string                      `json:"icon"`               // 图标
	LangData     string                      `json:"-"`                  // 语言
	LangDataJson []map[string]any            `json:"lang_data" gorm:"-"` // 语言
	Children     []GetMenuListByRoleResponse `json:"children" gorm:"-"`  // 子菜单
}
