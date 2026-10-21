package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type MerchantDayaccount struct{}

// ListValidate 商户输赢表列表验证
func (v *MerchantDayaccount) ListValidate(req *request.MerchantDayaccountListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// SubListValidate 获取主商户下子商户的输赢报表验证
func (v *MerchantDayaccount) SubListValidate(req *request.GetSubMerchantDayaccountRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
