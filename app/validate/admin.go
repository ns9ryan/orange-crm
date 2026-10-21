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
)

type Admin struct{}

// LoginAdminValidate 管理员登录验证
func (v *Admin) LoginAdminValidate(req *request.LoginAdminRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// SaveValidate 创建、编辑管理员验证
func (v *Admin) SaveValidate(req *request.SaveAdminRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":       v.checkId,
		"checkUsername": v.checkUsername,
		"checkRoleId":   v.checkRoleId,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 管理员列表验证
func (v *Admin) ListValidate(req *request.AdminListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// EditCurrentPasswordValidate 编辑当前登录的管理员密码
func (v *Admin) EditCurrentPasswordValidate(req *request.EditCurrentPasswordRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证管理员是否存在
func (v *Admin) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var adminModel = model.Admin{ID: in.Value.Uint()}
		if err := global.DB.First(&adminModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#admin.id.no.exist}`)
		}
	}
	return nil
}

// 验证用户名是否重复
func (v *Admin) checkUsername(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveAdminRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where("username = ?", in.Value.String())
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var adminModel model.Admin
	db.First(&adminModel)
	if adminModel.ID != 0 {
		return gerror.Newf(`{#admin.username.exist}`)
	}

	return nil
}

// 验证角色Id
func (v *Admin) checkRoleId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var roleModel = model.Role{ID: in.Value.Uint()}
		if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#role.id.no.exist}`)
		}
		if roleModel.Status == 0 {
			return gerror.Newf(`{#role.disable}`)
		}
	}
	return nil
}
