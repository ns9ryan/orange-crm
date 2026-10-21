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

type SystemNotice struct{}

// List 系统公告列表
func (h *SystemNotice) List(c *gin.Context) {
	var req request.SystemNoticeListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemNotice.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.SystemNotice.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建系统公告
func (h *SystemNotice) Add(c *gin.Context) {
	var req request.SaveSystemNoticeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemNotice.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = service.Service.SystemNotice.Add(&req, userInfo.Username); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建系统公告",
		Content:                 fmt.Sprintf("创建系统公告:%s", req.Title),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationSystemNotice,
	})
}

// Edit 编辑系统公告
func (h *SystemNotice) Edit(c *gin.Context) {
	var req request.SaveSystemNoticeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemNotice.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.SystemNotice.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑系统公告",
		Content:                 fmt.Sprintf("编辑系统公告:%s", req.Title),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationSystemNotice,
	})
}

// Detail 系统公告详情
func (h *SystemNotice) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.SystemNotice.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除系统公告
func (h *SystemNotice) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	title, err := service.Service.SystemNotice.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除系统公告",
		Content:                 fmt.Sprintf("删除系统公告:%s", title),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationSystemNotice,
	})
}
