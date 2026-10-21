package service

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
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
)

type TemplatePlan struct{}

// List 限红方案列表
func (s *TemplatePlan) List(req *request.TemplatePlanListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.TemplatePlanListResponse{},
	}
	db := global.DB.Model(&model.TemplatePlan{}).Preload("GameCategory")

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
	var templatePlans []model.TemplatePlan
	if err = db.Find(&templatePlans).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.TemplatePlanListResponse
	for _, templatePlan := range templatePlans {
		tmp := response.TemplatePlanListResponse{
			Id:                  templatePlan.ID,                       // ID
			CategorySub:         templatePlan.CategorySub,              // 游戏子分类
			CategoryDescription: templatePlan.GameCategory.Description, // 分类描述
			Name:                templatePlan.Name,                     // 方案名称
			PlanId:              templatePlan.PlanId,                   // 限红方案ID
			MinValue:            templatePlan.MinValue,                 // 方案的最小投注值
			MaxValue:            templatePlan.MaxValue,                 // 方案的最大投注值
			Digits:              templatePlan.Digits,                   // 方案投注的最小单位倍数
		}

		// 已使用次数
		var usedCountForm int64
		global.DB.Model(&model.TemplatePlanFormArea{}).Where(map[string]interface{}{
			model.TemplatePlanFormAreaColumns.PlanId: templatePlan.PlanId,
		}).Count(&usedCountForm)

		var usedCountMerchant int64
		global.DB.Model(&model.MerchantTemplateBetlimitArea{}).Where(map[string]interface{}{
			model.MerchantTemplateBetlimitAreaColumns.PlanId: templatePlan.PlanId,
		}).Count(&usedCountMerchant)

		tmp.UsedCount = uint(usedCountForm) + uint(usedCountMerchant)

		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建限红方案
func (s *TemplatePlan) Add(c *gin.Context, req *request.SaveTemplatePlanRequest) error {
	var category model.GameCategory
	if err := global.DB.Where(map[string]interface{}{
		model.GameCategoryColumns.SubCode: req.CategorySub,
	}).First(&category).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#game_category.sub.code.no.exist}`)
	}

	// 获取系统下注区域数据
	var betAreas []model.BetArea
	global.DB.Where(map[string]interface{}{
		model.BetAreaColumns.GameCategoryCode: category.GameCategoryCode,
	}).Find(&betAreas)

	planId := guid.S() // 限红方案ID
	templatePlan := model.TemplatePlan{
		CategorySub: req.CategorySub,             // 游戏子分类
		Name:        strings.TrimSpace(req.Name), // 方案名称
		PlanId:      planId,                      // 限红方案ID
	}

	areaMap := make(map[string]request.TemplatePlanArea)
	for _, v := range req.TemplatePlanAreas {
		areaMap[v.AreaId] = v
	}

	// 记录所有投注金额
	betValueArr := make([]uint, 0)

	var templatePlanAreas []model.TemplatePlanArea
	for _, betArea := range betAreas {
		v, ok := areaMap[betArea.AreaId]
		if !ok {
			// 限红方案区域不正确
			return gerror.Newf(`{#template_plan.area.incorrect}`)
		}

		tmp := model.TemplatePlanArea{
			CategorySub:  req.CategorySub,      // 游戏子分类
			AreaId:       betArea.AreaId,       // 区域ID
			AreaName:     betArea.AreaName,     // 区域名称
			MinBetAmount: v.MinBetAmount,       // 最小投注金额
			MaxBetAmount: v.MaxBetAmount,       // 最大投注金额
			Digits:       model.PlanAreaDigits, // 投注的最小单位倍数,100
			PlanId:       planId,               // 限红方案ID
			PlanName:     req.Name,             // 方案名称
		}

		betValueArr = append(betValueArr, v.MinBetAmount)
		betValueArr = append(betValueArr, v.MaxBetAmount)
		templatePlanAreas = append(templatePlanAreas, tmp)
	}

	slice.SortBy(betValueArr, func(a, b uint) bool {
		return a > b
	})

	if len(betValueArr) < 2 {
		return errors.New("系统错误")
	}

	templatePlan.MaxValue = betValueArr[0]                  // 方案的最小投注值
	templatePlan.MinValue = betValueArr[len(betValueArr)-1] // 方案的最大投注值
	templatePlan.Digits = model.PlanAreaDigits              // 方案投注的最小单位倍数

	tx := global.DB.Begin()

	if err := tx.Create(&templatePlan).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(templatePlanAreas).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 存入redis
	if _, err := g.Redis().Do(gctx.New(), "set", fmt.Sprintf("template_plan:%s", planId), templatePlanAreas); err != nil {
		return err
	}

	// 异步发送发送限红方案操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendPlanOptHttp(c, templatePlan, 0)
	}()
	return nil
}

// Edit 编辑限红方案
func (s *TemplatePlan) Edit(c *gin.Context, req *request.SaveTemplatePlanRequest) error {
	var category model.GameCategory
	if err := global.DB.Where(map[string]interface{}{
		model.GameCategoryColumns.SubCode: req.CategorySub,
	}).First(&category).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#game_category.sub.code.no.exist}`)
	}

	// 获取系统下注区域数据
	var betAreas []model.BetArea
	global.DB.Where(map[string]interface{}{
		model.BetAreaColumns.GameCategoryCode: category.GameCategoryCode,
	}).Find(&betAreas)

	var templatePlan model.TemplatePlan
	if err := global.DB.First(&templatePlan, req.Id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	templatePlan.Name = strings.TrimSpace(req.Name) // 方案名称

	var usedCountForm int64
	global.DB.Model(&model.TemplatePlanFormArea{}).Where(map[string]interface{}{
		model.TemplatePlanFormAreaColumns.PlanId: templatePlan.PlanId,
	}).Count(&usedCountForm)
	if usedCountForm > 0 {
		return gerror.Newf(`{#template_plan.used}`) // 限红方案已被使用不能修改/删除
	}

	var usedCountMerchant int64
	global.DB.Model(&model.MerchantTemplateBetlimitArea{}).Where(map[string]interface{}{
		model.MerchantTemplateBetlimitAreaColumns.PlanId: templatePlan.PlanId,
	}).Count(&usedCountMerchant)
	if usedCountMerchant > 0 {
		return gerror.Newf(`{#template_plan.used}`) // 限红方案已被使用不能修改/删除
	}

	areaMap := make(map[string]request.TemplatePlanArea)
	for _, v := range req.TemplatePlanAreas {
		areaMap[v.AreaId] = v
	}

	// 记录所有投注金额
	betValueArr := make([]uint, 0)

	var templatePlanAreas []model.TemplatePlanArea
	for _, betArea := range betAreas {
		v, ok := areaMap[betArea.AreaId]
		if !ok {
			return gerror.Newf(`{#template_plan.area.incorrect}`)
		}

		var tmp model.TemplatePlanArea
		err := global.DB.Where(map[string]interface{}{
			model.TemplatePlanAreaColumns.PlanId: templatePlan.PlanId, // 限红方案ID
			model.TemplatePlanAreaColumns.AreaId: betArea.AreaId,      // 区域ID
		}).First(&tmp).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if tmp.ID == 0 { // 有新的区域
			tmp = model.TemplatePlanArea{
				CategorySub:  req.CategorySub,      // 游戏子分类
				AreaId:       betArea.AreaId,       // 区域ID
				AreaName:     betArea.AreaName,     // 区域名称
				MinBetAmount: v.MinBetAmount,       // 最小投注金额
				MaxBetAmount: v.MaxBetAmount,       // 最大投注金额
				Digits:       model.PlanAreaDigits, // 投注的最小单位倍数,100
				PlanId:       templatePlan.PlanId,  // 限红方案ID
				PlanName:     req.Name,             // 方案名称
			}
		} else {
			tmp.MinBetAmount = v.MinBetAmount // 最小投注金额
			tmp.MaxBetAmount = v.MaxBetAmount // 最大投注金额
			tmp.Digits = model.PlanAreaDigits // 投注的最小单位倍数,100
			tmp.PlanName = req.Name           // 方案名称
		}
		betValueArr = append(betValueArr, v.MinBetAmount)
		betValueArr = append(betValueArr, v.MaxBetAmount)
		templatePlanAreas = append(templatePlanAreas, tmp)
	}

	slice.SortBy(betValueArr, func(a, b uint) bool {
		return a > b
	})

	if len(betValueArr) < 2 {
		return errors.New("系统错误")
	}

	templatePlan.MaxValue = betValueArr[0]                  // 方案的最小投注值
	templatePlan.MinValue = betValueArr[len(betValueArr)-1] // 方案的最大投注值

	tx := global.DB.Begin()

	if err := tx.Save(&templatePlan).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(templatePlanAreas).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 存入redis
	if _, err := g.Redis().Do(gctx.New(), "set", fmt.Sprintf("template_plan:%s", templatePlan.PlanId), templatePlanAreas); err != nil {
		return err
	}

	// 异步发送发送限红方案操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendPlanOptHttp(c, templatePlan, 1)
	}()
	return nil
}

// Detail 限红方案详情
func (s *TemplatePlan) Detail(id uint) (*response.TemplatePlanDetailResponse, error) {
	var templatePlan model.TemplatePlan

	err := global.DB.Model(&model.TemplatePlan{}).Preload("TemplatePlanAreas").Where("id = ?", id).First(&templatePlan).Error

	res := response.TemplatePlanDetailResponse{
		Id:                templatePlan.ID,                       // ID
		CategorySub:       templatePlan.CategorySub,              // 游戏子分类
		Name:              templatePlan.Name,                     // 方案名称
		PlanId:            templatePlan.PlanId,                   // 限红方案ID
		MinValue:          templatePlan.MinValue,                 // 方案的最小投注值
		MaxValue:          templatePlan.MaxValue,                 // 方案的最大投注值
		Digits:            templatePlan.Digits,                   // 方案投注的最小单位倍数
		TemplatePlanAreas: []response.TemplatePlanAreaResponse{}, // 投注区域
	}

	for _, v := range templatePlan.TemplatePlanAreas {
		tmp := response.TemplatePlanAreaResponse{
			Id:           v.ID,           // ID
			AreaId:       v.AreaId,       // 区域ID
			MinBetAmount: v.MinBetAmount, // 最小投注金额
			MaxBetAmount: v.MaxBetAmount, // 最大投注金额
			Digits:       v.Digits,       // 投注的最小单位倍数
			PlanId:       v.PlanId,       // 限红方案ID
		}
		res.TemplatePlanAreas = append(res.TemplatePlanAreas, tmp)
	}
	return &res, err
}

// Del 删除限红方案
func (s *TemplatePlan) Del(c *gin.Context, id uint) (templatePlan model.TemplatePlan, err error) {
	if errors.Is(global.DB.Where("id = ?", id).First(&templatePlan).Error, gorm.ErrRecordNotFound) {
		return model.TemplatePlan{}, gorm.ErrRecordNotFound
	}

	var usedCountForm int64
	global.DB.Model(&model.TemplatePlanFormArea{}).Where(map[string]interface{}{
		model.TemplatePlanFormAreaColumns.PlanId: templatePlan.PlanId,
	}).Count(&usedCountForm)
	if usedCountForm > 0 {
		return model.TemplatePlan{}, gerror.Newf(`{#template_plan.used}`) // 限红方案已被使用不能修改/删除
	}

	var usedCountMerchant int64
	global.DB.Model(&model.MerchantTemplateBetlimitArea{}).Where(map[string]interface{}{
		model.MerchantTemplateBetlimitAreaColumns.PlanId: templatePlan.PlanId,
	}).Count(&usedCountMerchant)
	if usedCountMerchant > 0 {
		return model.TemplatePlan{}, gerror.Newf(`{#template_plan.used}`) // 限红方案已被使用不能修改/删除
	}

	tx := global.DB.Begin()
	tx.Delete(&templatePlan)

	if err = tx.Commit().Error; err != nil {
		return model.TemplatePlan{}, err
	}

	// 异步发送发送限红方案操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendPlanOptHttp(c, templatePlan, 2)
	}()
	return templatePlan, nil
}

// GetBoxByCategorySub 根据子分类获取下注区域的下拉框数据
func (s *TemplatePlan) GetBoxByCategorySub(req *request.GetTemplatePlanBoxByCategoryRequest) (list []response.GetTemplatePlanBoxByCategoryResponse, err error) {

	db := global.DB.Model(&model.TemplatePlan{})

	// 游戏分类
	if req.CategorySub != "" {
		db.Where(map[string]interface{}{
			model.TemplatePlanColumns.CategorySub: req.CategorySub,
		})
	}

	// 获取数据
	var templatePlans []model.TemplatePlan
	if err = db.Find(&templatePlans).Error; err != nil {
		return
	}

	// 列表赋值
	for _, templatePlan := range templatePlans {
		tmp := response.GetTemplatePlanBoxByCategoryResponse{
			Id:          templatePlan.ID,          // ID
			CategorySub: templatePlan.CategorySub, // 游戏子分类
			Name:        templatePlan.Name,        // 方案名称
			PlanId:      templatePlan.PlanId,      // 限红方案ID
			MinValue:    templatePlan.MinValue,    // 方案的最小投注值
			MaxValue:    templatePlan.MaxValue,    // 方案的最大投注值
			Digits:      templatePlan.Digits,      // 方案投注的最小单位倍数
		}
		list = append(list, tmp)
	}

	return
}
