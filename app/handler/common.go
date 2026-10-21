package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
	"time"
)

type Common struct{}

// GetLangBox 语言下拉框
func (h *Common) GetLangBox(c *gin.Context) {
	res, err := service.Service.Common.GetLangBox()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// GetSystemConfig 获取系统配置数据
func (h *Common) GetSystemConfig(c *gin.Context) {
	res, err := service.Service.Common.GetSystemConfig()
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// GameServiceName 游戏服务名称
func (h *Common) GameServiceName(c *gin.Context) {
	utils.Response.SuccessWithData(c, global.CONFIG.GameServiceName)
}

// getUploadImgSubDir 根据 EventType 获取上传图片对应的子目录路径
func (h *Common) getUploadImgSubDir(eventType string) string {
	// 默认文件夹
	baseDir := "default"
	dateStr := time.Now().Format("20060102")
	subDir := fmt.Sprintf("/%s/%s/", baseDir, dateStr)
	// 根据不同的 EventType 修改文件夹名称
	switch eventType {
	case global.UploadAdminAvatar:
		subDir = fmt.Sprintf("/%s/%s/", "admin", dateStr)
	case global.UploadDealerAvatar:
		subDir = fmt.Sprintf("/%s/%s/", "dealer", dateStr)
	case global.UploadMerchantAccountAvatar:
		subDir = fmt.Sprintf("/%s/%s/", "merchant_account", dateStr)
	case global.UploadPlayerAvatar:
		subDir = fmt.Sprintf("/%s/%s/", "player", dateStr)
	}
	return subDir
}

// UploadImg 通用上传图片
func (h *Common) UploadImg(c *gin.Context) {
	var req request.UploadImgRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Common.UploadImgValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取文件
	file, err := c.FormFile("img")
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	// 获取文件夹路径
	subDir := h.getUploadImgSubDir(req.EventType)

	// 调用通用上传图片服务
	src, err := service.Service.Common.UploadImg(file, subDir, c)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	res := response.UploadImgResponse{
		Src: src,
	}

	utils.Response.SuccessWithData(c, res)
}

// UploadChunk 通用上传切片
func (h *Common) UploadChunk(c *gin.Context) {
	var req request.UploadChunkRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Common.UploadChunkValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// LoginAdmin 登录
func (h *Common) LoginAdmin(c *gin.Context) {
	var req request.LoginAdminRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Admin.LoginAdminValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Admin.LoginAdmin(&req, c)
	if err != nil {
		utils.Response.Fail(c, 0, err.Error(), nil)
		return
	}

	utils.Response.Success(c, utils.TT("http.login.success"), res)
}

// LoginMerchant 登录
func (h *Common) LoginMerchant(c *gin.Context) {
	var req request.LoginMerchantRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Merchant.LoginValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Merchant.Login(&req, c)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// LoginPlayer 登录-测试
func (h *Common) LoginPlayer(c *gin.Context) {
	var req request.LoginPlayerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	// 查询用户
	var player model.Player
	global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Username: req.Username,
	}).First(&player)
	if player.ID == 0 {
		utils.Response.Fail(c, 0, "会员不存在", nil)
		return
	}

	// 密码检查
	if ok := utils.BcryptCheck(req.Password, player.Password); !ok {
		utils.Response.Fail(c, 0, "密码错误", nil)
		return
	}

	utils.Response.SuccessSimple(c)
}
