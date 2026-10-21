package service

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/global"
)

type CasbinCustom struct {
	*casbin.CachedEnforcer
}

func NewCasbinCustom(e *casbin.CachedEnforcer) *CasbinCustom {
	return &CasbinCustom{e}
}

// RemovePolicy 删除策略 - 分配权限和删除api的时候调用
func (c *CasbinCustom) RemovePolicy(i int, v ...string) (ok bool) {
	ok, _ = c.RemoveFilteredPolicy(i, v...)
	_ = c.InvalidateCache()
	return
}

// UpdatePolicy 根据账号更新策略
func (c *CasbinCustom) UpdatePolicy(sub string, apiModels []model.Api) error {
	// 删除对应角色的策略
	c.RemovePolicy(0, sub)
	rules := make([][]string, 0, len(apiModels))
	for _, v := range apiModels {
		rules = append(rules, []string{sub, v.Path, v.Method}) // 增加Id前端要渲染
	}

	// 设置空权限无需调用 AddPolicies 方法
	if len(rules) == 0 {
		return nil
	}

	if ok, err := c.AddPolicies(rules); !ok {
		return gerror.Wrap(err, "存在相同api,添加失败,请联系管理员1")
	}
	return nil
}

// BatchRemovePolicies 批量删除策略函数-多个
func (c *CasbinCustom) BatchRemovePolicies(subs []string) error {
	if len(subs) == 0 {
		return nil
	}

	err := global.DB.Where("v0 IN ?", subs).Where("ptype = ?", "p").Delete(&model.CasbinRule{}).Error
	if err != nil {
		return err
	}

	return nil
}

// BatchUpdatePolicies 批量更新策略函数-多个
func (c *CasbinCustom) BatchUpdatePolicies(subs []string, apiModels []model.Api) error {
	if len(subs) == 0 {
		return nil
	}

	rules := make([][]string, 0, len(subs)*len(apiModels))
	for _, sub := range subs {
		// 构建新策略规则
		for _, api := range apiModels {
			rules = append(rules, []string{sub, api.Path, api.Method})
		}

	}

	// 设置空权限无需调用 AddPolicies 方法
	if len(rules) == 0 {
		return nil
	}

	// 批量删除所有subs对应的策略
	if err := c.BatchRemovePolicies(subs); err != nil {
		return err
	}

	_ = c.Enforcer.LoadPolicy() // 强制从数据库加载

	// 批量插入规则
	ok, err := c.AddPolicies(rules)

	//if slice.Contain(subs, "merchant_60") {
	//	utils.DD(err)
	//	panic(111)
	//}
	if !ok || err != nil {
		return errors.New(fmt.Sprintf("存在重复 api，批量添加失败，请联系管理员:%v", err))
	}

	// 统一清空缓存
	return c.InvalidateCache()
}

// BatchUpdateSubjectsByRole 批量更新某角色所有 admin 与 merchant 的权限
func (c *CasbinCustom) BatchUpdateSubjectsByRole(roleId uint, apiModels []model.Api) error {
	var subs []string

	var adminModels []model.Admin
	global.DB.Where("role_id = ?", roleId).Find(&adminModels)
	for _, admin := range adminModels {
		subs = append(subs, "admin_"+gconv.String(admin.ID))
	}

	var merchantModels []model.Merchant
	global.DB.Where("role_id = ?", roleId).Find(&merchantModels)
	for _, merchant := range merchantModels {
		subs = append(subs, "merchant_"+gconv.String(merchant.ID))
	}

	// 默认同步执行（事务外执行）
	if err := c.BatchUpdatePolicies(subs, apiModels); err != nil {
		return err
	}

	// 选：异步执行（开启此功能时注释上面那一段）
	/*
		go func(subs []string, apiModels []model.Api) {
			if err := c.BatchUpdatePolicies(subs, apiModels); err != nil {
				log.Printf("[WARN] 异步更新 Casbin 策略失败: %v", err)
			}
		}(subs, apiModels)
	*/

	return nil
}

// UpdateCasbinApi 更新角色Api - 更新api的时候调用
func (c *CasbinCustom) UpdateCasbinApi(beforeObj, beforeAct, afterObj, afterAct string) error {
	if err := global.DB.Model(&model.CasbinRule{}).Where("v1 = ? AND v2 = ?", beforeObj, beforeAct).Updates(map[string]any{
		"v1": afterObj,
		"v2": afterAct,
	}).Error; err != nil {
		return err
	}
	if err := c.InvalidateCache(); err != nil {
		return err
	}
	return nil
}
