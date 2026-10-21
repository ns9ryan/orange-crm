package response

// CurrencyListResponse 币种列表返回结构体
type CurrencyListResponse struct {
	Id             uint   `json:"id"`               // 币种ID
	Code           string `json:"code"`             // 币种编码
	Name           string `json:"name"`             // 币种名称
	Symbol         string `json:"symbol"`           // 币种符号
	Digits         uint   `json:"digits"`           // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
	MerchantCount  int64  `json:"merchant_count"`   // 商户数量
	PlayerCount    int64  `json:"player_count"`     // 玩家数量
	DealerCount    int64  `json:"dealer_count"`     // 荷官数量
	MaxTip         uint   `json:"max_tip"`          // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip         uint   `json:"min_tip"`          // 荷官可接收的最小打赏金额，存储为最小单位
	DefaultChipIds []uint `json:"default_chip_ids"` // 默认筹码ID列表
	DefaultTips    []uint `json:"default_tips"`     // 默认打赏面值
}

// CurrencyExchangeRateResponse 币种汇率返回结构体
type CurrencyExchangeRateResponse struct {
	FromCode     string `json:"from_code"`     // 原币种编码
	ToCode       string `json:"to_code"`       // 目标币种编码
	ExchangeRate string `json:"exchange_rate"` // 汇率
}

// CurrencyDetailResponse 币种详情返回结构体
type CurrencyDetailResponse struct {
	Id                       uint                           `json:"id"`                          // 币种ID
	Code                     string                         `json:"code"`                        // 币种编码
	Name                     string                         `json:"name"`                        // 币种名称
	Symbol                   string                         `json:"symbol"`                      // 币种符号
	Digits                   uint                           `json:"digits"`                      // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
	MaxTip                   uint                           `json:"max_tip"`                     // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip                   uint                           `json:"min_tip"`                     // 荷官可接收的最小打赏金额，存储为最小单位
	DefaultChipIds           []uint                         `json:"default_chip_ids"`            // 默认筹码ID列表
	CurrencyExchangeRateList []CurrencyExchangeRateResponse `json:"currency_exchange_rate_list"` // 汇率列表
}

// CurrencyExchangeRateHistoryListResponse 币种汇率历史列表返回结构体
type CurrencyExchangeRateHistoryListResponse struct {
	Id                 uint   `json:"id"`                   // 历史记录ID
	FromCode           string `json:"from_code"`            // 原币种编码
	ToCode             string `json:"to_code"`              // 目标币种编码
	ExchangeRate       string `json:"exchange_rate"`        // 汇率
	EffectiveStartTime string `json:"effective_start_time"` // 汇率的生效开始时间
	EffectiveEndTime   string `json:"effective_end_time"`   // 汇率生效结束时间
	RateSource         uint   `json:"rate_source"`          // 来源:1系统
}
