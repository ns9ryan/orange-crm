package service

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
	"time"
)

type TemplatePlanForm struct{}

// List 限红模板列表
func (s *TemplatePlanForm) List(req *request.TemplatePlanFormListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.TemplatePlanFormListResponse{},
	}
	db := global.DB.Model(&model.TemplatePlanForm{})

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
	var templatePlanForms []model.TemplatePlanForm
	if err = db.Find(&templatePlanForms).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.TemplatePlanFormListResponse
	for _, templatePlanForm := range templatePlanForms {
		tmp := response.TemplatePlanFormListResponse{
			Id:          templatePlanForm.ID,                               // ID
			Name:        templatePlanForm.Name,                             // 模板名称
			PlanFormId:  templatePlanForm.PlanFormId,                       // 限红模板ID
			CreatedTime: templatePlanForm.CreatedTime.Format(time.RFC3339), // 创建时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建限红模板
func (s *TemplatePlanForm) Add(req *request.SaveTemplatePlanFormRequest) (err error) {
	planFormId := guid.S() // 限红模板ID
	_, areaData, err := s.makeContent(req, planFormId, strings.TrimSpace(req.Name))
	if err != nil {
		return err
	}

	templatePlanForm := model.TemplatePlanForm{
		PlanFormId: planFormId,                  // 限红模板ID
		Name:       strings.TrimSpace(req.Name), // 模板名称
	}
	// 详细内容
	//templatePlanForm.Content, err = json.Marshal(content)
	//if err != nil {
	//	return
	//}

	tx := global.DB.Begin()

	if err = tx.Create(&templatePlanForm).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&areaData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Edit 编辑限红模板
func (s *TemplatePlanForm) Edit(req *request.SaveTemplatePlanFormRequest) (err error) {
	var templatePlanForm model.TemplatePlanForm
	if err = global.DB.First(&templatePlanForm, req.Id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	templatePlanForm.Name = strings.TrimSpace(req.Name) // 模板名称

	_, areaData, err := s.makeContent(req, templatePlanForm.PlanFormId, templatePlanForm.Name)
	if err != nil {
		return err
	}

	// 详细内容
	//templatePlanForm.Content, err = json.Marshal(content)
	//if err != nil {
	//	return
	//}

	tx := global.DB.Begin()

	if err = tx.Save(&templatePlanForm).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 先删除
	tx.Where(map[string]interface{}{
		model.TemplatePlanFormAreaColumns.PlanFormId: templatePlanForm.PlanFormId,
	}).Delete(&model.TemplatePlanFormArea{})
	// 再重新创建
	if err = tx.Create(&areaData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Detail 限红模板详情
func (s *TemplatePlanForm) Detail(id uint) (*response.TemplatePlanFormDetailResponse, error) {
	var templatePlanForm model.TemplatePlanForm
	if err := global.DB.First(&templatePlanForm, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//var content []map[string]interface{}
	//err := json.Unmarshal(templatePlanForm.Content, &content)
	//if err != nil {
	//	return nil, err
	//}
	res := response.TemplatePlanFormDetailResponse{
		Id:                templatePlanForm.ID,               // ID
		Name:              templatePlanForm.Name,             // 模板名称
		PlanFormId:        templatePlanForm.PlanFormId,       // 限红模板ID
		GameTemplatePlans: []response.GameTemplatePlanArea{}, // 游戏限红模板
		//Content:           content,                           // 游戏限红模板
	}

	var areas []model.TemplatePlanFormArea
	global.DB.Where(map[string]interface{}{
		model.TemplatePlanFormAreaColumns.PlanFormId: templatePlanForm.PlanFormId,
	}).Preload("TemplatePlan").Find(&areas)

	for _, area := range areas {
		tmp := response.GameTemplatePlanArea{
			Id:       area.TemplatePlan.ID,       // ID
			GameId:   area.GameId,                // 游戏ID
			PlanId:   area.PlanId,                // 方案ID
			MinValue: area.TemplatePlan.MinValue, // 方案的最小投注值
			MaxValue: area.TemplatePlan.MaxValue, // 方案的最大投注值
			Digits:   area.TemplatePlan.Digits,   // 投注的最小单位倍数
		}
		res.GameTemplatePlans = append(res.GameTemplatePlans, tmp)
	}

	var err error
	res.GameTemplatePlans, err = slice.UniqueByField(res.GameTemplatePlans, "GameId")
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Del 删除限红模板
func (s *TemplatePlanForm) Del(id uint) (name string, err error) {
	var templatePlanForm model.TemplatePlanForm
	if errors.Is(global.DB.First(&templatePlanForm, id).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	return templatePlanForm.Name, global.DB.Delete(&templatePlanForm).Error
}

// 构建详情数据表字段content存储的值
func (s *TemplatePlanForm) makeContent(req *request.SaveTemplatePlanFormRequest, planFormId, name string) (content []model.GameTemplatePlan, areaData []model.TemplatePlanFormArea, err error) {
	/** 根据前端上传的数据构建map
	[
		"BACX001" => [
			"1001" => [
				{"ID": 14,"CategorySub": "CLASSIC_BAC","AreaId": "1001","AreaName": "闲","MinBetAmount": 100000,"MaxBetAmount": 200000000,"Digits": 100,"PlanId": "9ou42208hw0d8vc1tjsw4t4300pga4zi","PlanName": "1000至200万","CreatedTime": "2025-04-02T16:31:29+08:00","UpdatedTime": "2025-04-02T16:31:29+08:00"}
			]
		]
	]
	*/
	gameMap := make(map[string]map[string]model.TemplatePlanArea)
	for _, gamePlan := range req.GameTemplatePlans {
		// 每个游戏的下注区域其中key为下注区域ID
		gameAreaReq := make(map[string]model.TemplatePlanArea)

		// 先从redis读取数据
		var dataRedis *gvar.Var
		dataRedis, err = g.Redis().Do(gctx.New(), "GET", fmt.Sprintf("template_plan:%s", gamePlan.PlanId))
		if err != nil {
			return
		}

		// 转化成结构体
		var data []model.TemplatePlanArea
		if err = dataRedis.Struct(&data); err != nil {
			return
		}

		// 如果redis没有数据从数据库中读取
		if len(data) == 0 {
			global.DB.Where(map[string]interface{}{
				model.TemplatePlanAreaColumns.PlanId: gamePlan.PlanId,
			}).Find(&data)
		}

		for _, area := range data {
			gameAreaReq[area.AreaId] = area
		}
		gameMap[gamePlan.GameId] = gameAreaReq
	}

	// 获取所有游戏
	var gameTables []model.GameTable
	global.DB.Find(&gameTables)

	// 获取下注区域并根据分类构建map map[string][]model.BetArea
	betAreaGroup := Service.BetArea.GetBetAreaWithGroup()

	for _, gameTable := range gameTables {
		// gameReq为前端上传的每个游戏
		gameReq, ok := gameMap[gameTable.GameTableCode]
		if !ok {
			return content, areaData, gerror.Newf(`{#template_plan_form.game.no.exist}`) // 限红模板游戏不存在
		}

		// 根据游戏分类找到对应的下注区域
		betAreaData, ok := betAreaGroup[gameTable.GameCategoryCode]
		if !ok {
			return content, areaData, gerror.Newf(`{#template_plan_form.betArea.categorySub.incorrect}`) // 限红模板下注区域分类不正确
		}

		gameTemplatePlan := model.GameTemplatePlan{
			GameId:                gameTable.GameTableCode,        // 游戏ID
			GameTemplatePlanAreas: []model.GameTemplatePlanArea{}, // 限红区域
		}

		// 构建限红区域
		for _, betArea := range betAreaData {
			area, ok1 := gameReq[betArea.AreaId]
			if !ok1 {
				return content, areaData, gerror.Newf(`{#template_plan_form.area.incorrect}`) // 限红模板区域不正确
			}

			tmp := model.GameTemplatePlanArea{
				AreaId:       betArea.AreaId,    // 区域ID
				MinBetAmount: area.MinBetAmount, // 最小投注金额
				MaxBetAmount: area.MaxBetAmount, // 最大投注金额
				Digits:       area.Digits,       // 投注的最小单位倍数
			}
			gameTemplatePlan.GameTemplatePlanAreas = append(gameTemplatePlan.GameTemplatePlanAreas, tmp)

			tmp1 := model.TemplatePlanFormArea{
				PlanFormId:   planFormId,              // 限红模板ID
				Name:         name,                    // 模板名称
				GameId:       gameTemplatePlan.GameId, // 游戏ID
				AreaId:       betArea.AreaId,          // 区域ID
				MinBetAmount: area.MinBetAmount,       // 最小投注金额
				MaxBetAmount: area.MaxBetAmount,       // 最大投注金额
				Digits:       area.Digits,             // 投注的最小单位倍数
				PlanId:       area.PlanId,             // 限红方案ID
			}
			areaData = append(areaData, tmp1)
		}

		content = append(content, gameTemplatePlan)
	}

	return content, areaData, err
}

// GetBox 限红模板下拉框数据
func (s *TemplatePlanForm) GetBox() (list []response.TemplatePlanResponse, err error) {
	var planForms []model.TemplatePlanForm
	if err = global.DB.Find(&planForms).Error; err != nil {
		return
	}

	// 赋值
	for _, planForm := range planForms {
		tmp := response.TemplatePlanResponse{
			Id:         planForm.ID,         // ID
			PlanFormId: planForm.PlanFormId, // 限红模板ID
			Name:       planForm.Name,       // 模板名称
		}

		list = append(list, tmp)
	}
	return list, err
}

//func (s *TemplatePlanForm) makeContent1(req *request.SaveTemplatePlanFormRequest) (content []model.GameTemplatePlan, err error) {
//	// 获取下注区域并根据分类构建map map[string][]model.BetArea
//	betAreaGroup := Service.BetArea.GetBetAreaWithGroup()
//
//	/** 根据前端上传的数据构建map
//	[
//		"BACX001" => [
//			"1001" => [
//				{"area_id":"1001","min_bet":"10.11","max_bet":"1000.80"}
//			]
//		]
//	]
//	*/
//	gameMap := make(map[string]map[string]request.GameTemplatePlanArea)
//	for _, gamePlan := range req.GameTemplatePlans {
//		// 每个游戏的下注区域其中key为下注区域ID
//		gameAreaReq := make(map[string]request.GameTemplatePlanArea)
//		for _, gameArea := range gamePlan.GameTemplatePlanAreas {
//			gameAreaReq[gameArea.AreaId] = gameArea
//		}
//		gameMap[gamePlan.GameId] = gameAreaReq
//	}
//
//	// 获取所有游戏
//	var gameTables []model.GameTable
//	global.DB.Find(&gameTables)
//
//	for _, gameTable := range gameTables {
//		// gameReq为前端上传的每个游戏
//		gameReq, ok := gameMap[gameTable.GameTableCode]
//		if !ok {
//			// 限红模板游戏不存在
//			return content, gerror.Newf(`{#template_plan_form.game.no.exist}`)
//		}
//
//		// 根据游戏分类找到对应的下注区域
//		betAreaData, ok := betAreaGroup[gameTable.GameCategoryCode]
//		if !ok {
//			// 限红模板下注区域分类不正确
//			return content, gerror.Newf(`{#template_plan_form.betArea.categorySub.incorrect}`)
//		}
//
//		gameTemplatePlan := model.GameTemplatePlan{
//			GameId:                gameTable.GameTableCode,        // 游戏ID
//			GameTemplatePlanAreas: []model.GameTemplatePlanArea{}, // 限红区域
//		}
//
//		// 构建限红区域
//		for _, betArea := range betAreaData {
//			area, ok1 := gameReq[betArea.AreaId]
//			if !ok1 {
//				// 限红模板区域不正确
//				return content, gerror.Newf(`{#template_plan_form.area.incorrect}`)
//			}
//
//			// 解析浮点数:返回的是数部分,小数部分,换算成整数,小数倍率,错误
//			_, _, minBetInt, minDigits, errTmp := utils.ParseDecimalWithZero(area.MinBet)
//			if errTmp != nil {
//				// 限红模板最小投注金额格式不正确
//				return content, gerror.Newf(`{#template_plan_form.min_bet.incorrect}`)
//			}
//
//			// 解析浮点数:返回的是数部分,小数部分,换算成整数,小数倍率,错误
//			_, _, maxBetInt, maxDigits, errTmp := utils.ParseDecimalWithZero(area.MaxBet)
//			if errTmp != nil {
//				// 限红模板最大投注金额格式不正确
//				return content, gerror.Newf(`{#template_plan_form.max_bet.incorrect}`)
//			}
//
//			if minDigits != maxDigits {
//				// 限红模板最大和最小投注金额小数位必须一样
//				return content, gerror.Newf(`{#template_plan_form.digits.incorrect}`)
//			}
//
//			tmp := model.GameTemplatePlanArea{
//				AreaId:       betArea.AreaId,  // 区域ID
//				MinBetAmount: uint(minBetInt), // 最小投注金额
//				MaxBetAmount: uint(maxBetInt), // 最大投注金额
//				Digits:       uint(minDigits), // 投注的最小单位倍数
//			}
//			gameTemplatePlan.GameTemplatePlanAreas = append(gameTemplatePlan.GameTemplatePlanAreas, tmp)
//		}
//
//		content = append(content, gameTemplatePlan)
//	}
//
//	return content, err
//}
