package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
)

type Player struct{}

// List 玩家列表
func (h *Player) List(c *gin.Context) {
	var req request.PlayerListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Player.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Player.List(&req, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建玩家
func (h *Player) Add(c *gin.Context) {
	var req request.SavePlayerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Player.SaveValidate(&req); err != nil {
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
		model.MerchantColumns.Username: req.MerchantId,
	}).First(&merchant).Error; err != nil {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if userInfo != nil { // 商户端
		if userInfo.ParentId == global.MerchantRootName { // 主商户：可以创建自己和底下的子商户的会员
			if req.MerchantId != userInfo.Username { // 不是创建自己的
				if merchant.ParentId != userInfo.Username { // 也不是创建子商户的
					utils.Response.FailValidate(c, gerror.Newf(`{#play.add.self}`)) // 只能创建自己或子商户的会员
					return
				}
			}
		} else { // 子商户：只能创建自己的会员
			utils.DD(req.Username)
			if req.MerchantId != userInfo.Username { // 不是创建自己的
				utils.Response.FailValidate(c, gerror.Newf(`{#play.add.self}`)) // 只能创建自己或子商户的会员
				return
			}
		}
	}

	if err = service.Service.Player.Add(&req, merchant); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建玩家",
		Content:                 fmt.Sprintf("创建玩家:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationPlayer,
	})
}

// Edit 编辑玩家
func (h *Player) Edit(c *gin.Context) {
	var req request.SavePlayerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Player.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	// 查询玩家
	var player model.Player
	if err = global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Id:         req.Id,
		model.PlayerColumns.MerchantId: req.MerchantId,
	}).First(&player).Error; err != nil {
		utils.Response.FailValidate(c, gerror.Newf(`{#play.id.no.exist}`)) // 玩家不存在
		return
	}

	if userInfo != nil { // 商户端
		if userInfo.ParentId == global.MerchantRootName { // 主商户：可以编辑自己和底下的子商户的会员
			if req.MerchantId != userInfo.Username { // 不是编辑自己的
				// 查询商户
				var merchant model.Merchant
				if err = global.DB.Where(map[string]interface{}{
					model.MerchantColumns.Username: req.MerchantId,
				}).First(&merchant).Error; err != nil {
					utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
					return
				}

				if merchant.ParentId != userInfo.Username { // 也不是创建子商户的
					utils.Response.FailValidate(c, gerror.Newf(`{#play.edit.self}`)) // 只能编辑自己或子商户的会员
					return
				}
			}
		} else {
			// 子商户：只能创建自己的会员
			if req.MerchantId != userInfo.Username { // 不是编辑自己的
				utils.Response.FailValidate(c, gerror.Newf(`{#play.edit.self}`)) // 只能编辑自己或子商户的会员
				return
			}
		}
	}

	if err = service.Service.Player.Edit(&req, player); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑玩家",
		Content:                 fmt.Sprintf("编辑玩家:%s", req.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationPlayer,
	})
}

// Detail 玩家详情
func (h *Player) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	// 获取登录信息
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	res, err := service.Service.Player.Detail(id, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// ResetPassword 重置玩家密码
func (h *Player) ResetPassword(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	res, err := service.Service.Player.ResetPassword(req.Id, "")

	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "重置玩家密码",
		Content:                 fmt.Sprintf("重置玩家%s的密码", res.Username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationPlayer,
	})
}
