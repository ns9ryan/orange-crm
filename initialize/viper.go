package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"orange-crm/global"
)

// Viper 初始化viper
func Viper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")      // 配置文件路径
	v.SetConfigName("config") // 配置文件名称
	v.SetConfigType("yaml")   // 配置文件扩展名

	// 读取配置
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监视和重新读取配置文件
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err = v.Unmarshal(&global.CONFIG)
		if err != nil {
			panic(fmt.Errorf("unable to decode into struct, %v", err))
		}
	})
	v.WatchConfig()

	err = v.Unmarshal(&global.CONFIG)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return v
}
