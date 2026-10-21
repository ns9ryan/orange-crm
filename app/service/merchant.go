package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strconv"
	"strings"
	"time"
)

type Merchant struct{}

// List 商户列表
func (s *Merchant) List(req *request.MerchantListRequest) (*response.ListData, error) {
	var total int64
	var list []response.MerchantListResponse
	var res response.ListData
	var err error

	db := global.DB.Model(&model.Merchant{}).Preload("MerchantCurrency")

	if req.ParentId == "" {
		db.Where(map[string]interface{}{
			model.MerchantColumns.ParentId: global.MerchantRootName,
		})
	} else {
		db.Where(map[string]interface{}{
			model.MerchantColumns.ParentId: req.ParentId,
		})
	}

	order := "id asc"

	// 排序
	if req.SortField != nil && req.SortType != nil && slice.Contain(request.SortTypeEnum, strings.ToUpper(*req.SortType)) {
		// 根据ID排序
		if *req.SortField == model.MerchantColumns.Id {
			order = fmt.Sprintf("%s %s", model.MerchantColumns.Id, *req.SortType)
		}
	}
	db.Order(order)

	// 账号搜索
	if req.Username != "" {
		// 模糊搜索
		query := fmt.Sprintf("%s LIKE ?", model.MerchantColumns.Username)
		db.Where(query, "%"+req.Username+"%")

		// 完整搜索
		//db.Where(map[string]interface{}{
		//	model.MerchantColumns.Username: req.Username,
		//})
	}

	// 钱包类型搜索
	if req.WalletType != nil {
		db.Where(map[string]interface{}{
			model.MerchantColumns.WalletType: req.WalletType,
		})
	}

	db.Count(&total)

	if total <= 0 {
		res.List = []response.MerchantListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var merchants []model.Merchant
	if err = db.Find(&merchants).Error; err != nil {
		return &res, err
	}
	for _, merchant := range merchants {
		var loginTimeStr string
		if merchant.LoginTime != nil {
			loginTimeStr = merchant.LoginTime.Format("2006-01-02 15:04:05")
		}
		tmp := response.MerchantListResponse{
			Id:             merchant.ID,                        // 商户ID
			MerchantName:   merchant.MerchantName,              // 商户名称
			MerchantCode:   merchant.MerchantCode,              // 商户号
			ParentId:       merchant.ParentId,                  // 上级商户的username
			Username:       merchant.Username,                  // 账号
			Nickname:       merchant.Nickname,                  // 昵称
			Status:         merchant.Status,                    // 是否允许登录:取值说明在模型里
			CommissionRate: merchant.CommissionRate,            // 成数
			RollingRate:    merchant.RollingRate,               // 洗码比
			WalletType:     uint(merchant.WalletType),          // 钱包类型:取值说明在模型里
			LoginTime:      loginTimeStr,                       // 上次登录时间
			LoginIp:        merchant.LoginIp,                   // 登录IP
			Currency:       merchant.MerchantCurrency.Currency, // 商户币种
			MaxTip:         merchant.MerchantCurrency.MaxTip,   // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:         merchant.MerchantCurrency.MinTip,   // 荷官可接收的最小打赏金额，存储为最小单位
			Digits:         merchant.MerchantCurrency.Digits,   // 币种的最小单位倍数
			Children:       []response.MerchantListResponse{},  // 子账号信息
		}
		// 玩家数量
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.MerchantId: merchant.Username,
		}).Count(&tmp.PlayerCount)

		var subs []model.Merchant
		global.DB.Where(map[string]interface{}{
			model.MerchantColumns.ParentId: tmp.Username,
		}).Preload("MerchantCurrency").Find(&subs)
		for _, sub := range subs {
			var subLoginTimeStr string
			if sub.LoginTime != nil {
				subLoginTimeStr = sub.LoginTime.Format("2006-01-02 15:04:05")
			}
			subTmp1 := response.MerchantListResponse{
				Id:                sub.ID,                            // 商户ID
				MerchantName:      sub.MerchantName,                  // 商户名称
				MerchantCode:      sub.MerchantCode,                  // 商户号
				ParentId:          sub.ParentId,                      // 上级商户的username
				Username:          sub.Username,                      // 账号
				Nickname:          sub.Nickname,                      // 昵称
				Status:            sub.Status,                        // 是否允许登录:取值说明在模型里
				CommissionRate:    sub.CommissionRate,                // 成数
				MaxCommissionRate: merchant.CommissionRate,           // 上级成数
				RollingRate:       sub.RollingRate,                   // 洗码比
				MaxRollingRate:    merchant.RollingRate,              // 上级洗码比
				WalletType:        uint(sub.WalletType),              // 钱包类型:取值说明在模型里
				LoginTime:         subLoginTimeStr,                   // 上次登录时间
				LoginIp:           sub.LoginIp,                       // 登录IP
				Currency:          sub.MerchantCurrency.Currency,     // 商户币种
				Children:          []response.MerchantListResponse{}, // 子账号信息
			}

			// 玩家数量
			global.DB.Model(&model.Player{}).Where(map[string]interface{}{
				model.PlayerColumns.MerchantId: sub.Username,
			}).Count(&subTmp1.PlayerCount)

			tmp.PlayerCount += subTmp1.PlayerCount
			tmp.Children = append(tmp.Children, subTmp1)
		}
		list = append(list, tmp)
	}

	res.Total = total

	if list != nil {
		res.List = list
	} else {
		res.List = []response.MerchantListResponse{}
	}

	return &res, err
}

// Add 创建商户
func (s *Merchant) Add(c *gin.Context, req *request.SaveMerchantRequest) (err error) {
	if req.SubUsername == req.Username {
		return gerror.Newf(`{#merchant.add.username.eq.sub_username}`) // 主商户不能和子账号相同
	}
	// 1. 初始化总商户和子商户信息 =================================================================================

	// 总商户号
	merchantCode, err := s.makeCode()
	if err != nil {
		return err
	}
	// 子商户号
	subMerchantCode, err := s.makeCode()
	if err != nil {
		return err
	}

	var lastRecordId uint
	global.DB.Model(&model.Merchant{}).Order(fmt.Sprintf("%s desc", model.MerchantColumns.Id)).Limit(1).Pluck(model.MerchantColumns.Id, &lastRecordId)
	rootKey := lastRecordId + 1

	// 除了商户号、账号，其他都一样
	baseMerchant := model.Merchant{
		MerchantName: strings.TrimSpace(req.MerchantName), // 商户名称
		RootKey:      rootKey,                             // 顶级Key
		Nickname:     strings.TrimSpace(req.Nickname),     // 昵称
		Password:     s.encryptPassword(req.Password),     // 密码

		CommissionRate: req.CommissionRate, // 成数
		RollingRate:    req.RollingRate,    // 洗码比

		Status:     req.Status,                       // 是否允许登录:取值说明在模型里
		WalletType: model.WalletType(req.WalletType), // 钱包类型:取值说明在模型里

		Chips:         slice.Join(req.ChipIds, ","),         // 筹码
		ChipsTemplate: slice.Join(req.ChipTemplateIds, ","), // 筹码模板
		RoleId:        req.RoleId,                           // 角色ID

		AesKeySalt: s.makeAesKeySalt(),               // aesKey盐
		Md5Salt:    s.makeMd5Salt(),                  // md5盐
		WalletUrl:  strings.TrimSpace(req.WalletUrl), // 单钱包URL
		WhiteList:  strings.TrimSpace(req.WhiteList), // 白名单IP,逗号分隔
	}

	// 主商户
	merchant := baseMerchant
	merchant.MerchantCode = merchantCode                                 // 总商户号
	merchant.Username = strings.ToLower(strings.TrimSpace(req.Username)) // 账号
	merchant.ParentId = "root"                                           // 上级商户的username,总账号默认root
	merchant.OrTokenKey = s.makeOrTokenKey()                             // 生成token用的商户唯一key，永久有效
	merchant.OrToken, err = s.makeOrToken(merchantCode, []byte(merchant.OrTokenKey))
	if err != nil {
		return err
	}

	// 子商户
	subMerchant := baseMerchant
	subMerchant.MerchantCode = subMerchantCode                                 // 子商户号
	subMerchant.Username = strings.ToLower(strings.TrimSpace(req.SubUsername)) // 子账号
	subMerchant.ParentId = merchant.Username                                   // 上级商户的username,总账号的username
	subMerchant.OrTokenKey = s.makeOrTokenKey()                                // 生成token用的商户唯一key，永久有效
	subMerchant.OrToken, err = s.makeOrToken(subMerchantCode, []byte(subMerchant.OrTokenKey))
	if err != nil {
		return err
	}

	// 2. 验证币种是否存在 =================================================================================

	var currency model.Currency
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.CurrencyColumns.Code: req.Currency,
	}).First(&currency).Error, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#currency.code.no.exist}`)
	}

	// 3. 构建商户币种信息（总商户 + 子商户） =================================================================================

	// 打赏选项金额
	slice.Sort(req.TipConfig)
	if len(req.TipConfig) < 1 {
		return gerror.Newf(`{#merchant.add.tip.error}`) // 配置的打赏面值超出范围
	}
	first := req.TipConfig[0]
	last := req.TipConfig[len(req.TipConfig)-1]
	if first < req.MinTip || last > req.MaxTip {
		return gerror.Newf(`{#merchant.add.tip.error}`) // 配置的打赏面值超出范围
	}

	merchantCurrencies := []model.MerchantCurrency{
		{
			MerchantCurrencyId: s.makeMerchantCurrencyId(merchant.Username, currency.Code), // 唯一标识
			Username:           merchant.Username,                                          // 商户账号
			Currency:           currency.Code,                                              // 币种代码
			Digits:             currency.Digits,                                            // 币种的最小单位倍数
			Status:             1,                                                          // 是否启用:0否,1是
			BetPaused:          1,                                                          // 是否暂停投注（不开放）:0否,1是
			CommissionRate:     merchant.CommissionRate,                                    // 成数
			RollingRate:        merchant.RollingRate,                                       // 洗码比
			LimitPlayer:        global.CONFIG.Server.LimitPlayer,                           // 限制会员数量
			MaxTip:             req.MaxTip,                                                 // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:             req.MinTip,                                                 // 荷官可接收的最小打赏金额，存储为最小单位
			TipOptions:         slice.Join(req.TipConfig, ","),                             // 打赏选项金额，以逗号隔开，存储为最小单位
		},
		{
			MerchantCurrencyId: s.makeMerchantCurrencyId(subMerchant.Username, currency.Code), // 唯一标识
			Username:           subMerchant.Username,                                          // 商户账号
			Currency:           currency.Code,                                                 // 币种代码
			Digits:             currency.Digits,                                               // 币种的最小单位倍数
			Status:             1,                                                             // 是否启用:0否,1是
			BetPaused:          1,                                                             // 是否暂停投注（不开放）:0否,1是
			CommissionRate:     subMerchant.CommissionRate,                                    // 成数
			RollingRate:        subMerchant.RollingRate,                                       // 洗码比
			LimitPlayer:        global.CONFIG.Server.LimitPlayer,                              // 限制会员数量
			MaxTip:             req.MaxTip,                                                    // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:             req.MinTip,                                                    // 荷官可接收的最小打赏金额，存储为最小单位
			TipOptions:         slice.Join(req.TipConfig, ","),                                // 打赏选项金额，以逗号隔开，存储为最小单位
		},
	}

	// 4. 构建限红模板（总商户 + 子商户） =================================================================================

	var planForms []model.TemplatePlanForm
	global.DB.Where(map[string]interface{}{
		model.TemplatePlanFormColumns.PlanFormId: req.PlanFormIds,
	}).Preload("TemplatePlanFormAreas").Find(&planForms)
	var (
		merchantTemplateBetlimits     []model.MerchantTemplateBetlimit
		merchantTemplateBetlimitAreas []model.MerchantTemplateBetlimitArea
	)

	for _, planForm := range planForms {
		for _, username := range []string{merchant.Username, subMerchant.Username} {
			templateId := guid.S()
			merchantTemplateBetlimits = append(merchantTemplateBetlimits, model.MerchantTemplateBetlimit{
				Username:   username,      // 商户账号
				TemplateId: templateId,    // 商户限红模板ID
				Name:       planForm.Name, // 模板名称
				//Content:    planForm.Content,                                         // 详细内容
				IsDefault: utils.Btoi(planForm.PlanFormId == req.DefaultPlanFormId), // 是否默认模板
			})

			if username == merchant.Username && planForm.PlanFormId == req.DefaultPlanFormId {
				merchant.TemplateId = templateId
			}

			if username == subMerchant.Username && planForm.PlanFormId == req.DefaultPlanFormId {
				subMerchant.TemplateId = templateId
			}

			for _, area := range planForm.TemplatePlanFormAreas {
				merchantTemplateBetlimitAreas = append(merchantTemplateBetlimitAreas, model.MerchantTemplateBetlimitArea{
					Username:   username,   // 商户账号
					TemplateId: templateId, // 商户限红模板ID
					//Name:         planForm.Name,     // 模板名称
					GameId:       area.GameId,       // 游戏ID
					AreaId:       area.AreaId,       // 区域ID
					MinBetAmount: area.MinBetAmount, // 最小投注金额
					MaxBetAmount: area.MaxBetAmount, // 最大投注金额
					Digits:       area.Digits,       // 投注的最小单位倍数
					PlanId:       area.PlanId,       // 限红方案ID
				})
			}
		}
	}

	// 根据角色ID创建权限
	var apiIds []uint
	global.DB.Model(&model.RoleApi{}).Where("role_id = ?", req.RoleId).Pluck("api_id", &apiIds)

	var apiModel []model.Api
	global.DB.Where("id in ?", apiIds).Find(&apiModel)

	customCasbin := NewCasbinCustom(global.CASBIN)

	// 5. 写入数据库（使用事务） =================================================================================

	tx := global.DB.Begin()

	if err = tx.Create(&merchant).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = customCasbin.UpdatePolicy(fmt.Sprintf("merchant_%d", merchant.ID), apiModel); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&subMerchant).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = customCasbin.UpdatePolicy(fmt.Sprintf("merchant_%d", subMerchant.ID), apiModel); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantCurrencies).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantTemplateBetlimits).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantTemplateBetlimitAreas).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}

	// 异步发送创建商户通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendHttpCreateMerchant(c, merchant)
	}()

	return nil
}

// Edit 编辑商户
func (s *Merchant) Edit(c *gin.Context, req *request.SaveMerchantRequest) (username string, err error) {
	// 商户
	var merchant model.Merchant
	db := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: req.Id,
	})

	if errors.Is(db.First(&merchant).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	// 商户表相关操作===============================================================================>
	merchant.MerchantName = strings.TrimSpace(req.MerchantName) // 商户名称
	merchant.Nickname = strings.TrimSpace(req.Nickname)         // 昵称
	// 如果密码不为空，重新设置密码
	if len(strings.TrimSpace(req.Password)) != 0 {
		merchant.Password = s.encryptPassword(req.Password)
	}

	merchant.CommissionRate = req.CommissionRate // 成数
	merchant.RollingRate = req.RollingRate       // 洗码比

	merchant.Status = req.Status                           // 是否允许登录:取值说明在模型里
	merchant.WalletType = model.WalletType(req.WalletType) // 钱包类型:取值说明在模型里

	merchant.Chips = slice.Join(req.ChipIds, ",")                 // 筹码
	merchant.ChipsTemplate = slice.Join(req.ChipTemplateIds, ",") // 筹码模板
	merchant.RoleId = req.RoleId                                  // 角色ID

	merchant.WalletUrl = strings.TrimSpace(req.WalletUrl) // 单钱包URL
	merchant.WhiteList = strings.TrimSpace(req.WhiteList) // 白名单IP,逗号分隔

	// 币种表相关操作===============================================================================>
	var merchantCurrency model.MerchantCurrency
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantCurrencyColumns.Username: merchant.Username,
	}).First(&merchantCurrency).Error, gorm.ErrRecordNotFound) {
		return "", err
	}
	merchantCurrency.CommissionRate = merchant.CommissionRate // 成数
	merchantCurrency.RollingRate = merchant.RollingRate       // 洗码比
	merchantCurrency.MaxTip = req.MaxTip                      // 荷官可接收的最大打赏金额，存储为最小单位
	merchantCurrency.MinTip = req.MinTip                      // 荷官可接收的最小打赏金额，存储为最小单位

	if len(req.TipConfig) > 0 {
		slice.Sort(req.TipConfig)
		if len(req.TipConfig) < 1 {
			return merchant.Username, gerror.Newf(`{#merchant.add.tip.error}`) // 配置的打赏面值超出范围
		}
		first := req.TipConfig[0]
		last := req.TipConfig[len(req.TipConfig)-1]
		if first < req.MinTip || last > req.MaxTip {
			return merchant.Username, gerror.Newf(`{#merchant.add.tip.error}`) // 配置的打赏面值超出范围
		}
		merchantCurrency.TipOptions = slice.Join(req.TipConfig, ",") // 打赏选项金额，以逗号隔开，存储为最小单位
	}

	// 开启事务
	tx := global.DB.Begin()

	if err = tx.Save(&merchant).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if err = tx.Save(&merchantCurrency).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if err = tx.Commit().Error; err != nil {
		return "", err
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()

	return merchant.Username, err
}

// Detail 商户详情
func (s *Merchant) Detail(id uint, username string, userInfo *request.CustomClaims) (*response.MerchantDetailResponse, error) {
	var err error
	var merchant model.Merchant
	db := global.DB.Model(&model.Merchant{})

	if id != 0 {
		db.Where(map[string]interface{}{
			model.MerchantColumns.Id: id,
		})
	}

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
				model.MerchantColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.MerchantColumns.Username: userInfo.Username,
			})
		}
	}

	if username != "" {
		db.Where(map[string]interface{}{
			model.MerchantColumns.Username: username,
		})
	}

	err = db.Preload("MerchantCurrency").Preload("Role").First(&merchant).Error
	if err != nil {
		return nil, err
	}

	aesKeyBase64, err := s.getAesKeyBase64(merchant.MerchantCode, merchant.AesKeySalt)
	if err != nil {
		return nil, err
	}

	var parent model.Merchant
	global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: merchant.ParentId,
	}).First(&parent)

	res := response.MerchantDetailResponse{
		Id:           merchant.ID,           // 商户ID
		MerchantCode: merchant.MerchantCode, // 商户号
		MerchantName: merchant.MerchantName, // 商户名称
		ParentId:     merchant.ParentId,     // 上级商户的username
		Username:     merchant.Username,     // 账号
		Nickname:     merchant.Nickname,     // 昵称

		CommissionRate:    merchant.CommissionRate, // 成数
		MaxCommissionRate: parent.CommissionRate,   // 上级成数
		RollingRate:       merchant.RollingRate,    // 洗码比
		MaxRollingRate:    parent.RollingRate,      // 上级洗码比

		Status:     merchant.Status,           // 是否允许登录:取值说明在模型里
		WalletType: uint(merchant.WalletType), // 钱包类型:取值说明在模型里

		ChipIds:         []uint{},            // 筹码
		ChipTemplateIds: []uint{},            // 筹码模板
		RoleId:          merchant.RoleId,     // 角色ID
		TemplateId:      merchant.TemplateId, // 商户限红模板ID

		AesKeyBase64:  aesKeyBase64,                         // aesKey的base64字符串
		Md5SaltBase64: s.getMd5SaltBase64(merchant.Md5Salt), // md5盐的base64字符串
		OrToken:       merchant.OrToken,                     // 商户api调用验证or_token
		OrTokenKey:    merchant.OrTokenKey,                  // 生成token用的商户唯一key，永久有效
		WalletUrl:     merchant.WalletUrl,                   // 单钱包URL
		WhiteList:     merchant.WhiteList,                   // 白名单IP,逗号分隔
		Currency:      merchant.MerchantCurrency.Currency,   // 商户币种

		Tip:               merchant.MerchantCurrency.Tip,               // 总打赏金额
		TotalBet:          merchant.MerchantCurrency.TotalBet,          // 总投注金额
		TotalBetEffective: merchant.MerchantCurrency.TotalBetEffective, // 总有效投注金额
		TotalWater:        merchant.MerchantCurrency.TotalWater,        // 总洗码金额
		TotalWinLost:      merchant.MerchantCurrency.TotalWinLost,      // 总输赢金额

		LastBet:          merchant.MerchantCurrency.LastBet,          // 上次投注金额
		LastBetEffective: merchant.MerchantCurrency.LastBetEffective, // 上次有效投注金额
		LastWater:        merchant.MerchantCurrency.LastWater,        // 上次洗码金额
		LastWinLost:      merchant.MerchantCurrency.LastWinLost,      // 上次输赢金额
		LastTip:          merchant.MerchantCurrency.LastTip,          // 上次打赏金额

		RoleName:  merchant.Role.RoleName,           // 角色名称
		MaxTip:    merchant.MerchantCurrency.MaxTip, // 荷官可接收的最大打赏金额，存储为最小单位
		MinTip:    merchant.MerchantCurrency.MinTip, // 荷官可接收的最小打赏金额，存储为最小单位
		Digits:    merchant.MerchantCurrency.Digits, // 币种的最小单位倍数
		TipConfig: []uint{},                         // 商户打赏配置
	}

	// 筹码
	chipStr := strings.Split(merchant.Chips, ",")
	for _, p := range chipStr {
		var chip int
		chip, err = strconv.Atoi(p) // 转换为 int
		if err != nil {
			return nil, err
		}
		res.ChipIds = append(res.ChipIds, uint(chip))
	}

	// 筹码模板
	chipTemplateStr := strings.Split(merchant.ChipsTemplate, ",")
	for _, p := range chipTemplateStr {
		var chipTemplate int
		chipTemplate, err = strconv.Atoi(p) // 转换为 int
		if err != nil {
			return nil, err
		}
		res.ChipTemplateIds = append(res.ChipTemplateIds, uint(chipTemplate))
	}

	// 打赏选项金额，以逗号隔开，存储为最小单位
	if len(merchant.MerchantCurrency.TipOptions) > 0 {
		tipOptionsStr := strings.Split(merchant.MerchantCurrency.TipOptions, ",")
		for _, p := range tipOptionsStr {
			var tipOption int
			tipOption, err = strconv.Atoi(p) // 转换为 int
			if err != nil {
				return nil, err
			}
			res.TipConfig = append(res.TipConfig, uint(tipOption))
		}
	}

	return &res, err
}

// Del 删除商户
func (s *Merchant) Del(merchant model.Merchant) (err error) {
	var playerCount int64
	global.DB.Model(&model.Player{}).Where(map[string]interface{}{
		model.PlayerColumns.MerchantId: merchant.Username,
	}).Count(&playerCount)
	if playerCount > 0 {
		return gerror.Newf(`{#merchant.del.has.player}`) // 存在玩家的不能删除
	}

	var subCount int64
	global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.ParentId: merchant.Username,
	}).Count(&subCount)
	if subCount > 0 {
		return gerror.Newf(`{#merchant.del.has.sub}`) // 请先删除子账号
	}
	tx := global.DB.Begin()
	tx.Delete(&merchant)
	return tx.Commit().Error
}

// AddChildren 创建子商户
func (s *Merchant) AddChildren(c *gin.Context, req *request.AddSubMerchantRequest) (err error) {
	var merchant model.Merchant
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.ParentId,
	}).Preload("MerchantCurrency").First(&merchant).Error, gorm.ErrRecordNotFound) {
		return gerror.Newf(`{#merchant.parent_id.no.exist}`) // 上级商户不存在
	}

	if merchant.ParentId != "root" {
		return gerror.Newf(`{#merchant.addChildren.parent.no.root}`) // 只有主商户能添加子商户
	}

	if merchant.CommissionRate < req.CommissionRate {
		return gerror.Newf(`{#merchant.addChildren.commission_rate.no.gt.root}`) // 子商户的成数不能大于主商户
	}

	if merchant.RollingRate < req.RollingRate {
		return gerror.Newf(`{#merchant.addChildren.rolling_rate.no.gt.root}`) // 子商户的洗码比不能大于主商户
	}

	// 子商户号
	subMerchantCode, err := s.makeCode()
	if err != nil {
		return err
	}

	subMerchant := model.Merchant{
		MerchantCode: subMerchantCode,                                  // 商户号
		MerchantName: strings.TrimSpace(req.MerchantName),              // 商户名称
		ParentId:     merchant.Username,                                // 上级商户的username
		RootKey:      merchant.RootKey,                                 // 顶级Key
		Username:     strings.ToLower(strings.TrimSpace(req.Username)), // 账号
		Nickname:     strings.TrimSpace(req.Nickname),                  // 昵称
		Password:     s.encryptPassword(req.Password),                  // 密码

		CommissionRate: req.CommissionRate, // 成数
		RollingRate:    req.RollingRate,    // 洗码比

		Status:     req.Status,          // 是否允许登录:取值说明在模型里
		WalletType: merchant.WalletType, // 钱包类型:取值说明在模型里

		Chips:         slice.Join(req.ChipIds, ","),         // 筹码
		ChipsTemplate: slice.Join(req.ChipTemplateIds, ","), // 筹码模板
		RoleId:        merchant.RoleId,                      // 角色ID

		AesKeySalt: merchant.AesKeySalt, // aesKey盐
		Md5Salt:    merchant.Md5Salt,    // md5盐
		WalletUrl:  merchant.WalletUrl,  // 单钱包URL
		WhiteList:  merchant.WhiteList,  // 白名单IP,逗号分隔
	}
	subMerchant.OrTokenKey = s.makeOrTokenKey() // 生成token用的商户唯一key，永久有效
	subMerchant.OrToken, err = s.makeOrToken(subMerchantCode, []byte(subMerchant.OrTokenKey))
	if err != nil {
		return err
	}

	// 3. 构建商户币种信息（总商户 + 子商户） =================================================================================

	merchantCurrency := model.MerchantCurrency{
		MerchantCurrencyId: s.makeMerchantCurrencyId(subMerchant.Username, merchant.MerchantCurrency.Currency), // 唯一标识
		Username:           subMerchant.Username,                                                               // 商户账号
		Currency:           merchant.MerchantCurrency.Currency,                                                 // 币种代码
		Digits:             merchant.MerchantCurrency.Digits,                                                   // 币种的最小单位倍数
		Status:             1,                                                                                  // 是否启用:0否,1是
		BetPaused:          1,                                                                                  // 是否暂停投注（不开放）:0否,1是
		CommissionRate:     subMerchant.CommissionRate,                                                         // 成数
		RollingRate:        subMerchant.RollingRate,                                                            // 洗码比
		LimitPlayer:        global.CONFIG.Server.LimitPlayer,                                                   // 限制会员数量
		MaxTip:             merchant.MerchantCurrency.MaxTip,                                                   // 荷官可接收的最大打赏金额，存储为最小单位
		MinTip:             merchant.MerchantCurrency.MinTip,                                                   // 荷官可接收的最小打赏金额，存储为最小单位
	}

	// 4. 构建限红模板（总商户 + 子商户） =================================================================================

	var planForms []model.TemplatePlanForm
	global.DB.Where(map[string]interface{}{
		model.TemplatePlanFormColumns.PlanFormId: req.PlanFormIds,
	}).Preload("TemplatePlanFormAreas").Find(&planForms)
	var (
		merchantTemplateBetlimits     []model.MerchantTemplateBetlimit
		merchantTemplateBetlimitAreas []model.MerchantTemplateBetlimitArea
	)

	for _, planForm := range planForms {
		templateId := guid.S()
		merchantTemplateBetlimits = append(merchantTemplateBetlimits, model.MerchantTemplateBetlimit{
			Username:   subMerchant.Username, // 商户账号
			TemplateId: templateId,           // 商户限红模板ID
			Name:       planForm.Name,        // 模板名称
			//Content:    planForm.Content,                                         // 详细内容
			IsDefault: utils.Btoi(planForm.PlanFormId == req.DefaultPlanFormId), // 是否默认模板
		})
		if planForm.PlanFormId == req.DefaultPlanFormId {
			subMerchant.TemplateId = templateId
		}
		for _, area := range planForm.TemplatePlanFormAreas {
			merchantTemplateBetlimitAreas = append(merchantTemplateBetlimitAreas, model.MerchantTemplateBetlimitArea{
				Username:   subMerchant.Username, // 商户账号
				TemplateId: templateId,           // 商户限红模板ID
				//Name:         planForm.Name,        // 模板名称
				GameId:       area.GameId,       // 游戏ID
				AreaId:       area.AreaId,       // 区域ID
				MinBetAmount: area.MinBetAmount, // 最小投注金额
				MaxBetAmount: area.MaxBetAmount, // 最大投注金额
				Digits:       area.Digits,       // 投注的最小单位倍数
				PlanId:       area.PlanId,       // 限红方案ID
			})
		}
	}

	// 根据角色ID创建权限
	var apiIds []uint
	global.DB.Model(&model.RoleApi{}).Where("role_id = ?", req.RoleId).Pluck("api_id", &apiIds)

	var apiModel []model.Api
	global.DB.Where("id in ?", apiIds).Find(&apiModel)
	customCasbin := NewCasbinCustom(global.CASBIN)
	// 5. 写入数据库（使用事务） =================================================================================

	tx := global.DB.Begin()

	if err = tx.Create(&subMerchant).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = customCasbin.UpdatePolicy(fmt.Sprintf("merchant_%d", subMerchant.ID), apiModel); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantCurrency).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantTemplateBetlimits).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&merchantTemplateBetlimitAreas).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}

	// 异步发送创建商户通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendHttpCreateMerchant(c, subMerchant)
	}()
	return nil
}

// Login 登录账号信息
func (s *Merchant) Login(req *request.LoginMerchantRequest, c *gin.Context) (*response.LoginMerchantResponse, error) {
	var merchant model.Merchant
	if err := global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).Preload("Role.DefaultMenu").Preload("MerchantCurrency").First(&merchant).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else {
			// 用户名或密码错误
			return nil, errors.New(utils.TT("http.login.merchant.no.exit"))
		}
	}

	// 是否禁止登录
	if merchant.Status == global.CannotLogin {
		return nil, errors.New(utils.TT("http.login.merchant.prohibited"))
	}

	// 密码检查
	if ok := utils.BcryptCheck(req.Password, merchant.Password); !ok {
		// 登录日志
		log := model.LoginLog{
			UserType:  model.LoginLogMerchant, // 用户类型:取值说明在模型里
			Username:  merchant.Username,      // 账号
			LoginTime: merchant.LoginTime,     // 登录时间
			LoginIp:   merchant.LoginIp,       // 登录IP
			IsOk:      0,                      // 是否成功:取值说明在模型里
			Reason:    "密码错误",             // 失败原因
		}
		global.DB.Create(&log)

		global.LOG.Info("密码错误")
		return nil, errors.New(utils.TT("http.login.merchant.password.error"))
	}

	token, claims, err := LoginMerchantToken(merchant)
	if err != nil {
		global.LOG.Error("获取token失败：", zap.Error(err))
		return nil, errors.New(utils.TT("http.login.merchant.token.get.error") + ":" + err.Error())
	}

	// 更新登录时间
	currentTime := time.Now()
	merchant.LoginTime = &currentTime
	// 更新登录IP地址
	merchant.LoginIp = c.ClientIP()
	// 更新登录IP地址
	merchant.IsActive = 1
	global.DB.Save(&merchant)

	// 登录日志
	log := model.LoginLog{
		UserType:  model.LoginLogMerchant, // 用户类型:取值说明在模型里
		Username:  merchant.Username,      // 账号
		LoginTime: merchant.LoginTime,     // 登录时间
		LoginIp:   merchant.LoginIp,       // 登录IP
		IsOk:      1,                      // 是否成功:取值说明在模型里
	}
	global.DB.Create(&log)

	res := &response.LoginMerchantResponse{
		Id:          merchant.ID,                                                         // 商户账号ID
		ParentId:    merchant.ParentId,                                                   // 上级商户的username
		Username:    merchant.Username,                                                   // 账号
		Nickname:    merchant.Nickname,                                                   // 昵称
		Status:      merchant.Status,                                                     // 是否允许登录:取值说明在模型里
		Token:       token,                                                               // 登录token
		ExpiresAt:   utils.TimestampToDate(claims.RegisteredClaims.ExpiresAt.Unix(), ""), // token过期时间
		DefaultPath: merchant.Role.DefaultMenu.Path,
		Currency:    merchant.MerchantCurrency.Currency, // 币种
	}

	return res, nil
}

// CheckUsername 检查账号是否可用
func (s *Merchant) CheckUsername(req *request.CheckUsernameRequest) (*response.CheckUsernameResponse, error) {
	var count int64
	if err := global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.Username: req.Username,
	}).Count(&count).Error; err != nil {
		return nil, err
	}
	res := &response.CheckUsernameResponse{
		Available: count == 0, // 是否可用
	}
	return res, nil
}

// RegenerateAesKeySalt 重新生成AesKeySalt
func (s *Merchant) RegenerateAesKeySalt(c *gin.Context, merchantId uint) (res *response.RegenerateAesKeySaltResponse, err error) {
	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: merchantId,
	}).First(&merchant).Error; err != nil {
		return nil, err
	}
	merchant.AesKeySalt = s.makeAesKeySalt() // aesKey盐
	if err = global.DB.Save(&merchant).Error; err != nil {
		return nil, err
	}
	aesKeyBase64, err := s.getAesKeyBase64(merchant.MerchantCode, merchant.AesKeySalt)
	if err != nil {
		return nil, err
	}
	res = &response.RegenerateAesKeySaltResponse{
		Id:           merchant.ID,  // ID
		AesKeyBase64: aesKeyBase64, // aesKey的base64字符串
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()
	return
}

// RegenerateMd5Salt 重新生成md5盐
func (s *Merchant) RegenerateMd5Salt(c *gin.Context, merchantId uint) (res *response.RegenerateMd5SaltResponse, err error) {
	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: merchantId,
	}).First(&merchant).Error; err != nil {
		return nil, err
	}
	merchant.Md5Salt = s.makeMd5Salt() // md5盐
	if err = global.DB.Save(&merchant).Error; err != nil {
		return nil, err
	}
	res = &response.RegenerateMd5SaltResponse{
		Id:            merchant.ID,                          // ID
		Md5SaltBase64: s.getMd5SaltBase64(merchant.Md5Salt), // md5盐的base64字符串
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()
	return res, nil
}

// RegenerateTokenKey 重新生成商户tokenKey
func (s *Merchant) RegenerateTokenKey(c *gin.Context, merchantId uint) (res *response.RegenerateTokenKeyResponse, err error) {
	var merchant model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: merchantId,
	}).First(&merchant).Error; err != nil {
		return nil, err
	}

	merchant.OrTokenKey = s.makeOrTokenKey() // 生成token用的商户唯一key，永久有效
	// 商户api调用验证or_token
	if merchant.OrToken, err = s.makeOrToken(merchant.MerchantCode, []byte(merchant.OrTokenKey)); err != nil {
		return nil, err
	}
	if err = global.DB.Save(&merchant).Error; err != nil {
		return nil, err
	}
	res = &response.RegenerateTokenKeyResponse{
		Id:         merchant.ID,         // ID
		OrToken:    merchant.OrToken,    // 商户api调用验证or_token
		OrTokenKey: merchant.OrTokenKey, // 生成token用的商户唯一key，永久有效
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()
	return res, nil
}

// ResetPassword 重置商户密码
func (s *Merchant) ResetPassword(c *gin.Context, id uint, userInfo *request.CustomClaims) (*response.ResetMerchantPasswordResponse, error) {
	var merchant model.Merchant
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: id,
	}).First(&merchant).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	// 商户端：不是重置自己的密码
	if userInfo != nil && userInfo.Username != merchant.Username {
		if userInfo.ParentId != global.MerchantRootName {
			return nil, errors.New("操作失败")
		} else {
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			if len(usernames) == 0 {
				return nil, errors.New("操作失败")
			}
			if !slice.Contain(usernames, merchant.Username) {
				return nil, errors.New("操作失败")
			}
		}
	}
	resetPassword := global.CONFIG.Server.ResetPassword
	if len(resetPassword) == 0 {
		resetPassword = "123456"
	}

	merchant.Password = utils.BcryptHash(strings.TrimSpace(resetPassword))

	if err := global.DB.Save(&merchant).Error; err != nil {
		return nil, err
	}
	res := response.ResetMerchantPasswordResponse{
		Id:       merchant.ID,       // 商户账号ID
		Username: merchant.Username, // 账号
		Password: resetPassword,     // 密码
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()
	return &res, nil
}

// GetBox 商户下拉框数据
func (s *Merchant) GetBox() (list []response.MerchantBoxResponse, err error) {
	var merchants []model.Merchant
	if err = global.DB.Find(&merchants).Error; err != nil {
		return
	}

	// 赋值
	for _, merchant := range merchants {
		tmp := response.MerchantBoxResponse{
			Id:         merchant.ID,         // ID
			Username:   merchant.Username,   // 账号
			TemplateId: merchant.TemplateId, // 商户限红模板ID
		}

		list = append(list, tmp)
	}
	return list, err
}

// GetApiKey 获取api相关密钥
func (s *Merchant) GetApiKey(id uint) (*response.MerchantApiKeyResponse, error) {
	var err error
	var merchant model.Merchant
	db := global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.Id: id,
	})

	err = db.First(&merchant).Error
	if err != nil {
		return nil, err
	}

	aesKeyBase64, err := s.getAesKeyBase64(merchant.MerchantCode, merchant.AesKeySalt)
	if err != nil {
		return nil, err
	}

	res := response.MerchantApiKeyResponse{
		Id:            merchant.ID,                          // 商户ID
		MerchantCode:  merchant.MerchantCode,                // 商户号
		Username:      merchant.Username,                    // 账号
		AesKeyBase64:  aesKeyBase64,                         // aesKey的base64字符串
		Md5SaltBase64: s.getMd5SaltBase64(merchant.Md5Salt), // md5盐的base64字符串
		OrToken:       merchant.OrToken,                     // 商户api调用验证or_token
		OrTokenKey:    merchant.OrTokenKey,                  // 生成token用的商户唯一key，永久有效

	}
	return &res, err
}

// GetParentBox 主商户下拉框数据
func (s *Merchant) GetParentBox() (list []response.MerchantBoxResponse, err error) {
	var merchants []model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.ParentId: "root",
	}).Find(&merchants).Error; err != nil {
		return
	}

	// 赋值
	for _, merchant := range merchants {
		tmp := response.MerchantBoxResponse{
			Id:         merchant.ID,         // ID
			Username:   merchant.Username,   // 账号
			TemplateId: merchant.TemplateId, // 商户限红模板ID
		}
		list = append(list, tmp)
	}
	return list, err
}

// GetSubBox 子商户下拉框数据
func (s *Merchant) GetSubBox(username string) (list []response.MerchantBoxResponse, err error) {
	var merchants []model.Merchant
	if err = global.DB.Where(map[string]interface{}{
		model.MerchantColumns.ParentId: username,
	}).Find(&merchants).Error; err != nil {
		return
	}

	// 赋值
	for _, merchant := range merchants {
		tmp := response.MerchantBoxResponse{
			Id:         merchant.ID,         // ID
			Username:   merchant.Username,   // 账号
			TemplateId: merchant.TemplateId, // 商户限红模板ID
		}
		list = append(list, tmp)
	}
	return list, err
}

// EditCurrentPassword 编辑当前登录的商户密码
func (s *Merchant) EditCurrentPassword(c *gin.Context, req *request.EditCurrentMerchantPasswordRequest, id uint) error {
	var merchant model.Merchant
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantColumns.Id: id,
	}).Where("id = ?", id).First(&merchant).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	// 密码检查
	if ok := utils.BcryptCheck(req.OldPassword, merchant.Password); !ok {
		return errors.New(utils.TT("http.login.merchant.password.error"))
	}

	merchant.Password = utils.BcryptHash(strings.TrimSpace(req.NewPassword))

	if err := global.DB.Save(&merchant).Error; err != nil {
		return err
	}

	// 异步发送游戏配置操作通知，避免阻塞接口请求
	go func() {
		Service.GameNotifyLog.SendUpdateMerchantHttp(c, merchant)
	}()

	return nil
}

// ==========================================================================以下方法仅供内部调用⬇

// 生成商户号
func (s *Merchant) makeCode() (merchantCode string, err error) {
	for i := 0; i < 5; i++ {
		merchantCode = grand.S(8, false)
		var merchant model.Merchant
		err = global.DB.Where(map[string]interface{}{
			model.MerchantColumns.MerchantCode: merchantCode,
		}).First(&merchant).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			break
		} else if err != nil {
			return "", err
		}
	}
	if merchantCode == "" {
		err = gerror.Newf(`{#merchant.add.make.code.error}`) // 生成商户号失败
		return
	}
	return merchantCode, nil
}

// 生成商户币种唯一标识
func (s *Merchant) makeMerchantCurrencyId(username, currencyCode string) string {
	return username + "_" + currencyCode
}

// 生成aesKey盐
func (s *Merchant) makeAesKeySalt() string {
	return grand.S(16, false)
}

// 生成md5盐
func (s *Merchant) makeMd5Salt() string {
	return grand.S(16, false)
}

// 生成or_token
func (s *Merchant) makeOrToken(merchantCode string, orTokenKey []byte) (string, error) {
	// 过期时间（纳秒）
	var expiresDuration time.Duration
	if global.CONFIG.Jwt.MerchantTokenExpiresTime == 0 {
		expiresDuration = 30 * 24 * time.Hour // 30天，单位是纳秒
	} else {
		expiresDuration = time.Duration(global.CONFIG.Jwt.MerchantTokenExpiresTime) * time.Second // 把秒转换成纳秒
	}
	claims := jwt.MapClaims{
		"merchant_code": merchantCode,                           // 商户号
		"ext":           time.Now().Add(expiresDuration).Unix(), // 过期时间
		"iat":           time.Now().Unix(),                      // 发行时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(orTokenKey)
}

// 生成or_token_key
func (s *Merchant) makeOrTokenKey() string {
	return grand.S(32, false)
}

// 获取aesKey的base64字符串
func (s *Merchant) getAesKeyBase64(MerchantCode string, AesKeySalt string) (aesKeyBase64 string, err error) {
	aesHash := sha256.New()                                   // 使用SHA-256加密
	_, err = aesHash.Write([]byte(MerchantCode + AesKeySalt)) // 商户号 + 动态盐值
	if err != nil {
		return "", err
	}
	aesKeyBase64 = base64.StdEncoding.EncodeToString(aesHash.Sum(nil)) // aesKey的base64字符串
	return
}

// 获取md5盐的base64字符串
func (s *Merchant) getMd5SaltBase64(md5Salt string) string {
	return base64.StdEncoding.EncodeToString([]byte(md5Salt)) // md5盐的base64字符串
}

// 生成密码
func (s *Merchant) encryptPassword(password string) string {
	return utils.BcryptHash(strings.TrimSpace(password))
}
