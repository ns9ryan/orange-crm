// =================================================================================
// 验证-荷官
// =================================================================================

package validate

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"orange-crm/utils"
)

type Dealer struct{}

// SaveValidate 创建、编辑荷官验证
func (v *Dealer) SaveValidate(req *request.SaveDealerRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":             v.checkId,
		"checkEmployeeNumber": v.checkEmployeeNumber,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 荷官列表验证
func (v *Dealer) ListValidate(req *request.DealerListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证荷官是否存在
func (v *Dealer) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var gameCategoryModel = model.Dealer{ID: in.Value.Uint()}
		if err := global.DB.First(&gameCategoryModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#dealer.id.no.exist}`)
		}
	}
	return nil
}

// 验证荷官是否存在
func (v *Dealer) checkEmployeeNumber(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveDealerRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.DealerColumns.EmployeeNumber: in.Value.String(),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var dealerModel model.Dealer
	db.First(&dealerModel)
	if dealerModel.ID != 0 {
		return gerror.Newf(`{#dealer.employee_number.exist}`)
	}

	return nil
}
