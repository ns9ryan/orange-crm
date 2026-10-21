package config

// Zap zap日志配置结构体
type Zap struct {
	Dir          string `mapstructure:"dir" json:"dir" yaml:"dir"`                                  // 日志文件夹
	Format       string `mapstructure:"format" json:"format" yaml:"format"`                         // 日志格式：console|json
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	EncodeLevel  string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`       // 编码器类型：LowercaseLevelEncoder 小写编码器(默认)|LowercaseColorLevelEncoder 小写编码器带颜色|CapitalLevelEncoder 大写编码器|CapitalColorLevelEncoder 大写编码器带颜色
	LogInConsole bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"` // 是否输出控制台
}
