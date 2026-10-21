package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type SystemChip struct{}

// ListValidate 系统筹码列表验证
func (v *SystemChip) ListValidate(req *request.SystemChipListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
