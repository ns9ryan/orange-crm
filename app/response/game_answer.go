package response

// GameAnswerListResponse 游戏开答列表返回结构体
type GameAnswerListResponse struct {
	ID          uint   `json:"id"`           // 投注明细ID
	GameId      string `json:"game_id"`      // 游戏id
	DeskNo      int    `json:"desk_no"`      // 桌号
	Type        string `json:"type"`         // 游戏类型
	Set         int    `json:"set"`          // 桌号
	Round       int    `json:"round"`        // 局号
	RoundId     string `json:"round_id"`     // 游戏局数全局ID
	Win         int    `json:"win"`          // 开奖结果
	Value1      int    `json:"value1"`       // 开奖参数1
	Value2      int    `json:"value2"`       // 开奖参数2
	Value3      int    `json:"value3"`       // 开奖参数3
	Dealer      string `json:"dealer"`       // 开奖的荷管
	AnswerTime  string `json:"answer_time"`  // 开奖时间
	Cards       any    `json:"cards"`        // 牌的内容
	CreatedTime string `json:"created_time"` // 创建时间
	UpdatedTime string `json:"updated_time"` // 更新时间
	Description string `json:"description"`  // 游戏分类描述
}
