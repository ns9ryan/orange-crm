package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type Menu struct{}

// Group 菜单分组
func (g *Menu) Group(r *gin.RouterGroup) {
	group := r.Group("menu")
	{
		group.GET("getMenuList", merchant.Entry.Menu.GetMenuList) // 获取菜单
	}
}
