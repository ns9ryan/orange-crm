package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type GameTable struct {
}

// List 游戏桌台列表
func (h *GameTable) List(c *gin.Context) {
	var req request.GameTableListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.GameTable.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.GameTable.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Detail 游戏桌台详情
func (h *GameTable) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.GameTable.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// GetBoxWithGroup 获取游戏桌台,分组下拉框
func (h *GameTable) GetBoxWithGroup(c *gin.Context) {
	res := service.Service.GameTable.GetBoxWithGroup()

	utils.Response.SuccessWithData(c, res)
}

// GetBox 游戏桌台下拉框
func (h *GameTable) GetBox(c *gin.Context) {
	res, err := service.Service.GameTable.GetBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}
