package request

// GameCategoryListRequest 游戏分类列表请求参数
type GameCategoryListRequest struct {
	PageRequest             // 分页参数
	GameCategoryCode string `json:"game_category_code" v:"max-length:32"` // 游戏分类
}
