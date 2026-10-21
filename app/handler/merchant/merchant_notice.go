package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type MerchantNotice struct{}

// List 商户公告列表
func (h *MerchantNotice) List(c *gin.Context) {
	var req request.MerchantNoticeListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantNotice.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.MerchantNotice.List(&req, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建商户公告
func (h *MerchantNotice) Add(c *gin.Context) {
	var req request.SaveMerchantNoticeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantNotice.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = service.Service.MerchantNotice.Add(&req, userInfo); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "创建商户公告",
		Content:                    fmt.Sprintf("创建商户公告:%s", req.Title),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantNotice,
	})
}

// Edit 编辑商户公告
func (h *MerchantNotice) Edit(c *gin.Context) {
	var req request.SaveMerchantNoticeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.MerchantNotice.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = service.Service.MerchantNotice.Edit(&req, userInfo); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "编辑商户公告",
		Content:                    fmt.Sprintf("编辑商户公告:%s", req.Title),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantNotice,
	})
}

// Detail 商户公告详情
func (h *MerchantNotice) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	res, err := service.Service.MerchantNotice.Detail(id, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除商户公告
func (h *MerchantNotice) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetAdminInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	title, err := service.Service.MerchantNotice.Del(req.Id, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "删除商户公告",
		Content:                    fmt.Sprintf("删除商户公告:%s", title),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchantNotice,
	})
}
