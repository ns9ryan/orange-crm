package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"net"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type Admin struct{}

// List 管理员列表
func (h *Admin) List(c *gin.Context) {
	var req request.AdminListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Admin.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Admin.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建管理员
func (h *Admin) Add(c *gin.Context) {
	var req request.SaveAdminRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Admin.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Admin.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建管理员",
		Content:                 fmt.Sprintf("创建管理员:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationAdmin,
	})
}

// Edit 编辑管理员
func (h *Admin) Edit(c *gin.Context) {
	var req request.SaveAdminRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Admin.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Admin.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑管理员",
		Content:                 fmt.Sprintf("编辑管理员:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationAdmin,
	})
}

// Detail 管理员详情
func (h *Admin) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Admin.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除管理员
func (h *Admin) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	username, err := service.Service.Admin.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除管理员",
		Content:                 fmt.Sprintf("删除管理员:%s", username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationAdmin,
	})
}

// EditCurrentPassword 编辑当前登录的管理员密码
func (h *Admin) EditCurrentPassword(c *gin.Context) {
	var req request.EditCurrentPasswordRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Admin.EditCurrentPasswordValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = service.Service.Admin.EditCurrentPassword(&req, userInfo.ID); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑密码",
		Content:                 "编辑密码",
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationAdmin,
	})
}

// ResetPassword 重置管理员密码
func (h *Admin) ResetPassword(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	res, err := service.Service.Admin.ResetPassword(req.Id)

	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "重置密码",
		Content:                 fmt.Sprintf("重置密码:%s", res.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationAdmin,
	})
}

// GetAdminInfo 获取管理员信息
func (h *Admin) GetAdminInfo(c *gin.Context) {
	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Admin.GetAdminInfo(userInfo.ID)
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// LogoutAdmin 登出
func (h *Admin) LogoutAdmin(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}

	utils.Response.SuccessSimple(c)
}
