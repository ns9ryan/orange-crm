package request

// SaveCurrencyRequest 创建、编辑币种请求参数
type SaveCurrencyRequest struct {
	SystemCurrencyId *uint `json:"system_currency_id" v:"required|min:1|checkSystemCurrencyId"` // 系统币种ID
}

// CurrencyListRequest 币种列表请求参数
type CurrencyListRequest struct {
	PageRequest // 分页参数
}

// SetCurrencyExchangeRateRequest 设置币种汇率请求参数
type SetCurrencyExchangeRateRequest struct {
	Id               uint                          `json:"id" v:"required|checkId"` // ID
	ExchangeRateList []CurrencyExchangeRateRequest `json:"exchange_rate_list"`      // 汇率数据
}

// CurrencyExchangeRateRequest 币种汇率请求参数
type CurrencyExchangeRateRequest struct {
	ToCode       string `json:"to_code" v:"required|checkToCode"` // 目标币种编码
	ExchangeRate string `json:"exchange_rate" v:"required"`       // 汇率
}

// CurrencyExchangeRateHistoryListRequest 币种汇率历史记录请求参数
type CurrencyExchangeRateHistoryListRequest struct {
	PageRequest               // 分页参数
	FromCode           string `json:"from_code"`            // 原币种编码
	ToCode             string `json:"to_code"`              // 目标币种编码
	EffectiveStartTime string `json:"effective_start_time"` // 汇率的生效开始时间
	EffectiveEndTime   string `json:"effective_end_time"`   // 汇率生效结束时间
}
