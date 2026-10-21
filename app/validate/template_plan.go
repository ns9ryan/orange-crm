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

type TemplatePlan struct{}

// SaveValidate 创建、编辑限红方案验证
func (v *TemplatePlan) SaveValidate(req *request.SaveTemplatePlanRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkName": v.checkName,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 限红方案列表验证
func (v *TemplatePlan) ListValidate(req *request.TemplatePlanListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证限红方案名称
func (v *TemplatePlan) checkName(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveTemplatePlanRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	// 判断名称是否重复
	db := global.DB.Where(map[string]interface{}{
		model.TemplatePlanColumns.CategorySub: strings.TrimSpace(req.CategorySub),
		model.TemplatePlanColumns.Name:        strings.TrimSpace(req.Name),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var templatePlan model.TemplatePlan
	db.First(&templatePlan)
	if templatePlan.ID != 0 {
		return gerror.Newf(`{#template_plan.name.exist}`)
	}

	return nil
}
