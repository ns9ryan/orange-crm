package response

// TemplatePlanFormListResponse 限红模板列表返回结构体
type TemplatePlanFormListResponse struct {
	Id          uint   `json:"id"`           // ID
	Name        string `json:"name"`         // 模板名称
	PlanFormId  string `json:"plan_form_id"` // 限红模板ID
	CreatedTime string `json:"created_time"` // 创建时间
}

// TemplatePlanFormDetailResponse 限红模板详情返回结构体
type TemplatePlanFormDetailResponse struct {
	Id                uint                   `json:"id"`                  // ID
	Name              string                 `json:"name"`                // 模板名称
	PlanFormId        string                 `json:"plan_form_id"`        // 限红模板ID
	GameTemplatePlans []GameTemplatePlanArea `json:"game_template_plans"` // 游戏限红模板
	Content           any                    `json:"content"`
}

// GameTemplatePlanArea 投注区域数据
type GameTemplatePlanArea struct {
	Id       uint   `json:"id"`        // ID
	GameId   string `json:"game_id"`   // 游戏ID
	PlanId   string `json:"plan_id"`   // 方案ID
	MinValue uint   `json:"min_value"` // 方案的最小投注值
	MaxValue uint   `json:"max_value"` // 方案的最大投注值
	Digits   uint   `json:"digits"`    // 投注的最小单位倍数
}

// TemplatePlanResponse 获取限红模板下拉框数据返回参数
type TemplatePlanResponse struct {
	Id         uint   `json:"id"`           // ID
	PlanFormId string `json:"plan_form_id"` // 限红模板ID
	Name       string `json:"name"`         // 模板名称
}
