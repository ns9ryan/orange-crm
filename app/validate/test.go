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

type Test struct{}

// PageValidate 创建、编辑角色验证
func (v *Test) PageValidate(req *request.PageRequest) error {
	rules := map[string]string{
		"Paging":   "in:0,1",
		"Page":     "required-if:paging,1",
		"PageSize": "required-if:paging,1|min:1",
	}
	messages := map[string]interface{}{
		"Paging": map[string]string{
			"in": utils.TT(`{#validate.in}`, utils.TT(`{#page.paging}`), req.Paging),
		},
		"Page": map[string]string{
			"required-if": utils.TT(`{#validate.required-if}`, utils.TT(`{#page.page}`)),
		},
		"PageSize": map[string]string{
			"required-if": utils.TT(`{#validate.required-if}`, utils.TT(`{#page.page_size}`)),
			"min":         utils.TT(`{#validate.min}`, utils.TT(`{#page.page_size}`), req.PageSize),
		},
	}

	if err := g.Validator().Bail().Messages(messages).Rules(rules).Assoc(req).Data(req).Run(global.I18N.Context); err != nil {
		utils.DD(err.Maps())
		return err
	}

	return nil
}

// SaveValidate 创建、编辑角色验证
func (v *Test) SaveValidate(req *request.SaveRoleRequest) error {
	rules := map[string]string{
		"Id":   "IdExist",
		"Pid":  "PidExist",
		"Name": "required|length:2,50|NameExist",
	}
	messages := map[string]interface{}{
		"Name": map[string]string{
			"required": utils.TT(`{#validate.require}`, utils.TT(`{#role.name}`)),
			"length":   utils.TT(`{#validate.length}`, utils.TT(`{#role.name}`), req.RoleName),
		},
	}

	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"IdExist":   v.IdExist,
		"PidExist":  v.PidExist,
		"NameExist": v.NameExist,
	}

	if err := g.Validator().Bail().Messages(messages).Rules(rules).RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// IdExist 验证角色是否存在
func (v *Test) IdExist(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var roleModel = model.Role{ID: in.Value.Uint()}
		if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#role.id.exist}`)
		}
	}
	return nil
}

// PidExist 验证父级角色是否存在
func (v *Test) PidExist(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var roleModel = model.Role{ID: in.Value.Uint()}
		if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#role.pid.exist}`)
		}
	}
	return nil
}

// NameExist 验证角色名称是否重复
func (v *Test) NameExist(ctx context.Context, in gvalid.RuleFuncInput) error {
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
