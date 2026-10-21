// =================================================================================
// api请求接收-荷官
// =================================================================================

package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/utils"
	"time"
)

type Dealer struct{}

// GenerateEmployeeNumber 生成工号
func (h *Dealer) GenerateEmployeeNumber(c *gin.Context) {
	employeeNumber, err := service.Service.Dealer.GenerateEmployeeNumber()
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	res := response.DealerEmployeeNumberResponse{
		EmployeeNumber: employeeNumber,
	}
	utils.Response.SuccessWithData(c, res)
}

// UploadAvatar 上传头像
func (h *Dealer) UploadAvatar(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("img")
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	// 文件夹名称
	subDir := fmt.Sprintf("/%s/%s/", "dealer", time.Now().Format("20060102"))

	// 调用通用上传头像服务
	src, err := service.Service.Common.UploadAvatar(file, subDir, c)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	res := response.DealerAvatarResponse{
		Src: src,
	}

	utils.Response.SuccessWithData(c, res)
}

// List 游戏荷官列表
func (h *Dealer) List(c *gin.Context) {
	var req request.DealerListRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Dealer.ListValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	res, err := service.Service.Dealer.List(&req)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}

// Add 创建荷官
func (h *Dealer) Add(c *gin.Context) {
	var req request.SaveDealerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Dealer.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Dealer.Add(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "创建荷官",
		Content:                 fmt.Sprintf("创建荷官:%s", req.EmployeeNumber),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationDealer,
	})
}

// Edit 编辑荷官
func (h *Dealer) Edit(c *gin.Context) {
	var req request.SaveDealerRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.Dealer.SaveValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.Dealer.Edit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "编辑荷官",
		Content:                 fmt.Sprintf("编辑荷官:%s", req.EmployeeNumber),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationDealer,
	})
}

// Detail 荷官详情
func (h *Dealer) Detail(c *gin.Context) {
	id := gconv.Uint(c.Param("id"))
	res, err := service.Service.Dealer.Detail(id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessWithData(c, res)
}

// Del 删除荷官
func (h *Dealer) Del(c *gin.Context) {
	var req request.IdRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}
	employeeNumber, err := service.Service.Dealer.Del(req.Id)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	utils.Response.SuccessSimple(c)

	c.Set("admin_log", model.AdminOperationLog{
		Action:                  "删除荷官",
		Content:                 fmt.Sprintf("删除荷官:%s", employeeNumber),
		IsSuccess:               1,
		AdminOperationEventType: model.AminOperationDealer,
	})
}
