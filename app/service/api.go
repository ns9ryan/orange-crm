package service

import (
	"errors"
	"github.com/duke-git/lancet/v2/slice"
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

type Api struct{}

// List api列表
func (s *Api) List(req *request.ApiListRequest) (*response.ListData, error) {
	var total int64
	var err error
	var list []response.ApiListResponse
	var res response.ListData

	db := global.DB.Model(&model.Api{})

	db.Count(&total)

	if total <= 0 {
		res.List = []response.ApiListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	err = db.Find(&list).Error

	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.ApiListResponse{}
	}

	return &res, err
}

// Add 创建api
func (s *Api) Add(req *request.SaveApiRequest) error {
	apiModel := model.Api{
		Path:        strings.TrimSpace(req.Path),                    // 路劲
		Method:      strings.TrimSpace(strings.ToUpper(req.Method)), // 请求方法
		Group:       strings.TrimSpace(req.Group),                   // 分组
		Description: req.Description,                                // 描述
		HasDbWrite:  req.HasDbWrite,                                 // 是否对数据库有新增或编辑操作:取值说明在模型里
		IsEditable:  global.CanEdit,                                 // 是否可编辑:取值说明在模型里
		IsDeletable: global.CanDelete,                               // 是否可删除:取值说明在模型里
	}
	if err := global.DB.Create(&apiModel).Error; err != nil {
		global.LOG.Error("创建api失败：", zap.Error(err))
		return err
	}
	return nil
}

// Edit 编辑api
func (s *Api) Edit(req *request.SaveApiRequest) error {
	var apiModel model.Api
	if errors.Is(global.DB.Where("id = ?", req.Id).First(&apiModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	if apiModel.IsEditable == global.CannotEdit {
		return errors.New(utils.MsgEditableErr)
	}

	afterPath := strings.TrimSpace(req.Path)
	afterMethod := strings.TrimSpace(req.Method)

	// 更新casbin
	customCasbin := NewCasbinCustom(global.CASBIN)
	customCasbin.UpdateCasbinApi(apiModel.Path, apiModel.Method, afterPath, afterMethod)

	apiModel.Path = afterPath                     // 路劲
	apiModel.Method = afterMethod                 // 请求方法
	apiModel.Group = strings.TrimSpace(req.Group) // 分组
	apiModel.Description = req.Description        // 描述
	apiModel.HasDbWrite = req.HasDbWrite          // 是否对数据库有新增或编辑操作:取值说明在模型里

	if err := global.DB.Save(&apiModel).Error; err != nil {
		global.LOG.Error("编辑api失败：", zap.Error(err))
		return err
	}

	return nil
}

// Detail api详情
func (s *Api) Detail(id uint) (*response.ApiDetailResponse, error) {
	var res response.ApiDetailResponse
	if errors.Is(global.DB.Model(&model.Api{}).Where("id = ?", id).First(&res).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &res, nil
}

// Del 删除api
func (s *Api) Del(id uint) error {
	var apiModel model.Api
	if errors.Is(global.DB.Where("id = ?", id).First(&apiModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	if apiModel.IsDeletable == global.CannotDelete {
		return errors.New(utils.MsgDeletableErr)
	}

	// 有角色正在使用的菜单不能删除
	var roleCount int64
	global.DB.Model(&model.RoleApi{}).Where("api_id = ?", id).Count(&roleCount)
	if roleCount > 0 {
		return gerror.Newf(`{#api.del.has.role}`)
	}

	// 更新casbin
	customCasbin := NewCasbinCustom(global.CASBIN)
	customCasbin.RemovePolicy(1, apiModel.Path, apiModel.Method)

	if err := global.DB.Delete(&apiModel).Error; err != nil {
		global.LOG.Error("删除失败：", zap.Error(err))
		return err
	}

	return nil
}

// GetApiGroupBox 获取api分组选择框
func (s *Api) GetApiGroupBox() ([]string, error) {
	var nameList []string
	err := global.DB.Model(&model.Api{}).Pluck("group", &nameList).Error
	return slice.Unique(nameList), err
}
