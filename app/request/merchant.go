package request

// SaveMerchantRequest 创建、编辑商户请求参数
type SaveMerchantRequest struct {
	Id           uint   `json:"id" v:"checkId"`                                              // ID
	MerchantName string `json:"merchant_name" v:"required-if:Id,0|length:2,30"`              // 商户名称
	Username     string `json:"username" v:"required-if:Id,0|length:2,30|checkUsername"`     // 账号
	SubUsername  string `json:"sub_username" v:"required-if:Id,0|length:2,30|checkUsername"` // 子账号
	Nickname     string `json:"nickname" v:"length:2,50"`                                    // 昵称
	Password     string `json:"password" v:"required-if:Id,0|length:6,30"`                   // 密码

	CommissionRate uint `json:"commission_rate" v:"required|between:0,100000"` // 成数
	RollingRate    uint `json:"rolling_rate" v:"required|between:0,100000"`    // 洗码比

	Status     uint `json:"status" v:"required|in:0,1"`      // 是否允许登录:取值说明在模型里
	WalletType uint `json:"wallet_type" v:"required|in:0,1"` // 钱包类型:取值说明在模型里

	ChipIds           []uint   `json:"chip_ids" v:"required|checkChipIds"`                               // 筹码ID数组
	ChipTemplateIds   []uint   `json:"chip_template_ids" v:"required|checkChipTemplateIds"`              // 筹码模板ID数组
	RoleId            uint     `json:"role_id" v:"required|min:1|checkRoleId"`                           // 角色ID
	PlanFormIds       []string `json:"plan_form_ids" v:"required-if:Id,0|checkPlanFormIds"`              // 系统限红模板ID数组
	DefaultPlanFormId string   `json:"default_plan_form_id" v:"required-if:Id,0|checkDefaultPlanFormId"` // 默认系统限红模板ID
	WalletUrl         string   `json:"wallet_url" v:"length:1,255"`                                      // 单钱包URL
	WhiteList         string   `json:"white_list" v:"length:1,255"`                                      // 白名单IP,逗号分隔

	Currency  string `json:"currency" v:"required-if:Id,0"`   // 币种
	MaxTip    uint   `json:"max_tip" v:"required|min:1"`      // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip    uint   `json:"min_tip" v:"required|min:1"`      // 荷官可接收的最小打赏金额，存储为最小单位
	TipConfig []uint `json:"tip_config" v:"required-if:Id,0"` // 商户打赏配置
}

// AddSubMerchantRequest 创建子商户请求参数
type AddSubMerchantRequest struct {
	ParentId     string `json:"parent_id" v:"required|checkParentId"`            // 上级商户ID
	MerchantName string `json:"merchant_name" v:"required|length:2,30"`          // 商户名称
	Username     string `json:"username" v:"required|length:2,30|checkUsername"` // 账号
	Nickname     string `json:"nickname" v:"length:2,50"`                        // 昵称
	Password     string `json:"password" v:"required|length:6,30"`               // 密码

	CommissionRate uint `json:"commission_rate" v:"required|between:0,100000"` // 成数
	RollingRate    uint `json:"rolling_rate" v:"required|between:1,1000"`      // 洗码比

	Status uint `json:"status" v:"required|in:0,1"` // 是否允许登录:取值说明在模型里

	ChipIds           []uint   `json:"chip_ids" v:"required|checkChipIds"`                       // 筹码ID数组
	ChipTemplateIds   []uint   `json:"chip_template_ids" v:"required|checkChipTemplateIds"`      // 筹码模板ID数组
	RoleId            uint     `json:"role_id" v:"required|min:1|checkRoleId"`                   // 角色ID
	PlanFormIds       []string `json:"plan_form_ids" v:"required|checkPlanFormIds"`              // 系统限红模板ID数组
	DefaultPlanFormId string   `json:"default_plan_form_id" v:"required|checkDefaultPlanFormId"` // 默认系统限红模板ID
	MaxTip            uint     `json:"max_tip" v:"required"`                                     // 荷官可接收的最大打赏金额，存储为最小单位
	MinTip            uint     `json:"min_tip" v:"required"`                                     // 荷官可接收的最小打赏金额，存储为最小单位
}

// MerchantListRequest 商户列表请求参数
type MerchantListRequest struct {
	PageRequest        // 分页参数
	SortRequest        // 排序参数
	Username    string `json:"username" v:"length:2,30"` // 账号
	// todo 验证要改
	WalletType *uint  `json:"wallet_type" v:"in:0,1"` // 钱包类型:取值说明在模型里
	ParentId   string `json:"parent_id"`              // 上级账号
}

// CheckUsernameRequest 判断账号是否可用请求参数
type CheckUsernameRequest struct {
	Username string `json:"username" v:"required|length:2,30"` // 账号
}

// LoginMerchantRequest 商户登录请求参数
type LoginMerchantRequest struct {
	Username string `json:"username" v:"required|length:2,30"` // 账号
	Password string `json:"password" v:"required|length:6,30"` // 密码
}

// EditCurrentMerchantPasswordRequest 编辑当前登录的商户密码请求参数
type EditCurrentMerchantPasswordRequest struct {
	OldPassword string `json:"old_password" v:"required|length:6,30"` // 原密码
	NewPassword string `json:"new_password" v:"required|length:6,30"` // 新密码
}
