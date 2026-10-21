package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type PlayerTransaction struct{}

// CreditAndDebitListValidate 上分/下分记录列表验证
func (v *PlayerTransaction) CreditAndDebitListValidate(req *request.CreditAndDebitListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
