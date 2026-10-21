package request

// SaveControlAccountRequest 创建、编辑控端账号请求参数
type SaveControlAccountRequest struct {
	Id       uint   `json:"id" v:"checkId"`                                          // ID
	Username string `json:"username" v:"required-if:Id,0|length:2,30|checkUsername"` // 账号,创建的时候必传
	Nickname string `json:"nickname" v:"required|length:2,50"`                       // 昵称
	Level    uint   `json:"level" v:"required|min:1|checkLevel"`                     // 等级
	Status   uint   `json:"status" v:"required|in:0,1"`                              // 是否允许登录:取值说明在模型里
	Password string `json:"password" v:"required-if:Id,0|length:6,30"`               // 密码,创建的时候必传
}

// ControlAccountListRequest 控端账号列表请求参数
type ControlAccountListRequest struct {
	PageRequest // 分页参数
}
