package request

// SaveMerchantTemplateBetlimitRequest 创建、编辑商户限红模板请求参数
type SaveMerchantTemplateBetlimitRequest struct {
	Id                uint                       `json:"id"`                                // ID
	Username          string                     `json:"username" v:"required|length:2,30"` // 商户账号
	Name              string                     `json:"name" v:"required|length:2,100"`    // 模板名称
	IsDefault         uint                       `json:"is_default" v:"required|in:0,1"`    // 是否默认模板:取值说明在模型里
	GameTemplatePlans []MerchantGameTemplatePlan `json:"game_template_plans" v:"required"`  // 游戏限红模板
}

// MerchantGameTemplatePlan 商户游戏限红模板数据
type MerchantGameTemplatePlan struct {
	GameId string `json:"game_id" v:"required"` // 游戏ID
	PlanId string `json:"plan_id" v:"required"` // 方案ID
}

// MerchantTemplateBetlimitListRequest 限红模板列表请求参数
type MerchantTemplateBetlimitListRequest struct {
	PageRequest        // 分页参数
	Username    string `json:"username"` // 商户账号
}

// GetTemplateBoxByMerchantRequest 获取商户限红模板下拉框数据请求参数
type GetTemplateBoxByMerchantRequest struct {
	Username string `json:"username" v:"required"` // 商户账号
}
