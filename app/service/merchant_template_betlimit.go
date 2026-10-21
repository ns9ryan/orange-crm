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

type MerchantTemplateBetlimit struct{}

// List 商户限红模板列表
func (s *MerchantTemplateBetlimit) List(req *request.MerchantTemplateBetlimitListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.MerchantTemplateBetlimitListResponse{},
	}
	db := global.DB.Model(&model.MerchantTemplateBetlimit{})

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的模板
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: userInfo.Username,
			})
		}
	}
	// 商户账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.MerchantTemplateBetlimitColumns.Username: req.Username,
		})
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
	var merchantTemplateBetlimits []model.MerchantTemplateBetlimit
	if err = db.Find(&merchantTemplateBetlimits).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.MerchantTemplateBetlimitListResponse
	for _, merchantTemplateBetlimit := range merchantTemplateBetlimits {
		tmp := response.MerchantTemplateBetlimitListResponse{
			Id:          merchantTemplateBetlimit.ID,                               // ID
			Username:    merchantTemplateBetlimit.Username,                         // 商户账号
			Name:        merchantTemplateBetlimit.Name,                             // 模板名称
			TemplateId:  merchantTemplateBetlimit.TemplateId,                       // 商户限红模板ID
			IsDefault:   merchantTemplateBetlimit.IsDefault,                        // 是否默认模板:取值说明在模型里
			CreatedTime: merchantTemplateBetlimit.CreatedTime.Format(time.RFC3339), // 创建时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建商户限红模板
func (s *MerchantTemplateBetlimit) Add(req *request.SaveMerchantTemplateBetlimitRequest, merchant model.Merchant) (err error) {
	templateId := guid.S() // 限红模板ID
	_, areaData, err := s.makeContent(req, templateId)
	if err != nil {
		return err
	}

	merchantTemplateBetlimit := model.MerchantTemplateBetlimit{
		Username:   req.Username,                // 商户账号
		TemplateId: templateId,                  // 商户限红模板ID
		Name:       strings.TrimSpace(req.Name), // 模板名称
		IsDefault:  req.IsDefault,               // 是否默认模板:取值说明在模型里
	}

	// 详细内容
	//merchantTemplateBetlimit.Content, err = json.Marshal(content)
	//if err != nil {
	//	return
	//}

	tx := global.DB.Begin()

	// 设置默认模板
	if req.IsDefault == 1 {
		merchant.TemplateId = templateId // 赋值默认模板
		if err = tx.Save(&merchant).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 如果已存在默认的模板要设置成非默认
		var hasDefault model.MerchantTemplateBetlimit
		global.DB.Where(map[string]interface{}{
			model.MerchantTemplateBetlimitColumns.Username:  merchant.Username,
			model.MerchantTemplateBetlimitColumns.IsDefault: 1,
		}).First(&hasDefault)
		if hasDefault.ID != 0 {
			hasDefault.IsDefault = 0
			if err = tx.Save(&hasDefault).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err = tx.Create(&merchantTemplateBetlimit).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&areaData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Edit 编辑商户限红模板
func (s *MerchantTemplateBetlimit) Edit(req *request.SaveMerchantTemplateBetlimitRequest, merchant model.Merchant) (err error) {
	var merchantTemplateBetlimit model.MerchantTemplateBetlimit
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitColumns.Id:       req.Id,
		model.MerchantTemplateBetlimitColumns.Username: req.Username,
	}).First(&merchantTemplateBetlimit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	merchantTemplateBetlimit.Name = strings.TrimSpace(req.Name) // 模板名称
	merchantTemplateBetlimit.IsDefault = req.IsDefault          // 是否默认模板:取值说明在模型里

	_, areaData, err := s.makeContent(req, merchantTemplateBetlimit.TemplateId)
	if err != nil {
		return err
	}

	// 详细内容
	//merchantTemplateBetlimit.Content, err = json.Marshal(content)
	//if err != nil {
	//	return
	//}

	tx := global.DB.Begin()

	// 设置默认模板
	if req.IsDefault == 1 {
		merchant.TemplateId = merchantTemplateBetlimit.TemplateId // 赋值默认模板
		if err = tx.Save(&merchant).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 如果已存在默认的模板要设置成非默认
		var hasDefault model.MerchantTemplateBetlimit
		global.DB.Where(map[string]interface{}{
			model.MerchantTemplateBetlimitColumns.Username:  merchant.Username,
			model.MerchantTemplateBetlimitColumns.IsDefault: 1,
		}).Where(fmt.Sprintf("%s != ?", model.MerchantTemplateBetlimitColumns.TemplateId), merchantTemplateBetlimit.TemplateId).First(&hasDefault)
		if hasDefault.ID != 0 {
			hasDefault.IsDefault = 0
			if err = tx.Save(&hasDefault).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		if merchant.TemplateId == merchantTemplateBetlimit.TemplateId {
			return gerror.Newf(`{#merchant_template_betlimit.edit.default.one}`) // 至少保留一个模板是默认
		}
	}

	// 更新商户限红模板主表
	if err = tx.Save(&merchantTemplateBetlimit).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 先删除商户限红模板投注区域的相关数据
	tx.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitAreaColumns.TemplateId: merchantTemplateBetlimit.TemplateId,
	}).Delete(&model.MerchantTemplateBetlimitArea{})

	// 再重新创建
	if err = tx.Create(&areaData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Detail 商户限红模板详情
func (s *MerchantTemplateBetlimit) Detail(id uint, userInfo *request.CustomClaims) (*response.MerchantTemplateBetlimitDetailResponse, error) {
	var merchantTemplateBetlimit model.MerchantTemplateBetlimit
	db := global.DB.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitColumns.Id: id,
	})

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的模板
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: userInfo.Username,
			})
		}
	}

	if err := db.First(&merchantTemplateBetlimit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//var content []map[string]interface{}
	//err := json.Unmarshal(merchantTemplateBetlimit.Content, &content)
	//if err != nil {
	//	return nil, err
	//}

	res := response.MerchantTemplateBetlimitDetailResponse{
		Id:                merchantTemplateBetlimit.ID,               // ID
		Username:          merchantTemplateBetlimit.Username,         // 商户账号
		Name:              merchantTemplateBetlimit.Name,             // 模板名称
		TemplateId:        merchantTemplateBetlimit.TemplateId,       // 商户限红模板ID
		IsDefault:         merchantTemplateBetlimit.IsDefault,        // 是否默认模板:取值说明在模型里
		GameTemplatePlans: []response.MerchantGameTemplatePlanArea{}, // 游戏限红模板
		//Content:           content,                                   // 游戏限红模板
	}

	var areas []model.MerchantTemplateBetlimitArea
	global.DB.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitAreaColumns.TemplateId: merchantTemplateBetlimit.TemplateId,
	}).Preload("TemplatePlan").Find(&areas)

	for _, area := range areas {
		tmp := response.MerchantGameTemplatePlanArea{
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

// Del 删除商户限红模板
func (s *MerchantTemplateBetlimit) Del(id uint, userInfo *request.CustomClaims) (name string, err error) {
	var merchantTemplateBetlimit model.MerchantTemplateBetlimit
	if err = global.DB.First(&merchantTemplateBetlimit, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	// 商户端：如果不是删除自己的限红模板
	if userInfo != nil && merchantTemplateBetlimit.Username != userInfo.Username {
		if userInfo.ParentId == global.MerchantRootName { // 如果是主商户
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			if !slice.Contain(usernames, merchantTemplateBetlimit.Username) {
				return "", gerror.Newf(`{#merchant_template_betlimit.del.self}`) // 只能删除自己或子商户的模板
			}
		} else {
			return "", gerror.Newf(`{#merchant_template_betlimit.del.self}`) // 只能删除自己或子商户的模板
		}
	}

	if merchantTemplateBetlimit.IsDefault == 1 {
		return "", gerror.Newf(`{#merchant_template_betlimit.del.default}`) // 默认模板不能删除
	}

	var count int64
	global.DB.Model(&model.MerchantTemplateBetlimit{}).Where(map[string]interface{}{
		model.MerchantTemplateBetlimitColumns.Username: merchantTemplateBetlimit.Username,
	}).Count(&count)
	if count == 1 {
		return "", gerror.Newf(`{#merchant_template_betlimit.del.one}`) // 至少保留一个模板
	}
	return merchantTemplateBetlimit.Name, global.DB.Delete(&merchantTemplateBetlimit).Error
}

// 构建详情数据表字段content存储的值
func (s *MerchantTemplateBetlimit) makeContent(req *request.SaveMerchantTemplateBetlimitRequest, templateId string) (content []model.GameTemplatePlan, areaData []model.MerchantTemplateBetlimitArea, err error) {
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
			return content, areaData, gerror.Newf(`{#merchant_template_betlimit.game.no.exist}`) // 商户限红模板游戏不存在
		}

		// 根据游戏分类找到对应的下注区域
		betAreaData, ok := betAreaGroup[gameTable.GameCategoryCode]
		if !ok {
			return content, areaData, gerror.Newf(`{#merchant_template_betlimit.betArea.categorySub.incorrect}`) // 商户限红模板下注区域分类不正确
		}

		gameTemplatePlan := model.GameTemplatePlan{
			GameId:                gameTable.GameTableCode,        // 游戏ID
			GameTemplatePlanAreas: []model.GameTemplatePlanArea{}, // 限红区域
		}

		// 构建限红区域
		for _, betArea := range betAreaData {
			area, ok1 := gameReq[betArea.AreaId]
			if !ok1 {
				return content, areaData, gerror.Newf(`{#merchant_template_betlimit.area.incorrect}`) // 商户限红模板区域不正确
			}

			tmp := model.GameTemplatePlanArea{
				AreaId:       betArea.AreaId,    // 区域ID
				MinBetAmount: area.MinBetAmount, // 最小投注金额
				MaxBetAmount: area.MaxBetAmount, // 最大投注金额
				Digits:       area.Digits,       // 投注的最小单位倍数
			}
			gameTemplatePlan.GameTemplatePlanAreas = append(gameTemplatePlan.GameTemplatePlanAreas, tmp)

			tmp1 := model.MerchantTemplateBetlimitArea{
				Username:     req.Username,            // 商户账号
				TemplateId:   templateId,              // 商户限红模板ID
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

// GetTemplateBoxByMerchant 获取商户限红模板下拉框数据
func (s *MerchantTemplateBetlimit) GetTemplateBoxByMerchant(req *request.GetTemplateBoxByMerchantRequest, userInfo *request.CustomClaims) (res []response.GetTemplateBoxByMerchantResponse, err error) {

	var list []model.MerchantTemplateBetlimit
	db := global.DB.Where(map[string]interface{}{
		model.MerchantTemplateBetlimitColumns.Username: req.Username,
	})

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的模板
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.MerchantTemplateBetlimitColumns.Username: userInfo.Username,
			})
		}
	}

	db.Find(&list)

	if len(list) == 0 {
		return make([]response.GetTemplateBoxByMerchantResponse, 0), nil
	}

	for _, v := range list {
		res = append(res, response.GetTemplateBoxByMerchantResponse{
			Id:         v.ID,         // ID
			Username:   v.Username,   // 商户账号
			Name:       v.Name,       // 模板名称
			TemplateId: v.TemplateId, // 商户限红模板ID
			IsDefault:  v.IsDefault,  // 是否默认模板:取值说明在模型里
		})
	}
	return
}
