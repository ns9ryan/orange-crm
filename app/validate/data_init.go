// =================================================================================
// 验证-商户
// =================================================================================

package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type DataInit struct {
}

// AddMerchantsValidate 创批量添加商户验证
func (v *DataInit) AddMerchantsValidate(req *request.AddMerchantsRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}

// AddPlayersValidate 批量添加玩家
func (v *DataInit) AddPlayersValidate(req *request.AddPlayersRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}

// PlayersDepositValidate 批量存款上分
func (v *DataInit) PlayersDepositValidate(req *request.PlayersDepositRequest) error {
	return g.Validator().Bail().Data(req).Run(global.I18N.Context)
}
