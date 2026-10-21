package service

import (
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
	"strconv"
	"strings"
)

type Player struct{}

// List 玩家列表
func (s *Player) List(req *request.PlayerListRequest, userInfo *request.CustomClaims) (*response.ListData, error) {
	var total int64
	var list []response.PlayerListResponse
	var res response.ListData
	var err error

	db := global.DB.Model(&model.Player{})

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的会员
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: userInfo.Username,
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

	// 商户账号筛选
	if req.MerchantId != "" {
		// 模糊搜索
		query := fmt.Sprintf("%s LIKE ?", model.PlayerColumns.MerchantId)
		db.Where(query, "%"+req.MerchantId+"%")

		// 完整搜索
		//db.Where(map[string]interface{}{
		//	model.PlayerColumns.MerchantId: req.MerchantId,
		//})
	}

	// 商户ID筛选
	if req.Username != "" {
		// 模糊搜索
		query := fmt.Sprintf("%s LIKE ?", model.PlayerColumns.Username)
		db.Where(query, "%"+req.Username+"%")

		// 完整搜索
		//db.Where(map[string]interface{}{
		//	model.PlayerColumns.Username: req.Username,
		//})
	}

	// 商户ID筛选
	if req.Nickname != "" {
		query := fmt.Sprintf("%s LIKE ?", model.PlayerColumns.Nickname)
		db.Where(query, "%"+req.Nickname+"%")
	}

	// 币种code筛选
	if req.Currency != "" {
		db.Where(map[string]interface{}{
			model.PlayerColumns.Currency: req.Currency,
		})
	}

	// 是否冻结:取值说明在模型里
	if req.Status != nil {
		db.Where(map[string]interface{}{
			model.PlayerColumns.Status: req.Status,
		})
	}

	// 是否暂停投注:取值说明在模型里
	if req.BetPaused != nil {
		db.Where(map[string]interface{}{
			model.PlayerColumns.BetPaused: req.BetPaused,
		})
	}

	// 是否在线:取值说明在模型里
	if req.IsOnline != nil {
		db.Where(map[string]interface{}{
			model.PlayerColumns.IsOnline: req.IsOnline,
		})
	}

	db.Count(&total)

	if total <= 0 {
		res.List = []response.PlayerListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var players []model.Player
	if err = db.Preload("Merchant").Preload("PlayerWallet").Find(&players).Error; err != nil {
		return &res, err
	}

	for _, player := range players {
		var loginTimeStr string
		if player.LoginTime != nil {
			loginTimeStr = player.LoginTime.Format("2006-01-02 15:04:05")
		}

		var loginTimeFirstStr string
		if player.LoginTimeFirst != nil {
			loginTimeFirstStr = player.LoginTimeFirst.Format("2006-01-02 15:04:05")
		}

		tmp := response.PlayerListResponse{
			Id:                 player.ID,                    // 玩家ID
			MerchantName:       player.Merchant.MerchantName, // 商户商户名称
			MerchantId:         player.MerchantId,            // 商户账号
			Currency:           player.Currency,              // 币种
			MerchantCurrencyId: player.MerchantCurrencyId,    // 商户币种唯一标识
			Username:           player.Username,              // 初始账号
			Nickname:           player.Nickname,              // 昵称
			ParentId:           player.ParentId,              // 上级玩家的的username

			LoginTimeFirst: loginTimeFirstStr,   // 首次次登录时间
			LoginIpFirst:   player.LoginIpFirst, // 首次次登录IP
			LoginTime:      loginTimeStr,        // 上次登录时间
			LoginIp:        player.LoginIp,      // 上次登录IP

			Status:     player.Status,           // 是否冻结:取值说明在模型里
			BetPaused:  player.BetPaused,        // 是否暂停投注:取值说明在模型里
			IsOnline:   player.IsOnline,         // 是否在线:取值说明在模型里
			PlayerType: uint(player.PlayerType), // 玩家类型:取值说明在模型里

			Digits:            player.PlayerWallet.Digits,            // 币种的最小单位倍数
			CreditSum:         player.PlayerWallet.CreditSum,         // 累计充值额度
			Balance:           player.PlayerWallet.Balance,           // 账户余额
			Tip:               player.PlayerWallet.Tip,               // 打赏金额
			TotalBet:          player.PlayerWallet.TotalBet,          // 总投注金额
			TotalBetEffective: player.PlayerWallet.TotalBetEffective, // 总有效投注金额
			TotalWater:        player.PlayerWallet.TotalWater,        // 总洗码金额
			TotalWinLost:      player.PlayerWallet.TotalWinLost,      // 总输赢金额

			LastBetEffective: player.PlayerWallet.LastBetEffective, // 上次有效投注金额
		}
		list = append(list, tmp)
	}

	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.PlayerListResponse{}
	}

	return &res, err
}

// Add 创建玩家 - 完成
func (s *Player) Add(req *request.SavePlayerRequest, merchant model.Merchant) (err error) {
	// 查询商户钱包
	var merchantCurrency model.MerchantCurrency
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: req.MerchantId,
	}).First(&merchantCurrency).Error; err != nil {
		return gerror.Newf(`{#merchant.currency.id.no.exist}`) // 商户币种不存在
	}

	// 开启事务
	tx := global.DB.Begin()

	// 玩家主表
	player := model.Player{
		RootKey:            merchant.RootKey,                                 // 顶级Key
		MerchantId:         merchantCurrency.Username,                        // 商户账号
		Currency:           merchantCurrency.Currency,                        // 币种
		MerchantCurrencyId: merchantCurrency.MerchantCurrencyId,              // 商户币种唯一标识
		ApiUser:            strings.TrimSpace(req.Username),                  // 存储api会员原始账号
		Username:           strings.ToLower(strings.TrimSpace(req.Username)), // 账号
		Nickname:           strings.TrimSpace(req.Nickname),                  // 昵称
		ParentId:           strings.ToLower(strings.TrimSpace(req.ParentId)), // 上级玩家账号

		Status:    req.Status,                   // 是否允许登录:取值说明在模型里
		BetPaused: req.BetPaused,                // 是否暂停投注:取值说明在模型里
		MaxWin:    req.MaxWin,                   // 最高赢额,0表示不限制
		Level:     model.PlayerLevel(req.Level), // 等级

		Email:      strings.TrimSpace(req.Email),                      // 邮箱
		Phone:      strings.TrimSpace(req.Phone),                      // 手机号
		Password:   utils.BcryptHash(strings.TrimSpace(req.Password)), // 密码
		Chips:      slice.Join(req.ChipIds, ","),                      // 筹码
		TemplateId: req.TemplateId,                                    // 商户限红模板ID
	}

	// 玩家类型:取值说明在模型里
	switch req.PlayerType {
	case 1:
		// 现金代理
		player.PlayerType = model.CashAgent
	case 2:
		// 佣金代理
		player.PlayerType = model.CommissionAgent
	default:
		return gerror.Newf(`{#play.player_type.error}`) // 玩家类型错误
	}

	// 玩家模式类型:取值说明在模型里
	switch req.MoneyType {
	case 1:
		// 真钱
		player.MoneyType = model.RealMoney
	case 2:
		// 试玩
		player.MoneyType = model.Trial
	default:
		return gerror.Newf(`{#play.money_type.error}`) // 玩家模式错误
	}

	if err = tx.Create(&player).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 玩家钱包表
	{
		playerWallet := model.PlayerWallet{
			MerchantId:         merchantCurrency.Username,           // 商户账号
			Currency:           merchantCurrency.Currency,           // 币种
			MerchantCurrencyId: merchantCurrency.MerchantCurrencyId, // 商户币种唯一标识
			Username:           player.Username,                     // 玩家key
			Digits:             merchantCurrency.Digits,             // 币种的最小单位倍数
		}
		if err = tx.Create(&playerWallet).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Edit 编辑玩家
func (s *Player) Edit(req *request.SavePlayerRequest, player model.Player) (err error) {
	player.Nickname = strings.TrimSpace(req.Nickname)                  // 昵称
	player.ParentId = strings.ToLower(strings.TrimSpace(req.ParentId)) // 上级玩家账号

	player.Status = req.Status                              // 是否允许登录:取值说明在模型里
	player.BetPaused = req.BetPaused                        // 是否暂停投注:取值说明在模型里
	player.MaxWin = req.MaxWin                              // 最高赢额,0表示不限制
	player.MoneyType = model.PlayerMoneyType(req.MoneyType) // 玩家模式类型:取值说明在模型里
	player.Level = model.PlayerLevel(req.Level)             // 等级

	player.Email = strings.TrimSpace(req.Email) // 邮箱
	player.Phone = strings.TrimSpace(req.Phone) // 手机号
	// 如果密码不为空，重新设置密码
	if len(strings.TrimSpace(req.Password)) != 0 {
		player.Password = utils.BcryptHash(strings.TrimSpace(req.Password))
	}
	player.Chips = slice.Join(req.ChipIds, ",") // 筹码
	player.TemplateId = req.TemplateId          // 商户限红模板ID

	// 玩家类型:取值说明在模型里
	switch req.PlayerType {
	case 1:
		// 现金代理
		player.PlayerType = model.CashAgent
	case 2:
		// 佣金代理
		player.PlayerType = model.CommissionAgent
	default:
		return gerror.Newf(`{#play.player_type.error}`) // 玩家类型错误
	}

	// 玩家模式类型:取值说明在模型里
	switch req.MoneyType {
	case 1:
		// 真钱
		player.MoneyType = model.RealMoney
	case 2:
		// 试玩
		player.MoneyType = model.Trial
	default:
		return gerror.Newf(`{#play.money_type.error}`) // 玩家模式错误
	}

	return global.DB.Save(&player).Error
}

// Detail 玩家详情
func (s *Player) Detail(id uint, userInfo *request.CustomClaims) (*response.PlayerDetailResponse, error) {
	var player model.Player

	db := global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Id: id,
	}).Preload("Merchant").Preload("PlayerWallet")

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的会员
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: userInfo.Username,
			})
		}
	}

	err := db.First(&player).Error

	var loginTimeStr string
	if player.LoginTime != nil {
		loginTimeStr = player.LoginTime.Format("2006-01-02 15:04:05")
	}

	var loginTimeFirstStr string
	if player.LoginTimeFirst != nil {
		loginTimeFirstStr = player.LoginTimeFirst.Format("2006-01-02 15:04:05")
	}

	res := response.PlayerDetailResponse{
		Id:                 player.ID,                    // 玩家ID
		MerchantName:       player.Merchant.MerchantName, // 商户商户名称
		MerchantId:         player.MerchantId,            // 商户账号
		Currency:           player.Currency,              // 币种
		MerchantCurrencyId: player.MerchantCurrencyId,    // 商户币种唯一标识
		ApiUser:            player.ApiUser,               // 存储api会员原始账号
		Username:           player.Username,              // 初始账号
		Nickname:           player.Nickname,              // 昵称
		ParentId:           player.ParentId,              // 上级玩家的的username

		LoginTimeFirst: loginTimeFirstStr,   // 首次次登录时间
		LoginIpFirst:   player.LoginIpFirst, // 首次次登录IP
		LoginTime:      loginTimeStr,        // 上次登录时间
		LoginIp:        player.LoginIp,      // 上次登录IP

		Status:     player.Status,           // 是否冻结:取值说明在模型里
		BetPaused:  player.BetPaused,        // 是否暂停投注:取值说明在模型里
		CreateType: uint(player.CreateType), // 玩家来源:取值说明在模型里
		MaxWin:     player.MaxWin,           // 最高赢额:取值说明在模型里
		DeviceType: uint(player.DeviceType), // 设备类型:取值说明在模型里
		IsOnline:   player.IsOnline,         // 是否在线:取值说明在模型里
		PlayerType: uint(player.PlayerType), // 玩家类型:取值说明在模型里
		MoneyType:  uint(player.MoneyType),  // 玩家模式类型:取值说明在模型里
		Level:      uint(player.Level),      // 等级

		Email:      player.Email,      // 邮箱
		Phone:      player.Phone,      // 手机号
		ChipIds:    []uint{},          // 筹码
		TemplateId: player.TemplateId, // 商户限红模板ID

		Digits:            player.PlayerWallet.Digits,            // 币种的最小单位倍数
		CreditSum:         player.PlayerWallet.CreditSum,         // 累计充值额度
		Balance:           player.PlayerWallet.Balance,           // 账户余额
		Tip:               player.PlayerWallet.Tip,               // 总打赏金额
		TotalBet:          player.PlayerWallet.TotalBet,          // 总投注金额
		TotalBetEffective: player.PlayerWallet.TotalBetEffective, // 总有效投注金额
		TotalWater:        player.PlayerWallet.TotalWater,        // 总洗码金额
		TotalWinLost:      player.PlayerWallet.TotalWinLost,      // 总输赢金额

		LastBet:          player.PlayerWallet.LastBet,          // 上次投注金额
		LastBetEffective: player.PlayerWallet.LastBetEffective, // 上次有效投注金额
		LastWater:        player.PlayerWallet.LastWater,        // 上次洗码金额
		LastWinLost:      player.PlayerWallet.LastWinLost,      // 上次输赢金额
		LastTip:          player.PlayerWallet.LastTip,          // 上次打赏金额
	}

	// 筹码
	chipStr := strings.Split(player.Chips, ",")
	for _, p := range chipStr {
		var chip int
		chip, err = strconv.Atoi(p) // 转换为 int
		if err != nil {
			return nil, err
		}
		res.ChipIds = append(res.ChipIds, uint(chip))
	}

	return &res, err
}

// Del 删除玩家
func (s *Player) Del(id uint) (username string, err error) {
	var playerModel model.Player
	if errors.Is(global.DB.First(&playerModel, id).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}
	return playerModel.Username, global.DB.Delete(&playerModel).Error
}

// ResetPassword 重置玩家密码
func (s *Player) ResetPassword(id uint, merchantKey string) (*response.ResetPlayerPasswordResponse, error) {
	var playerModel model.Player

	db := global.DB.Where(map[string]interface{}{
		model.PlayerColumns.Id: id,
	})
	if len(merchantKey) > 0 {
		db.Where(map[string]interface{}{
			model.PlayerColumns.MerchantId: merchantKey,
		})
	}
	if errors.Is(db.First(&playerModel).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	resetPassword := global.CONFIG.Server.ResetPassword
	if len(resetPassword) == 0 {
		resetPassword = "123456"
	}

	playerModel.Password = utils.BcryptHash(strings.TrimSpace(resetPassword))

	if err := global.DB.Save(&playerModel).Error; err != nil {
		return nil, err
	}
	res := response.ResetPlayerPasswordResponse{
		Id:       playerModel.ID,       // 玩家ID
		Username: playerModel.Username, // 账号
		Password: resetPassword,        // 密码
	}
	return &res, nil
}
