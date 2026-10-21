package response

// BetRecordListResponse 投注明细列表返回结构体
type BetRecordListResponse struct {
	ID           uint   `json:"id"`            // 投注明细ID
	RecordId     string `json:"record_id"`     // 投注明细表的ID
	UserId       string `json:"user_id"`       // 帐号的ID
	MerchantId   string `json:"merchant_id"`   // 商户ID
	MerchantCode string `json:"merchant_code"` // 商户码
	ParentId     string `json:"parent_id"`     // 上级的ID
	RootKey      uint   `json:"root_key"`      // 上级的ID
	GameId       string `json:"game_id"`       // 游戏id
	DeskNo       int    `json:"desk_no"`       // 桌号
	Type         string `json:"type"`          // 游戏类型
	SubType      string `json:"sub_type"`      // 游戏类型二级
	Set          int    `json:"set"`           // 桌号
	Round        int    `json:"round"`         // 局号
	RoundId      string `json:"round_id"`      // 游戏局数全局ID
	Currency     string `json:"currency"`      // 货币种类
	BetTime      string `json:"bet_time"`      // 投注时间
	BetMode      int    `json:"bet_mode"`      // 投注的模式
	BetAreaId    int    `json:"bet_area_id"`   // 投注区域ID
	BetAreaName  string `json:"bet_area_name"` // 投注区域名称
	BetAmount    int    `json:"bet_amount"`    // 投注额
	BetEffective int    `json:"bet_effective"` // 有效投注额
	WinLost      int    `json:"win_lost"`      // 胜负金额
	Digits       int    `json:"digits"`        // 投注额的小数位
	Winner       int    `json:"winner"`        // 开奖结果,游戏种类不同
	Value1       int    `json:"value1"`        // 开奖参数1
	Value2       int    `json:"value2"`        // 开奖参数2
	Value3       int    `json:"value3"`        // 开奖参数3
	AnswerTime   string `json:"answer_time"`   // 开奖时间
	Answer       any    `json:"answer"`        // 开奖的详细内容
	CreatedTime  string `json:"created_time"`  // 创建时间
	UpdatedTime  string `json:"updated_time"`  // 更新时间
	Description  string `json:"description"`   // 游戏分类描述
}
