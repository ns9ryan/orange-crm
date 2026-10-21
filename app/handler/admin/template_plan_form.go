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

type TemplatePlanForm struct{}

// List 限红模板列表
func (h *TemplatePlanForm) List(c *gin.Context) {
	var req request.TemplatePlanFormListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlanForm.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.TemplatePlanForm.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建限红模板
func (h *TemplatePlanForm) Add(c *gin.Context) {
	var req request.SaveTemplatePlanFormRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlanForm.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.TemplatePlanForm.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建限红模板",
		Content:                 fmt.Sprintf("创建限红模板:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlanForm,
	})
}

// Edit 编辑限红模板
func (h *TemplatePlanForm) Edit(c *gin.Context) {
	var req request.SaveTemplatePlanFormRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlanForm.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.TemplatePlanForm.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑限红模板",
		Content:                 fmt.Sprintf("编辑限红模板:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlanForm,
	})
}

// Detail 限红模板详情
func (h *TemplatePlanForm) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.TemplatePlanForm.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除限红模板
func (h *TemplatePlanForm) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	name, err := service.Service.TemplatePlanForm.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除限红模板",
		Content:                 fmt.Sprintf("删除限红模板:%s", name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlanForm,
	})
}

// GetBox 限红模板下拉框数据
func (h *TemplatePlanForm) GetBox(c *gin.Context) {
	res, err := service.Service.TemplatePlanForm.GetBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}
