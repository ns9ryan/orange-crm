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
	"orange-crm/app/response"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type Merchant struct{}

// List 商户列表
func (h *Merchant) List(c *gin.Context) {
	var req request.MerchantListRequest
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

	if err = validate.Validate.Merchant.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	req.ParentId = userInfo.Username

	res, err := service.Service.Merchant.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// AddChildren 创建子商户
func (h *Merchant) AddChildren(c *gin.Context) {
	var req request.AddSubMerchantRequest
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
	req.ParentId = userInfo.Username

	if err = validate.Validate.Merchant.AddChildrenValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err = service.Service.Merchant.AddChildren(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "创建子商户",
		Content:                    fmt.Sprintf("创建子商户:%s", req.Username),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchant,
	})
}

// Edit 编辑商户
func (h *Merchant) Edit(c *gin.Context) {
	var req request.SaveMerchantRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Merchant.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	var merchant model.Merchant
	if errors.Is(global.DB.First(&merchant, req.Id).Error, gorm.ErrRecordNotFound) {
		utils.Response.FailSystem(c, err)
		return
	}
	canEdit := true

	if userInfo.ParentId == "root" {
		// 主商户：可以编辑自己或直接下级
		canEdit = merchant.Username == userInfo.Username || merchant.ParentId == userInfo.Username
	} else {
		// 子商户：只能编辑自己
		canEdit = merchant.Username == userInfo.Username
	}

	if !canEdit {
		utils.Response.FailSystem(c, gerror.Newf(`{#merchant.can.not.edit.this.account}`)) // 你不能编辑这个账号
		return
	}

	username, err := service.Service.Merchant.Edit(c, &req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "编辑商户",
		Content:                    fmt.Sprintf("编辑商户:%s", username),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchant,
	})
}

// Del 删除商户
func (h *Merchant) Del(c *gin.Context) {
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

	// 子商户直接返回
	if userInfo.ParentId != global.MerchantRootName {
		utils.Response.SuccessSimple(c)
		return
	}

	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id:       req.Id,
		model.MerchantColumns.ParentId: userInfo.Username,
	}).First(&merchant).Error; err != nil {
		utils.Response.FailBadRequest(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if err = service.Service.Merchant.Del(merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "删除子商户",
		Content:                    fmt.Sprintf("删除子商户:%s", merchant.Username),
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchant,
	})
}

// Detail 商户详情
func (h *Merchant) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Merchant.Detail(id, "", userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// DetailByMerchantKey 通过key获取详情
func (h *Merchant) DetailByMerchantKey(c *gin.Context) {
	username := c.Param("username")

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	// 不是查看自己的,判断是否是自己的自己
	if userInfo.Username != username && userInfo.ParentId != global.MerchantRootName {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Merchant.Detail(0, username, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// CheckUsername 检查账号是否可用
func (h *Merchant) CheckUsername(c *gin.Context) {
	var req request.CheckUsernameRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Merchant.CheckUsernameValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Merchant.CheckUsername(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// RegenerateAesKeySalt 重新生成AesKeySalt
func (h *Merchant) RegenerateAesKeySalt(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))

	res, err := service.Service.Merchant.RegenerateAesKeySalt(c, id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// RegenerateMd5Salt 重新生成Md5Salt
func (h *Merchant) RegenerateMd5Salt(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Merchant.RegenerateMd5Salt(c, id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// RegenerateTokenKey 重新生成商户tokenKey
func (h *Merchant) RegenerateTokenKey(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Merchant.RegenerateTokenKey(c, id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// GetApiKey 获取api相关密钥
func (h *Merchant) GetApiKey(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Merchant.GetApiKey(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// GetSubBox 子商户下拉框数据
func (h *Merchant) GetSubBox(c *gin.Context) {
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}
	if userInfo.ParentId != "root" {
		list := []response.MerchantBoxResponse{}
		utils.Response.SuccessWithData(c, list)
		return
	}

	res, err := service.Service.Merchant.GetSubBox(userInfo.Username)
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// EditCurrentPassword 编辑当前登录的商户密码
func (h *Merchant) EditCurrentPassword(c *gin.Context) {
	var req request.EditCurrentMerchantPasswordRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Merchant.EditCurrentPasswordValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	if err = service.Service.Merchant.EditCurrentPassword(c, &req, userInfo.ID); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set(global.MerchantLogName, model.MerchantOperationLog{
		Action:                     "编辑密码",
		Content:                    "编辑密码",
		IsSuccess:                  1,
		MerchantOperationEventType: model.MerchantOperationMerchant,
	})
}

// ResetPassword 重置商户密码
func (h *Merchant) ResetPassword(c *gin.Context) {
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

	res, err := service.Service.Merchant.ResetPassword(c, req.Id, userInfo)

	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "重置商户密码",
		Content:                 fmt.Sprintf("重置商户%s的密码", res.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationMerchant,
	})
}
