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
	"strings"
	"time"
)

type PlayerDayaccount struct{}

// List 玩家输赢表列表
func (s *PlayerDayaccount) List(req *request.PlayerDayaccountListRequest, userInfo *request.CustomClaims) (*response.ListData, error) {
	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的会员的投注明细
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)

			global.DB.Model(&model.Player{}).Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: usernames,
			}).Pluck("username", &req.PlayerIds)
		} else {
			// 子商户：只能看到自己的
			global.DB.Model(&model.Player{}).Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: userInfo.Username,
			}).Pluck("username", &req.PlayerIds)
		}
		// 如果商户没有玩家直接返回
		if len(req.PlayerIds) == 0 {
			return &response.ListData{
				Total: 0,
				List:  []response.PlayerDayaccountListResponse{},
			}, nil
		}
	}

	var err error
	var startTime, endTime time.Time

	// 示例：2025-03-04 15:42:20
	if startTime, err = time.ParseInLocation(time.DateTime, req.StartTime, time.Local); err != nil {
		return nil, err
	}
	// 示例：2025-04-04 13:42:20
	if endTime, err = time.ParseInLocation(time.DateTime, req.EndTime, time.Local); err != nil {
		return nil, err
	}
	if endTime.Sub(startTime) > 31*24*time.Hour {
		return nil, gerror.Newf(`{#time.cannot.exceed.one.month}`) // 时间间隔不能超过一个月
	}

	// Time: 当前时间
	now := time.Now()
	// Time: 今天的零点 2025-04-09 00:00:00.000
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 只查询今天
	onlyToday := !startTime.Before(todayStart) && endTime.After(todayStart)
	// 只查询之前
	onlyBefore := endTime.Before(todayStart)
	// 查询今天和之前
	mixRange := startTime.Before(todayStart) && endTime.After(todayStart)

	switch {
	case onlyToday:
		return s.queryToday(req, startTime, endTime)
	case onlyBefore:
		return s.queryBefore(req, startTime, endTime)
	case mixRange:
		return s.queryRange(req, startTime, endTime)
	default:
		return nil, gerror.Newf(`{#query.type.exception}`) // 查询类型异常
	}
}

// 查询今天数据
func (s *PlayerDayaccount) queryToday(req *request.PlayerDayaccountListRequest, startTime, endTime time.Time) (res *response.ListData, err error) {
	utils.DD("查询今天")
	res = &response.ListData{
		Total: 0,
		List:  []response.PlayerDayaccountListResponse{},
	}

	// 直接分页查询 or_player_wallet 表
	db := global.DB.Model(&model.PlayerWallet{}).Where(s.getWalletWhere())

	// 玩家账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.PlayerWalletColumns.Username: req.Username,
		})
	}

	// 商户账号
	if req.MerchantId != "" {
		db.Where(map[string]interface{}{
			model.PlayerWalletColumns.MerchantId: req.MerchantId,
		})
	}

	// 商户端查询限制
	if len(req.PlayerIds) != 0 {
		db.Where(map[string]interface{}{
			model.PlayerWalletColumns.Username: req.PlayerIds,
		})
	}

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
	var wallets []model.PlayerWallet
	if err = db.Find(&wallets).Error; err != nil {
		return
	}
	// 列表赋值
	var list []response.PlayerDayaccountListResponse
	for _, wallet := range wallets {
		tmp := response.PlayerDayaccountListResponse{
			Username:         wallet.Username,                       // 玩家账号,每个玩家的唯一值
			MerchantId:       wallet.MerchantId,                     // 对应商户表的username
			Digits:           wallet.Digits,                         // 币种的最小单位倍数
			LastBet:          gconv.String(wallet.LastBet),          // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(wallet.LastBetEffective), // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(wallet.LastWater),        // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(wallet.LastWinLost),      // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(wallet.LastTip),          // 上次打赏金额，存储为最小单位
			StartTime:        startTime.Format(time.RFC3339),        // 开始时间
			EndTime:          endTime.Format(time.RFC3339),          // 结束时间
			Currency:         wallet.Currency,                       // 币种
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// 查询之前不包含今天数据
func (s *PlayerDayaccount) queryBefore(req *request.PlayerDayaccountListRequest, startTime, endTime time.Time) (res *response.ListData, err error) {
	utils.DD("查询之前")
	res = &response.ListData{
		Total: 0,
		List:  []response.PlayerDayaccountListResponse{},
	}

	// 直接分页查询 or_player_dayaccount 表
	whereStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.PlayerDayaccountColumns.AccountDate)
	db := global.DB.Model(&model.PlayerDayaccount{}).Where(whereStr, startTime, endTime).Group(model.PlayerDayaccountColumns.Username)

	// 玩家账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.Username: req.Username,
		})
	}

	// 商户账号
	if req.MerchantId != "" {
		db.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.ParentId: req.MerchantId,
		})
	}

	// 商户端查询限制
	if len(req.PlayerIds) != 0 {
		db.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.Username: req.PlayerIds,
		})
	}

	db.Count(&res.Total)

	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	// 分页
	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 统计sum语句 sum()....
	selectStr := s.getDayAccountSelect()
	var accounts []model.PlayerDayaccount
	if err = db.Select(selectStr).Find(&accounts).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.PlayerDayaccountListResponse
	for _, account := range accounts {
		tmp := response.PlayerDayaccountListResponse{
			Username:         account.Username,                       // 玩家账号,每个玩家的唯一值
			MerchantId:       account.MerchantId,                     // 对应商户表的username
			Digits:           account.Digits,                         // 币种的最小单位倍数
			LastBet:          gconv.String(account.LastBet),          // 上次投注金额，存储为最小单位
			LastBetEffective: gconv.String(account.LastBetEffective), // 上次有效投注金额，存储为最小单位
			LastWater:        gconv.String(account.LastWater),        // 上次洗码金额，存储为最小单位
			LastWinLost:      gconv.String(account.LastWinLost),      // 上次输赢金额，存储为最小单位
			LastTip:          gconv.String(account.LastTip),          // 上次打赏金额，存储为最小单位
			StartTime:        startTime.Format(time.RFC3339),         // 开始时间
			EndTime:          endTime.Format(time.RFC3339),           // 结束时间
			Currency:         account.Currency,                       // 币种
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// 查询之前包含今天数据
func (s *PlayerDayaccount) queryRange(req *request.PlayerDayaccountListRequest, startTime, endTime time.Time) (res *response.ListData, err error) {
	utils.DD("查询之前和今天")
	res = &response.ListData{
		Total: 0,
		List:  []response.PlayerDayaccountListResponse{},
	}

	// 表 or_player_wallet 满足条件的username
	var walletUsernames []string
	// 表 or_player_dayaccount 满足条件的username
	var accountUsernames []string

	dayaccountStr := fmt.Sprintf("%s BETWEEN ? AND ?", model.PlayerDayaccountColumns.AccountDate)

	// 钱包表有数据的查询条件
	dbWallet := global.DB.Model(&model.PlayerWallet{}).Where(s.getWalletWhere())
	// 历史表有数据的查询条件
	dbAccount := global.DB.Model(&model.PlayerDayaccount{}).Where(dayaccountStr, startTime, endTime)

	// 玩家账号
	if req.Username != "" {
		dbWallet.Where(map[string]interface{}{
			model.PlayerWalletColumns.Username: req.Username,
		})
		dbAccount.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.Username: req.Username,
		})
	}
	// 商户账号
	if req.MerchantId != "" {
		dbWallet.Where(map[string]interface{}{
			model.PlayerWalletColumns.MerchantId: req.MerchantId,
		})
		dbAccount.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.ParentId: req.MerchantId,
		})
	}

	// 商户端查询限制
	if len(req.PlayerIds) != 0 {
		dbWallet.Where(map[string]interface{}{
			model.PlayerWalletColumns.Username: req.PlayerIds,
		})
		dbAccount.Where(map[string]interface{}{
			model.PlayerDayaccountColumns.Username: req.PlayerIds,
		})
	}

	dbWallet.Pluck(model.PlayerWalletColumns.Username, &walletUsernames)
	dbAccount.Pluck(model.PlayerDayaccountColumns.Username, &accountUsernames)

	// 所有满足条件的已经去重的username
	usernames := slice.Union(walletUsernames, accountUsernames)

	// 查询所有满足条件的玩家钱包,肯定会找的到，但不一定有值还需要查询玩家输赢表
	db := global.DB.Model(&model.PlayerWallet{}).Where(map[string]interface{}{
		model.PlayerWalletColumns.Username: usernames,
	})
	db.Count(&res.Total)
	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}
	// 分页
	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}
	var wallets []model.PlayerWallet
	if err = db.Find(&wallets).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.PlayerDayaccountListResponse
	for _, wallet := range wallets {
		// 查询统计sum语句 sum()....
		selectStr := s.getDayAccountSelect()
		var account model.PlayerDayaccount
		global.DB.Model(&model.PlayerDayaccount{}).Group(model.PlayerDayaccountColumns.Username).Select(selectStr).Where(dayaccountStr, startTime, endTime).Where(map[string]interface{}{
			model.PlayerDayaccountColumns.Username: wallet.Username,
		}).Find(&account)

		// 累加金额（最小单位）
		lastBet := wallet.LastBet + account.LastBet                            // 累计投注金额，存储为最小单位
		lastBetEffective := wallet.LastBetEffective + account.LastBetEffective // 累计有效投注金额，存储为最小单位
		lastWater := wallet.LastWater + account.LastWater                      // 累计输赢金额，存储为最小单位
		lastWinLost := wallet.LastWinLost + account.LastWinLost                // 累计输赢金额，存储为最小单位
		lastTip := wallet.LastTip + account.LastTip                            // 累计打赏金额，存储为最小单位

		// 初始化临时结构体
		tmp := response.PlayerDayaccountListResponse{
			Username:         wallet.Username,                // 玩家账号,每个玩家的唯一值
			MerchantId:       wallet.MerchantId,              // 对应商户表的username
			Digits:           wallet.Digits,                  // 币种的最小单位倍数
			LastBet:          gconv.String(lastBet),          // 累计投注金额，存储为最小单位
			LastBetEffective: gconv.String(lastBetEffective), // 累计有效投注金额，存储为最小单位
			LastWater:        gconv.String(lastWater),        // 累计洗码金额，存储为最小单位
			LastWinLost:      gconv.String(lastWinLost),      // 累计输赢金额，存储为最小单位
			LastTip:          gconv.String(lastTip),          // 累计打赏金额，存储为最小单位
			StartTime:        startTime.Format(time.RFC3339), // 开始时间
			EndTime:          endTime.Format(time.RFC3339),   // 结束时间
			Currency:         wallet.Currency,                // 币种
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// 获取输赢历史表的统计select语句-ANY_VALUE(username) AS username, ANY_VALUE(merchant_id) AS merchant_id, ANY_VALUE(digits) AS digits, SUM(last_bet) AS last_bet, SUM(last_bet_effective) AS last_bet_effective, SUM(last_water) AS last_water, SUM(last_win_lost) AS last_win_lost, SUM(last_tip) AS last_tip, ANY_VALUE(currency) AS currency
func (s *PlayerDayaccount) getDayAccountSelect() (selectStr string) {
	type columnAgg struct {
		Column string
		Agg    string
	}

	columns := []columnAgg{
		{model.PlayerDayaccountColumns.Username, "ANY_VALUE"},
		{model.PlayerDayaccountColumns.MerchantId, "ANY_VALUE"},
		{model.PlayerDayaccountColumns.Digits, "ANY_VALUE"},
		{model.PlayerDayaccountColumns.LastBet, "SUM"},
		{model.PlayerDayaccountColumns.LastBetEffective, "SUM"},
		{model.PlayerDayaccountColumns.LastWater, "SUM"},
		{model.PlayerDayaccountColumns.LastWinLost, "SUM"},
		{model.PlayerDayaccountColumns.LastTip, "SUM"},
		{model.PlayerDayaccountColumns.Currency, "ANY_VALUE"},
	}

	var selects []string
	for _, c := range columns {
		selects = append(selects, fmt.Sprintf("%s(%s) AS %s", c.Agg, c.Column, c.Column))
	}
	selectStr = strings.Join(selects, ", ")
	return
}

// 返回钱包表的过滤条件
func (s *PlayerDayaccount) getWalletWhere() (walletWhere string) {
	return fmt.Sprintf("%s != 0 OR %s != 0 OR %s != 0 OR %s != 0 OR %s != 0", model.PlayerWalletColumns.LastBet, model.PlayerWalletColumns.LastBetEffective, model.PlayerWalletColumns.LastWater, model.PlayerWalletColumns.LastWinLost, model.PlayerWalletColumns.LastTip)
}
