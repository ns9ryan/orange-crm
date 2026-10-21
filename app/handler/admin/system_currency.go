package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type SystemCurrency struct{}

// GetSystemCurrencyNoUse 获取未使用的系统币种
func (h *SystemCurrency) GetSystemCurrencyNoUse(c *gin.Context) {
	res, err := service.Service.SystemCurrency.GetSystemCurrencyNoUse()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// List 系统币种列表
func (h *SystemCurrency) List(c *gin.Context) {
	var req request.SystemCurrencyListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemCurrency.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.SystemCurrency.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
