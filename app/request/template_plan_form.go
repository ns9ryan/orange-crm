package request

// SaveTemplatePlanFormRequest 创建、编辑限红模板请求参数
type SaveTemplatePlanFormRequest struct {
	Id                uint                   `json:"id"`                                       // ID
	Name              string                 `json:"name" v:"required|length:2,100|checkName"` // 模板名称
	GameTemplatePlans []GameTemplatePlanArea `json:"game_template_plans" v:"required"`         // 游戏限红模板
}

// GameTemplatePlanArea 投注区域数据
type GameTemplatePlanArea struct {
	GameId string `json:"game_id" v:"required"` // 游戏ID
	PlanId string `json:"plan_id" v:"required"` // 方案ID
}

// TemplatePlanFormListRequest 限红模板列表请求参数
type TemplatePlanFormListRequest struct {
	PageRequest // 分页参数
}
