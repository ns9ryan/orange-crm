package service

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type BetRecord struct{}

// List 投注明细列表
func (s *BetRecord) List(req *request.BetRecordListRequest, userInfo *request.CustomClaims) (*response.ListData, error) {
	var total int64
	var list []response.BetRecordListResponse
	var res response.ListData
	var err error

	db := global.DB.Model(&model.BetRecord{}).Preload("BetArea").Preload("GameCategory")

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的玩家投注明细
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.BetRecordColumns.ParentId: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.BetRecordColumns.ParentId: userInfo.Username,
			})
		}
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

	// 游戏id筛选
	if req.GameTableCode != "" {
		db.Where(map[string]interface{}{
			model.BetRecordColumns.GameId: req.GameTableCode,
		})
	}

	// 玩家账号筛选
	if req.PlayerUsername != "" {
		db.Where(map[string]interface{}{
			model.BetRecordColumns.UserId: req.PlayerUsername,
		})
	}

	// 商户账号筛选
	if req.MerchantUsername != "" {
		db.Where(map[string]interface{}{
			model.BetRecordColumns.ParentId: req.MerchantUsername,
		})
	}

	// 币种筛选
	if req.Currency != "" {
		db.Where(map[string]interface{}{
			model.BetRecordColumns.Currency: req.Currency,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.BetRecordColumns.BetTime), req.StartTime)
	}

	// 开始时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.BetRecordColumns.BetTime), req.EndTime)
	}

	db.Count(&total)

	if total <= 0 {
		res.List = []response.BetRecordListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var betRecords []model.BetRecord
	if err = db.Find(&betRecords).Error; err != nil {
		return &res, err
	}

	for _, betRecord := range betRecords {
		answerData := make(map[string]interface{})
		_ = json.Unmarshal(betRecord.Answer, &answerData)
		tmp := response.BetRecordListResponse{
			ID:           betRecord.ID,                                          // 投注明细ID
			RecordId:     betRecord.RecordId,                                    // 投注明细表的ID
			UserId:       betRecord.UserId,                                      // 帐号的ID
			MerchantId:   betRecord.MerchantId,                                  // 商户ID
			MerchantCode: betRecord.MerchantCode,                                // 商户码
			ParentId:     betRecord.ParentId,                                    // 上级的ID
			RootKey:      betRecord.RootKey,                                     // 顶级Key
			GameId:       betRecord.GameId,                                      // 游戏id
			DeskNo:       betRecord.DeskNo,                                      // 桌号
			Type:         betRecord.Type,                                        // 游戏类型
			SubType:      betRecord.SubType,                                     // 游戏类型二级
			Set:          betRecord.Set,                                         // 桌号
			Round:        betRecord.Round,                                       // 局号
			RoundId:      betRecord.RoundId,                                     // 游戏局数全局ID
			Currency:     betRecord.Currency,                                    // 货币种类
			BetTime:      betRecord.BetTime.Format(global.TIME_STANDARD_MS),     // 投注时间
			BetMode:      betRecord.BetMode,                                     // 投注的模式
			BetAreaId:    betRecord.BetAreaId,                                   // 投注区域ID
			BetAreaName:  betRecord.BetArea.AreaName,                            // 投注区域名称
			BetAmount:    betRecord.BetAmount,                                   // 投注额
			BetEffective: betRecord.BetEffective,                                // 有效投注额
			WinLost:      betRecord.WinLost,                                     // 胜负金额
			Digits:       betRecord.Digits,                                      // 投注额的小数位
			Winner:       betRecord.Winner,                                      // 开奖结果,游戏种类不同
			Value1:       betRecord.Value1,                                      // 开奖参数1
			Value2:       betRecord.Value2,                                      // 开奖参数2
			Value3:       betRecord.Value3,                                      // 开奖参数3
			AnswerTime:   betRecord.AnswerTime.Format(global.TIME_STANDARD_MS),  // 开奖时间
			Answer:       answerData,                                            // 开奖的详细内容
			CreatedTime:  betRecord.CreatedTime.Format(global.TIME_STANDARD_MS), // 创建时间
			UpdatedTime:  betRecord.UpdatedTime.Format(global.TIME_STANDARD_MS), // 更新时间
			Description:  betRecord.GameCategory.Description,                    // 游戏分类描述
		}

		list = append(list, tmp)
	}

	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.BetRecordListResponse{}
	}

	return &res, err
}
