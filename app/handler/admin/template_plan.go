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

type TemplatePlan struct{}

// List 限红方案列表
func (h *TemplatePlan) List(c *gin.Context) {
	var req request.TemplatePlanListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlan.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.TemplatePlan.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建限红方案
func (h *TemplatePlan) Add(c *gin.Context) {
	var req request.SaveTemplatePlanRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlan.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.TemplatePlan.Add(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建限红方案",
		Content:                 fmt.Sprintf("创建限红方案:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlan,
	})
}

// Edit 编辑限红方案
func (h *TemplatePlan) Edit(c *gin.Context) {
	var req request.SaveTemplatePlanRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.TemplatePlan.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.TemplatePlan.Edit(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑限红方案",
		Content:                 fmt.Sprintf("编辑限红方案:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlan,
	})
}

// Detail 限红方案详情
func (h *TemplatePlan) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.TemplatePlan.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除限红方案
func (h *TemplatePlan) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	templatePlan, err := service.Service.TemplatePlan.Del(c, req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑限红方案",
		Content:                 fmt.Sprintf("编辑限红方案:%s", templatePlan.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationTemplatePlan,
	})
}

// GetBoxByCategorySub 根据子分类获取下拉框数据
func (h *TemplatePlan) GetBoxByCategorySub(c *gin.Context) {
	var req request.GetTemplatePlanBoxByCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	res, err := service.Service.TemplatePlan.GetBoxByCategorySub(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
