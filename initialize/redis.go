package initialize

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/redis/go-redis/v9"
	"orange-crm/global"
)

// Redis 初始化redis
func Redis() *redis.Client {
	config := global.CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // 没有密码，默认值
		DB:       config.DB,       // 默认DB 0
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("cache drive init is panic: %s", err))
	}

	gConfig := gredis.Config{
		Address: config.Addr,
		Db:      config.DB,       // 默认DB 0
		Pass:    config.Password, // 没有密码，默认值
	}

	gredis.SetConfig(&gConfig)

	return client
}
