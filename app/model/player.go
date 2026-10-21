package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
	"time"
)

// playerColumns defines and stores column names for the table or_player.
type playerColumns struct {
	Id                 string // ID
	RootKey            string // 顶级Key
	MerchantId         string // 对应商户表的username
	Currency           string // 币种
	MerchantCurrencyId string // 商户币种唯一标识
	ApiUser            string // 存储api会员原始账号
	Username           string // 账号
	Nickname           string // 昵称
	ParentId           string // 上级玩家的的username
	LoginTimeFirst     string // 首次次登录时间
	LoginIpFirst       string // 首次次登录IP
	LoginTime          string // 上次登录时间
	LoginIp            string // 上次登录IP
	Status             string // 是否允许登录:0冻结,1正常
	BetPaused          string // 是否暂停投注:0正常,1暂停
	CreateType         string // 玩家来源:1后台创建,2API创建
	MaxWin             string // 最高赢额,0表示不限制
	DeviceType         string // 设备类型:0桌面H5横屏,1手机H5竖屏
	IsOnline           string // 是否在线:0否,1是
	PlayerType         string // 玩家类型:1现金代理,2佣金代理
	MoneyType          string // 玩家模式类型:1真钱,2试玩
	Level              string // 玩家级别
	Email              string // 邮箱
	Phone              string // 手机号
	Password           string // 密码
	Chips              string // 筹码
	TemplateId         string // 商户限红模板ID
	CreatedTime        string // 创建时间
	UpdatedTime        string // 更新时间
}

// PlayerColumns holds the columns for the table or_player.
var PlayerColumns = playerColumns{
	Id:                 "id",
	RootKey:            "root_key",
	MerchantId:         "merchant_id",
	Currency:           "currency",
	MerchantCurrencyId: "merchant_currency_id",
	ApiUser:            "api_user",
	Username:           "username",
	Nickname:           "nickname",
	ParentId:           "parent_id",
	LoginTimeFirst:     "login_time_first",
	LoginIpFirst:       "login_ip_first",
	LoginTime:          "login_time",
	LoginIp:            "login_ip",
	Status:             "status",
	BetPaused:          "bet_paused",
	CreateType:         "create_type",
	MaxWin:             "max_win",
	DeviceType:         "device_type",
	IsOnline:           "is_online",
	PlayerType:         "player_type",
	MoneyType:          "money_type",
	Level:              "level",
	Email:              "email",
	Phone:              "phone",
	Password:           "password",
	Chips:              "chips",
	TemplateId:         "template_id",
	CreatedTime:        "created_time",
	UpdatedTime:        "updated_time",
}

// Player 玩家表结构体
type Player struct {
	ID                 uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	RootKey            uint   `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`
	MerchantId         string `gorm:"size:32;not null;comment:对应商户表的username"`
	Currency           string `gorm:"type:varchar(10);not null;comment:币种"`
	MerchantCurrencyId string `gorm:"size:50;not null;comment:商户币种唯一标识"`
	ApiUser            string `gorm:"not null;comment:存储api会员原始账号"`
	Username           string `gorm:"uniqueindex;size:32;not null;comment:账号"`
	Nickname           string `gorm:"size:32;not null;default:'';comment:昵称"`
	ParentId           string `gorm:"size:32;not null;default:'';comment:上级玩家的的username"`

	LoginTimeFirst *time.Time `gorm:"default:NULL;comment:首次次登录时间"`
	LoginIpFirst   string     `gorm:"size:45;comment:首次次登录IP"`
	LoginTime      *time.Time `gorm:"default:NULL;comment:上次登录时间"`
	LoginIp        string     `gorm:"size:45;comment:上次登录IP"`

	Status     uint             `gorm:"type:tinyint unsigned;not null;default:1;comment:是否允许登录:0冻结,1正常"`
	BetPaused  uint             `gorm:"type:tinyint unsigned;not null;default:0;comment:是否暂停投注:0正常,1暂停"`
	CreateType PlayerCreateType `gorm:"type:tinyint unsigned;not null;default:1;comment:玩家来源:1后台创建,2API创建"`
	MaxWin     uint             `gorm:"type:int unsigned;not null;default:0;comment:最高赢额,0表示不限制"`
	DeviceType PlayerDeviceType `gorm:"type:tinyint unsigned;not null;default:0;comment:设备类型:0桌面H5横屏,1手机H5竖屏"`
	IsOnline   uint             `gorm:"type:tinyint unsigned;not null;default:0;comment:是否在线:0否,1是"`
	PlayerType PlayerType       `gorm:"type:tinyint unsigned;not null;default:1;comment:玩家类型:1现金代理,2佣金代理"`
	MoneyType  PlayerMoneyType  `gorm:"type:tinyint unsigned;not null;default:1;comment:玩家模式类型:1真钱,2试玩"`
	Level      PlayerLevel      `gorm:"type:tinyint unsigned;not null;default:1;comment:玩家级别"`

	Email      string `gorm:"size:100;not null;default:'';comment:邮箱"`
	Phone      string `gorm:"size:20;not null;default:'';comment:手机号"`
	Password   string `gorm:"not null;default:'';comment:密码"`
	Chips      string `gorm:"not null;comment:筹码"`
	TemplateId string `gorm:"size:36;not null;comment:商户限红模板ID"`

	BaseModel
	Merchant     Merchant     `gorm:"foreignKey:MerchantId;references:Username"`
	PlayerWallet PlayerWallet `gorm:"foreignKey:Username;references:Username"`
}

// PlayerType 玩家类型
type PlayerType uint

// PlayerLevel 玩家级别
type PlayerLevel uint

// PlayerCreateType 玩家来源
type PlayerCreateType uint

// PlayerMoneyType 玩家模式类型（真钱/试玩）
type PlayerMoneyType uint

// PlayerDeviceType 设备类型
type PlayerDeviceType uint

const (
	CashAgent       PlayerType = 1 // 现金代理
	CommissionAgent PlayerType = 2 // 佣金代理

	PlayerBackendCreate PlayerCreateType = 1 // 后台创建
	PlayerAPICreate     PlayerCreateType = 2 // API创建

	RealMoney PlayerMoneyType = 1 // 真钱
	Trial     PlayerMoneyType = 2 // 试玩

	DesktopH5Horizontal PlayerDeviceType = 0 // 桌面H5横屏
	MobileH5Vertical    PlayerDeviceType = 1 // 手机H5竖屏

	PlayerLevel1 PlayerLevel = iota + 1
	PlayerLevel2
	PlayerLevel3
	PlayerLevel4
	PlayerLevel5
)

// RegisterTable 数据迁移
func (m *Player) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='玩家主表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		//password := "123456"
		//hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		//// 填充数据
		//data := []Player{
		//	{ID: 1, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player001", Username: "player001", Nickname: "玩家001", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 2, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player002", Username: "player002", Nickname: "玩家002", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 3, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player003", Username: "player003", Nickname: "玩家003", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 4, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player004", Username: "player004", Nickname: "玩家004", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 5, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player005", Username: "player005", Nickname: "玩家005", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 6, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player006", Username: "player006", Nickname: "玩家006", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 7, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player007", Username: "player007", Nickname: "玩家007", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 8, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player008", Username: "player008", Nickname: "玩家008", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 9, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "sub_merchant001_USD", ApiUser: "player009", Username: "player009", Nickname: "玩家009", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//	{ID: 10, MerchantId: "merchant001", Currency: "USD", MerchantCurrencyId: "merchant001_USD", ApiUser: "player010", Username: "player010", Nickname: "玩家010", Password: string(hash), Chips: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", TemplateId: "9ou4220a780d8wogjw4v20o200z6n8bc"},
		//}
		//
		//if err := global.DB.Create(&data).Error; err != nil {
		//	return err
		//}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *Player) AfterDelete(tx *gorm.DB) (err error) {
	// 删除商户币种
	tx.Where(map[string]interface{}{
		PlayerWalletColumns.Username: m.Username,
	}).Delete(&PlayerWallet{})

	// 删除交易流水
	tx.Where(map[string]interface{}{
		PlayerWalletColumns.Username: m.Username,
	}).Delete(&PlayerTransaction{})

	return
}
