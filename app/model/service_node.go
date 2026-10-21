package model

import (
	"orange-crm/global"
)

// serviceNodeColumns defines and stores column names for the table or_service_node.
type serviceNodeColumns struct {
	Id              string // ID
	Ip              string // IP地址
	Port            string // 端口
	OnlineCount     string // 当前在线玩家数（针对连接数）
	Disabled        string // 是否禁用:0否,1是
	Note            string // 备注
	ServiceKey      string // 服务标识，例如 login_ws 或 wallet_tcp 或 chat_ws等
	ServiceConnType string // 服务类型，ws=WebSocket，tcp=TCP
	CreatedTime     string // 创建时间
	UpdatedTime     string // 更新时间
}

// ServiceNodeColumns holds the columns for the table or_service_node.
var ServiceNodeColumns = serviceNodeColumns{
	Id:              "id",
	Ip:              "ip",
	Port:            "port",
	OnlineCount:     "online_count",
	Disabled:        "disabled",
	Note:            "note",
	ServiceKey:      "service_key",
	ServiceConnType: "service_conn_type",
	CreatedTime:     "created_time",
	UpdatedTime:     "updated_time",
}

// ServiceNode  服务节点池表结构体
type ServiceNode struct {
	ID              uint                  `gorm:"primarykey;type:int unsigned;comment:ID"`
	Ip              string                `gorm:"size:50;comment:IP地址"`
	Port            uint                  `gorm:"type:int unsigned;not null;comment:端口"`
	OnlineCount     uint                  `gorm:"type:int unsigned;not null;default:0;comment:当前在线玩家数（针对连接数）"`
	Disabled        uint                  `gorm:"type:tinyint unsigned;not null;default:0;comment:是否禁用:0否,1是"`
	Note            string                `gorm:"not null;default:'';comment:备注"`
	ServiceKey      NodeServiceKey        `gorm:"size:64;not null;default:'login_ws';comment:服务标识，例如 login_ws 或 wallet_tcp 或 chat_ws等"`
	ServiceConnType ServiceConnectionType `gorm:"size:10;not null;default:'ws';comment:服务类型，ws=WebSocket，tcp=TCP"`
	BaseModel
}

type NodeServiceKey string

const (
	NodeServiceKeyLoginWS   NodeServiceKey = "login_ws"
	NodeServiceKeyWalletTCP NodeServiceKey = "wallet_tcp"
)

// ServiceConnectionType 定义具体针对服务节点连接类型的自定义类型
type ServiceConnectionType string

// 定义常量
const (
	ServiceConnectionTypeWS  ServiceConnectionType = "ws"
	ServiceConnectionTypeTCP ServiceConnectionType = "tcp"
)

// RegisterTable 数据迁移
func (m *ServiceNode) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='服务节点池表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 填充数据
		data := []ServiceNode{
			{Ip: "t2.ds16688.com", Port: 9444, Note: "gate1", ServiceKey: NodeServiceKeyLoginWS, ServiceConnType: ServiceConnectionTypeWS},
			{Ip: "t2.ds16688.com", Port: 9445, Note: "gate2", ServiceKey: NodeServiceKeyLoginWS, ServiceConnType: ServiceConnectionTypeWS},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
