package response

// BetAreaResponse 获取下注区域返回参数
type BetAreaResponse struct {
	Id               uint   `json:"id"`                 // ID
	GameCategoryCode string `json:"game_category_code"` // 游戏分类
	AreaId           string `json:"area_id"`            // 区域ID
	AreaKey          string `json:"area_key"`           // 区域标识
	AreaName         string `json:"area_name"`          // 区域名称
	Description      string `json:"description"`        // 下注区域说明
	Odds             string `json:"odds"`               // 赔率
}
