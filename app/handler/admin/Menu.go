package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type Menu struct{}

// List 菜单列表
func (h *Menu) List(c *gin.Context) {
	var req request.MenuListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Menu.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Menu.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建菜单
func (h *Menu) Add(c *gin.Context) {
	var req request.SaveMenuRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Menu.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Menu.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// Edit 编辑菜单
func (h *Menu) Edit(c *gin.Context) {
	var req request.SaveMenuRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Menu.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Menu.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// Detail 菜单详情
func (h *Menu) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Menu.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除菜单
func (h *Menu) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := service.Service.Menu.Del(req.Id); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}
