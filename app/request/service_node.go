package request

// SaveServiceNodeRequest 创建、编辑服务节点请求参数
type SaveServiceNodeRequest struct {
	Id       uint   `json:"id" v:"checkId"`               // ID
	Ip       string `json:"ip" v:"required|length:1,50"`  // IP地址
	Port     uint   `json:"port" v:"required|min:10"`     // 端口
	Disabled uint   `json:"disabled" v:"required|in:0,1"` // 是否禁用:取值说明在模型里
	Note     string `json:"note" v:"max-length:255"`      // 备注
}

// ServiceNodeListRequest 服务节点列表请求参数
type ServiceNodeListRequest struct {
	PageRequest // 分页参数
}
