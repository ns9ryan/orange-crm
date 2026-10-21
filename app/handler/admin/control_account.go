// =================================================================================
// api请求接收-控端账号
// =================================================================================

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

type ControlAccount struct{}

// GetLevelBox 获取控端账号等级选择框
func (h *ControlAccount) GetLevelBox(c *gin.Context) {
	data := service.Service.ControlAccount.GetLevelBox()
	utils.Response.SuccessWithData(c, data)
}

// List 控端账号列表
func (h *ControlAccount) List(c *gin.Context) {
	var req request.ControlAccountListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ControlAccount.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.ControlAccount.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建控端账号
func (h *ControlAccount) Add(c *gin.Context) {
	var req request.SaveControlAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ControlAccount.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.ControlAccount.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建控端账号",
		Content:                 fmt.Sprintf("创建控端账号:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationControlAccount,
	})
}

// Edit 编辑控端账号
func (h *ControlAccount) Edit(c *gin.Context) {
	var req request.SaveControlAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ControlAccount.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.ControlAccount.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑控端账号",
		Content:                 fmt.Sprintf("编辑控端账号:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationControlAccount,
	})
}

// Detail 控端账号详情
func (h *ControlAccount) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.ControlAccount.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// Del 删除控端账号
func (h *ControlAccount) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}
	username, err := service.Service.ControlAccount.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除控端账号",
		Content:                 fmt.Sprintf("删除控端账号:%s", username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationControlAccount,
	})
}

// ResetPassword 重置控端账号密码
func (h *ControlAccount) ResetPassword(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	res, err := service.Service.ControlAccount.ResetPassword(req.Id)

	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// GetBox 获取控端账号下拉框数据
func (h *ControlAccount) GetBox(c *gin.Context) {
	res, err := service.Service.ControlAccount.GetBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}
