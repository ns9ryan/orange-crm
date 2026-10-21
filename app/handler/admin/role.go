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

type Role struct{}

// List 角色列表
func (h *Role) List(c *gin.Context) {
	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	var req request.RoleListRequest
	if err = c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err = validate.Validate.Role.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Role.List(&req, userInfo.RoleId)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建角色
func (h *Role) Add(c *gin.Context) {
	var req request.SaveRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Role.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Role.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建角色",
		Content:                 fmt.Sprintf("创建角色:%s", req.RoleName),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationRole,
	})
}

// Edit 编辑角色
func (h *Role) Edit(c *gin.Context) {
	var req request.SaveRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Role.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Role.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑角色",
		Content:                 fmt.Sprintf("编辑角色:%s", req.RoleName),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationRole,
	})
}

// Detail 角色详情
func (h *Role) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Role.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除角色
func (h *Role) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	roleName, err := service.Service.Role.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除角色",
		Content:                 fmt.Sprintf("删除角色:%s", roleName),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationRole,
	})
}

// SetRoleMenus 设置角色菜单
func (h *Role) SetRoleMenus(c *gin.Context) {
	var req request.SetRoleMenusRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Role.SetRoleMenusValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	roleName, err := service.Service.Role.SetRoleMenus(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "设置角色菜单",
		Content:                 fmt.Sprintf("设置角色%s的菜单", roleName),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationRole,
	})
}

// SetRoleApis 设置角色Apis
func (h *Role) SetRoleApis(c *gin.Context) {
	var req request.SetRoleApisRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Role.SetRoleApisValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	roleName, err := service.Service.Role.SetRoleApis(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "设置角色菜单",
		Content:                 fmt.Sprintf("设置角色%s的Api权限", roleName),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationRole,
	})
}

// GetMenuListByRole 根据角色获取菜单列表
func (h *Role) GetMenuListByRole(c *gin.Context) {
	var req request.GetMenuListByRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}
	if err := validate.Validate.Role.GetMenuListByRoleValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	var roleId uint
	if req.RoleId > 0 {
		roleId = req.RoleId
	} else {
		// 如果没有传就获取当前登录的角色ID
		userInfo, err := service.GetAdminInfo(c)
		if err != nil {
			utils.Response.FailUnauthorized(c)
			return
		}
		roleId = userInfo.RoleId
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Role.GetMenuListByRole(roleId, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
