package request

// BetAreaListRequest 下注区域列表请求参数
type BetAreaListRequest struct {
	PageRequest             // 分页参数
	GameCategoryCode string `json:"game_category_code"` // 游戏分类
}

// GetBoxByCategoryRequest 根据子分类获取下注区域的下拉框数据请求参数
type GetBoxByCategoryRequest struct {
	CategorySub string `json:"category_sub" v:"required|length:2,32"` // 游戏子分类
}
