package service

import (
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
)

type BetArea struct{}

// List 下注区域列表
func (s *BetArea) List(req *request.BetAreaListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.BetAreaResponse{},
	}
	db := global.DB.Model(&model.BetArea{})

	// 游戏分类
	if req.GameCategoryCode != "" {
		db.Where(map[string]interface{}{
			model.BetAreaColumns.GameCategoryCode: req.GameCategoryCode,
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
	var betAreas []model.BetArea
	if err = db.Find(&betAreas).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.BetAreaResponse
	for _, betArea := range betAreas {
		tmp := response.BetAreaResponse{
			Id:               betArea.ID,               // ID
			GameCategoryCode: betArea.GameCategoryCode, // 游戏分类
			AreaId:           betArea.AreaId,           // 区域ID
			AreaKey:          betArea.AreaKey,          // 区域标识
			AreaName:         betArea.AreaName,         // 区域名称
			Description:      betArea.Description,      // 下注区域说明
			Odds:             betArea.Odds,             // 赔率
		}

		list = append(list, tmp)
	}
	res.List = list

	return
}

// GetBetAreaBox 下注区域下拉框
func (s *BetArea) GetBetAreaBox() (list []response.BetAreaResponse, err error) {
	// 先从redis读取数据
	data, err := g.Redis().Do(gctx.New(), "GET", "system:bet_area")
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
	if err = global.DB.Model(&model.BetArea{}).Find(&list).Error; err != nil {
		return nil, err
	}
	// 存入redis
	_, err = g.Redis().Do(gctx.New(), "set", "system:bet_area", list)

	return list, err
}

// GetBoxByCategorySub 根据子分类获取下注区域的下拉框数据
func (s *BetArea) GetBoxByCategorySub(req *request.GetBoxByCategoryRequest) (list []response.BetAreaResponse, err error) {

	db := global.DB.Model(&model.BetArea{})

	// 游戏分类
	if req.CategorySub != "" {
		var category model.GameCategory
		if err = global.DB.Where(map[string]interface{}{
			model.GameCategoryColumns.SubCode: req.CategorySub,
		}).First(&category).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gerror.Newf(`{#game_category.sub.code.no.exist}`)
		}

		db.Where(map[string]interface{}{
			model.BetAreaColumns.GameCategoryCode: category.GameCategoryCode,
		})
	}

	// 获取数据
	var betAreas []model.BetArea
	if err = db.Find(&betAreas).Error; err != nil {
		return
	}

	// 列表赋值
	for _, betArea := range betAreas {
		tmp := response.BetAreaResponse{
			Id:               betArea.ID,               // ID
			GameCategoryCode: betArea.GameCategoryCode, // 游戏分类
			AreaId:           betArea.AreaId,           // 区域ID
			AreaKey:          betArea.AreaKey,          // 区域标识
			AreaName:         betArea.AreaName,         // 区域名称
			Description:      betArea.Description,      // 下注区域说明
			Odds:             betArea.Odds,             // 赔率
		}

		list = append(list, tmp)
	}

	return
}

// GetBetAreaWithGroup 获取下注区域并根据分类构建map
func (s *BetArea) GetBetAreaWithGroup() (betAreaGroup map[string][]model.BetArea) {
	// 获取系统下注区域数据
	var betAreas []model.BetArea
	global.DB.Find(&betAreas)

	// 下注区域分类函数
	betAreaGroupFn := func(betArea model.BetArea) string {
		return betArea.GameCategoryCode
	}

	// key是分类code
	return slice.GroupWith(betAreas, betAreaGroupFn) //map[string][]model.BetArea
}
