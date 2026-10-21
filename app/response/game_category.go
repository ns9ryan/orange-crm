package response

// GameCategoryResponse 获取游戏分类信息返回参数
type GameCategoryResponse struct {
	Id               uint   `json:"id"`                 // ID
	GameCategoryCode string `json:"game_category_code"` // 分类
	SubCode          string `json:"sub_code"`           // 子分类
	Description      string `json:"description"`        // 描述
	Status           uint   `json:"status"`             // 是否启用:取值说明在模型里
	IsHide           uint   `json:"is_hide"`            // 是否隐藏:取值说明在模型里
}
