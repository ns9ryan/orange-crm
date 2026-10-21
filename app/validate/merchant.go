// =================================================================================
// 验证-商户
// =================================================================================

package validate

import (
	"context"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
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

type Merchant struct {
}

// SaveValidate 创建、编辑商户验证
func (v *Merchant) SaveValidate(req *request.SaveMerchantRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":                v.checkId,
		"checkUsername":          v.checkUsername,
		"checkChipIds":           v.checkChipIds,
		"checkChipTemplateIds":   v.checkChipTemplateIds,
		"checkRoleId":            v.checkRoleId,
		"checkPlanFormIds":       v.checkPlanFormIds,
		"checkDefaultPlanFormId": v.checkDefaultPlanFormId,
	}
	return g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context)
}

// AddChildrenValidate 创建子商户
func (v *Merchant) AddChildrenValidate(req *request.AddSubMerchantRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkParentId":          v.checkParentId,
		"checkUsername":          v.checkUsername,
		"checkChipIds":           v.checkChipIds,
		"checkChipTemplateIds":   v.checkChipTemplateIds,
		"checkRoleId":            v.checkRoleId,
		"checkPlanFormIds":       v.checkPlanFormIds,
		"checkDefaultPlanFormId": v.checkDefaultPlanFormId,
	}
	return g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context)
}

// ListValidate 商户列表验证
func (v *Merchant) ListValidate(req *request.MerchantListRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}

// LoginValidate 商户登录验证
func (v *Merchant) LoginValidate(req *request.LoginMerchantRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}

// CheckUsernameValidate 用户名校验
func (v *Merchant) CheckUsernameValidate(req *request.CheckUsernameRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}

// EditCurrentPasswordValidate 编辑当前登录的商户密码
func (v *Merchant) EditCurrentPasswordValidate(req *request.EditCurrentMerchantPasswordRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证商户是否存在
func (v *Merchant) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var gameTableModel = model.Merchant{ID: in.Value.Uint()}
		if err := global.DB.First(&gameTableModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#merchant.id.no.exist}`) // 商户不存在
		}
	}
	return nil
}

// 验证账号
func (v *Merchant) checkUsername(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveMerchantRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	// 如果是编辑不需要校验,编辑不会修改username
	if req.Id != 0 {
		return nil
	}

	var merchant model.Merchant
	global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: strings.ToLower(in.Value.String()),
	}).First(&merchant)
	if merchant.ID != 0 {
		return gerror.Newf(`{#merchant.username.exist}`) // 商户账号已存在
	}

	return nil
}

// 验证筹码
func (v *Merchant) checkChipIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	chipIds := in.Value.Uints()
	var count int64
	global.DB.Model(&model.SystemChip{}).Where(map[string]interface{}{
		model.SystemChipColumns.Id: chipIds,
	}).Count(&count)

	if int(count) != len(chipIds) {
		return gerror.Newf(`{#merchant.add.chip.no.exist}`) // 筹码不存在或已被禁用
	}

	return nil
}

// 验证筹码模板
func (v *Merchant) checkChipTemplateIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveMerchantRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}
	isContain := slice.ContainSubSlice(req.ChipIds, req.ChipTemplateIds)
	if !isContain {
		return gerror.Newf(`{#play.add.chip.merchant.template.no.exist}`) // 筹码模板里只能包含商户筹码
	}
	return nil
}

// 验证角色Id
func (v *Merchant) checkRoleId(ctx context.Context, in gvalid.RuleFuncInput) error {
	var roleModel = model.Role{ID: in.Value.Uint()}
	if err := global.DB.First(&roleModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#role.id.no.exist}`)
	}
	if roleModel.Status == 0 {
		return gerror.Newf(`{#role.disable}`)
	}
	return nil
}

// 验证系统限红模板是否存在
func (v *Merchant) checkPlanFormIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	planFormIds := in.Value.Strings()
	var count int64
	global.DB.Model(&model.TemplatePlanForm{}).Where(map[string]interface{}{
		model.TemplatePlanFormColumns.PlanFormId: planFormIds,
	}).Count(&count)
	if int(count) != len(planFormIds) {
		return gerror.Newf(`{#template_plan_form.id.no.exist}`) // 系统限红模板不存在
	}
	return nil
}

// 验证系统限红模板是否存在
func (v *Merchant) checkDefaultPlanFormId(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveMerchantRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	isContain := slice.Contain(req.PlanFormIds, req.DefaultPlanFormId)

	if !isContain {
		return gerror.Newf(`{#merchant.add.DefaultPlanForm.error}`) // 默认系统限红模板错误
	}
	return nil
}

// 验证上级商户账号
func (v *Merchant) checkParentId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.String() == "root" {
		return nil
	}
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: strings.ToLower(in.Value.String()),
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#merchant.parent_id.no.exist}`) // 上级商户不存在
	}
	return nil
}
