package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/service"
	"orange-crm/utils"
)

type Menu struct{}

// GetMenuList 商户菜单
func (h *Menu) GetMenuList(c *gin.Context) {
	// 如果没有传就获取当前登录的角色ID
	userInfo, err := service.GetMerchantInfo(c)
	if err != nil {
		utils.Response.FailUnauthorized(c)
		return
	}

	var roleId uint
	roleId = userInfo.RoleId

	res, err := service.Service.Role.GetMenuListByRole(roleId, userInfo)
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessWithData(c, res)
}
