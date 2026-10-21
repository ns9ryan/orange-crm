package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
)

type SystemCurrency struct{}

// GetSystemCurrencyNoUse 获取未使用的系统币种
func (s *SystemCurrency) GetSystemCurrencyNoUse() ([]response.SystemCurrencyResponse, error) {
	// 获取平台币种
	var currencyCodes []string
	global.DB.Model(&model.Currency{}).Pluck(model.CurrencyColumns.Code, &currencyCodes)

	var systemCurrencyData []response.SystemCurrencyResponse
	if err := global.DB.Model(&model.SystemCurrency{}).Where(fmt.Sprintf("%s NOT IN ?", model.SystemCurrencyColumns.Code), currencyCodes).Find(&systemCurrencyData).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return []response.SystemCurrencyResponse{}, err
	}
	return systemCurrencyData, nil
}

// List 系统币种列表
func (s *SystemCurrency) List(req *request.SystemCurrencyListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.SystemCurrencyResponse{},
	}
	db := global.DB.Model(&model.SystemCurrency{})

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
	var systemCurrencies []model.SystemCurrency
	if err = db.Find(&systemCurrencies).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.SystemCurrencyResponse
	for _, currency := range systemCurrencies {
		tmp := response.SystemCurrencyResponse{
			Id:          currency.ID,          // ID
			Code:        currency.Code,        // 字母编码
			Name:        currency.Name,        // 名称
			Symbol:      currency.Symbol,      // 符号
			Digits:      currency.Digits,      // 币种的最小单位倍数
			Number:      currency.Number,      // 数字编码
			Region:      currency.Region,      // 国家/地区/组织
			MaxTip:      currency.MaxTip,      // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:      currency.MinTip,      // 荷官可接收的最小打赏金额，存储为最小单位
			Description: currency.Description, // 备注
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// GetSystemCurrencyBox 系统币种下拉框
func (s *SystemCurrency) GetSystemCurrencyBox() (list []response.SystemCurrencyResponse, err error) {
	// 先从redis读取数据
	data, err := g.Redis().Do(gctx.New(), "GET", "system:system_currency")
	if err != nil {
		return nil, err
	}
	if err = data.Struct(&list); err != nil {
		return nil, err
	}
	// 如果存在直接返回
	if len(list) != 0 {
		return list, err
	}

	// 从数据库种读取数据
	if err = global.DB.Model(&model.SystemCurrency{}).Find(&list).Error; err != nil {
		return nil, err
	}
	// 存入redis
	_, err = g.Redis().Do(gctx.New(), "set", "system:system_currency", list)

	return list, err
}
