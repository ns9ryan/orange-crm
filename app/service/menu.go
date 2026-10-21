package service

import (
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type Menu struct{}

// ConvertMenusToResponse 把菜单模型字段转换成接口返回字段
func (s *Menu) ConvertMenusToResponse(menus []model.Menu) []response.MenuListResponse {
	var responses []response.MenuListResponse
	for _, menu := range menus {
		// 构造单个 MenuListResponse
		tmp := response.MenuListResponse{
			Id:          menu.ID,                                 // 菜单ID
			Pid:         menu.Pid,                                // 父级ID
			MenuName:    menu.MenuName,                           // 菜单名称
			Status:      menu.Status,                             // 状态:取值说明在模型里
			Sort:        menu.Sort,                               // 排序,值越小越前面
			MenuType:    menu.MenuType,                           // 类型:取值说明在模型里
			Description: menu.Description,                        // 描述
			Path:        menu.Path,                               // 跳转路径
			Icon:        menu.Icon,                               // 图标
			LangData:    []response.MenuListLangResponse{},       // 语言
			Children:    s.ConvertMenusToResponse(menu.Children), // 递归处理子菜单
		}
		json.Unmarshal(menu.LangData, &tmp.LangData)
		// 添加到结果列表
		responses = append(responses, tmp)
	}
	if responses == nil {
		responses = []response.MenuListResponse{}
	}

	return responses
}

// PreloadChildren 递归调用获取子集
func (s *Menu) PreloadChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", func(db *gorm.DB) *gorm.DB {
		return s.PreloadChildren(db) // 递归调用自己
	})
}

// List 菜单列表
func (s *Menu) List(req *request.MenuListRequest) (*response.ListData, error) {
	var total int64
	var err error
	var list []response.MenuListResponse
	var res response.ListData

	db := global.DB.Model(&model.Menu{})

	// 状态筛选
	if req.Status != nil {
		db = db.Where("status = ?", req.Status)
	}

	db.Count(&total)

	if total <= 0 {
		res.List = []response.MenuListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	var menus []model.Menu

	err = s.PreloadChildren(db).Where("pid = ?", 0).Order("sort ASC").Find(&menus).Error

	list = s.ConvertMenusToResponse(menus)
	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.MenuListResponse{}
	}

	return &res, err
}

// Add 创建菜单
func (s *Menu) Add(req *request.SaveMenuRequest) error {
	menuModel := model.Menu{
		Pid:         req.Pid,                         // 父级ID
		MenuName:    strings.TrimSpace(req.MenuName), // 菜单名称
		Status:      req.Status,                      // 状态:取值说明在模型里
		Sort:        req.Sort,                        // 排序,值越小越前面
		Description: req.Description,                 // 描述
		Path:        req.Path,                        // 跳转路径
		Icon:        req.Icon,                        // 图标
	}

	// 菜单类型
	if req.MenuType == 1 {
		menuModel.MenuType = model.PageMenu // 页面
	} else {
		menuModel.MenuType = model.BtnMenu // 按钮
	}

	// 菜单语言
	menuModel.LangData, _ = json.Marshal(req.LangData)

	// 排序,值越小越前面
	//lastMenu := model.Menu{}
	//global.DB.Where("pid = ?", req.Pid).Order("sort desc").First(&lastMenu)
	//menuModel.Sort = lastMenu.Sort + 1

	return global.DB.Create(&menuModel).Error
}

// Edit 编辑菜单
func (s *Menu) Edit(req *request.SaveMenuRequest) error {
	var menuModel model.Menu
	if errors.Is(global.DB.Where("id = ?", req.Id).First(&menuModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	menuModel.MenuName = strings.TrimSpace(req.MenuName) // 菜单名称
	menuModel.Status = req.Status                        // 状态:取值说明在模型里
	menuModel.Sort = req.Sort                            // 排序,值越小越前面
	menuModel.Description = req.Description              // 描述
	menuModel.Path = req.Path                            // 跳转路径
	menuModel.Icon = req.Icon                            // 图标
	// 菜单类型
	if req.MenuType == 1 {
		menuModel.MenuType = model.PageMenu // 页面
	} else {
		menuModel.MenuType = model.BtnMenu // 按钮
	}

	menuModel.LangData, _ = json.Marshal(req.LangData)

	return global.DB.Save(&menuModel).Error
}

// Detail 菜单详情
func (s *Menu) Detail(id uint) (*response.MenuDetailResponse, error) {
	var menuModel model.Menu
	if errors.Is(global.DB.Where("id = ?", id).Find(&menuModel).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	res := &response.MenuDetailResponse{
		Id:          menuModel.ID,                      // 菜单ID
		Pid:         menuModel.Pid,                     // 父级ID
		MenuName:    menuModel.MenuName,                // 菜单名称
		Status:      menuModel.Status,                  // 状态:取值说明在模型里
		Sort:        menuModel.Sort,                    // 排序,值越小越前面
		MenuType:    menuModel.MenuType,                // 类型:取值说明在模型里
		Description: menuModel.Description,             // 描述
		Path:        menuModel.Path,                    // 跳转路径
		Icon:        menuModel.Icon,                    // 图标
		LangData:    []response.MenuListLangResponse{}, // 语言
	}
	json.Unmarshal(menuModel.LangData, &res.LangData)
	return res, nil
}

// Del 删除菜单
func (s *Menu) Del(id uint) error {
	var menuModel model.Menu
	if errors.Is(global.DB.Where("id = ?", id).First(&menuModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	// 存在子菜单不能删除
	var sonMenu model.Menu
	global.DB.Where("pid = ?", id).First(&sonMenu)
	if sonMenu.ID != 0 {
		return gerror.Newf(`{#menu.del.submenu.first}`)
	}

	// 有角色正在使用的菜单不能删除
	var roleCount int64
	global.DB.Model(&model.RoleMenu{}).Where("menu_id = ?", id).Count(&roleCount)
	if roleCount > 0 {
		return gerror.Newf(`{#menu.del.has.role}`)
	}

	tx := global.DB.Begin()

	if err := global.DB.Delete(&menuModel).Error; err != nil {
		tx.Rollback()
		global.LOG.Error("删除失败：", zap.Error(err))
		return err
	}

	return tx.Commit().Error
}
