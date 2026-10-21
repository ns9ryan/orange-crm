package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type GameCategory struct{}

// GetBox 获取游戏分类选择框
func (h *GameCategory) GetBox(c *gin.Context) {
	res, err := service.Service.GameCategory.GetGameCategoryBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// List 游戏分类列表
func (h *GameCategory) List(c *gin.Context) {
	var req request.GameCategoryListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.GameCategory.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.GameCategory.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
