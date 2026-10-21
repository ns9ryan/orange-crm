package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
)

type ServiceNode struct{}

// List 服务节点列表
func (h *ServiceNode) List(c *gin.Context) {
	var req request.ServiceNodeListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ServiceNode.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.ServiceNode.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建服务节点
func (h *ServiceNode) Add(c *gin.Context) {
	var req request.SaveServiceNodeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ServiceNode.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.ServiceNode.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建服务节点",
		Content:                 fmt.Sprintf("创建服务节点:%s:%d", req.Ip, req.Port),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationServiceNode,
	})
}

// Edit 编辑服务节点
func (h *ServiceNode) Edit(c *gin.Context) {
	var req request.SaveServiceNodeRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.ServiceNode.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	beforeAddr, err := service.Service.ServiceNode.Edit(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑服务节点",
		Content:                 fmt.Sprintf("修改服务节点:%s:%d为%s", req.Ip, req.Port, beforeAddr),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationServiceNode,
	})
}

// Detail 服务节点详情
func (h *ServiceNode) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.ServiceNode.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Del 删除服务节点
func (h *ServiceNode) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	addr, err := service.Service.ServiceNode.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除服务节点",
		Content:                 fmt.Sprintf("删除服务节点:%s", addr),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationServiceNode,
	})
}
