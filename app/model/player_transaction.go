package model

import (
	"orange-crm/global"
)

// playerTransactionColumns defines and stores column names for the table or_player_transaction.
type playerTransactionColumns struct {
	Id              string // ID
	RootKey         string // 顶级Key
	MerchantId      string // 对应商户表的username
	Username        string // 玩家账号,每个玩家的唯一值
	Currency        string // 币种
	Digits          string // 币种的最小单位倍数
	TransactionType string // 交易类型:1充值,2提现,3下注,4中奖,5返水,6调整
	WalletField     string // 钱包字段
	Before          string // 变动前余额,存储为最小单位
	Amount          string // 变动金额,存储为最小单位
	After           string // 变动后余额,存储为最小单位
	OrangeTransId   string // 系统产生的订单号,全局唯一
	MerchantTransId string // 商户产生的订单号,api商户传过来的订单号
	ReferenceId     string // 关联的业务ID(如订单号/投注ID),用来以后扩展关联具体的详细信息表
	Remark          string // 备注
	CreatedTime     string // 创建时间
	UpdatedTime     string // 更新时间
}

// PlayerTransactionColumns holds the columns for the table or_player_transaction.
var PlayerTransactionColumns = playerTransactionColumns{
	Id:              "id",
	RootKey:         "root_key",
	MerchantId:      "merchant_id",
	Username:        "username",
	Currency:        "currency",
	Digits:          "digits",
	TransactionType: "transaction_type",
	WalletField:     "wallet_field",
	Before:          "before",
	Amount:          "amount",
	After:           "after",
	OrangeTransId:   "orange_trans_id",
	MerchantTransId: "merchant_trans_id",
	ReferenceId:     "reference_id",
	Remark:          "remark",
	CreatedTime:     "created_time",
	UpdatedTime:     "updated_time",
}

// PlayerTransaction 玩家交易流水表结构体
type PlayerTransaction struct {
	ID              uint            `gorm:"primarykey;type:int unsigned;comment:ID"`
	RootKey         uint            `gorm:"type:int unsigned;not null;default:0;comment:顶级Key"`
	MerchantId      string          `gorm:"size:32;not null;comment:对应商户表的username"`
	Username        string          `gorm:"index;not null;comment:玩家账号,每个玩家的唯一值"`
	Currency        string          `gorm:"type:varchar(10);not null;comment:币种"`
	Digits          uint            `gorm:"type:int unsigned;not null;comment:币种的最小单位倍数"`
	TransactionType TransactionType `gorm:"type:tinyint unsigned;not null;comment:交易类型:1充值,2提现,3下注,4中奖,5返水,6调整"`
	WalletField     string          `gorm:"type:varchar(64);not null;comment:钱包字段"`
	Before          uint            `gorm:"type:int unsigned;not null;default:0;comment:变动前余额,存储为最小单位"`
	Amount          int             `gorm:"type:int;not null;default:0;comment:变动金额,存储为最小单位"`
	After           uint            `gorm:"type:int unsigned;not null;default:0;comment:变动后余额,存储为最小单位"`
	OrangeTransId   string          `gorm:"size:32;not null;comment:系统产生的订单号,全局唯一"`
	MerchantTransId string          `gorm:"size:32;not null;default:'';comment:商户产生的订单号,api商户传过来的订单号"`
	ReferenceId     string          `gorm:"size:32;not null;default:'';comment:关联的业务ID(如订单号/投注ID),用来以后扩展关联具体的详细信息表"`
	Remark          string          `gorm:"not null;default:'';comment:备注"`
	BaseModel
	Player Player `gorm:"foreignKey:Username;references:Username"`
}

// TransactionType 交易类型
type TransactionType uint

const (
	TransactionDeposit  TransactionType = 1 // 存款上分
	TransactionWithdraw TransactionType = 2 // 取现下分
	TransactionBet      TransactionType = 3 // 下注
	TransactionWin      TransactionType = 4 // 中奖
	TransactionRebate   TransactionType = 5 // 返水
	TransactionAdjust   TransactionType = 6 // 调整
)

// RegisterTable 数据迁移
func (m *PlayerTransaction) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='玩家交易流水表'").Migrator().CreateTable(m); err != nil {
			return err
		}
	}
	return nil
}
