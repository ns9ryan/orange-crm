package response

// SystemChipResponse 获取系统筹码信息返回参数
type SystemChipResponse struct {
	Id     uint   `json:"id"`     // ID
	Name   string `json:"name"`   // 筹码名称
	Amount string `json:"amount"` // 面额
	Status uint   `json:"digits"` // 是否启用:0否,1是
}
