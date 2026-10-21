// =================================================================================
// 模型层-荷官
// =================================================================================

package model

import (
	"fmt"
	"orange-crm/global"
	"orange-crm/utils"
)

// dealerColumns defines and stores column names for the table or_dealer.
type dealerColumns struct {
	Id             string // ID
	Nickname       string // 昵称
	EmployeeNumber string // 工号
	BalanceAmount  string // 账户余额，存储为最小单位
	CurrencyCode   string // 币种代码
	CurrencyName   string // 币种名称
	CurrencySymbol string // 币种符号
	Digits         string // 币种的最小单位倍数
	DealerType     string // 荷官类型:1操作者,2观察者
	Level          string // 荷官级别
	LeadName       string // 上级名称
	Status         string // 状态:待定
	Avatar         string // 头像
	Qrcode         string // qrcode
	CreatedTime    string // 创建时间
	UpdatedTime    string // 更新时间
}

// DealerColumns holds the columns for the table or_dealer.
var DealerColumns = dealerColumns{
	Id:             "id",
	Nickname:       "nickname",
	EmployeeNumber: "employee_number",
	BalanceAmount:  "balance_amount",
	CurrencyCode:   "currency_code",
	CurrencyName:   "currency_name",
	CurrencySymbol: "currency_symbol",
	Digits:         "digits",
	DealerType:     "dealer_type",
	Level:          "level",
	LeadName:       "lead_name",
	Status:         "status",
	Avatar:         "avatar",
	Qrcode:         "qrcode",
	CreatedTime:    "created_time",
	UpdatedTime:    "updated_time",
}

// Dealer 荷官表结构体
type Dealer struct {
	ID             uint        `gorm:"primarykey;type:int unsigned;comment:ID"`
	Nickname       string      `gorm:"size:50;not null;default:'';comment:昵称"`
	EmployeeNumber string      `gorm:"uniqueindex;size:50;not null;comment:工号"`
	BalanceAmount  uint        `gorm:"type:int unsigned;not null;default:0;comment:账户余额，存储为最小单位"`
	CurrencyCode   string      `gorm:"type:varchar(10);not null;comment:币种代码"`
	CurrencyName   string      `gorm:"type:varchar(32);not null;comment:币种名称"`
	CurrencySymbol string      `gorm:"type:varchar(10);not null;comment:币种符号"`
	Digits         uint        `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	DealerType     DealerType  `gorm:"type:tinyint unsigned;not null;comment:荷官类型:1操作者,2观察者"`
	Level          DealerLevel `gorm:"type:int unsigned;not null;default:1;comment:荷官级别"`
	LeadName       string      `gorm:"size:50;not null;default:'';comment:上级名称"`
	Status         uint        `gorm:"type:tinyint unsigned;not null;default:0;comment:状态:待定"`
	Avatar         string      `gorm:"not null;default:'';comment:头像"`
	Qrcode         string      `gorm:"not null;default:'';comment:qrcode"`
	BaseModel
}

// DealerType 荷官类型
type DealerType uint

// DealerLevel 荷官级别
type DealerLevel uint

const (
	DealerType1 DealerType = 1 // 操作者
	DealerType2 DealerType = 2 // 观察者
)

const (
	DealerLevel1 DealerLevel = iota + 1
	DealerLevel2
	DealerLevel3
	DealerLevel4
	DealerLevel5
)

// RegisterTable 数据迁移
func (m *Dealer) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='荷官表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		employees := []string{
			"D00000001",
			"D00000002",
			"D00000003",
			"D00000004",
			"D00000005",
			"D00000006",
			"D00000007",
			"D00000008",
		}
		var data []Dealer
		for k, employee := range employees {
			tmp := Dealer{
				Nickname: fmt.Sprintf("测试荷官%03d", k), EmployeeNumber: employee, CurrencyCode: "USD", CurrencyName: "美元", CurrencySymbol: "$", Digits: 100, DealerType: DealerType1, Level: DealerLevel1,
			}
			// 生成qrcode
			if _, err := utils.MakeQrcode(employee); err != nil {
				return err
			}
			tmp.Qrcode = fmt.Sprintf("/%s/%s", global.CONFIG.Upload.Qrcode.SavePath, fmt.Sprintf("%s.png", employee))

			data = append(data, tmp)
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
