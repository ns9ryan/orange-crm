package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type GameCategory struct{}

// ListValidate 游戏分类列表验证
func (v *GameCategory) ListValidate(req *request.GameCategoryListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
