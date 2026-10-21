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
)

type Api struct{}

// SaveValidate 创建、编辑api验证
func (v *Api) SaveValidate(req *request.SaveApiRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":     v.checkId,
		"checkPath":   v.checkPath,
		"checkMethod": v.checkMethod,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// 验证api是否存在
func (v *Api) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var apiModel = model.Api{ID: in.Value.Uint()}
		if err := global.DB.First(&apiModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#api.id.no.exist}`)
		}
	}
	return nil
}

// checkPath 验证path
func (v *Api) checkPath(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveApiRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}
	db := global.DB.Where("path = ?", in.Value.String())
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}
	var apiModel model.Api
	db.First(&apiModel)
	if apiModel.ID != 0 {
		return gerror.Newf(`{#api.path.exist}`)
	}
	return nil
}

// 验证请求方法
func (v *Api) checkMethod(ctx context.Context, in gvalid.RuleFuncInput) error {
	has := slice.Contain([]string{"GET", "POST", "PUT", "DELETE"}, in.Value.String())
	if !has {
		return gerror.Newf(`{#api.params.method.incorrect}`)
	}
	return nil
}

// ListValidate api列表验证
func (v *Api) ListValidate(req *request.ApiListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
