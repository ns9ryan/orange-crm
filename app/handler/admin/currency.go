package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type Currency struct{}

// List 币种列表
func (h *Currency) List(c *gin.Context) {
	var req request.CurrencyListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Currency.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Currency.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建币种
func (h *Currency) Add(c *gin.Context) {
	var req request.SaveCurrencyRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Currency.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	code, err := service.Service.Currency.Add(&req, userInfo.Username)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建币种",
		Content:                 fmt.Sprintf("创建币种:%s", code),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationCurrency,
	})
}

// Detail 币种详情
func (h *Currency) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Currency.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除币种
func (h *Currency) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := service.Service.Currency.Del(req.Id); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// SetExchangeRate 设置币种汇率
func (h *Currency) SetExchangeRate(c *gin.Context) {
	var req request.SetCurrencyExchangeRateRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Currency.SetExchangeRateValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	code, err := service.Service.Currency.SetExchangeRate(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "设置币种汇率",
		Content:                 fmt.Sprintf("设置币种%s的汇率:", code),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationCurrency,
	})
}

// ExchangeRateHistoryList 币种汇率历史记录
func (h *Currency) ExchangeRateHistoryList(c *gin.Context) {
	var req request.CurrencyExchangeRateHistoryListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Currency.ExchangeRateHistoryListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Currency.ExchangeRateHistoryList(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
