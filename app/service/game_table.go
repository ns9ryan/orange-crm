package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type GameTable struct{}

// List 游戏桌台列表
func (s *GameTable) List(req *request.GameTableListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.GameTableListResponse{},
	}
	db := global.DB.Model(&model.GameTable{}).Order("id desc")

	// 游戏子分类
	if req.CategorySub != "" {
		db.Where(map[string]interface{}{
			model.GameTableColumns.CategorySub: req.CategorySub,
		})
	}

	// 游戏桌台标识符筛选
	if req.GameTableCode != "" {
		whereStr := fmt.Sprintf("%s LIKE ?", model.GameTableColumns.GameTableCode)
		db.Where(map[string]interface{}{}).Where(whereStr, "%"+req.GameTableCode+"%")
	}

	// 游戏服务名筛选
	if req.GameServiceName != "" {
		whereStr := fmt.Sprintf("%s LIKE ?", model.GameTableColumns.GameServiceName)
		db.Where(whereStr, "%"+req.GameServiceName+"%")
	}

	// 总条数
	db.Count(&res.Total)
	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var gameTables []model.GameTable
	if err = db.Preload("GameCategory").Find(&gameTables).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.GameTableListResponse
	for _, gameTable := range gameTables {
		tmp := response.GameTableListResponse{
			GameTableInfoResponse: s.getGameTableInfoResponse(gameTable),
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建游戏桌台
func (s *GameTable) Add(c *gin.Context, req *request.SaveGameTableRequest) error {
	var category model.GameCategory
	if err := global.DB.Where(map[string]interface{}{
		model.GameCategoryColumns.SubCode: req.CategorySub,
	}).First(&category).Error; err != nil {
		return err
	}

	if len(req.Description) == 0 {
		req.Description = fmt.Sprintf("%s%03d", category.Description, req.DeskNo)
	}

	gameTable := model.GameTable{
		GameTableCode:    strings.TrimSpace(req.GameTableCode),   // 游戏桌台标识符
		GameCategoryCode: category.GameCategoryCode,              // 游戏分类
		CategorySub:      req.CategorySub,                        // 游戏子分类
		DeskNo:           req.DeskNo,                             // 桌号
		BetInterval:      req.BetInterval,                        // 投注间隔时间,秒
		UnbetRounds:      req.UnbetRounds,                        // 允许未下注局数,局,0表示不限制
		Status:           req.Status,                             // 是否启用:取值说明在模型里
		GameServiceName:  strings.TrimSpace(req.GameServiceName), // 游戏服务名
		StreamUrl:        strings.TrimSpace(req.StreamUrl),       // 视频流地址
		Description:      req.Description,                        // 游戏说明
	}

	if err := global.DB.Create(&gameTable).Error; err != nil {
		return err
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendGameConfigOptHttp(c, gameTable, 0)
	}()

	return nil
}

// Edit 编辑游戏桌台
func (s *GameTable) Edit(c *gin.Context, req *request.SaveGameTableRequest) error {
	var gameTable model.GameTable
	if errors.Is(global.DB.First(&gameTable, req.Id).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	gameTable.BetInterval = req.BetInterval                // 投注间隔时间,秒
	gameTable.UnbetRounds = req.UnbetRounds                // 允许未下注局数,局,0表示不限制
	gameTable.Status = req.Status                          // 是否启用:取值说明在模型里
	gameTable.StreamUrl = strings.TrimSpace(req.StreamUrl) // 视频流地址
	gameTable.Description = req.Description                // 游戏说明

	if err := global.DB.Save(&gameTable).Error; err != nil {
		return err
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendGameConfigOptHttp(c, gameTable, 1)
	}()

	return nil
}

// Detail 游戏桌台详情
func (s *GameTable) Detail(id uint) (*response.GameTableDetailResponse, error) {
	var gameTable model.GameTable

	err := global.DB.Model(&model.GameTable{}).Preload("GameCategory").Where("id = ?", id).First(&gameTable).Error

	res := response.GameTableDetailResponse{
		GameTableInfoResponse: s.getGameTableInfoResponse(gameTable),
	}

	return &res, err
}

// Del 删除游戏桌台
func (s *GameTable) Del(c *gin.Context, id uint) (gameTable model.GameTable, err error) {
	// todo 这里到时候要放开
	//return "", errors.New("该功能暂不开放")
	if errors.Is(global.DB.Where("id = ?", id).First(&gameTable).Error, gorm.ErrRecordNotFound) {
		return model.GameTable{}, gorm.ErrRecordNotFound
	}
	if err = global.DB.Delete(&gameTable).Error; err != nil {
		return gameTable, err
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendGameConfigOptHttp(c, gameTable, 2)
	}()

	return gameTable, global.DB.Delete(&gameTable).Error
}

func (s *GameTable) getGameTableInfoResponse(gameTable model.GameTable) response.GameTableInfoResponse {

	return response.GameTableInfoResponse{
		Id:               gameTable.ID,                                        // 游戏桌台ID
		GameTableCode:    gameTable.GameTableCode,                             // 游戏桌台标识符
		GameCategoryCode: gameTable.GameCategoryCode,                          // 游戏分类
		CategorySub:      gameTable.CategorySub,                               // 游戏子分类
		DeskNo:           gameTable.DeskNo,                                    // 桌号
		BetInterval:      gameTable.BetInterval,                               // 投注间隔时间,秒
		UnbetRounds:      gameTable.UnbetRounds,                               // 允许未下注局数,局,0表示不限制
		Status:           gameTable.Status,                                    // 是否启用:取值说明在模型里
		GameServiceName:  gameTable.GameServiceName,                           // 游戏服务名
		StreamUrl:        gameTable.StreamUrl,                                 // 视频流地址
		Description:      gameTable.Description,                               // 游戏说明
		CreatedTime:      gameTable.CreatedTime.Format("2006-01-02 15:04:05"), // 创建时间
		UpdatedTime:      gameTable.UpdatedTime.Format("2006-01-02 15:04:05"), // 更新时间
		CategoryData: response.GameTableDetailCategoryResponse{
			Id:               gameTable.GameCategory.ID,               // 游戏分类ID
			GameCategoryCode: gameTable.GameCategory.GameCategoryCode, // 游戏分类
			SubCode:          gameTable.GameCategory.SubCode,          // 游戏子分类
			Description:      gameTable.GameCategory.Description,      // 游戏分类描述
		},
	}
}

// GetBoxWithGroup 获取游戏桌台,分组下拉框
func (s *GameTable) GetBoxWithGroup() (list []response.GetGameBoxWithGroupResponse) {
	var categories []model.GameCategory
	global.DB.Preload("GameTables").Preload("TemplatePlans").Find(&categories)

	for _, category := range categories {
		if len(category.GameTables) == 0 {
			continue
		}

		tmp := response.GetGameBoxWithGroupResponse{
			GameCategoryCode: category.GameCategoryCode,                         // 分类
			SubCode:          category.SubCode,                                  // 子分类
			Description:      category.Description,                              // 描述
			GameData:         []response.GameBoxResponse{},                      // 游戏数据
			TemplatePlanData: []response.GetTemplatePlanBoxByCategoryResponse{}, // 限红方案数据
		}

		// 游戏数据
		for _, game := range category.GameTables {
			gameBox := response.GameBoxResponse{
				Id:            game.ID,            // Id
				GameTableCode: game.GameTableCode, // 游戏桌台标识符
				DeskNo:        game.DeskNo,        // 桌号
				Description:   game.Description,   // 游戏说明
			}
			tmp.GameData = append(tmp.GameData, gameBox)
		}

		// 限红方案数据
		for _, plan := range category.TemplatePlans {
			templatePlanBox := response.GetTemplatePlanBoxByCategoryResponse{
				Id:          plan.ID,          // ID
				CategorySub: plan.CategorySub, // 游戏子分类
				Name:        plan.Name,        // 方案名称
				PlanId:      plan.PlanId,      // 限红方案ID
				MinValue:    plan.MinValue,    // 方案的最小投注值
				MaxValue:    plan.MaxValue,    // 方案的最大投注值
				Digits:      plan.Digits,      // 方案投注的最小单位倍数
			}
			tmp.TemplatePlanData = append(tmp.TemplatePlanData, templatePlanBox)
		}

		list = append(list, tmp)
	}

	return
}

// GetBox 游戏桌台下拉框
func (s *GameTable) GetBox() (list []response.GameBoxResponse, err error) {
	var gameTables []model.GameTable
	if err = global.DB.Find(&gameTables).Error; err != nil {
		return
	}
	for _, game := range gameTables {
		tmp := response.GameBoxResponse{
			Id:            game.ID,            // Id
			GameTableCode: game.GameTableCode, // 游戏桌台标识符
			DeskNo:        game.DeskNo,        // 桌号
			Description:   game.Description,   // 游戏说明
		}
		list = append(list, tmp)
	}
	return list, err
}
