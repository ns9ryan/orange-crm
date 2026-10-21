package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
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

// GetBox 限红模板下拉框数据
func (h *TemplatePlanForm) GetBox(c *gin.Context) {
	res, err := service.Service.TemplatePlanForm.GetBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}
