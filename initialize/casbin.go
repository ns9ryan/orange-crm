package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"orange-crm/global"
	"strings"
)

// CasbinModel 访问控制模型模板
const CasbinModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && MyMatch(r.obj, p.obj) && r.act == p.act
`

func Casbin() *casbin.CachedEnforcer {

	// 创建casbin的gorm适配器指定表前缀
	a, err := gormadapter.NewAdapterByDBUseTableName(global.DB, global.CONFIG.Mysql.Prefix, "")
	if err != nil {
		panic(err)
	}

	// 创建访问控制模型
	m, err := model.NewModelFromString(CasbinModel)
	if err != nil {
		panic(err)
	}

	// 创建带缓存的执行器
	cacheEnforce, err := casbin.NewCachedEnforcer(m, a)
	if err != nil {
		panic(err)
	}

	// 设置过期时间 1小时
	cacheEnforce.SetExpireTime(60 * 60)

	// 自定义函数
	cacheEnforce.AddFunction("MyMatch", KeyMatchFunc)

	return cacheEnforce
}

func KeyMatchFunc(args ...any) (any, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return KeyMatch(name1, name2), nil
}

func KeyMatch(key1 string, key2 string) bool {
	// 例子：
	// 请求key1  /user/123/profile
	// 策略key2  /user/123/profile
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	// 例子：
	// 请求key1  /user/123/profile
	// 策略key2  /user/*/profile
	if len(key1) > i {
		return key1[:i] == key2[:i]
	}

	// 例子：
	// 请求key1  /user/
	// 策略key2  /user/*/profile
	//return key1 == key2[:i]

	// 例子：
	// 请求key1  /del
	// 策略key2  /del/*
	return key1 == key2[:i-1]
}
