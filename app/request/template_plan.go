package request

// SaveTemplatePlanRequest 创建、编辑限红方案请求参数
type SaveTemplatePlanRequest struct {
	Id                uint               `json:"id"`                                       // ID
	CategorySub       string             `json:"category_sub" v:"required|length:2,32"`    // 游戏子分类
	Name              string             `json:"name" v:"required|length:2,100|checkName"` // 方案名称
	TemplatePlanAreas []TemplatePlanArea `json:"template_plan_areas" v:"required"`         // 投注区域数据
}

// TemplatePlanArea 投注区域数据
type TemplatePlanArea struct {
	AreaId       string `json:"area_id" v:"required"`        // 区域ID
	MinBetAmount uint   `json:"min_bet_amount" v:"required"` // 最小投注金额
	MaxBetAmount uint   `json:"max_bet_amount" v:"required"` // 最大投注金额
}

// TemplatePlanListRequest 限红方案列表请求参数
type TemplatePlanListRequest struct {
	PageRequest // 分页参数
}

// GetTemplatePlanBoxByCategoryRequest 根据子分类获取限红方案的下拉框数据请求参数
type GetTemplatePlanBoxByCategoryRequest struct {
	CategorySub string `json:"category_sub" v:"required|length:2,32"` // 游戏子分类
}
