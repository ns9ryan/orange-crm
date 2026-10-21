package validate

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type Menu struct{}

// SaveValidate 创建、编辑菜单验证
func (v *Menu) SaveValidate(req *request.SaveMenuRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":       v.checkId,
		"checkPid":      v.checkPid,
		"checkMenuName": v.checkMenuName,
		"checkLang":     v.checkLang,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// ListValidate 菜单列表验证
func (v *Menu) ListValidate(req *request.MenuListRequest) error {

	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// 验证菜单是否存在
func (v *Menu) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var menuModel = model.Menu{ID: in.Value.Uint()}
		if err := global.DB.First(&menuModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#menu.id.no.exist}`)
		}
	}
	return nil
}

// 验证父级菜单是否存在
func (v *Menu) checkPid(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var menuModel = model.Menu{ID: in.Value.Uint()}
		if err := global.DB.First(&menuModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#menu.pid.no.exist}`)
		}
	}
	return nil
}

// 验证菜单名称是否重复
func (v *Menu) checkMenuName(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveMenuRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	sql := global.DB.Where(map[string]any{
		"menu_name": strings.TrimSpace(req.MenuName),
	})

	if req.Id != 0 {
		sql.Where("id <> ?", req.Id)
	}

	var menuModel model.Menu
	res := sql.Find(&menuModel)
	if res.RowsAffected == 1 {
		return gerror.Newf(`{#menu.menu_name.exist}`)
	}

	return nil
}

// 验证语言是否存在
func (v *Menu) checkLang(ctx context.Context, in gvalid.RuleFuncInput) error {
	var langModel = model.Lang{Lang: in.Value.String()}
	if err := global.DB.First(&langModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#lang.no.exist}`)
	}
	return nil
}
