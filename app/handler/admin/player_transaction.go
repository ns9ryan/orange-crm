package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type PlayerTransaction struct{}

// CreditAndDebitList 上分/下分记录列表
func (h *PlayerTransaction) CreditAndDebitList(c *gin.Context) {
	var req request.CreditAndDebitListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.PlayerTransaction.CreditAndDebitListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.PlayerTransaction.CreditAndDebitList(&req, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
