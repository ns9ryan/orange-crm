package validate

import (
	"context"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"orange-crm/app/request"
	"orange-crm/global"
)

type Common struct{}

// PageValidate 通用分页验证
func (v *Common) PageValidate(req *request.PageRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// UploadImgValidate 通用上传图片验证
func (v *Common) UploadImgValidate(req *request.UploadImgRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkUploadImgEventType": v.checkUploadImgEventType,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// UploadChunkValidate 通用上传切片验证
func (v *Common) UploadChunkValidate(req *request.UploadChunkRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkUploadImgEventType": v.checkUploadImgEventType,
	}

	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}

	return nil
}

// 验证上传图片事件类型
func (v *Common) checkUploadImgEventType(ctx context.Context, in gvalid.RuleFuncInput) error {
	has := slice.Contain(global.UploadImgEventType, in.Value.String())
	if !has {
		return gerror.Newf(`{#common.upload.img.event_type.incorrect}`)
	}
	return nil
}
