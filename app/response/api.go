package response

// ApiListResponse api列表返回的结构体
type ApiListResponse struct {
	Id          uint   `json:"id"`           // ID
	Path        string `json:"path"`         // 路径
	Method      string `json:"method"`       // 方法,"GET", "POST", "PUT", "DELETE"
	Group       string `json:"group"`        // 分组
	IsEditable  uint   `json:"is_editable"`  // 是否可编辑:取值说明在模型里
	IsDeletable uint   `json:"is_deletable"` // 是否可删除:取值说明在模型里
	Description string `json:"description"`  // api简介
}

// ApiDetailResponse api详情返回的结构体
type ApiDetailResponse struct {
	Id          uint   `json:"id"`           // ID
	Path        string `json:"path"`         // 路径
	Method      string `json:"method"`       // 方法,"GET", "POST", "PUT", "DELETE"
	Group       string `json:"group"`        // 分组
	IsEditable  uint   `json:"is_editable"`  // 是否可编辑:取值说明在模型里
	IsDeletable uint   `json:"is_deletable"` // 是否可删除:取值说明在模型里
	Description string `json:"description"`  // api简介
}
