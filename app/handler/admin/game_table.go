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

// Add 创建游戏桌台
func (h *GameTable) Add(c *gin.Context) {
	var req request.SaveGameTableRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.GameTable.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		c.Set("admin_log", model.AdminOperationLog{
			Action:                  "创建游戏桌台",
			Content:                 fmt.Sprintf("创建游戏桌台:%s", req.GameTableCode),
			ErrorMessage:            err.Error(),
			AdminOperationEventType: model.AminOperationGameTable,
		})
		return
	}

	if err := service.Service.GameTable.Add(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		c.Set("admin_log", model.AdminOperationLog{
			Action:                  "创建游戏桌台",
			Content:                 fmt.Sprintf("创建游戏桌台:%s", req.GameTableCode),
			ErrorMessage:            err.Error(),
			AdminOperationEventType: model.AminOperationGameTable,
		})
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建游戏桌台",
		Content:                 fmt.Sprintf("创建游戏桌台:%s", req.GameTableCode),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationGameTable,
	})
}

// Edit 编辑游戏桌台
func (h *GameTable) Edit(c *gin.Context) {
	var req request.SaveGameTableRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.GameTable.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.GameTable.Edit(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑游戏桌台",
		Content:                 fmt.Sprintf("编辑游戏桌台:%s", req.GameTableCode),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationGameTable,
	})
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

// Del 删除游戏桌台
func (h *GameTable) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	gameTable, err := service.Service.GameTable.Del(c, req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除游戏桌台",
		Content:                 fmt.Sprintf("删除游戏桌台:%s", gameTable.GameTableCode),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationGameTable,
	})
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
