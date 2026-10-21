package response

type GameTableInfoResponse struct {
	Id               uint                            `json:"id"`                 // 游戏桌台ID
	GameTableCode    string                          `json:"game_table_code"`    // 游戏桌台标识符
	GameCategoryCode string                          `json:"game_category_code"` // 游戏分类
	CategorySub      string                          `json:"category_sub"`       // 游戏子分类
	DeskNo           uint                            `json:"desk_no"`            // 桌号
	BetInterval      uint                            `json:"bet_interval"`       // 投注间隔时间,秒
	UnbetRounds      uint                            `json:"unbet_rounds"`       // 允许未下注局数,局,0表示不限制
	Status           uint                            `json:"status"`             // 是否启用:取值说明在模型里
	GameServiceName  string                          `json:"game_service_name"`  // 游戏服务名
	StreamUrl        string                          `json:"stream_url"`         // 视频流地址
	Description      string                          `json:"description"`        // 游戏说明
	CreatedTime      string                          `json:"created_time"`       // 创建时间
	UpdatedTime      string                          `json:"updated_time"`       // 更新时间
	CategoryData     GameTableDetailCategoryResponse `json:"category_data"`      // 分类数据
}

// GameTableDetailCategoryResponse 详情返回中的分类结构体
type GameTableDetailCategoryResponse struct {
	Id               uint   `json:"game_category_id"`   // 游戏分类ID
	GameCategoryCode string `json:"game_category_code"` // 游戏分类
	SubCode          string `json:"sub_code"`           // 游戏子分类
	Description      string `json:"description"`        // 游戏分类描述
}

// GameTableListResponse 游戏桌台列表返回结构体
type GameTableListResponse struct {
	GameTableInfoResponse
}

// GameTableDetailResponse 游戏桌台详情返回结构体
type GameTableDetailResponse struct {
	GameTableInfoResponse
}

// GetGameBoxWithGroupResponse 获取游戏桌台,分组下拉框返回参数
type GetGameBoxWithGroupResponse struct {
	GameCategoryCode string                                 `json:"game_category_code"` // 分类
	SubCode          string                                 `json:"sub_code"`           // 子分类
	Description      string                                 `json:"description"`        // 描述
	GameData         []GameBoxResponse                      `json:"game_data"`          // 游戏数据
	TemplatePlanData []GetTemplatePlanBoxByCategoryResponse `json:"template_plan_data"` // 限红方案数据
}

// GameBoxResponse 获取游戏桌台,分组下拉框返回中的游戏数据
type GameBoxResponse struct {
	Id            uint   `json:"id"`              // Id
	GameTableCode string `json:"game_table_code"` // 游戏桌台标识符
	DeskNo        uint   `json:"desk_no"`         // 桌号
	Description   string `json:"description"`     // 游戏说明
}
