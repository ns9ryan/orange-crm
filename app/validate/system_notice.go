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

type SystemNotice struct{}

// SaveValidate 创建、编辑游戏系统公告验证
func (v *SystemNotice) SaveValidate(req *request.SaveSystemNoticeRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":    v.checkId,
		"checkTitle": v.checkTitle,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 系统公告列表验证
func (v *SystemNotice) ListValidate(req *request.SystemNoticeListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListForSelfValidate 商户收到的系统公告列表验证
func (v *SystemNotice) ListForSelfValidate(req *request.SystemNoticeListForSelfRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// MarkReadValidate 标记已读验证
func (v *SystemNotice) MarkReadValidate(req *request.SystemNoticeMarkReadRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证公告是否存在
func (v *SystemNotice) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var notice = model.SystemNotice{ID: in.Value.Uint()}
		if err := global.DB.First(&notice).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#system_notice.id.no.exist}`)
		}
	}
	return nil
}

// 验证标题
func (v *SystemNotice) checkTitle(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveSystemNoticeRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.SystemNoticeColumns.Title: in.Value.String(),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var notice model.SystemNotice
	db.First(&notice)
	if notice.ID != 0 {
		return gerror.Newf(`{#system_notice.title.exist}`)
	}

	return nil
}
