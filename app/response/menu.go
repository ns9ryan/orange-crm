package response

import "orange-crm/app/model"

// MenuListResponse 菜单列表返回结构体
type MenuListResponse struct {
	Id          uint                   `json:"id"`          // 菜单ID
	Pid         uint                   `json:"pid"`         // 父级的菜单ID，0表示没有父级为顶级菜单
	MenuName    string                 `json:"menu_name"`   // 菜单名称
	Status      uint                   `json:"status"`      // 状态:取值说明在模型里
	Sort        uint                   `json:"sort"`        // 排序,值越小越前面，相同的pid才能进行排序
	MenuType    model.MenuType         `json:"menu_type"`   // 类型:取值说明在模型里
	Description string                 `json:"description"` // 描述
	Path        string                 `json:"path"`        // 跳转路径
	Icon        string                 `json:"icon"`        // 图标
	LangData    []MenuListLangResponse `json:"lang_data"`   // 语言
	Children    []MenuListResponse     `json:"children"`    // 子菜单
}

// MenuListLangResponse 菜单列表返回的语言结构体
type MenuListLangResponse struct {
	Lang     string `json:"lang"`      // 语言编码
	MenuName string `json:"menu_name"` // 菜单名称翻译
}

// MenuDetailResponse 菜单详情返回的结构体
type MenuDetailResponse struct {
	Id          uint                   `json:"id"`          // 菜单ID
	Pid         uint                   `json:"pid"`         // 父级的菜单ID，0表示没有父级为顶级菜单
	MenuName    string                 `json:"menu_name"`   // 菜单名称
	Status      uint                   `json:"status"`      // 状态:取值说明在模型里
	Sort        uint                   `json:"sort"`        // 排序,值越小越前面，相同的pid才能进行排序
	MenuType    model.MenuType         `json:"menu_type"`   // 类型:取值说明在模型里
	Description string                 `json:"description"` // 描述
	Path        string                 `json:"path"`        // 跳转路径
	Icon        string                 `json:"icon"`        // 图标
	LangData    []MenuListLangResponse `json:"lang_data"`   // 语言
}
