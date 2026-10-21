package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type MerchantDayaccount struct{}

// List 获取商户自身的输赢报表
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
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	res, err := service.Service.MerchantDayaccount.GetMerchantReportForSelf(&req, userInfo.Username)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// SubList 获取主商户下子商户的输赢报表（仅主商户可用）
func (h *MerchantDayaccount) SubList(c *gin.Context) {
	var req request.GetSubMerchantDayaccountRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	// 如果是子商户访问了这个接口直接返回
	if userInfo.ParentId != global.MerchantRootName {
		utils.Response.SuccessSimple(c)
		return
	}

	req.ParentId = userInfo.Username

	if err = validate.Validate.MerchantDayaccount.SubListValidate(&req); err != nil {
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
