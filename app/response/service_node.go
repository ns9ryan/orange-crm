package response

// ServiceNodeListResponse 服务节点列表返回结构体
type ServiceNodeListResponse struct {
	Id          uint   `json:"id"`           // ID
	Ip          string `json:"ip"`           // IP地址
	Port        uint   `json:"port"`         // 端口
	OnlineCount uint   `json:"online_count"` // 当前在线玩家数
	Disabled    uint   `json:"disabled"`     // 是否禁用:取值说明在模型里
	Note        string `json:"note"`         // 备注
}

// ServiceNodeDetailResponse 服务节点详情返回结构体
type ServiceNodeDetailResponse struct {
	Id          uint   `json:"id"`           // ID
	Ip          string `json:"ip"`           // IP地址
	Port        uint   `json:"port"`         // 端口
	OnlineCount uint   `json:"online_count"` // 当前在线玩家数
	Disabled    uint   `json:"disabled"`     // 是否禁用:取值说明在模型里
	Note        string `json:"note"`         // 备注
}
