package response

// LangResponse 语言下拉框返回结构体
type LangResponse struct {
	Id          uint   `json:"id"`          // ID
	Lang        string `json:"lang"`        // 语言编码
	Description string `json:"description"` // 备注
}

// ListData 分页返回参数
type ListData struct {
	Total int64 `json:"total"` // 总条数
	List  any   `json:"list"`  // 列表数据
}

// SystemConfigResponse 获取系统配置数据返回参数
type SystemConfigResponse struct {
	LangList           []LangResponse           `json:"lang_list"`            // 语言
	SystemCurrencyList []SystemCurrencyResponse `json:"system_currency_list"` // 系统币种
	BetAreaList        []BetAreaResponse        `json:"bet_area_list"`        // 下注区域
	SystemChipList     []SystemChipResponse     `json:"system_chip_list"`     // 系统筹码
	GameCategoryList   []GameCategoryResponse   `json:"game_category_list"`   // 游戏分类
}

// UploadImgResponse 通用上传图片返回结构体
type UploadImgResponse struct {
	Src string `json:"src"`
}

// AdminOperationLogEventResponse 管理员操作日志事件下拉框
type AdminOperationLogEventResponse struct {
	EventType   string `json:"event_type"`  // 事件类型
	Description string `json:"description"` // 描述
}
