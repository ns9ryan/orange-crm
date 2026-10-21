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
	"strings"
)

type ControlAccount struct{}

// SaveValidate 创建、编辑控端账号验证
func (v *ControlAccount) SaveValidate(req *request.SaveControlAccountRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":       v.checkId,
		"checkUsername": v.checkUsername,
		"checkLevel":    v.checkLevel,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 控端账号列表验证
func (v *ControlAccount) ListValidate(req *request.ControlAccountListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证荷官是否存在
func (v *ControlAccount) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var controlAccountModel = model.ControlAccount{ID: in.Value.Uint()}
		if err := global.DB.First(&controlAccountModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#control_account.id.no.exist}`)
		}
	}
	return nil
}

// 验证用户名是否重复
func (v *ControlAccount) checkUsername(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveControlAccountRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.ControlAccountColumns.Username: strings.ToLower(in.Value.String()),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var controlAccountModel model.ControlAccount
	db.First(&controlAccountModel)
	if controlAccountModel.ID != 0 {
		return gerror.Newf(`{#control_account.username.exist}`)
	}
	return nil
}

// 验证控端账号等级
func (v *ControlAccount) checkLevel(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		if _, ok := model.ControlAccountLevelMap[model.ControlAccountLevel(in.Value.Uint())]; !ok {
			return gerror.Newf(`{#control_account.level.no.exist}`)
		}
	}
	return nil
}
