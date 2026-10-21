package request

// GameAnswerListRequest 游戏开答列表请求参数
type GameAnswerListRequest struct {
	PageRequest             // 分页参数
	SortRequest             // 排序参数
	GameCategoryCode string `json:"game_category_code"`                     // 游戏分类
	GameTableCode    string `json:"game_table_code" v:"required"`           // 游戏id
	DeskNo           string `json:"desk_no"`                                // 桌号
	Set              uint   `json:"set"`                                    // 场次
	Dealer           string `json:"dealer"`                                 // 荷管
	StartTime        string `json:"start_time" v:"date-format:Y-m-d H:i:s"` // 开始时间
	EndTime          string `json:"end_time" v:"date-format:Y-m-d H:i:s"`   // 结束时间
}
