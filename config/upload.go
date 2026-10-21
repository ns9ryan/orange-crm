package config

// Upload 上传配置结构体
type Upload struct {
	Img    Img    `mapstructure:"img" json:"img" yaml:"img"`          // 图片
	Avatar Avatar `mapstructure:"avatar" json:"avatar" yaml:"avatar"` // 头像
	Qrcode Qrcode `mapstructure:"qrcode" json:"qrcode" yaml:"qrcode"` // 二维码
	Video  Video  `mapstructure:"video" json:"video" yaml:"video"`    // 视频
}

// Img 图片配置结构体
type Img struct {
	Size     int64    `mapstructure:"size" json:"size" yaml:"size"`                // 允许上传的文件大小
	Ext      []string `mapstructure:"ext" json:"ext" yaml:"ext"`                   // 允许上传的文件格式
	SavePath string   `mapstructure:"save_path" json:"save_path" yaml:"save_path"` // 保存路径
}

// Avatar 头像配置结构体
type Avatar struct {
	Size     int64    `mapstructure:"size" json:"size" yaml:"size"`                // 允许上传的文件大小
	Ext      []string `mapstructure:"ext" json:"ext" yaml:"ext"`                   // 允许上传的文件格式
	SavePath string   `mapstructure:"save_path" json:"save_path" yaml:"save_path"` // 保存路径
}

// Qrcode 二维码配置结构体
type Qrcode struct {
	Size     int64    `mapstructure:"size" json:"size" yaml:"size"`                // 允许上传的文件大小
	Ext      []string `mapstructure:"ext" json:"ext" yaml:"ext"`                   // 允许上传的文件格式
	SavePath string   `mapstructure:"save_path" json:"save_path" yaml:"save_path"` // 保存路径
}

// Video 头像配置结构体
type Video struct {
	Size     int64    `mapstructure:"size" json:"size" yaml:"size"`                // 允许上传的文件大小
	Ext      []string `mapstructure:"ext" json:"ext" yaml:"ext"`                   // 允许上传的文件格式
	SavePath string   `mapstructure:"save_path" json:"save_path" yaml:"save_path"` // 保存路径
}
