package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
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
