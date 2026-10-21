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

type GameTable struct{}

// SaveValidate 创建、编辑游戏桌台验证
func (v *GameTable) SaveValidate(req *request.SaveGameTableRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":            v.checkId,
		"checkGameTableCode": v.checkGameTableCode,
		"checkCategorySub":   v.checkCategorySub,
		"checkDeskNo":        v.checkDeskNo,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 游戏桌台列表验证
func (v *GameTable) ListValidate(req *request.GameTableListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证游戏桌台是否存在
func (v *GameTable) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var gameTableModel = model.GameTable{ID: in.Value.Uint()}
		if err := global.DB.First(&gameTableModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#game_table.id.no.exist}`)
		}
	}
	return nil
}

// 验证游戏桌台标识是否重复
func (v *GameTable) checkGameTableCode(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveGameTableRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.GameTableColumns.GameTableCode: in.Value.String(),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}
	var gameTableModel model.GameTable
	db.First(&gameTableModel)
	if gameTableModel.ID != 0 {
		return gerror.Newf(`{#game_table.code.exist}`)
	}
	return nil
}

// 验证游戏分类是否存在
func (v *GameTable) checkCategorySub(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveGameTableRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	var gameCategoryModel model.GameCategory
	if err := global.DB.Where(map[string]interface{}{
		model.GameCategoryColumns.SubCode: req.CategorySub,
	}).First(&gameCategoryModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#game_category.sub.code.no.exist}`)
	}
	return nil
}

// 验证桌号
func (v *GameTable) checkDeskNo(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveGameTableRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Model(&model.GameTable{}).Where(map[string]interface{}{
		model.GameTableColumns.DeskNo: in.Value.String(),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}
	var count int64
	db.Count(&count)
	if count > 0 {
		return gerror.Newf(`{#game_table.desk_no.exist}`)
	}
	return nil
}
