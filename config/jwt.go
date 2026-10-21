package config

// Jwt jwt鉴权配置结构体
type Jwt struct {
	ExpiresTime              uint   `mapstructure:"expires_time" json:"expires_time" yaml:"expires_time"`                                              // 过期时间(单位：秒)
	SecretKey                string `mapstructure:"secret_key" json:"SecretKey" yaml:"secret_key"`                                                     // 密钥,用于加盐的字符串
	Issuer                   string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                                                                // 发布者
	MerchantTokenExpiresTime uint   `mapstructure:"merchant_token_expires_time" json:"merchant_token_expires_time" yaml:"merchant_token_expires_time"` // 商户token过期时间(单位：秒)
}
