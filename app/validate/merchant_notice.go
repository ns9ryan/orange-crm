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
)

type MerchantNotice struct{}

// SaveValidate 创建、编辑商户公告验证
func (v *MerchantNotice) SaveValidate(req *request.SaveMerchantNoticeRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkTitle": v.checkTitle,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 商户公告列表验证
func (v *MerchantNotice) ListValidate(req *request.MerchantNoticeListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证标题
func (v *MerchantNotice) checkTitle(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveMerchantNoticeRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.MerchantNoticeColumns.Title: in.Value.String(),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var notice model.MerchantNotice
	db.First(&notice)
	if notice.ID != 0 {
		return gerror.Newf(`{#merchant_notice.title.exist}`)
	}

	return nil
}
