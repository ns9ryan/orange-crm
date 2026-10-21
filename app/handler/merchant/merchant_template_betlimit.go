package merchant

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
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	res, err := service.Service.MerchantTemplateBetlimit.List(&req, userInfo)
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

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = validate.Validate.MerchantTemplateBetlimit.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 查询商户
	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if userInfo != nil { // 商户端
		if userInfo.ParentId == global.MerchantRootName { // 主商户：可以创建自己和底下的子商户的限红模板
			if req.Username != userInfo.Username { // 不是创建自己的
				if merchant.ParentId != userInfo.Username { // 也不是创建子商户的
					utils.Response.FailValidate(c, gerror.Newf(`{#merchant_template_betlimit.add.self}`)) // 只能创建自己或子商户的模板
					return
				}
			}
		} else { // 子商户：只能创建自己的限红模板
			if req.Username != userInfo.Username { // 不是创建自己的
				utils.Response.FailValidate(c, gerror.Newf(`{#merchant_template_betlimit.add.self}`)) // 只能创建自己或子商户的模板
				return
			}
		}
	}

	if err = service.Service.MerchantTemplateBetlimit.Add(&req, merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "创建限红模板",
		Content:                    fmt.Sprintf("创建限红模板:%s", req.Name),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantTemplateBetlimit,
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

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	// 查询商户
	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).First(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if userInfo != nil { // 商户端
		if userInfo.ParentId == global.MerchantRootName { // 主商户：可以编辑自己和底下的子商户的限红模板
			if req.Username != userInfo.Username { // 不是编辑自己的
				if merchant.ParentId != userInfo.Username { // 也不是编辑子商户的
					utils.Response.FailValidate(c, gerror.Newf(`{#merchant_template_betlimit.edit.self}`)) // 只能编辑自己或子商户的模板
					return
				}
			}
		} else { // 子商户：只能创建自己的限红模板
			if req.Username != userInfo.Username { // 不是创建自己的
				utils.Response.FailValidate(c, gerror.Newf(`{#merchant_template_betlimit.edit.self}`)) // 只能编辑自己或子商户的模板
				return
			}
		}
	}

	if err = service.Service.MerchantTemplateBetlimit.Edit(&req, merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "编辑限红模板",
		Content:                    fmt.Sprintf("编辑限红模板:%s", req.Name),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantTemplateBetlimit,
	})
}

// Detail 商户限红模板详情
func (h *MerchantTemplateBetlimit) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.MerchantTemplateBetlimit.Detail(id, userInfo)
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

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	name, err := service.Service.MerchantTemplateBetlimit.Del(req.Id, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "删除限红模板",
		Content:                    fmt.Sprintf("删除限红模板:%s", name),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantTemplateBetlimit,
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

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	res, err := service.Service.MerchantTemplateBetlimit.GetTemplateBoxByMerchant(&req, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
