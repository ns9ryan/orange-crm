package request

// SaveGameTableRequest 创建、编辑游戏桌台请求参数
type SaveGameTableRequest struct {
	Id              uint   `json:"id" v:"checkId"`                                                      // ID
	GameTableCode   string `json:"game_table_code" v:"required-if:Id,0|length:2,32|checkGameTableCode"` // 游戏桌台标识符
	CategorySub     string `json:"category_sub" v:"required-if:Id,0|length:2,32|checkCategorySub"`      // 游戏子分类
	DeskNo          uint   `json:"desk_no" v:"required-if:Id,0|checkDeskNo"`                            // 桌号
	BetInterval     uint   `json:"bet_interval" v:"min:5|max:1800"`                                     // 投注间隔时间,秒
	UnbetRounds     uint   `json:"unbet_rounds" v:"max:1000"`                                           // 允许未下注局数,局,0表示不限制
	Status          uint   `json:"status" v:"required|in:0,1"`                                          // 是否启用:取值说明在模型里
	GameServiceName string `json:"game_service_name" v:"required-if:Id,0|length:2,32"`                  // 游戏服务名
	StreamUrl       string `json:"stream_url" v:"max-length:255"`                                       // 视频流地址
	Description     string `json:"description" v:"max-length:100"`                                      // 游戏说明
}

// GameTableListRequest 游戏桌台列表请求参数
type GameTableListRequest struct {
	PageRequest            // 分页参数
	CategorySub     string `json:"category_sub"`                        // 游戏子分类
	GameTableCode   string `json:"game_table_code" v:"max-length:32"`   // 游戏桌台标识符
	GameServiceName string `json:"game_service_name" v:"max-length:32"` // 游戏服务名
}
