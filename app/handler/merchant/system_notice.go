package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type SystemNotice struct{}

// ListForNew 商户新消息
func (h *SystemNotice) ListForNew(c *gin.Context) {
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.SystemNotice.ListForNew(userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// ListForSelf 商户收到的系统公共列表
func (h *SystemNotice) ListForSelf(c *gin.Context) {
	var req request.SystemNoticeListForSelfRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemNotice.ListForSelfValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.SystemNotice.ListForSelf(&req, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// MarkRead 标记已读
func (h *SystemNotice) MarkRead(c *gin.Context) {
	var req request.SystemNoticeMarkReadRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.SystemNotice.MarkReadValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	var merchantNotice model.SystemNoticeMerchant
	if err = global.DB.Where(map[string]interface{}{
		model.SystemNoticeMerchantColumns.NoticeId:   req.NoticeId,
		model.SystemNoticeMerchantColumns.MerchantId: userInfo.Username,
	}).Preload("SystemNotice").First(&merchantNotice).Error; err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err = service.Service.SystemNotice.MarkRead(merchantNotice, userInfo); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "标记已读系统消息",
		Content:                    fmt.Sprintf("标记已读系统消息:%s", merchantNotice.SystemNotice.Title),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationSystemNotice,
	})
}
