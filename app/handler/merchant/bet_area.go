package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type BetArea struct{}

// List 下注区域列表
func (h *BetArea) List(c *gin.Context) {
	var req request.BetAreaListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.BetArea.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.BetArea.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// GetBoxByCategorySub 根据子分类获取下注区域的下拉框数据
func (h *BetArea) GetBoxByCategorySub(c *gin.Context) {
	var req request.GetBoxByCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	res, err := service.Service.BetArea.GetBoxByCategorySub(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
