package config

import "fmt"

// Mysql mysql配置结构体
type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`    // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // 密码
	Hostname string `mapstructure:"hostname" json:"hostname" yaml:"hostname"`    // 服务器地址
	HostPort int    `mapstructure:"host_port" json:"host_port" yaml:"host_port"` // 数据库端口
	Database string `mapstructure:"database" json:"database" yaml:"database"`    // 数据库名
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`          // 数据库表前缀
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Hostname,
		m.HostPort,
		m.Database,
	)
}
