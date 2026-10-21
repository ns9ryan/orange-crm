package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type BetRecord struct{}

// List 投注明细列表
func (h *BetRecord) List(c *gin.Context) {
	var req request.BetRecordListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.BetRecord.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 如果没有传就获取当前登录的角色ID
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.BetRecord.List(&req, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
