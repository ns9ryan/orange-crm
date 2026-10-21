package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type GameAnswer struct{}

// List 游戏开答列表
func (s *GameAnswer) List(req *request.GameAnswerListRequest) (*response.ListData, error) {
	var total int64
	var list []response.GameAnswerListResponse
	var res response.ListData
	var err error

	var game model.GameTable
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.GameTableColumns.GameTableCode: req.GameTableCode,
	}).Preload("GameCategory").First(&game).Error, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var _ *gorm.DB
	switch game.GameCategoryCode {
	case "BACX":
		// 百家乐
		total, list, err = fetchAnswerList[*model.AnswerBacx](req, game, &model.AnswerBacx{})
		//total, list, err = s.bacxList(req, game)
	case "DTX":
		// 龙虎
		total, list, err = fetchAnswerList[*model.AnswerDtx](req, game, &model.AnswerDtx{})
		//total, list, err = s.dtxList(req, game)
	case "ROT":
		// 轮盘
		total, list, err = fetchAnswerList[*model.AnswerRot](req, game, &model.AnswerRot{})
		//total, list, err = s.rotList(req, game)
	case "SICBO":
		// 骰宝
		total, list, err = fetchAnswerList[*model.AnswerSicbo](req, game, &model.AnswerSicbo{})
		//total, list, err = s.sicboList(req, game)
	case "FTAN":
		// 番摊
		total, list, err = fetchAnswerList[*model.AnswerFtan](req, game, &model.AnswerFtan{})
	case "SDIE":
		// 色碟
		total, list, err = fetchAnswerList[*model.AnswerSdie](req, game, &model.AnswerSdie{})
	case "FSC":
		// 鱼虾蟹
		total, list, err = fetchAnswerList[*model.AnswerFsc](req, game, &model.AnswerFsc{})
	default:
		return nil, gerror.Newf(`{#game_answer.game_type.incorrect}`)
	}

	res.Total = total
	res.List = list
	return &res, err
}

func (s *GameAnswer) getCommonWhere(db *gorm.DB, req *request.GameAnswerListRequest) *gorm.DB {
	db.Where(map[string]interface{}{
		model.SGAnswerColumns.GameId: req.GameTableCode,
	})

	// 桌号
	if req.DeskNo != "" {
		db.Where(map[string]interface{}{
			model.SGAnswerColumns.DeskNo: req.DeskNo,
		})
	}

	// 场次
	if req.Set != 0 {
		db.Where(map[string]interface{}{
			model.SGAnswerColumns.Set: req.Set,
		})
	}

	// 荷管
	if req.Dealer != "" {
		whereStr := fmt.Sprintf("%s LIKE ?", model.SGAnswerColumns.Dealer)
		db.Where(whereStr, "%"+req.Dealer+"%")
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.SGAnswerColumns.AnswerTime), req.StartTime)
	}

	// 开始时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.SGAnswerColumns.AnswerTime), req.EndTime)
	}

	// 排序
	order := "id desc"
	if req.SortField != nil && req.SortType != nil && slice.Contain(request.SortTypeEnum, strings.ToUpper(*req.SortType)) {
		// 根据ID排序
		if *req.SortField == model.MerchantColumns.Id {
			order = fmt.Sprintf("%s %s", model.MerchantColumns.Id, *req.SortType)
		}
	}
	db.Order(order)

	return db
}

// 通用泛型函数
func fetchAnswerList[T model.IGameAnswerModel](req *request.GameAnswerListRequest, game model.GameTable, modelPtr interface{}) (int64, []response.GameAnswerListResponse, error) {
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(modelPtr)

	// 查询公共搜索条件
	db = (&GameAnswer{}).getCommonWhere(db, req)

	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}
	if req.Paging > 0 {
		db = db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answers []T
	if err := db.Find(&answers).Error; err != nil {
		return total, list, nil
	}

	for _, item := range answers {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(item.GetCards(), &cardData)
		answer := item.ToDTO()
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                     // 投注明细ID
			GameId:      answer.GameId,                 // 游戏id
			DeskNo:      answer.DeskNo,                 // 桌号
			Type:        answer.Type,                   // 游戏类型
			Set:         answer.Set,                    // 桌号
			Round:       answer.Round,                  // 局号
			RoundId:     answer.RoundId,                // 游戏局数全局ID
			Win:         answer.Win,                    // 开奖结果
			Value1:      answer.Value1,                 // 开奖参数1
			Value2:      answer.Value2,                 // 开奖参数2
			Value3:      answer.Value3,                 // 开奖参数3
			Dealer:      answer.Dealer,                 // 开奖的荷管
			AnswerTime:  answer.AnswerTime,             // 开奖时间
			Cards:       cardData,                      // 开奖的荷管
			CreatedTime: answer.CreatedTime,            // 创建时间
			UpdatedTime: answer.UpdatedTime,            // 更新时间
			Description: game.GameCategory.Description, // 游戏分类描述
		}
		list = append(list, tmp)
	}

	return total, list, nil
}

// 百家乐
func (s *GameAnswer) bacxList(req *request.GameAnswerListRequest, game model.GameTable) (int64, []response.GameAnswerListResponse, error) {
	var err error
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(&model.AnswerBacx{})

	// 查询公共搜索条件
	s.getCommonWhere(db, req)
	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answerBacx []model.AnswerBacx
	if err = db.Find(&answerBacx).Error; err != nil {
		return total, list, nil
	}
	for _, answer := range answerBacx {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(answer.Cards, &cardData)
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                                          // 投注明细ID
			GameId:      answer.GameId,                                      // 游戏id
			DeskNo:      answer.DeskNo,                                      // 桌号
			Type:        answer.Type,                                        // 游戏类型
			Set:         answer.Set,                                         // 桌号
			Round:       answer.Round,                                       // 局号
			RoundId:     answer.RoundId,                                     // 游戏局数全局ID
			Win:         answer.Win,                                         // 开奖结果
			Value1:      answer.Value1,                                      // 开奖参数1
			Value2:      answer.Value2,                                      // 开奖参数2
			Value3:      answer.Value3,                                      // 开奖参数3
			Dealer:      answer.Dealer,                                      // 开奖的荷管
			AnswerTime:  answer.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Cards:       cardData,                                           // 开奖的荷管
			CreatedTime: answer.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime: answer.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description: game.GameCategory.Description,                      // 游戏分类描述
		}

		list = append(list, tmp)
	}

	return total, list, nil
}

// 龙虎
func (s *GameAnswer) dtxList(req *request.GameAnswerListRequest, game model.GameTable) (int64, []response.GameAnswerListResponse, error) {
	var err error
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(&model.AnswerDtx{})

	// 查询公共搜索条件
	s.getCommonWhere(db, req)

	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answerDtx []model.AnswerDtx
	if err = db.Find(&answerDtx).Error; err != nil {
		return total, list, nil
	}

	for _, answer := range answerDtx {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(answer.Cards, &cardData)
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                                          // 投注明细ID
			GameId:      answer.GameId,                                      // 游戏id
			DeskNo:      answer.DeskNo,                                      // 桌号
			Type:        answer.Type,                                        // 游戏类型
			Set:         answer.Set,                                         // 桌号
			Round:       answer.Round,                                       // 局号
			RoundId:     answer.RoundId,                                     // 游戏局数全局ID
			Win:         answer.Win,                                         // 开奖结果
			Value1:      answer.Value1,                                      // 开奖参数1
			Value2:      answer.Value2,                                      // 开奖参数2
			Value3:      answer.Value3,                                      // 开奖参数3
			Dealer:      answer.Dealer,                                      // 开奖的荷管
			AnswerTime:  answer.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Cards:       cardData,                                           // 开奖的荷管
			CreatedTime: answer.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime: answer.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description: game.GameCategory.Description,                      // 游戏分类描述
		}

		list = append(list, tmp)
	}

	return total, list, nil
}

// 轮盘
func (s *GameAnswer) rotList(req *request.GameAnswerListRequest, game model.GameTable) (int64, []response.GameAnswerListResponse, error) {
	var err error
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(&model.AnswerRot{})

	// 查询公共搜索条件
	s.getCommonWhere(db, req)
	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answerRot []model.AnswerRot
	if err = db.Find(&answerRot).Error; err != nil {
		return total, list, nil
	}

	for _, answer := range answerRot {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(answer.Cards, &cardData)
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                                          // 投注明细ID
			GameId:      answer.GameId,                                      // 游戏id
			DeskNo:      answer.DeskNo,                                      // 桌号
			Type:        answer.Type,                                        // 游戏类型
			Set:         answer.Set,                                         // 桌号
			Round:       answer.Round,                                       // 局号
			RoundId:     answer.RoundId,                                     // 游戏局数全局ID
			Win:         answer.Win,                                         // 开奖结果
			Value1:      answer.Value1,                                      // 开奖参数1
			Value2:      answer.Value2,                                      // 开奖参数2
			Value3:      answer.Value3,                                      // 开奖参数3
			Dealer:      answer.Dealer,                                      // 开奖的荷管
			AnswerTime:  answer.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Cards:       cardData,                                           // 开奖的荷管
			CreatedTime: answer.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime: answer.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description: game.GameCategory.Description,                      // 游戏分类描述
		}

		list = append(list, tmp)
	}

	return total, list, nil
}

// 番摊
func (s *GameAnswer) ftanList(req *request.GameAnswerListRequest, game model.GameTable) (int64, []response.GameAnswerListResponse, error) {
	var err error
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(&model.AnswerFtan{})

	// 查询公共搜索条件
	s.getCommonWhere(db, req)

	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answerFtan []model.AnswerFtan
	if err = db.Find(&answerFtan).Error; err != nil {
		return total, list, nil
	}
	for _, answer := range answerFtan {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(answer.Cards, &cardData)
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                                          // 投注明细ID
			GameId:      answer.GameId,                                      // 游戏id
			DeskNo:      answer.DeskNo,                                      // 桌号
			Type:        answer.Type,                                        // 游戏类型
			Set:         answer.Set,                                         // 桌号
			Round:       answer.Round,                                       // 局号
			RoundId:     answer.RoundId,                                     // 游戏局数全局ID
			Win:         answer.Win,                                         // 开奖结果
			Value1:      answer.Value1,                                      // 开奖参数1
			Value2:      answer.Value2,                                      // 开奖参数2
			Value3:      answer.Value3,                                      // 开奖参数3
			Dealer:      answer.Dealer,                                      // 开奖的荷管
			AnswerTime:  answer.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Cards:       cardData,                                           // 开奖的荷管
			CreatedTime: answer.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime: answer.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description: game.GameCategory.Description,                      // 游戏分类描述
		}

		list = append(list, tmp)
	}

	return total, list, nil
}

// 骰宝
func (s *GameAnswer) sicboList(req *request.GameAnswerListRequest, game model.GameTable) (int64, []response.GameAnswerListResponse, error) {
	var err error
	var total int64
	list := make([]response.GameAnswerListResponse, 0)

	db := global.DB.Model(&model.AnswerSicbo{})

	// 查询公共搜索条件
	s.getCommonWhere(db, req)

	// 总条数
	db.Count(&total)
	if total <= 0 {
		return total, list, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var answerSicbo []model.AnswerSicbo
	if err = db.Find(&answerSicbo).Error; err != nil {
		return total, list, nil
	}
	for _, answer := range answerSicbo {
		cardData := make(map[string]interface{})
		_ = json.Unmarshal(answer.Cards, &cardData)
		tmp := response.GameAnswerListResponse{
			ID:          answer.ID,                                          // 投注明细ID
			GameId:      answer.GameId,                                      // 游戏id
			DeskNo:      answer.DeskNo,                                      // 桌号
			Type:        answer.Type,                                        // 游戏类型
			Set:         answer.Set,                                         // 桌号
			Round:       answer.Round,                                       // 局号
			RoundId:     answer.RoundId,                                     // 游戏局数全局ID
			Win:         answer.Win,                                         // 开奖结果
			Value1:      answer.Value1,                                      // 开奖参数1
			Value2:      answer.Value2,                                      // 开奖参数2
			Value3:      answer.Value3,                                      // 开奖参数3
			Dealer:      answer.Dealer,                                      // 开奖的荷管
			AnswerTime:  answer.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Cards:       cardData,                                           // 开奖的荷管
			CreatedTime: answer.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime: answer.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description: game.GameCategory.Description,                      // 游戏分类描述
		}

		list = append(list, tmp)
	}

	return total, list, nil
}
