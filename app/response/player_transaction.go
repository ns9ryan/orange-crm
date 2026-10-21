package response

// CreditAndDebitListResponse 上分/下分记录列表返回参数
type CreditAndDebitListResponse struct {
	ID              uint   `json:"id"`                // ID
	RootKey         uint   `json:"root_key"`          // 顶级Key
	MerchantId      string `json:"merchant_id"`       // 对应商户表的username
	Username        string `json:"username"`          // 玩家账号,每个玩家的唯一值
	Currency        string `json:"currency"`          // 币种
	Digits          uint   `json:"digits"`            // 币种的最小单位倍数
	TransactionType uint   `json:"transaction_type"`  // 交易类型:取值说明在模型里
	WalletField     string `json:"wallet_field"`      // 钱包字段
	Amount          int    `json:"amount"`            // 变动金额,存储为最小单位
	OrangeTransId   string `json:"orange_trans_id"`   // 系统产生的订单号,全局唯一
	MerchantTransId string `json:"merchant_trans_id"` // 商户产生的订单号,api商户传过来的订单号
	Remark          string `json:"remark"`            // 备注
	CreatedTime     string `json:"created_time"`      // 创建时间
}
