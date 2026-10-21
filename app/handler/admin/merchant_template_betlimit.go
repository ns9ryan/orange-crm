package admin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type MerchantTemplateBetlimit struct{}

// List 商户限红模板列表
func (h *MerchantTemplateBetlimit) List(c *gin.Context) {
	var req request.MerchantTemplateBetlimitListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantTemplateBetlimit.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.MerchantTemplateBetlimit.List(&req, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建商户限红模板
func (h *MerchantTemplateBetlimit) Add(c *gin.Context) {
	var req request.SaveMerchantTemplateBetlimitRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantTemplateBetlimit.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 查询商户
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if err := service.Service.MerchantTemplateBetlimit.Add(&req, merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建商户限红模板",
		Content:                 fmt.Sprintf("创建商户限红模板:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationMerchantTemplateBetlimit,
	})
}

// Edit 编辑商户限红模板
func (h *MerchantTemplateBetlimit) Edit(c *gin.Context) {
	var req request.SaveMerchantTemplateBetlimitRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantTemplateBetlimit.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 查询商户
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if err := service.Service.MerchantTemplateBetlimit.Edit(&req, merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑商户限红模板",
		Content:                 fmt.Sprintf("编辑商户限红模板:%s", req.Name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationMerchantTemplateBetlimit,
	})
}

// Detail 商户限红模板详情
func (h *MerchantTemplateBetlimit) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.MerchantTemplateBetlimit.Detail(id, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除商户限红模板
func (h *MerchantTemplateBetlimit) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	name, err := service.Service.MerchantTemplateBetlimit.Del(req.Id, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除商户限红模板",
		Content:                 fmt.Sprintf("删除商户限红模板:%s", name),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationMerchantTemplateBetlimit,
	})
}

// GetTemplateBoxByMerchant 获取商户限红模板下拉框数据
func (h *MerchantTemplateBetlimit) GetTemplateBoxByMerchant(c *gin.Context) {
	var req request.GetTemplateBoxByMerchantRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantTemplateBetlimit.GetTemplateBoxByMerchantValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.MerchantTemplateBetlimit.GetTemplateBoxByMerchant(&req, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
