package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type SystemCurrency struct{}

// ListValidate 系统币种列表验证
func (v *SystemCurrency) ListValidate(req *request.SystemCurrencyListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
