package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type MerchantDayaccount struct{}

// List 获取主商户输赢报表
func (h *MerchantDayaccount) List(c *gin.Context) {
	var req request.MerchantDayaccountListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantDayaccount.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}
	res, err := service.Service.MerchantDayaccount.GetMerchantReportForAll(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// SubList 获取主商户下子商户的输赢报表
func (h *MerchantDayaccount) SubList(c *gin.Context) {
	var req request.GetSubMerchantDayaccountRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantDayaccount.SubListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}
	res, err := service.Service.MerchantDayaccount.GetSubMerchantDayaccount(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
