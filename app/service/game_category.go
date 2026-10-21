package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
)

type GameCategory struct{}

// List 游戏分类列表
func (s *GameCategory) List(req *request.GameCategoryListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.GameCategoryResponse{},
	}
	db := global.DB.Model(&model.GameCategory{})

	if req.GameCategoryCode != "" {
		db.Where(map[string]interface{}{
			model.GameCategoryColumns.GameCategoryCode: req.GameCategoryCode,
		})
	}

	// 总条数
	db.Count(&res.Total)
	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	// 分页
	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var categories []model.GameCategory
	if err = db.Find(&categories).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.GameCategoryResponse
	for _, category := range categories {
		tmp := response.GameCategoryResponse{
			Id:               category.ID,               // ID
			GameCategoryCode: category.GameCategoryCode, // 分类
			SubCode:          category.SubCode,          // 子分类
			Description:      category.Description,      // 描述
			Status:           category.Status,           // 是否启用:取值说明在模型里
			IsHide:           category.IsHide,           // 是否隐藏:取值说明在模型里
		}

		list = append(list, tmp)
	}
	res.List = list

	return
}

// GetGameCategoryBox 游戏分类下拉框
func (s *GameCategory) GetGameCategoryBox() (list []response.GameCategoryResponse, err error) {
	// 先从redis读取数据
	data, err := g.Redis().Do(gctx.New(), "GET", "system:game_category")
	if err != nil {
		return nil, err
	}
	if err = data.Struct(&list); err != nil {
		return nil, err
	}
	// 如果存在直接返回
	if len(list) != 0 {
		return list, err
	}

	// 从数据库种读取数据
	if err = global.DB.Model(&model.GameCategory{}).Find(&list).Error; err != nil {
		return nil, err
	}
	// 存入redis
	_, err = g.Redis().Do(gctx.New(), "set", "system:game_category", list)

	return list, err
}
