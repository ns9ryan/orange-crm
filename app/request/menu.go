package request

// SaveMenuRequest 创建、编辑菜单请求参数
type SaveMenuRequest struct {
	Id          uint       `json:"id" v:"checkId"`                                   // 菜单ID
	Pid         uint       `json:"pid" v:"checkPid"`                                 // 父级ID
	MenuName    string     `json:"menu_name" v:"required|length:2,50|checkMenuName"` // 菜单名称
	Status      uint       `json:"status" v:"required|in:0,1"`                       // 状态:取值说明在模型里
	Sort        uint       `json:"sort" v:"required|min:1"`                          // 排序,值越小越前面
	MenuType    uint       `json:"menu_type" v:"required|in:1,2"`                    // 类型:取值说明在模型里
	Description string     `json:"description" v:"max-length:255"`                   // 描述
	Path        string     `json:"path" v:"max-length:255"`                          // 跳转路径
	Icon        string     `json:"icon" v:"max-length:255"`                          // 图标
	LangData    []LangData `json:"lang_data" v:""`                                   // 语言
}

// LangData 语言包结构体
type LangData struct {
	Lang     string `json:"lang" v:"required|length:2,10|checkLang"` // 语言代码
	MenuName string `json:"menu_name" v:"required|max-length:255"`   // 菜单名称翻译
}

// MenuListRequest 菜单列表请求参数
type MenuListRequest struct {
	PageRequest       // 分页参数
	Status      *uint `json:"status" v:"in:0,1"` // 状态:取值说明在模型里
}
