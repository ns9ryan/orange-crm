package response

// MerchantTemplateBetlimitListResponse 商户限红模板列表返回结构体
type MerchantTemplateBetlimitListResponse struct {
	Id          uint   `json:"id"`           // ID
	Username    string `json:"username"`     // 商户账号
	Name        string `json:"name"`         // 模板名称
	TemplateId  string `json:"template_id"`  // 商户限红模板ID
	IsDefault   uint   `json:"is_default"`   // 是否默认模板:取值说明在模型里
	CreatedTime string `json:"created_time"` // 创建时间
}

// MerchantTemplateBetlimitDetailResponse 限红模板详情返回结构体
type MerchantTemplateBetlimitDetailResponse struct {
	Id                uint                           `json:"id"`                  // ID
	Username          string                         `json:"username"`            // 商户账号
	Name              string                         `json:"name"`                // 模板名称
	TemplateId        string                         `json:"template_id"`         // 商户限红模板ID
	IsDefault         uint                           `json:"is_default"`          // 是否默认模板:取值说明在模型里
	GameTemplatePlans []MerchantGameTemplatePlanArea `json:"game_template_plans"` // 游戏限红模板
	Content           any                            `json:"content"`
}

// MerchantGameTemplatePlanArea 投注区域数据
type MerchantGameTemplatePlanArea struct {
	Id       uint   `json:"id"`        // ID
	GameId   string `json:"game_id"`   // 游戏ID
	PlanId   string `json:"plan_id"`   // 方案ID
	MinValue uint   `json:"min_value"` // 方案的最小投注值
	MaxValue uint   `json:"max_value"` // 方案的最大投注值
	Digits   uint   `json:"digits"`    // 投注的最小单位倍数
}

// GetTemplateBoxByMerchantResponse 获取商户限红模板下拉框数据返回结构体
type GetTemplateBoxByMerchantResponse struct {
	Id         uint   `json:"id"`          // ID
	Username   string `json:"username"`    // 商户账号
	Name       string `json:"name"`        // 模板名称
	TemplateId string `json:"template_id"` // 商户限红模板ID
	IsDefault  uint   `json:"isDefault"`   // 是否默认模板:取值说明在模型里
}
