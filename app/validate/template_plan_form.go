package validate

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type TemplatePlanForm struct{}

// SaveValidate 创建、编辑限红模板验证
func (v *TemplatePlanForm) SaveValidate(req *request.SaveTemplatePlanFormRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkName": v.checkName,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 限红模板列表验证
func (v *TemplatePlanForm) ListValidate(req *request.TemplatePlanFormListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证限红模板名称
func (v *TemplatePlanForm) checkName(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveTemplatePlanFormRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	// 判断名称是否重复
	db := global.DB.Where(map[string]interface{}{
		model.TemplatePlanFormColumns.Name: strings.TrimSpace(req.Name),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var templatePlanForm model.TemplatePlanForm
	db.First(&templatePlanForm)
	if templatePlanForm.ID != 0 {
		return gerror.Newf(`{#template_plan_form.name.exist}`)
	}

	return nil
}
