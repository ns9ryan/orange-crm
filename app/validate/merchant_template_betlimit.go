package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type MerchantTemplateBetlimit struct{}

// SaveValidate 创建、编辑商户限红模板验证
func (v *MerchantTemplateBetlimit) SaveValidate(req *request.SaveMerchantTemplateBetlimitRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 获取商户限红模板下拉框数据验证
func (v *MerchantTemplateBetlimit) ListValidate(req *request.MerchantTemplateBetlimitListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// GetTemplateBoxByMerchantValidate 获取商家限红模板列表验证
func (v *MerchantTemplateBetlimit) GetTemplateBoxByMerchantValidate(req *request.GetTemplateBoxByMerchantRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
