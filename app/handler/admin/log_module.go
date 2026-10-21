package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type LogModule struct{}

// AdminLoginLogList 管理员登录日志列表
func (h *LogModule) AdminLoginLogList(c *gin.Context) {
	var req request.AdminLoginLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.LogModule.AdminLoginLogListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.LogModule.AdminLoginLogList(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// MerchantLoginLogList 商户登录日志列表
func (h *LogModule) MerchantLoginLogList(c *gin.Context) {
	var req request.MerchantLoginLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.LogModule.MerchantLoginLogListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.LogModule.MerchantLoginLogList(&req, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// AdminOperationLogList 管理员操作日志列表
func (h *LogModule) AdminOperationLogList(c *gin.Context) {
	var req request.AdminOperationLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.LogModule.AdminOperationLogListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.LogModule.AdminOperationLogList(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// GetAdminOperationLogEventBox 管理员操作日志事件下拉框
func (h *LogModule) GetAdminOperationLogEventBox(c *gin.Context) {
	res := model.AdminOperationLogEventMap
	utils.Response.SuccessWithData(c, res)
}

// PlayerLoginLogList 会员登录日志列表
func (h *LogModule) PlayerLoginLogList(c *gin.Context) {
	var req request.PlayerLoginLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.LogModule.PlayerLoginLogListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.LogModule.PlayerLoginLogList(&req, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// ControlLoginLogList 控端登录日志列表
func (h *LogModule) ControlLoginLogList(c *gin.Context) {
	var req request.ControlLoginLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.LogModule.ControlLoginLogListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.LogModule.ControlLoginLogList(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
