package config

// Server 系统服务配置结构体
type Server struct {
	Addr               int      `mapstructure:"addr" json:"addr" yaml:"addr"`                                                 // http运行端口
	Mode               string   `mapstructure:"mode" json:"mode" yaml:"mode"`                                                 // 运行模式
	GormLogLevel       int      `mapstructure:"gorm_log_level" json:"gorm_log_level" yaml:"gorm_log_level"`                   // gorm日志打印级别：Silent = 1, Error = 2, Warn = 3, Info = 4
	Lang               string   `mapstructure:"lang" json:"lang" yaml:"lang"`                                                 // 默认语言：zh 中文(默认)|en 英语|es 西班牙语|pt 葡萄牙语
	SuperPassword      string   `mapstructure:"super_password" json:"super_password" yaml:"super_password"`                   // 初始超级管理员密码,为空则默认123456
	ResetPassword      string   `mapstructure:"reset_password" json:"reset_password" yaml:"reset_password"`                   // 重置密码,为空则默认123456
	DealerCurrencyCode string   `mapstructure:"dealer_currency_code" json:"dealer_currency_code" yaml:"dealer_currency_code"` // 创建荷官默认币种,为空默认USD
	LimitPlayer        uint     `mapstructure:"limit_player" json:"limit_player" yaml:"limit_player"`                         // 限制会员数量
	OriginPathPrefix   []string `mapstructure:"origin_path_prefix" json:"origin_path_prefix" yaml:"origin_path_prefix"`       // origin路径前缀
}
