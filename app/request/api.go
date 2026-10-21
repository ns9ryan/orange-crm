package request

// SaveApiRequest 创建、编辑api请求参数
type SaveApiRequest struct {
	Id          uint   `json:"id" v:"checkId"`                                                                       // ID
	Path        string `json:"path" v:"required|regex:^/([a-zA-Z0-9_-]+/)*[a-zA-Z0-9_-]+$|max-length:256|checkPath"` // 路劲
	Method      string `json:"method" v:"required|checkMethod"`                                                      // 请求方法
	Group       string `json:"group" v:"required|max-length:255"`                                                    // 分组
	Description string `json:"description" v:"max-length:256"`                                                       // 描述
	HasDbWrite  uint   `json:"has_db_write" v:"required|in:0,1"`                                                     // 是否对数据库有新增或编辑操作:取值说明在模型里
}

// ApiListRequest api列表请求参数
type ApiListRequest struct {
	PageRequest // 分页参数
}
