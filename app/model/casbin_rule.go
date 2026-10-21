package model

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"orange-crm/global"
)

// CasbinRule casbin表结构体
type CasbinRule struct {
	gormadapter.CasbinRule
}

// RegisterTable 数据迁移
func (m *CasbinRule) RegisterTable() error {
	var count int64
	global.DB.Model(&CasbinRule{}).Count(&count)
	// 获取所有的api
	var apiModels []Api
	global.DB.Find(&apiModels)

	if count == 0 && len(apiModels) > 0 {
		rules := make([][]string, 0, len(apiModels))
		subs := []string{
			"admin_1", "admin_2", "admin_3",
		}
		// 超级账号权限
		for _, sub := range subs {
			for _, v2 := range apiModels {
				rules = append(rules, []string{sub, v2.Path, v2.Method}) // 增加Id前端要渲染
			}
		}

		// 添加到casbin缓存中
		if ok, err := global.CASBIN.AddPolicies(rules); !ok {
			return gerror.Wrap(err, "存在相同api,添加失败,请联系管理员2")
		}
	}

	// 获取所有的api
	//var apiMerchantModels []Api
	//global.DB.Where("can_assign_merchant = ?", 1).Find(&apiMerchantModels)
	//
	//ruleMerchants := make([][]string, 0, len(apiMerchantModels))
	//
	//sub := "merchant_1"
	//for _, v := range apiMerchantModels {
	//	ruleMerchants = append(ruleMerchants, []string{sub, v.Path, v.Method})
	//}
	//
	//sub = "merchant_2"
	//for _, v := range apiMerchantModels {
	//	ruleMerchants = append(ruleMerchants, []string{sub, v.Path, v.Method})
	//}
	//// 添加到casbin缓存中
	//if ok, err := global.CASBIN.AddPolicies(ruleMerchants); !ok {
	//	return gerror.Wrap(err, "存在相同api,添加失败,请联系管理员3")
	//}

	return nil
}

// TableName 重写 TableName 方法来指定自定义的表前缀
func (m *CasbinRule) TableName() string {
	// 返回带有前缀的表名
	return global.CONFIG.Mysql.Prefix + "casbin_rule"
}
