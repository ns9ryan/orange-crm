package response

// SystemCurrencyResponse 获取系统币种信息返回参数
type SystemCurrencyResponse struct {
	Id          uint   `json:"id"`          // ID
	Code        string `json:"code"`        // 字母编码
	Name        string `json:"name"`        // 名称
	Symbol      string `json:"symbol"`      // 符号
	Digits      uint   `json:"digits"`      // 币种的最小单位倍数
	Number      string `json:"number"`      // 数字编码
	Region      string `json:"region"`      // 国家/地区/组织
	MaxTip      uint   `json:"max_tip"`     // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip      uint   `json:"min_tip"`     // 荷官可接收的最小打赏金额，存储为最小单位
	Description string `json:"description"` // 备注
}
