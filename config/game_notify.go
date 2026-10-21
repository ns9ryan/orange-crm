package config

type GameNotify struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 请求主机名
	Port uint   `mapstructure:"port" json:"port" yaml:"port"` // 端口号
}
