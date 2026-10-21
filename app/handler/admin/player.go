package admin

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

	res, err := service.Service.Player.List(&req, nil)
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

	// 查询商户
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.MerchantId,
	}).First(&merchant).Error; err != nil {
		utils.Response.FailValidate(c, gerror.Newf(`{#merchant.id.no.exist}`)) // 商户不存在
		return
	}

	if err := service.Service.Player.Add(&req, merchant); err != nil {
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

	// 查询玩家
	var player model.Player
	if err := global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Id: req.Id,
	}).First(&player).Error; err != nil {
		utils.Response.FailValidate(c, gerror.Newf(`{#play.id.no.exist}`)) // 玩家不存在
		return
	}

	if err := service.Service.Player.Edit(&req, player); err != nil {
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
	res, err := service.Service.Player.Detail(id, nil)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除玩家
func (h *Player) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	username, err := service.Service.Player.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除玩家",
		Content:                 fmt.Sprintf("删除玩家:%s", username),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationPlayer,
	})
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
