package service

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"sort"
	"time"
)

type MerchantDayaccount struct{}

// 返回钱包表的过滤条件
func (s *MerchantDayaccount) getWalletWhere() (walletWhere string) {
	return fmt.Sprintf("%s != 0 OR %s != 0 OR %s != 0 OR %s != 0 OR %s != 0", model.MerchantCurrencyColumns.LastBet, model.MerchantCurrencyColumns.LastBetEffective, model.MerchantCurrencyColumns.LastWater, model.MerchantCurrencyColumns.LastWinLost, model.MerchantCurrencyColumns.LastTip)
}

// =====================================================================================================================================================================主商户输赢报表

// GetMerchantReportForAll 总网端查询主商户输赢报表
func (s *MerchantDayaccount) GetMerchantReportForAll(req *request.MerchantDayaccountListRequest) (*response.ListData, error) {
	var merchantIds []string
	var dayAccountData, currencyData []response.MerchantDayaccountListResponse
	var err error

	// 判断查询时间属于哪种类型
	rangeType, startTime, endTime, err := utils.DetectTimeRangeType(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	hasChildren := uint(1)
	// 根据时间范围类型选择相应的查询逻辑
	switch rangeType {
	case utils.TimeRangeToday:
		// 第一种情况: 仅查询今天及之后的数据
		if merchantIds, err = s.getMerchantIdsFromCurrency(startTime, endTime, req.MerchantId); err != nil {
			return nil, err
		}
		if currencyData, err = s.getCurrencyData(startTime, endTime, merchantIds, hasChildren); err != nil {
			return nil, err
		}
	case utils.TimeRangeHistory:
		// 第二种情况: 仅查询今天之前的数据
		if merchantIds, err = s.getMerchantIdsFromDayaccount(startTime, endTime, req.MerchantId); err != nil {
			return nil, err
		}
		if dayAccountData, err = s.getDayaccountData(startTime, endTime, merchantIds, hasChildren); err != nil {
			return nil, err
		}
	case utils.TimeRangeMixed:
		// 第三种情况: 包含今天之前和今天及之后
		var fromDayaccount, fromCurrency []string
		// 查询 merchant_dayaccount 和 merchant_currency 两张表
		fromDayaccount, err = s.getMerchantIdsFromDayaccount(startTime, endTime, req.MerchantId)
		if err != nil {
			return nil, err
		}
		fromCurrency, err = s.getMerchantIdsFromCurrency(startTime, endTime, req.MerchantId)
		if err != nil {
			return nil, err
		}
		// 合并去重
		merchantIds = slice.Union(fromCurrency, fromDayaccount)

		// 查询 merchant_currency（今天及之后）和 merchant_dayaccount（今天之前）的数据
		if currencyData, err = s.getCurrencyData(startTime, endTime, merchantIds, hasChildren); err != nil {
			return nil, err
		}
		if dayAccountData, err = s.getDayaccountData(startTime, endTime, merchantIds, hasChildren); err != nil {
			return nil, err
		}
	default:
		return nil, gerror.Newf(`{#query.type.exception}`) // 查询类型异常
	}

	if len(merchantIds) == 0 {
		return &response.ListData{Total: 0, List: []response.MerchantDayaccountListResponse{}}, nil
	}

	// 合并并聚合数据（按商户和成数分组）
	finalList := s.mergeAndGroupByRate(dayAccountData, currencyData)

	res := &response.ListData{
		Total: int64(len(finalList)),
		List:  []response.MerchantDayaccountListResponse{},
	}

	// 分页处理
	if req.Paging > 0 {
		// 把list分割赋值给res.List
		sliceList := slice.Chunk(finalList, *req.PageSize)
		index := *req.Page - 1
		if index >= 0 && index < len(sliceList) {
			res.List = sliceList[index]
		}
	}

	return res, nil
}

// GetMerchantReportForSelf 商户端查询自己的输赢报表
func (s *MerchantDayaccount) GetMerchantReportForSelf(req *request.MerchantDayaccountListRequest, username string) (*response.ListData, error) {
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: username,
	}).First(&merchant).Error; err != nil {
		return nil, err
	}

	var (
		dayAccountData []response.MerchantDayaccountListResponse
		currencyData   []response.MerchantDayaccountListResponse
	)

	// 判断查询时间属于哪种类型
	rangeType, startTime, endTime, err := utils.DetectTimeRangeType(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	// 主商户底下有子商户，子商户没有
	var hasChildren uint
	if merchant.ParentId == global.MerchantRootName {
		hasChildren = 1
	}
	switch rangeType {
	case utils.TimeRangeToday:
		// 第一种情况: 仅查询今天及之后的数据
		if currencyData, err = s.getCurrencyDataForSelf(startTime, endTime, username, hasChildren); err != nil {
			return nil, err
		}
	case utils.TimeRangeHistory:
		// 第二种情况: 仅查询今天之前的数据
		if dayAccountData, err = s.getDayaccountDataForSelf(startTime, endTime, username, hasChildren); err != nil {
			return nil, err
		}
	case utils.TimeRangeMixed:
		// 第三种情况: 包含今天之前和今天及之后
		// 查询 merchant_currency（今天及之后）和 merchant_dayaccount（今天之前）的数据
		if currencyData, err = s.getCurrencyDataForSelf(startTime, endTime, username, hasChildren); err != nil {
			return nil, err
		}
		if dayAccountData, err = s.getDayaccountDataForSelf(startTime, endTime, username, hasChildren); err != nil {
			return nil, err
		}
	default:
		return nil, gerror.Newf(`{#query.type.exception}`) // 查询类型异常
	}

	if len(dayAccountData) == 0 && len(currencyData) == 0 {
		return &response.ListData{Total: 0, List: []response.MerchantDayaccountListResponse{}}, nil
	}

	// 合并并聚合数据（按商户和成数分组）
	finalList := s.mergeAndGroupByRate(dayAccountData, currencyData)

	// 其实商户就一条数据是自己的
	res := &response.ListData{
		Total: int64(len(finalList)),
		List:  []response.MerchantDayaccountListResponse{},
	}

	// 分页处理
	if req.Paging > 0 {
		// 把list分割赋值给res.List
		sliceList := slice.Chunk(finalList, *req.PageSize)
		index := *req.Page - 1
		if index >= 0 && index < len(sliceList) {
			res.List = sliceList[index]
		}
	}

	return res, nil
}

// 总网端：获取商户ID列表：查询今天及之后的数据（来自 merchant_currency），根据merchant_currency表获取主商户账号
func (s *MerchantDayaccount) getMerchantIdsFromCurrency(startTime, endTime time.Time, merchantId string) (ids []string, err error) {
	// 查询 merchant 表找到主商户
	var merchantIds []string
	tx := global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.ParentId: global.MerchantRootName, // 主商户
	})
	// 单个商户查询
	if merchantId != "" {
		tx.Where(map[string]interface{}{
			model.MerchantColumns.Username: merchantId,
		})
	}
	tx.Pluck(model.MerchantColumns.Username, &merchantIds)
	if len(merchantIds) == 0 {
		return
	}

	// 查询 merchant_currency 中至少有一个字段不为0
	global.DB.Model(&model.MerchantCurrency{}).Where(s.getWalletWhere()).Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: merchantIds,
	}).Pluck(model.MerchantCurrencyColumns.Username, &ids)

	return
}

// 总网端：获取商户ID列表：查询今天之前的数据（来自 merchant_dayaccount），根据merchant_dayaccount表获取主商户账号
func (s *MerchantDayaccount) getMerchantIdsFromDayaccount(startTime, endTime time.Time, merchantId string) (ids []string, err error) {
	// 历史表有数据的查询条件
	dayaccountStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.MerchantDayaccountColumns.AccountDate)
	tx := global.DB.Model(&model.MerchantDayaccount{}).Where(dayaccountStr, startTime, endTime).Where(map[string]interface{}{
		model.MerchantDayaccountColumns.ParentId: global.MerchantRootName, // 主商户
	})
	// 单个商户查询
	if merchantId != "" {
		tx = tx.Where(map[string]interface{}{
			model.MerchantDayaccountColumns.MerchantId: merchantId,
		})
	}

	tx.Pluck(model.MerchantDayaccountColumns.MerchantId, &ids)
	return
}

// 总网端：（查询 merchant_dayaccount 数据）
func (s *MerchantDayaccount) getDayaccountData(startTime, endTime time.Time, merchantIds []string, hasChildren uint) (data []response.MerchantDayaccountListResponse, err error) {
	var dayDataRaw []model.MerchantDayaccount
	dayaccountStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.MerchantDayaccountColumns.AccountDate)
	err = global.DB.Where(dayaccountStr, startTime, endTime).Where(map[string]interface{}{
		model.MerchantDayaccountColumns.MerchantId: merchantIds,
	}).Find(&dayDataRaw).Error

	for _, v := range dayDataRaw {
		data = append(data, response.MerchantDayaccountListResponse{
			MerchantId:       v.MerchantId,                                                                                                                              // 对应商户表的username
			ParentId:         v.ParentId,                                                                                                                                // 上级商户的userName
			Currency:         v.Currency,                                                                                                                                // 币种
			Digits:           v.Digits,                                                                                                                                  // 币种的最小单位倍数
			LastBet:          gconv.String(v.Bet),                                                                                                                       // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.BetEffective),                                                                                                              // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.Water),                                                                                                                     // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.WinLost),                                                                                                                   // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.Tip),                                                                                                                       // 累计打赏金额，存储为最小单位
			StartTime:        v.AccountDate.Format(time.RFC3339),                                                                                                        // 开始时间
			EndTime:          time.Date(v.AccountDate.Year(), v.AccountDate.Month(), v.AccountDate.Day(), 23, 59, 59, 0, v.AccountDate.Location()).Format(time.RFC3339), // 结束时间
			HasChildren:      hasChildren,                                                                                                                               // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                                                                                                                          // 成数
			RollingRate:      v.RollingRate,                                                                                                                             // 洗码比
			AccountDate:      v.AccountDate.Format(time.RFC3339),                                                                                                        // 结算归属日期
		})
	}

	return
}

// 总网端：（查询 merchant_currency 数据）
func (s *MerchantDayaccount) getCurrencyData(startTime, endTime time.Time, merchantIds []string, hasChildren uint) (data []response.MerchantDayaccountListResponse, err error) {
	// 实际上这个函数的 startTime/endTime 不起作用，因为 merchant_currency 是当前实时数据，仅用 merchantIds 即可
	// 获取当天的开始时间和结束时间
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 今天的 00:00:00
	todayEndTime := todayStartTime.Add(24*time.Hour - time.Second)                              // 今天的 23:59:59

	var currencyDataRaw []model.MerchantCurrency
	// 查询 merchant_currency 中至少有一个字段不为0
	err = global.DB.Model(&model.MerchantCurrency{}).Where(s.getWalletWhere()).Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: merchantIds,
	}).Find(&currencyDataRaw).Error

	for _, v := range currencyDataRaw {
		var parentId string
		global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
			model.MerchantColumns.Username: v.Username,
		}).Pluck(model.MerchantColumns.ParentId, &parentId)

		data = append(data, response.MerchantDayaccountListResponse{
			MerchantId:       v.Username,                          // 对应商户表的username
			ParentId:         parentId,                            // 上级商户的userName
			Currency:         v.Currency,                          // 币种
			Digits:           v.Digits,                            // 币种的最小单位倍数
			LastBet:          gconv.String(v.LastBet),             // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.LastBetEffective),    // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.LastWater),           // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.LastWinLost),         // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.LastTip),             // 累计打赏金额，存储为最小单位
			StartTime:        todayStartTime.Format(time.RFC3339), // 开始时间
			EndTime:          todayEndTime.Format(time.RFC3339),   // 结束时间
			HasChildren:      hasChildren,                         // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                    // 成数
			RollingRate:      v.RollingRate,                       // 洗码比
			AccountDate:      time.Now().Format(time.DateOnly),    // 补上这个字段
		})
	}
	return
}

// 商户端：（查询 merchant_dayaccount 数据）
func (s *MerchantDayaccount) getDayaccountDataForSelf(startTime, endTime time.Time, username string, hasChildren uint) (data []response.MerchantDayaccountListResponse, err error) {
	var dayDataRaw []model.MerchantDayaccount
	dayaccountStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.MerchantDayaccountColumns.AccountDate)
	err = global.DB.Where(dayaccountStr, startTime, endTime).Where(map[string]interface{}{
		model.MerchantDayaccountColumns.MerchantId: username,
	}).Find(&dayDataRaw).Error

	for _, v := range dayDataRaw {
		data = append(data, response.MerchantDayaccountListResponse{
			MerchantId:       v.MerchantId,                                                                                                                              // 对应商户表的username
			ParentId:         v.ParentId,                                                                                                                                // 上级商户的userName
			Currency:         v.Currency,                                                                                                                                // 币种
			Digits:           v.Digits,                                                                                                                                  // 币种的最小单位倍数
			LastBet:          gconv.String(v.Bet),                                                                                                                       // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.BetEffective),                                                                                                              // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.Water),                                                                                                                     // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.WinLost),                                                                                                                   // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.Tip),                                                                                                                       // 累计打赏金额，存储为最小单位
			StartTime:        v.AccountDate.Format(time.RFC3339),                                                                                                        // 开始时间
			EndTime:          time.Date(v.AccountDate.Year(), v.AccountDate.Month(), v.AccountDate.Day(), 23, 59, 59, 0, v.AccountDate.Location()).Format(time.RFC3339), // 结束时间
			HasChildren:      hasChildren,                                                                                                                               // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                                                                                                                          // 成数
			RollingRate:      v.RollingRate,                                                                                                                             // 洗码比
			AccountDate:      v.AccountDate.Format(time.RFC3339),                                                                                                        // 结算归属日期
		})
	}
	return
}

// 商户端：（查询 merchant_currency 数据）
func (s *MerchantDayaccount) getCurrencyDataForSelf(startTime, endTime time.Time, username string, hasChildren uint) (data []response.MerchantDayaccountListResponse, err error) {
	// 实际上这个函数的 startTime/endTime 不起作用，因为 merchant_currency 是当前实时数据，仅用 merchantIds 即可
	// 获取当天的开始时间和结束时间
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 今天的 00:00:00
	todayEndTime := todayStartTime.Add(24*time.Hour - time.Second)                              // 今天的 23:59:59

	var currencyDataRaw []model.MerchantCurrency
	// 查询 merchant_currency 中至少有一个字段不为0
	err = global.DB.Model(&model.MerchantCurrency{}).Where(s.getWalletWhere()).Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: username,
	}).Find(&currencyDataRaw).Error

	for _, v := range currencyDataRaw {
		var parentId string
		global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
			model.MerchantColumns.Username: v.Username,
		}).Pluck(model.MerchantColumns.ParentId, &parentId)

		data = append(data, response.MerchantDayaccountListResponse{
			MerchantId:       v.Username,                          // 对应商户表的username
			ParentId:         parentId,                            // 上级商户的userName
			Currency:         v.Currency,                          // 币种
			Digits:           v.Digits,                            // 币种的最小单位倍数
			LastBet:          gconv.String(v.LastBet),             // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.LastBetEffective),    // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.LastWater),           // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.LastWinLost),         // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.LastTip),             // 累计打赏金额，存储为最小单位
			StartTime:        todayStartTime.Format(time.RFC3339), // 开始时间
			EndTime:          todayEndTime.Format(time.RFC3339),   // 结束时间
			HasChildren:      hasChildren,                         // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                    // 成数
			RollingRate:      v.RollingRate,                       // 洗码比
			AccountDate:      time.Now().Format(time.DateOnly),    // 补上这个字段
		})
	}
	return
}

// 合并并聚合数据（按商户和成数分组）
func (s *MerchantDayaccount) mergeAndGroupByRate(dayData, currencyData []response.MerchantDayaccountListResponse) []response.MerchantDayaccountListResponse {
	// 1. 对 dayData 按商户分组并按日期升序排列
	dayGroup := make(map[string][]response.MerchantDayaccountListResponse)
	for _, d := range dayData {
		dayGroup[d.MerchantId] = append(dayGroup[d.MerchantId], d)
	}
	for merchantId := range dayGroup {
		sort.Slice(dayGroup[merchantId], func(i, j int) bool {
			return dayGroup[merchantId][i].AccountDate < dayGroup[merchantId][j].AccountDate
		})
	}

	// 2. 对每个商户的 dayData 进行“按连续日期 + 相同成数”聚合
	groupedMap := make(map[string][]response.MerchantDayaccountListResponse)
	for merchantId, list := range dayGroup {

		var grouped []response.MerchantDayaccountListResponse // 已分组好的结果
		// 同个商户底下的所有dayData数据
		for _, item := range list {
			n := len(grouped)
			if n == 0 {
				grouped = append(grouped, item)
				continue
			}
			last := &grouped[n-1] // 上一条已分组的数据
			// 如果成数一致，合并数据
			if last.CommissionRate == item.CommissionRate {
				// 合并逻辑
				last.LastBet = utils.SumStringInt(last.LastBet, item.LastBet)
				last.LastBetEffective = utils.SumStringInt(last.LastBetEffective, item.LastBetEffective)
				last.LastWater = utils.SumStringInt(last.LastWater, item.LastWater)
				last.LastWinLost = utils.SumStringInt(last.LastWinLost, item.LastWinLost)
				last.LastTip = utils.SumStringInt(last.LastTip, item.LastTip)
				//last.AccountDate = item.AccountDate // 更新最新日期
				// 合并开始时间和结束时间
				if last.StartTime > item.StartTime {
					last.StartTime = item.StartTime // 最早开始时间
				}
				if last.EndTime < item.EndTime {
					last.EndTime = item.EndTime // 最晚结束时间
				}
			} else {
				grouped = append(grouped, item)
			}
		}
		groupedMap[merchantId] = grouped
	}
	// 3. 合并 currencyData：如果商户在 dayData 中有最近记录，成数一致就合并，否则单独添加
	for _, cur := range currencyData {
		dayList := groupedMap[cur.MerchantId]
		if len(dayList) == 0 {
			// 没有 dayData，直接新增
			groupedMap[cur.MerchantId] = append(groupedMap[cur.MerchantId], cur)
			continue
		}
		last := &dayList[len(dayList)-1] // 最近的一条
		if last.CommissionRate == cur.CommissionRate {
			last.LastBet = utils.SumStringInt(last.LastBet, cur.LastBet)
			last.LastBetEffective = utils.SumStringInt(last.LastBetEffective, cur.LastBetEffective)
			last.LastWater = utils.SumStringInt(last.LastWater, cur.LastWater)
			last.LastWinLost = utils.SumStringInt(last.LastWinLost, cur.LastWinLost)
			last.LastTip = utils.SumStringInt(last.LastTip, cur.LastTip)
			last.AccountDate = cur.AccountDate // 更新最新日期
			// 合并开始时间和结束时间
			if last.StartTime > cur.StartTime {
				last.StartTime = cur.StartTime // 最早开始时间
			}
			if last.EndTime < cur.EndTime {
				last.EndTime = cur.EndTime // 最晚结束时间
			}
		} else {
			groupedMap[cur.MerchantId] = append(groupedMap[cur.MerchantId], cur)
		}
	}
	// 4. 汇总所有结果
	var result []response.MerchantDayaccountListResponse
	for _, list := range groupedMap {
		result = append(result, list...)
	}

	// 排序：因为mysql的in查询是无序的会导致结果也是无序，先按商户ID排序，再按日期排序
	sort.Slice(result, func(i, j int) bool {
		if result[i].MerchantId == result[j].MerchantId {
			return result[i].AccountDate < result[j].AccountDate // 时间倒序
		}
		return result[i].MerchantId < result[j].MerchantId // 倒序
	})

	return result
}

// =====================================================================================================================================================================主商户下子商户的输赢报表

// GetSubMerchantDayaccount 获取主商户下子商户的输赢报表
func (s *MerchantDayaccount) GetSubMerchantDayaccount(req *request.GetSubMerchantDayaccountRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.MerchantDayaccountListResponse{},
	}
	var dayAccountData, currencyData []response.GetSubMerchantDayaccountResponse

	// 判断查询时间属于哪种类型
	rangeType, startTime, endTime, err := utils.DetectTimeRangeType(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	// 根据时间范围类型选择相应的查询逻辑
	switch rangeType {
	case utils.TimeRangeToday:
		// 第一种情况: 仅查询今天及之后的数据
		if currencyData, err = s.getCurrencyDataForSub(startTime, endTime, req.ParentId); err != nil {
			return nil, err
		}
	case utils.TimeRangeHistory:
		// 第二种情况: 仅查询今天之前的数据
		if dayAccountData, err = s.getDayaccountDataForSub(startTime, endTime, req.ParentId); err != nil {
			return nil, err
		}
	case utils.TimeRangeMixed:
		// 第三种情况: 包含今天之前和今天及之后
		// 查询 merchant_currency（今天及之后）和 merchant_dayaccount（今天之前）的数据
		if currencyData, err = s.getCurrencyDataForSub(startTime, endTime, req.ParentId); err != nil {
			return nil, err
		}
		if dayAccountData, err = s.getDayaccountDataForSub(startTime, endTime, req.ParentId); err != nil {
			return nil, err
		}
	default:
		return nil, gerror.Newf(`{#query.type.exception}`) // 查询类型异常
	}

	// 合并并聚合数据（按商户和成数分组）
	finalList := s.mergeAndGroupByRateForSub(dayAccountData, currencyData)
	if len(finalList) == 0 {
		return
	}
	res.Total = int64(len(finalList))
	res.List = finalList

	return res, nil
}

// 下级：（查询 merchant_dayaccount 数据）
func (s *MerchantDayaccount) getDayaccountDataForSub(startTime, endTime time.Time, parentId string) (data []response.GetSubMerchantDayaccountResponse, err error) {
	var dayDataRaw []model.MerchantDayaccount
	dayaccountStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.MerchantDayaccountColumns.AccountDate)
	err = global.DB.Where(dayaccountStr, startTime, endTime).Where(map[string]interface{}{
		model.MerchantDayaccountColumns.ParentId: parentId,
	}).Find(&dayDataRaw).Error

	for _, v := range dayDataRaw {
		data = append(data, response.GetSubMerchantDayaccountResponse{
			MerchantId:       v.MerchantId,                                                                                                                              // 对应商户表的username
			ParentId:         parentId,                                                                                                                                  // 对应商户表的username
			Currency:         v.Currency,                                                                                                                                // 币种
			Digits:           v.Digits,                                                                                                                                  // 币种的最小单位倍数
			LastBet:          gconv.String(v.Bet),                                                                                                                       // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.BetEffective),                                                                                                              // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.Water),                                                                                                                     // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.WinLost),                                                                                                                   // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.Tip),                                                                                                                       // 累计打赏金额，存储为最小单位
			StartTime:        v.AccountDate.Format(time.RFC3339),                                                                                                        // 开始时间
			EndTime:          time.Date(v.AccountDate.Year(), v.AccountDate.Month(), v.AccountDate.Day(), 23, 59, 59, 0, v.AccountDate.Location()).Format(time.RFC3339), // 结束时间
			HasChildren:      0,                                                                                                                                         // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                                                                                                                          // 成数
			RollingRate:      v.RollingRate,                                                                                                                             // 洗码比
			AccountDate:      v.AccountDate.Format(time.RFC3339),                                                                                                        // 结算归属日期
		})
	}
	return
}

// 下级：（查询 merchant_currency 数据）
func (s *MerchantDayaccount) getCurrencyDataForSub(startTime, endTime time.Time, parentId string) (data []response.GetSubMerchantDayaccountResponse, err error) {
	// 实际上这个函数的 startTime/endTime 不起作用，因为 merchant_currency 是当前实时数据，仅用 merchantIds 即可
	// 获取当天的开始时间和结束时间
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 今天的 00:00:00
	todayEndTime := todayStartTime.Add(24*time.Hour - time.Second)                              // 今天的 23:59:59

	// 查询 merchant 表找到所有子商户
	var merchantIds []string
	global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.ParentId: parentId,
	}).Pluck(model.MerchantColumns.Username, &merchantIds)
	if len(merchantIds) == 0 {
		return
	}

	var currencyDataRaw []model.MerchantCurrency
	// 查询 merchant_currency 中至少有一个字段不为0
	err = global.DB.Model(&model.MerchantCurrency{}).Where(s.getWalletWhere()).Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: merchantIds,
	}).Find(&currencyDataRaw).Error

	for _, v := range currencyDataRaw {
		data = append(data, response.GetSubMerchantDayaccountResponse{
			MerchantId:       v.Username,                          // 对应商户表的username
			ParentId:         parentId,                            // 对应商户表的username
			Currency:         v.Currency,                          // 币种
			Digits:           v.Digits,                            // 币种的最小单位倍数
			LastBet:          gconv.String(v.LastBet),             // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(v.LastBetEffective),    // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(v.LastWater),           // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(v.LastWinLost),         // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(v.LastTip),             // 累计打赏金额，存储为最小单位
			StartTime:        todayStartTime.Format(time.RFC3339), // 开始时间
			EndTime:          todayEndTime.Format(time.RFC3339),   // 结束时间
			HasChildren:      0,                                   // 是否存在子商户数据:0否,1是
			CommissionRate:   v.CommissionRate,                    // 成数
			RollingRate:      v.RollingRate,                       // 洗码比
			AccountDate:      time.Now().Format(time.DateOnly),    // 补上这个字段
		})
	}
	return
}

// 合并并聚合数据（按商户和成数分组）
func (s *MerchantDayaccount) mergeAndGroupByRateForSub(dayData, currencyData []response.GetSubMerchantDayaccountResponse) []response.GetSubMerchantDayaccountResponse {
	// 1. 对 dayData 按商户分组并按日期升序排列
	dayGroup := make(map[string][]response.GetSubMerchantDayaccountResponse)
	for _, d := range dayData {
		dayGroup[d.MerchantId] = append(dayGroup[d.MerchantId], d)
	}
	for merchantId := range dayGroup {
		sort.Slice(dayGroup[merchantId], func(i, j int) bool {
			return dayGroup[merchantId][i].AccountDate < dayGroup[merchantId][j].AccountDate
		})
	}

	// 2. 对每个商户的 dayData 进行“按连续日期 + 相同成数”聚合
	groupedMap := make(map[string][]response.GetSubMerchantDayaccountResponse)
	for merchantId, list := range dayGroup {

		var grouped []response.GetSubMerchantDayaccountResponse // 已分组好的结果
		// 同个商户底下的所有dayData数据
		for _, item := range list {
			n := len(grouped)
			if n == 0 {
				grouped = append(grouped, item)
				continue
			}
			last := &grouped[n-1] // 上一条已分组的数据
			// 如果成数一致，合并数据
			if last.CommissionRate == item.CommissionRate {
				// 合并逻辑
				last.LastBet = utils.SumStringInt(last.LastBet, item.LastBet)
				last.LastBetEffective = utils.SumStringInt(last.LastBetEffective, item.LastBetEffective)
				last.LastWater = utils.SumStringInt(last.LastWater, item.LastWater)
				last.LastWinLost = utils.SumStringInt(last.LastWinLost, item.LastWinLost)
				last.LastTip = utils.SumStringInt(last.LastTip, item.LastTip)
				//last.AccountDate = item.AccountDate // 更新最新日期
				// 合并开始时间和结束时间
				if last.StartTime > item.StartTime {
					last.StartTime = item.StartTime // 最早开始时间
				}
				if last.EndTime < item.EndTime {
					last.EndTime = item.EndTime // 最晚结束时间
				}
			} else {
				grouped = append(grouped, item)
			}
		}
		groupedMap[merchantId] = grouped
	}
	// 3. 合并 currencyData：如果商户在 dayData 中有最近记录，成数一致就合并，否则单独添加
	for _, cur := range currencyData {
		dayList := groupedMap[cur.MerchantId]
		if len(dayList) == 0 {
			// 没有 dayData，直接新增
			groupedMap[cur.MerchantId] = append(groupedMap[cur.MerchantId], cur)
			continue
		}
		last := &dayList[len(dayList)-1] // 最近的一条
		if last.CommissionRate == cur.CommissionRate {
			last.LastBet = utils.SumStringInt(last.LastBet, cur.LastBet)
			last.LastBetEffective = utils.SumStringInt(last.LastBetEffective, cur.LastBetEffective)
			last.LastWater = utils.SumStringInt(last.LastWater, cur.LastWater)
			last.LastWinLost = utils.SumStringInt(last.LastWinLost, cur.LastWinLost)
			last.LastTip = utils.SumStringInt(last.LastTip, cur.LastTip)
			last.AccountDate = cur.AccountDate // 更新最新日期
			// 合并开始时间和结束时间
			if last.StartTime > cur.StartTime {
				last.StartTime = cur.StartTime // 最早开始时间
			}
			if last.EndTime < cur.EndTime {
				last.EndTime = cur.EndTime // 最晚结束时间
			}
		} else {
			groupedMap[cur.MerchantId] = append(groupedMap[cur.MerchantId], cur)
		}
	}
	// 4. 汇总所有结果
	var result []response.GetSubMerchantDayaccountResponse
	for _, list := range groupedMap {
		result = append(result, list...)
	}

	// 排序：因为mysql的in查询是无序的会导致结果也是无序，先按商户ID排序，再按日期排序
	sort.Slice(result, func(i, j int) bool {
		if result[i].MerchantId == result[j].MerchantId {
			return result[i].AccountDate < result[j].AccountDate // 时间倒序
		}
		return result[i].MerchantId < result[j].MerchantId // 倒序
	})
	return result
}
