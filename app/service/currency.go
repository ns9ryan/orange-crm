package service

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strconv"
	"strings"
	"time"
)

type Currency struct{}

// List 币种列表
func (s *Currency) List(req *request.CurrencyListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.CurrencyListResponse{},
	}
	db := global.DB.Model(&model.Currency{}).Preload("SystemTips")

	// 总条数
	db.Count(&res.Total)
	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	// 分页
	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var currencies []model.Currency
	if err = db.Find(&currencies).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.CurrencyListResponse
	for _, currency := range currencies {

		var merchantCount int64
		global.DB.Model(&model.MerchantCurrency{}).Where(map[string]interface{}{
			model.MerchantCurrencyColumns.Currency: currency.Code,
		}).Count(&merchantCount)
		var playerCount int64
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.Currency: currency.Code,
		}).Count(&playerCount)

		var dealerCount int64
		global.DB.Model(&model.Dealer{}).Where(map[string]interface{}{
			model.DealerColumns.CurrencyCode: currency.Code,
		}).Count(&dealerCount)

		tmp := response.CurrencyListResponse{
			Id:             currency.ID,     // 币种D
			Code:           currency.Code,   // 币种编码
			Name:           currency.Name,   // 币种名称
			Symbol:         currency.Symbol, // 币种符号
			Digits:         currency.Digits, // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
			MerchantCount:  merchantCount,   // 商户数量
			PlayerCount:    playerCount,     // 商户数量
			DealerCount:    dealerCount,     // 商户数量
			MaxTip:         currency.MaxTip, // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:         currency.MinTip, // 荷官可接收的最小打赏金额，存储为最小单位
			DefaultChipIds: []uint{},        // 默认筹码ID列表
			DefaultTips:    []uint{},        // 默认打赏面值
		}

		// 默认筹码ID列表
		chipStr := strings.Split(currency.DefaultChipIds, ",")
		for _, p := range chipStr {
			var chip int
			chip, err = strconv.Atoi(p) // 转换为 int
			if err != nil {
				return nil, err
			}
			tmp.DefaultChipIds = append(tmp.DefaultChipIds, uint(chip))
		}

		// 默认打赏面值
		for _, tip := range currency.SystemTips {
			tmp.DefaultTips = append(tmp.DefaultTips, tip.Value)
		}

		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建币种
func (s *Currency) Add(req *request.SaveCurrencyRequest, operatorId string) (code string, err error) {
	var systemCurrency model.SystemCurrency
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.SystemCurrencyColumns.Id: req.SystemCurrencyId,
	}).First(&systemCurrency).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	currencyModel := model.Currency{
		Code:           systemCurrency.Code,           // 币种编码
		Name:           systemCurrency.Name,           // 币种名称
		Symbol:         systemCurrency.Symbol,         // 币种符号
		Digits:         systemCurrency.Digits,         // 币种的最小单位倍数
		MaxTip:         systemCurrency.MaxTip,         // 荷官可接收的最大打赏金额，存储为最小单位
		MinTip:         systemCurrency.MinTip,         // 荷官可接收的最小打赏金额，存储为最小单位
		DefaultChipIds: systemCurrency.DefaultChipIds, // 默认筹码ID列表
		OperatorId:     operatorId,                    // 操作者ID,当前登录账号的username
	}

	return systemCurrency.Code, global.DB.Create(&currencyModel).Error
}

// Detail 币种详情
func (s *Currency) Detail(id uint) (*response.CurrencyDetailResponse, error) {
	var currency model.Currency
	var res response.CurrencyDetailResponse

	err := global.DB.Model(&model.Currency{}).Where("id = ?", id).Preload("CurrencyExchangeRateList").First(&currency).Error

	res.Id = currency.ID         // 币种ID
	res.Code = currency.Code     // 币种编码
	res.Name = currency.Name     // 币种名称
	res.Symbol = currency.Symbol // 币种符号
	res.Digits = currency.Digits // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
	res.MaxTip = currency.MaxTip // 荷官可接收的最大打赏金额，存储为最小单位
	res.MinTip = currency.MinTip // 荷官可接收的最小打赏金额，存储为最小单位

	// 默认筹码ID列表
	chipStr := strings.Split(currency.DefaultChipIds, ",")
	for _, p := range chipStr {
		var chip int
		chip, err = strconv.Atoi(p) // 转换为 int
		if err != nil {
			return nil, err
		}
		res.DefaultChipIds = append(res.DefaultChipIds, uint(chip))
	}

	res.CurrencyExchangeRateList = []response.CurrencyExchangeRateResponse{} // 汇率列表

	for _, exchange := range currency.CurrencyExchangeRateList {
		rate := response.CurrencyExchangeRateResponse{
			FromCode:     exchange.FromCode,                                                                            // 原币种编码
			ToCode:       exchange.ToCode,                                                                              // 目标币种编码
			ExchangeRate: strconv.FormatFloat(float64(exchange.ExchangeRateInt)/float64(exchange.Digits), 'f', -1, 64), // 汇率
		}
		res.CurrencyExchangeRateList = append(res.CurrencyExchangeRateList, rate)
	}

	return &res, err
}

// Del 删除币种
func (s *Currency) Del(id uint) error {
	var currencyModel model.Currency
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.CurrencyColumns.Id: id,
	}).First(&currencyModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	var merchantCount int64
	global.DB.Model(&model.MerchantCurrency{}).Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Currency: currencyModel.Code,
	}).Count(&merchantCount)
	if merchantCount > 0 {
		return gerror.Newf(`{#currency.del.exist.merchant}`) // 币种已被商家使用不能删除
	}

	var playerCount int64
	global.DB.Model(&model.Player{}).Where(map[string]interface{}{
		model.PlayerColumns.Currency: currencyModel.Code,
	}).Count(&playerCount)
	if playerCount > 0 {
		return gerror.Newf(`{#currency.del.exist.player}`) // 币种已被会员使用不能删除
	}

	var dealerCount int64
	global.DB.Model(&model.Dealer{}).Where(map[string]interface{}{
		model.DealerColumns.CurrencyCode: currencyModel.Code,
	}).Count(&dealerCount)
	if dealerCount > 0 {
		return gerror.Newf(`{#currency.del.exist.dealer}`) // 币种已被荷官使用不能删除
	}
	tx := global.DB.Begin()
	tx.Delete(&currencyModel)
	return tx.Commit().Error
}

// SetExchangeRate 设置币种汇率
func (s *Currency) SetExchangeRate(req *request.SetCurrencyExchangeRateRequest) (code string, err error) {
	// 如果汇率为空不进行操作
	if len(req.ExchangeRateList) == 0 {
		return "", nil
	}
	var currency model.Currency
	if errors.Is(global.DB.First(&currency, req.Id).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	// 目标币种不能重复
	toCodeSlice, err := slice.UniqueByField(req.ExchangeRateList, "ToCode")
	if err != nil {
		return "", err
	}
	if len(toCodeSlice) != len(req.ExchangeRateList) {
		return "", gerror.Newf(`{#currency.to_code.duplicated}`)
	}

	tx := global.DB.Begin()

	nowTime := time.Now()

	// 更新历史记录汇率生效结束时间
	if err = tx.Model(&model.CurrencyExchangeRateHistory{}).Where(map[string]interface{}{
		model.CurrencyExchangeRateHistoryColumns.FromCode:         currency.Code,
		model.CurrencyExchangeRateHistoryColumns.EffectiveEndTime: nil,
	}).Updates(model.CurrencyExchangeRateHistory{EffectiveEndTime: &nowTime}).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 删除币种汇率
	if err = tx.Where(map[string]interface{}{
		model.CurrencyExchangeRateColumns.FromCode: currency.Code,
	}).Delete(&model.CurrencyExchangeRate{}).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 构建需要新建的数据,新的汇率和历史记录
	var currencyExchangeRateHistoryModels []model.CurrencyExchangeRateHistory
	var currencyExchangeRateModels []model.CurrencyExchangeRate

	for _, exchangeRate := range req.ExchangeRateList {
		// 解析浮点数:返回的是数部分,小数部分,换算成整数,小数倍率,错误
		_, _, exchangeRateInt, digits, errTmp := utils.ParseDecimal(exchangeRate.ExchangeRate)
		if errTmp != nil {
			tx.Rollback()
			return "", gerror.Newf(`{#currency.to_code.exchange_rate.incorrect}`)
		}

		// 历史记录
		currencyExchangeRateHistoryModel := model.CurrencyExchangeRateHistory{
			FromCode:           currency.Code,         // 原币种编码
			ToCode:             exchangeRate.ToCode,   // 目标币种编码
			ExchangeRateInt:    uint(exchangeRateInt), // 汇率整数表示,例如 1374
			Digits:             uint(digits),          // 汇率倍数,例如 10000
			EffectiveStartTime: nowTime,               // 汇率的生效开始时间
		}

		currencyExchangeRateHistoryModels = append(currencyExchangeRateHistoryModels, currencyExchangeRateHistoryModel)

		// 新的汇率
		exchangeRateModel := model.CurrencyExchangeRate{
			FromCode:        currency.Code,
			ToCode:          exchangeRate.ToCode,
			ExchangeRateInt: uint(exchangeRateInt),
			Digits:          uint(digits),
		}
		currencyExchangeRateModels = append(currencyExchangeRateModels, exchangeRateModel)
	}

	// 插入币种汇率历史记录
	if err = tx.Create(currencyExchangeRateHistoryModels).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 插入币种汇率
	if err = tx.Create(currencyExchangeRateModels).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	return currency.Code, tx.Commit().Error
}

// ExchangeRateHistoryList 币种汇率历史记录
func (s *Currency) ExchangeRateHistoryList(req *request.CurrencyExchangeRateHistoryListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.CurrencyExchangeRateHistoryListResponse{},
	}

	db := global.DB.Model(&model.CurrencyExchangeRateHistory{})

	// 原币种编码查询
	if req.FromCode != "" {
		db.Where(map[string]interface{}{
			model.CurrencyExchangeRateHistoryColumns.FromCode: req.FromCode,
		})
	}

	// 目标币种编码查询
	if req.ToCode != "" {
		db.Where(map[string]interface{}{
			model.CurrencyExchangeRateHistoryColumns.ToCode: req.ToCode,
		})
	}

	// 生效时间查询
	if req.EffectiveStartTime != "" {
		whereStr := fmt.Sprintf("%s >= ?", model.CurrencyExchangeRateHistoryColumns.EffectiveStartTime)
		db.Where(whereStr, req.EffectiveStartTime)
	}

	// 汇率生效结束时间
	if req.EffectiveEndTime != "" {
		whereStr := fmt.Sprintf("%s <=?", model.CurrencyExchangeRateHistoryColumns.EffectiveEndTime)
		db.Where(whereStr, req.EffectiveEndTime)
	}

	db.Count(&res.Total)

	if res.Total <= 0 {
		return
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var CurrencyExchangeRateHistories []model.CurrencyExchangeRateHistory
	if err = db.Find(&CurrencyExchangeRateHistories).Error; err != nil {
		return
	}
	// 列表赋值
	var list []response.CurrencyExchangeRateHistoryListResponse
	for _, history := range CurrencyExchangeRateHistories {
		tmp := response.CurrencyExchangeRateHistoryListResponse{
			Id:                 history.ID,                                                                                 // 历史记录ID
			FromCode:           history.FromCode,                                                                           // 原币种编码
			ToCode:             history.ToCode,                                                                             // 目标币种编码
			ExchangeRate:       strconv.FormatFloat(float64(history.ExchangeRateInt)/float64(history.Digits), 'f', -1, 64), // 汇率
			EffectiveStartTime: history.EffectiveStartTime.Format("2006-01-02 15:04:05"),                                   // 汇率的生效开始时间
			RateSource:         history.RateSource,                                                                         // 来源:1系统
		}

		// 汇率生效结束时间
		if history.EffectiveEndTime != nil {
			tmp.EffectiveEndTime = history.EffectiveEndTime.Format("2006-01-02 15:04:05")
		}

		list = append(list, tmp)
	}
	res.List = list

	return
}
