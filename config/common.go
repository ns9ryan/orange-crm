package config

type Common struct {
	Server          Server     `mapstructure:"server" json:"server" yaml:"server"`                                  // 系统服务配置
	Zap             Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`                                           // zap日志配置
	Mysql           Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`                                     // mysql配置
	Redis           Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`                                     // redis配置
	Jwt             Jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                                           // jwt鉴权配置
	Upload          Upload     `mapstructure:"upload" json:"upload" yaml:"upload"`                                  // 上传配置
	GameServiceName []string   `mapstructure:"game_service_name" json:"game_service_name" yaml:"game_service_name"` // 服务名称
	GameNotify      GameNotify `mapstructure:"game_notify" json:"game_notify" yaml:"game_notify"`                   // 游戏通知配置
}
