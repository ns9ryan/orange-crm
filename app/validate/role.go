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

type Role struct{}

// SaveValidate 创建、编辑角色验证
func (v *Role) SaveValidate(req *request.SaveRoleRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":       v.checkId,
		"checkPid":      v.checkPid,
		"checkRoleName": v.checkRoleName,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// ListValidate 菜单列表验证
func (v *Role) ListValidate(req *request.RoleListRequest) error {

	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// SetRoleMenusValidate 设置角色菜单验证
func (v *Role) SetRoleMenusValidate(req *request.SetRoleMenusRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":      v.checkId,
		"checkMenuIds": v.checkMenuIds,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Assoc(req).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// SetRoleApisValidate 设置角色Apis验证
func (v *Role) SetRoleApisValidate(req *request.SetRoleApisRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":     v.checkId,
		"checkApiIds": v.checkApiIds,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Assoc(req).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// GetMenuListByRoleValidate 根据角色获取菜单列表
func (v *Role) GetMenuListByRoleValidate(req *request.GetMenuListByRoleRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId": v.checkId,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// 验证角色是否存在
func (v *Role) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var roleModel = model.Role{ID: in.Value.Uint()}
		if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#role.id.no.exist}`)
		}
	}
	return nil
}

// 验证父级角色是否存在
func (v *Role) checkPid(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var roleModel = model.Role{ID: in.Value.Uint()}
		if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#role.pid.no.exist}`)
		}
	}
	return nil
}

// 验证角色名称是否重复
func (v *Role) checkRoleName(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveRoleRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	sql := global.DB.Where(map[string]any{
		"role_name": strings.TrimSpace(req.RoleName),
		"pid":       req.Pid, // 允许各自子类的名称相同
	})

	if req.Id != 0 {
		sql.Where("id <> ?", req.Id)
	}

	var roleModel model.Role
	res := sql.Find(&roleModel)
	if res.RowsAffected == 1 {
		return gerror.Newf(`{#role.name.exist}`)
	}

	return nil
}

// 验证api是否存在
func (v *Role) checkApiIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	if len(in.Value.Slice()) != 0 {
		var count int64
		global.DB.Model(&model.Api{}).Where("id in ?", in.Value.Slice()).Count(&count)
		if int(count) != len(in.Value.Slice()) {
			return gerror.Newf(`{#role.api.invalid}`)
		}
	}

	return nil
}

// 验证菜单是否存在
func (v *Role) checkMenuIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	var count int64
	global.DB.Model(&model.Menu{}).Where("id in ?", in.Value.Slice()).Count(&count)
	if int(count) != len(in.Value.Slice()) {
		return gerror.Newf(`{#role.menu.invalid}`)
	}
	return nil
}
