package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type BetArea struct{}

// ListValidate 下注区域列表验证
func (v *BetArea) ListValidate(req *request.BetAreaListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
