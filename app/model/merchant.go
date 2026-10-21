package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
	"time"
)

// merchantColumns is the data access object for the table or_merchant.
type merchantColumns struct {
	Id             string // ID
	MerchantCode   string // 商户号
	MerchantName   string // 商户名称
	ParentId       string // 上级商户的username
	RootKey        string // 顶级Key
	Username       string // 账号
	Nickname       string // 昵称
	Password       string // 密码
	CommissionRate string // 成数
	RollingRate    string // 洗码比
	Status         string // 是否允许登录:0否,1是
	IsActive       string // 是否激活(第一次登录成功后改成激活):0否,1是
	UseGoogle      string // 是否启用谷歌登录验证:0否,1是
	BetPaused      string // 是否暂停投注:0正常,1暂停
	IsMain         string // 是否主账号:0否,1是
	WalletPaused   string // 是否暂停上下分:0否,1是
	WalletType     string // 钱包类型:0多钱包,1单钱包
	LoginTime      string // 上次登录时间
	LoginIp        string // 登录IP
	Chips          string // 筹码
	ChipsTemplate  string // 筹码模板
	RoleId         string // 角色ID
	TemplateId     string // 商户限红模板ID,创建会员时默认使用
	AesKeySalt     string // aesKey盐
	Md5Salt        string // md5盐
	OrToken        string // 商户api调用验证or_token
	OrTokenKey     string // 生成token用的商户唯一key，永久有效
	WalletUrl      string // 单钱包URL
	WhiteList      string // 白名单IP,逗号分隔
	CreatedTime    string // 创建时间
	UpdatedTime    string // 更新时间
}

// MerchantColumns holds the columns for the table or_merchant.
var MerchantColumns = merchantColumns{
	Id:             "id",
	MerchantCode:   "merchant_code",
	MerchantName:   "merchant_name",
	ParentId:       "parent_id",
	RootKey:        "root_key",
	Username:       "username",
	Nickname:       "nickname",
	Password:       "password",
	CommissionRate: "commission_rate",
	RollingRate:    "rolling_rate",
	Status:         "status",
	IsActive:       "is_active",
	UseGoogle:      "use_google",
	BetPaused:      "bet_paused",
	IsMain:         "is_main",
	WalletPaused:   "wallet_paused",
	WalletType:     "wallet_type",
	LoginTime:      "login_time",
	LoginIp:        "login_ip",
	Chips:          "chips",
	ChipsTemplate:  "chips_template",
	RoleId:         "role_id",
	TemplateId:     "template_id",
	AesKeySalt:     "aes_key_salt",
	Md5Salt:        "md5_salt",
	OrToken:        "or_token",
	OrTokenKey:     "or_token_key",
	WalletUrl:      "wallet_url",
	WhiteList:      "white_list",
	CreatedTime:    "created_time",
	UpdatedTime:    "updated_time",
}

// Merchant 商户主表结构体
type Merchant struct {
	ID           uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	MerchantCode string `gorm:"uniqueindex;size:10;not null;comment:商户号"`
	MerchantName string `gorm:"size:32;not null;comment:商户名称"`
	ParentId     string `gorm:"size:32;not null;default:'';comment:上级商户的username"`
	RootKey      uint   `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`
	Username     string `gorm:"uniqueindex;size:32;not null;comment:账号"`
	Nickname     string `gorm:"size:32;not null;comment:昵称"`
	Password     string `gorm:"not null;comment:密码"`

	CommissionRate uint `gorm:"type:int unsigned;not null;default:0;comment:成数"`
	RollingRate    uint `gorm:"type:int unsigned;not null;default:0;comment:洗码比"`

	// 状态相关
	Status       uint       `gorm:"type:tinyint unsigned;not null;default:1;comment:是否允许登录:0否,1是"`
	IsActive     uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否激活(第一次登录成功后改成激活):0否,1是"`
	UseGoogle    uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否启用谷歌登录验证:0否,1是"`
	BetPaused    uint       `gorm:"type:tinyint unsigned;not null;default:0;comment:是否暂停投注:0正常,1暂停"`
	IsMain       uint       `gorm:"type:tinyint unsigned;not null;default:1;comment:是否主账号:0否,1是"`
	WalletPaused uint       `gorm:"type:tinyint unsigned;not null;default:1;comment:是否暂停上下分:0否,1是"`
	WalletType   WalletType `gorm:"type:tinyint unsigned;not null;default:0;comment:钱包类型:0多钱包,1单钱包"`

	LoginTime *time.Time `gorm:"default:NULL;comment:上次登录时间"`
	LoginIp   string     `gorm:"size:50;comment:登录IP"`

	Chips         string `gorm:"not null;comment:筹码"`
	ChipsTemplate string `gorm:"not null;comment:筹码模板"`
	RoleId        uint   `gorm:"type:int unsigned;not null;default:2;comment:角色ID"`
	TemplateId    string `gorm:"size:36;not null;default:'';comment:商户限红模板ID,创建会员时默认使用"`

	// api相关
	AesKeySalt string `gorm:"not null;comment:aesKey盐"`
	Md5Salt    string `gorm:"not null;comment:md5盐"`
	OrToken    string `gorm:"not null;comment:商户api调用验证or_token"`
	OrTokenKey string `gorm:"not null;comment:生成token用的商户唯一key，永久有效"`
	WalletUrl  string `gorm:"not null;default:'';comment:单钱包URL"`
	WhiteList  string `gorm:"not null;default:'';comment:白名单IP,逗号分隔"`

	BaseModel
	MerchantCurrency MerchantCurrency `gorm:"foreignKey:Username;references:Username"`
	Role             Role
}

// WalletType 钱包类型
type WalletType uint

const (
	SingleWallet WalletType = 1 // 单钱包
	MultiWallet  WalletType = 2 // 多钱包
)

// MerchantChip 商户筹码结构体
type MerchantChip struct {
	Id     uint // Id
	Value  uint // 面额
	Digits uint // 最小单位倍数
}

// MerchantBettingLimit 商户限红结构体
type MerchantBettingLimit struct {
	Id           uint // Id
	MinBetAmount uint // 最小投注金额
	MinDigits    uint // 最小投注的最小单位倍数
	MaxBetAmount uint // 最大投注金额
	MaxDigits    uint // 最大投注的最小单位倍数
}

// RegisterTable 数据迁移
func (m *Merchant) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='商户主表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		//password := "123456"
		//hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		//// 填充数据
		//data := []Merchant{
		//	{ID: 1, MerchantCode: "4OaVRf6e", MerchantName: "测试商户001", ParentId: "root", Username: "merchant001", Nickname: "测试商户001", CommissionRate: 1, RollingRate: 2, Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", ChipsTemplate: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", RoleId: 2, TemplateId: "173r3pru100d973rajw9rl81008h3d2o", AesKeySalt: "A8Kryew4UARph43I", Md5Salt: "mWTAOKu9dvfwQNLf", OrToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOi01NjA1MzIwMDUzLCJpYXQiOjE3NDQ3MDkwOTcsIm1lcmNoYW50X2NvZGUiOiJMamwxc0dYbiJ9.MQOnJ6shTUNmF9NYzdc_abDc4zXgkwJg8Pqizy_ufz0", OrTokenKey: "dbyDgMRB57Ay2AdjCQbizMbZFKMiIw3Z"},
		//
		//	{ID: 3, MerchantCode: "McinSkAh", MerchantName: "测试商户001", ParentId: "merchant001", Username: "sub_merchant001", Nickname: "测试商户001", CommissionRate: 1, RollingRate: 2, Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", ChipsTemplate: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", RoleId: 2, TemplateId: "173r3pru100d973rajw9rl8200on8ujc", AesKeySalt: "A8Kryew4UARph43I", Md5Salt: "mWTAOKu9dvfwQNLf", OrToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOi01NjA1MzIwMDUzLCJpYXQiOjE3NDQ3MDkwOTcsIm1lcmNoYW50X2NvZGUiOiJMamwxc0dYbiJ9.MQOnJ6shTUNmF9NYzdc_abDc4zXgkwJg8Pqizy_ufz0", OrTokenKey: "dbyDgMRB57Ay2AdjCQbizMbZFKMiIw3Z"},
		//
		//	{ID: 4, MerchantCode: "k4jhkcah", MerchantName: "测试商户002", ParentId: "root", Username: "merchant002", Nickname: "测试商户002", CommissionRate: 1, RollingRate: 2, Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", ChipsTemplate: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", RoleId: 2, TemplateId: "173r3pru100d973rdqlv11s300qnrvhf", AesKeySalt: "1LoOGKYfIRfC6CVg", Md5Salt: "AFBrVYLfxonBlYhY", OrToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOi01NjA1MzIwMDQ2LCJpYXQiOjE3NDQ3MDkxMDQsIm1lcmNoYW50X2NvZGUiOiJtWXpkN25oSyJ9.mreCNbvLV7ooo6LmD1Y3ukDy4psqyLYzgnY2eK5lPB0", OrTokenKey: "dY8C7FFnUpvm0lsT90gV0k3Bd8foFP4u"},
		//
		//	{ID: 5, MerchantCode: "ZfxvzvZn", MerchantName: "测试商户002", ParentId: "merchant002", Username: "sub_merchant002", Nickname: "测试商户002", CommissionRate: 1, RollingRate: 2, Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", ChipsTemplate: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26", RoleId: 2, TemplateId: "173r3pru100d973rdqlv11s400f0gox0", AesKeySalt: "1LoOGKYfIRfC6CVg", Md5Salt: "AFBrVYLfxonBlYhY", OrToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOi01NjA1MzIwMDQ2LCJpYXQiOjE3NDQ3MDkxMDQsIm1lcmNoYW50X2NvZGUiOiJtWXpkN25oSyJ9.mreCNbvLV7ooo6LmD1Y3ukDy4psqyLYzgnY2eK5lPB0", OrTokenKey: "dY8C7FFnUpvm0lsT90gV0k3Bd8foFP4u"},
		//}
		//
		//if err := global.DB.Create(&data).Error; err != nil {
		//	return err
		//}

		//// 导入 JSON 文件中的数据
		//if err := utils.ImportAndProcessJSON[Merchant](global.DB, "./data/init/merchant.json", "Merchant", func(m *Merchant) error {
		//	if m.Password == "" {
		//		m.Password = "123456"
		//	}
		//	hash, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
		//	if err != nil {
		//		return err
		//	}
		//	m.Password = string(hash)
		//	return nil
		//}); err != nil {
		//	return err
		//}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *Merchant) AfterDelete(tx *gorm.DB) (err error) {
	// 删除商户币种
	tx.Where(map[string]interface{}{
		MerchantCurrencyColumns.Username: m.Username,
	}).Delete(&MerchantCurrency{})

	// 删除商户登录记录
	tx.Where(map[string]interface{}{
		LoginLogColumns.Username: m.Username,
	}).Delete(&LoginLog{})

	// 删除商户限红模板
	var templates []MerchantTemplateBetlimit
	tx.Where(map[string]interface{}{
		MerchantTemplateBetlimitColumns.Username: m.Username,
	}).Find(&templates)
	for _, tpl := range templates {
		if err = tx.Delete(&tpl).Error; err != nil {
			return err
		}
	}

	return
}
