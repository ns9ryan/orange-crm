package service

import (
	"encoding/json"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type Role struct{}

// ConvertRolesToResponse 把角色数据转换成列表返回
func (s *Role) ConvertRolesToResponse(roles []model.Role) []response.RoleListResponse {
	var responses []response.RoleListResponse
	for _, role := range roles {
		// 构造单个 RoleListResponse
		tmp := response.RoleListResponse{
			Id:          role.ID,                                 // 角色ID
			Pid:         role.Pid,                                // 父级ID
			RoleName:    role.RoleName,                           // 角色名称
			Status:      role.Status,                             // 状态:取值说明在模型里
			IsEditable:  role.IsEditable,                         // 是否可编辑:取值说明在模型里
			IsDeletable: role.IsDeletable,                        // 是否可删除:取值说明在模型里
			Children:    s.ConvertRolesToResponse(role.Children), // 递归处理子菜单
		}
		// 添加到结果列表
		responses = append(responses, tmp)
	}
	if responses == nil {
		responses = []response.RoleListResponse{}
	}

	return responses
}

// PreloadChildren 递归获取子集
func (s *Role) PreloadChildren(db *gorm.DB, where map[string]interface{}) *gorm.DB {
	return db.Preload("Children", func(db *gorm.DB) *gorm.DB {
		for k, v := range where {
			db = db.Where(k, v)
		}
		return s.PreloadChildren(db, where) // 递归调用自己
	})
}

// List 菜单列表
func (s *Role) List(req *request.RoleListRequest, roleId uint) (*response.ListData, error) {
	var total int64
	var err error
	var list []response.RoleListResponse
	var res response.ListData

	db := global.DB.Model(&model.Role{})

	if req.Status != nil {
		db = db.Where("status = ?", req.Status)
	}

	db.Count(&total)

	if total <= 0 {
		res.List = []response.RoleListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	var roles []model.Role
	where := map[string]interface{}{}
	if req.Status != nil {
		db.Where("status = ?", req.Status)
		where["status = ?"] = req.Status
	}
	err = s.PreloadChildren(db, where).Where("pid = ?", 0).Find(&roles).Error

	list = s.ConvertRolesToResponse(roles)
	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.RoleListResponse{}
	}

	return &res, err
}

// Add 创建角色
func (s *Role) Add(req *request.SaveRoleRequest) error {
	adminModel := model.Role{
		Pid:         req.Pid,                         // 父级ID
		RoleName:    strings.TrimSpace(req.RoleName), // 角色名称
		Status:      req.Status,                      // 状态:取值说明在模型里
		IsEditable:  global.CanEdit,                  // 是否可编辑:取值说明在模型里
		IsDeletable: global.CanDelete,                // 是否可删除:取值说明在模型里
	}
	if err := global.DB.Create(&adminModel).Error; err != nil {
		global.LOG.Error("创建角色失败：", zap.Error(err))
		return err
	}
	return nil
}

// Edit 编辑角色
func (s *Role) Edit(req *request.SaveRoleRequest) error {
	var roleModel model.Role
	if errors.Is(global.DB.Where("id = ?", req.Id).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	if roleModel.IsEditable == global.CannotEdit {
		return errors.New(utils.MsgEditableErr)
	}

	roleModel.RoleName = strings.TrimSpace(req.RoleName) // 角色名称
	roleModel.Status = req.Status                        // 状态:取值说明在模型里

	// 如果角色底下有账号不能禁用角色
	if roleModel.Status == 0 {
		var adminModel model.Admin
		global.DB.Model(&model.Admin{}).Where("role_id = ?", roleModel.ID).First(&adminModel)
		if adminModel.ID != 0 {
			return gerror.Newf(`{#role.status.has.account.disable}`)
		}
	}
	if err := global.DB.Save(&roleModel).Error; err != nil {
		global.LOG.Error("编辑角色失败：", zap.Error(err))
		return err
	}

	return nil
}

// Detail 角色详情
func (s *Role) Detail(id uint) (*response.RoleDetailResponse, error) {
	var roleModel model.Role
	menuIds := make([]uint, 0)
	apiIds := make([]uint, 0)
	if errors.Is(global.DB.Where("id = ?", id).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	// 角色菜单ID数组
	if err := global.DB.Model(&model.RoleMenu{}).Where("role_id = ?", id).Pluck("menu_id", &menuIds).Error; err != nil {
		return nil, err
	}

	// 角色apiID数组
	if err := global.DB.Model(&model.RoleApi{}).Where("role_id = ?", id).Pluck("api_id", &apiIds).Error; err != nil {
		return nil, err
	}

	res := &response.RoleDetailResponse{
		Id:          roleModel.ID,          // 角色ID
		Pid:         roleModel.Pid,         // 父级ID
		RoleName:    roleModel.RoleName,    // 角色名称
		Status:      roleModel.Status,      // 状态:取值说明在模型里
		IsEditable:  roleModel.IsEditable,  // 是否可编辑:取值说明在模型里
		IsDeletable: roleModel.IsDeletable, // 是否可删除:取值说明在模型里
		MenuIds:     menuIds,               // 角色的菜单Id数组
		ApiIds:      apiIds,                // 角色的apiId数组
	}
	return res, nil
}

// Del 删除角色
func (s *Role) Del(id uint) (roleName string, err error) {
	var roleModel model.Role
	if errors.Is(global.DB.Where("id = ?", id).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	if roleModel.IsDeletable == global.CannotDelete {
		return "", errors.New(utils.MsgDeletableErr)
	}

	// 存在子角色的情况下不能删除
	var sonRole model.Role
	global.DB.Where("pid = ?", id).First(&sonRole)
	if sonRole.ID != 0 {
		return "", gerror.Newf(`{#role.del.subrole.first}`)
	}

	// 角色存在账号的情况下不能删除
	adminModel := model.Admin{}
	global.DB.Where("role_id = ?", id).First(&adminModel)
	if adminModel.ID != 0 {
		return "", gerror.Newf(`{#role.del.account.first}`)
	}

	if err = global.DB.Select("RoleMenus", "RoleApis").Delete(&roleModel).Error; err != nil {
		global.LOG.Error("删除失败：", zap.Error(err))
		return "", err
	}

	customCasbin := NewCasbinCustom(global.CASBIN)
	roleIdStr := gconv.String(id)
	customCasbin.RemovePolicy(0, roleIdStr)

	return roleModel.RoleName, nil
}

// SetRoleMenus 设置角色菜单
func (s *Role) SetRoleMenus(req *request.SetRoleMenusRequest) (roleName string, err error) {
	var roleModel model.Role
	if errors.Is(global.DB.Where("id = ?", req.RoleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	// 如果不是顶级角色且父级不是超级管理员的情况下要检查上级是否有菜单权限，没有的情况下不能设置该菜单
	if roleModel.Pid != 0 {
		// 获取父级
		var parentRoleModel model.Role
		if errors.Is(global.DB.Where("id = ?", roleModel.Pid).First(&parentRoleModel).Error, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		}

		// todo
		// 检查父级是否是超管
		//if parentRoleModel.IsSuper != 1 {
		// 父级角色的菜单ID数组
		var parentMenuIds []uint
		global.DB.Model(&model.RoleMenu{}).Where("role_id = ?", roleModel.Pid).Pluck("menu_id", &parentMenuIds)

		// 如果不包含在父级的菜单中，则返回错误
		if isContain := slice.ContainSubSlice(parentMenuIds, req.MenuIds); !isContain {
			return "", gerror.Newf(`{#role.parent.menu.no.exist}`)
		}
		//}
	}

	// 获取要分配的菜单
	var menus []model.Menu
	global.DB.Where("id in ?", req.MenuIds).Find(&menus)

	// 构建需要插入的数据
	var roleMenuModels []model.RoleMenu
	defaultMenuId := uint(0)
	for _, v := range menus {
		roleMenuModel := model.RoleMenu{
			RoleId: req.RoleId,
			MenuId: v.ID,
		}
		roleMenuModels = append(roleMenuModels, roleMenuModel)
		if defaultMenuId == 0 && v.Path != "" {
			defaultMenuId = v.ID
		}
	}

	tx := global.DB.Begin()

	// 设置默认菜单
	roleModel.DefaultMenuId = defaultMenuId
	if err = global.DB.Save(&roleModel).Error; err != nil {
		return "", err
	}

	// 如果更新的是父级，那么要去掉没有的子角色对于的菜单ID
	var sonModels []model.Role
	global.DB.Model(&model.Role{}).Where("pid = ?", req.RoleId).Find(&sonModels)

	if len(sonModels) != 0 {
		// 重新构建子角色的菜单
		for _, sonRole := range sonModels {

			// 获取子角色的菜单
			var sonMenuIds []uint
			global.DB.Model(&model.RoleMenu{}).Where("role_id = ?", sonRole.ID).Pluck("menu_id", &sonMenuIds)

			// 返回交集菜单ID赋值给子角色菜单
			sonMenuIds = slice.Intersection(sonMenuIds, req.MenuIds)

			// 获取要分配给子角色的菜单
			var sonMenus []model.Menu
			global.DB.Where("id in ?", sonMenuIds).Find(&sonMenus)

			defaultSonMenuId := uint(0)
			// 添加到需要插入的数据数组中
			for _, v := range sonMenus {
				roleMenuModel := model.RoleMenu{
					RoleId: sonRole.ID,
					MenuId: v.ID,
				}
				roleMenuModels = append(roleMenuModels, roleMenuModel)
				if defaultSonMenuId == 0 && v.Path != "" {
					defaultSonMenuId = v.ID
				}
			}

			// 删除子角色原来的菜单
			tx.Where("role_id = ?", sonRole.ID).Delete(&model.RoleMenu{})

			// 设置默认菜单
			sonRole.DefaultMenuId = defaultSonMenuId
			if err = global.DB.Save(&sonRole).Error; err != nil {
				return "", err
			}
		}
	}

	// 先删除
	tx.Where("role_id = ?", req.RoleId).Delete(&model.RoleMenu{})

	// 再创建
	tx.Create(&roleMenuModels)

	return roleModel.RoleName, tx.Commit().Error
}

// SetRoleApis 分配api权限
func (s *Role) SetRoleApis(req *request.SetRoleApisRequest) (roleName string, err error) {
	// ================查询需要的数据=========================================================================================

	// 当前角色
	var roleModel model.Role
	if errors.Is(global.DB.Where("id = ?", req.RoleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	// 需要分配的api
	var apiModels []model.Api
	global.DB.Where(map[string]interface{}{
		model.ApiColumns.Id: req.ApiIds,
	}).Find(&apiModels)

	// 子角色
	var sonRoleModels []model.Role
	global.DB.Where(map[string]interface{}{model.RoleColumns.Pid: req.RoleId}).Find(&sonRoleModels)

	customCasbin := NewCasbinCustom(global.CASBIN)

	// ================校验=========================================================================================

	// 有父级的只能设置父级的api
	if roleModel.Pid != 0 {
		// 父级角色的apiID数组
		var parentApiIds []uint
		global.DB.Model(&model.RoleApi{}).Where(map[string]interface{}{
			model.RoleApiColumns.RoleId: roleModel.Pid,
		}).Pluck("api_id", &parentApiIds)

		// 如果不包含在父级的api中，则返回错误
		if isContain := slice.ContainSubSlice(parentApiIds, req.ApiIds); !isContain {
			return "", gerror.Newf(`{#role.parent.api.no.exist}`)
		}
	}

	// ================构建数据=========================================================================================

	// 构建需要插入的数据
	var roleApiModels []model.RoleApi
	for _, v := range req.ApiIds {
		roleApiModel := model.RoleApi{
			RoleId: req.RoleId,
			ApiId:  v,
		}
		roleApiModels = append(roleApiModels, roleApiModel)
	}

	// ================数据库操作=========================================================================================

	tx := global.DB.Begin()

	// role_api表操作，先删除再创建
	tx.Where("role_id = ?", req.RoleId).Delete(&model.RoleApi{})
	if len(roleApiModels) > 0 {
		if err = tx.Create(&roleApiModels).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	// 重新构建子角色的api，要去掉没有的子角色对于的apiID
	for _, sonRoleModel := range sonRoleModels {
		// 查子角色API
		var sonApiIds []uint
		global.DB.Model(&model.RoleApi{}).Where(map[string]interface{}{
			model.RoleApiColumns.RoleId: sonRoleModel.ID,
		}).Pluck("api_id", &sonApiIds)

		// // 计算交集：返回交集apiID赋值给子角色api
		sonApiIds = slice.Intersection(sonApiIds, req.ApiIds)

		// 构造新的子角色API数据

		// 添加到需要插入的数据数组中
		var sonRoleApiModels []model.RoleApi
		for _, v := range sonApiIds {
			sonRoleApiModels = append(sonRoleApiModels, model.RoleApi{
				RoleId: sonRoleModel.ID,
				ApiId:  v,
			})
		}

		// 删除并重新创建子角色的API
		tx.Where(map[string]interface{}{
			model.RoleApiColumns.RoleId: sonRoleModel.ID,
		}).Delete(&model.RoleApi{})

		if len(sonRoleApiModels) > 0 {
			if err = tx.Create(&sonRoleApiModels).Error; err != nil {
				tx.Rollback()
				return "", err
			}
		}

		var sonApiModels []model.Api
		global.DB.Where(map[string]interface{}{model.ApiColumns.Id: sonApiIds}).Find(&sonApiModels)
		if err = customCasbin.BatchUpdateSubjectsByRole(sonRoleModel.ID, sonApiModels); err != nil {
			return "", err
		}
	}

	if err = tx.Commit().Error; err != nil {
		return "", err
	}

	// ================更新策略=========================================================================================

	// 事务提交后批量更新Casbin权限，避免阻塞事务
	if err = customCasbin.BatchUpdateSubjectsByRole(req.RoleId, apiModels); err != nil {
		return "", err
	}

	return roleModel.RoleName, err
}

// GetMenuListByRole 根据角色获取菜单列表
func (s *Role) GetMenuListByRole(roleId uint, userInfo *request.CustomClaims) (any, error) {
	var list []response.GetMenuListByRoleResponse
	var err error

	var roleModel model.Role
	global.DB.Model(&model.Role{}).Where("id = ?", roleId).Find(&roleModel)
	// 默认按sort字段排序,只显示未禁用的菜单
	db := global.DB.Model(&model.Menu{}).Order("sort")

	db.Where(map[string]interface{}{
		model.MenuColumns.Status: global.CanUse,
	})
	// 如果当前登录的账号不是超管,只展示启用的菜单
	//if userInfo.IsSuper != 1 {
	//	db.Where("status = ?", 1)
	//}

	var menuIds []uint
	global.DB.Model(&model.RoleMenu{}).Where("role_id = ?", roleModel.ID).Pluck("menu_id", &menuIds)
	if len(menuIds) == 0 {
		return make([]map[uint][]response.GetMenuListByRoleResponse, 0), nil
	}
	db.Where("id in ?", menuIds)
	err = db.Find(&list).Error

	// 按照parentId 分组
	treeMap := make(map[uint][]response.GetMenuListByRoleResponse)
	for _, v := range list {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}

	// 获取顶级父级
	superParentId := treeMap[0]
	// 遍历子集
	for i := 0; i < len(superParentId); i++ {
		err = json.Unmarshal([]byte(superParentId[i].LangData), &superParentId[i].LangDataJson)
		err = s.GetMenuListByRoleChildrenList(&superParentId[i], treeMap)
	}

	return &superParentId, err
}

func (s *Role) GetMenuListByRoleChildrenList(menu *response.GetMenuListByRoleResponse, treeMap map[uint][]response.GetMenuListByRoleResponse) (err error) {
	menu.Children = treeMap[menu.Id]
	if menu.Children == nil {
		menu.Children = make([]response.GetMenuListByRoleResponse, 0)
	}
	for i := 0; i < len(menu.Children); i++ {
		err = json.Unmarshal([]byte(menu.Children[i].LangData), &menu.Children[i].LangDataJson)
		err = s.GetMenuListByRoleChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
