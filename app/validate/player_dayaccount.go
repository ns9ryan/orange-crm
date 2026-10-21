package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type PlayerDayaccount struct{}

// ListValidate 玩家输赢表列表验证
func (v *PlayerDayaccount) ListValidate(req *request.PlayerDayaccountListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
