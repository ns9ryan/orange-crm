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

type Currency struct{}

// SaveValidate 创建、编辑币种验证
func (v *Currency) SaveValidate(req *request.SaveCurrencyRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkSystemCurrencyId": v.checkSystemCurrencyId,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 币种列表验证
func (v *Currency) ListValidate(req *request.CurrencyListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// SetExchangeRateValidate 设置币种汇率验证
func (v *Currency) SetExchangeRateValidate(req *request.SetCurrencyExchangeRateRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId":     v.checkId,
		"checkToCode": v.checkToCode,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ExchangeRateHistoryListValidate 币种汇率历史记录
func (v *Currency) ExchangeRateHistoryListValidate(req *request.CurrencyExchangeRateHistoryListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证币种是否存在
func (v *Currency) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var currencyModel = model.Currency{ID: in.Value.Uint()}
		if err := global.DB.First(&currencyModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#currency.id.no.exist}`)
		}
	}
	return nil
}

// 验证币种配置是否重复
func (v *Currency) checkSystemCurrencyId(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *request.SaveCurrencyRequest
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, utils.MsgValidateScanErr)
	}

	// 检查系统币种是否有该语言
	var systemCurrency = model.SystemCurrency{ID: in.Value.Uint()}
	if err := global.DB.First(&systemCurrency).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#system_currency.id.no.exist}`)
	}

	// 检查是否已经添加过改币种
	var currency model.Currency
	if err := global.DB.Where(map[string]interface{}{
		model.CurrencyColumns.Code: systemCurrency.Code,
	}).First(&currency).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#currency.code.exist}`)
	}

	return nil
}

// 验证目标币种代码是否存在
func (v *Currency) checkToCode(ctx context.Context, in gvalid.RuleFuncInput) error {
	db := global.DB.Where(map[string]interface{}{
		model.CurrencyColumns.Code: in.Value.String(),
	})
	var currencyModel model.Currency
	db.First(&currencyModel)
	if currencyModel.ID == 0 {
		return gerror.Newf(`{#currency.code.no.exist}`)
	}
	return nil
}
