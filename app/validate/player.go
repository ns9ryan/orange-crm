package validate

import (
	"context"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"orange-crm/utils"
	"strconv"
	"strings"
)

type Player struct{}

// SaveValidate 创建、编辑玩家验证
func (v *Player) SaveValidate(req *request.SavePlayerRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":         v.checkId,
		"checkMerchantId": v.checkMerchantId,
		"checkUsername":   v.checkUsername,
		"checkParentId":   v.checkParentId,
		"checkChipIds":    v.checkChipIds,
		"checkTemplateId": v.checkTemplateId,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 玩家列表验证
func (v *Player) ListValidate(req *request.PlayerListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证玩家是否存在
func (v *Player) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var playerModel = model.Player{ID: in.Value.Uint()}
		if err := global.DB.First(&playerModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#play.id.no.exist}`)
		}
	}
	return nil
}

// 验证商户账号
func (v *Player) checkMerchantId(ctx context.Context, in gvalid.RuleFuncInput) error {
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: strings.ToLower(in.Value.String()),
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#merchant.id.no.exist}`)
	}
	return nil
}

// 验证用户名是否重复
func (v *Player) checkUsername(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SavePlayerRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	db := global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Username: strings.ToLower(in.Value.String()),
	})
	if req.Id != 0 {
		db.Where("id <> ?", req.Id)
	}

	var playerModel model.Player
	db.First(&playerModel)
	if playerModel.ID != 0 {
		return gerror.Newf(`{#play.username.exist}`)
	}

	return nil
}

// 验证上级玩账号
func (v *Player) checkParentId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.String() == "root" {
		return nil
	}
	var player model.Player
	if err := global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Username: strings.ToLower(in.Value.String()),
	}).First(&player).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#play.parent_id.no.exist}`)
	}
	return nil
}

// 验证商户限红模板ID
func (v *Player) checkTemplateId(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SavePlayerRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	var template model.MerchantTemplateBetlimit
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitColumns.TemplateId: in.Value.String(),
		model.MerchantTemplateBetlimitColumns.Username:   req.MerchantId,
	}).First(&template).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#merchant_template_betlimit.id.no.exist}`)
	}
	return nil
}

// 验证商户限红模板ID
func (v *Player) checkChipIds(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SavePlayerRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	// 商户的筹码数据
	var chips []model.SystemChip
	global.DB.Where(map[string]interface{}{
		model.SystemChipColumns.Id: req.ChipIds,
	}).Find(&chips)
	if len(chips) != len(req.ChipIds) {
		return gerror.Newf(`{#play.add.chip.no.exist}`) // 筹码不存在或已被禁用
	}

	// 判断限红是否是商家已经拥有
	var merchant model.Merchant
	global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.MerchantId,
	}).Find(&merchant)

	chipStr := strings.Split(merchant.Chips, ",")
	var chipArr []uint
	for _, p := range chipStr {
		var chip int
		chip, err := strconv.Atoi(p) // 转换为 int
		if err != nil {
			return err
		}
		chipArr = append(chipArr, uint(chip))
	}
	isContain := slice.ContainSubSlice(chipArr, req.ChipIds)
	if !isContain {
		return gerror.Newf(`{#play.add.chip.merchant.no.exist}`) // 所选的筹码没有在商家筹码里
	}

	return nil
}
