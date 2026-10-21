package response

// TemplatePlanListResponse 限红方案列表返回结构体
type TemplatePlanListResponse struct {
	Id                  uint   `json:"id"`                   // ID
	CategorySub         string `json:"category_sub"`         // 游戏子分类
	CategoryDescription string `json:"category_description"` // 分类描述
	Name                string `json:"name"`                 // 方案名称
	PlanId              string `json:"plan_id"`              // 限红方案ID
	MinValue            uint   `json:"min_value"`            // 方案的最小投注值
	MaxValue            uint   `json:"max_value"`            // 方案的最大投注值
	Digits              uint   `json:"digits"`               // 方案投注的最小单位倍数
	UsedCount           uint   `json:"used_count"`           // 已被使用次数
}

// TemplatePlanDetailResponse 限红方案详情返回结构体
type TemplatePlanDetailResponse struct {
	Id                uint                       `json:"id"`                  // ID
	CategorySub       string                     `json:"category_sub"`        // 游戏子分类
	Name              string                     `json:"name"`                // 方案名称
	PlanId            string                     `json:"plan_id"`             // 限红方案ID
	MinValue          uint                       `json:"min_value"`           // 方案的最小投注值
	MaxValue          uint                       `json:"max_value"`           // 方案的最大投注值
	Digits            uint                       `json:"digits"`              // 方案投注的最小单位倍数
	TemplatePlanAreas []TemplatePlanAreaResponse `json:"template_plan_areas"` // 投注区域
}

// TemplatePlanAreaResponse 详情返回中的投注区域
type TemplatePlanAreaResponse struct {
	Id           uint   `json:"id"`             // ID
	AreaId       string `json:"area_id"`        // 区域ID
	MinBetAmount uint   `json:"min_bet_amount"` // 最小投注金额
	MaxBetAmount uint   `json:"max_bet_amount"` // 最大投注金额
	Digits       uint   `json:"digits"`         // 投注的最小单位倍数
	PlanId       string `json:"plan_id"`        // 限红方案ID
}

// GetTemplatePlanBoxByCategoryResponse 根据子分类获取下注区域的下拉框数据
type GetTemplatePlanBoxByCategoryResponse struct {
	Id          uint   `json:"id"`           // ID
	CategorySub string `json:"category_sub"` // 游戏子分类
	Name        string `json:"name"`         // 方案名称
	PlanId      string `json:"plan_id"`      // 限红方案ID
	MinValue    uint   `json:"min_value"`    // 方案的最小投注值
	MaxValue    uint   `json:"max_value"`    // 方案的最大投注值
	Digits      uint   `json:"digits"`       // 方案投注的最小单位倍数
}
