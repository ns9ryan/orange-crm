package request

// AddPlayersRequest 批量添加玩家请求参数
type AddPlayersRequest struct {
	Start          uint   `json:"start" v:"required"`
	End            uint   `json:"end" v:"required"`
	MerchantId     string `json:"merchant_id"`     // 商户账号
	UsernamePrefix string `json:"username_prefix"` // 账号前缀
}

// PlayersDepositRequest 批量存款上分请求参数
type PlayersDepositRequest struct {
	Amount uint `json:"amount" v:"required"`
}

// AddMerchantsRequest 批量添加商户请求参数
type AddMerchantsRequest struct {
	AddData []AddMerchantsInit `json:"add_data" v:"required"`
}
type AddMerchantsInit struct {
	ChipTemplate []uint                `json:"chip_template" v:"required"` // 筹码模板
	Currency     string                `json:"currency" v:"required"`      // 币种
	Username     string                `json:"username" v:"required"`      // 账号
	Password     string                `json:"password" v:"required"`      // 密码
	Sons         []AddMerchantsInitSon `json:"sons"`
}
type AddMerchantsInitSon struct {
	MerchantName string `json:"merchant_name"`         // 商户名称
	Username     string `json:"username"`              // 账号
	Password     string `json:"password" v:"required"` // 密码
}
