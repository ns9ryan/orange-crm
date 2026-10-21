package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"strconv"
	"strings"
)

type DataInit struct{}

// AddMerchants 批量添加商户
func (s *DataInit) AddMerchants(c *gin.Context, req *request.AddMerchantsRequest) error {

	// 系统筹码
	var chipIds []uint
	global.DB.Model(&model.SystemChip{}).Pluck(model.SystemChipColumns.Id, &chipIds)

	cStart := 1
	for _, v := range req.AddData {
		currency := v.Currency // 币种
		var template model.TemplatePlanForm
		global.DB.Where(map[string]interface{}{
			model.TemplatePlanFormColumns.Name: fmt.Sprintf("%s模板", currency),
		}).First(&template)
		if template.ID == 0 {
			return errors.New("限红模板错误")
		}

		var currencyModel model.Currency
		global.DB.Where(map[string]interface{}{
			model.CurrencyColumns.Code: currency,
		}).First(&currencyModel)
		if currencyModel.ID == 0 {
			return errors.New("币种不存在")
		}

		// 创建主商户==================================================================================================
		tmp := request.SaveMerchantRequest{
			MerchantName: fmt.Sprintf("测试商户%03d", cStart),         // 商户名称
			Username:     v.Username,                              // 账号
			SubUsername:  fmt.Sprintf("sub_merchant%03d", cStart), // 子账号
			Nickname:     fmt.Sprintf("测试商户%03d", cStart),         // 昵称
			Password:     v.Password,                              // 密码

			CommissionRate: 90000, // 成数
			RollingRate:    1000,  // 洗码比

			ChipIds:           chipIds,                       // 筹码ID数组
			ChipTemplateIds:   v.ChipTemplate,                // 筹码模板ID数组
			RoleId:            2,                             // 角色ID
			PlanFormIds:       []string{template.PlanFormId}, // 系统限红模板ID数组
			DefaultPlanFormId: template.PlanFormId,           // 默认系统限红模板ID
			Currency:          currency,                      // 币种
			MaxTip:            currencyModel.MaxTip,          // 荷官可接收的最大打赏金额，存储为最小单位
			MinTip:            currencyModel.MinTip,          // 荷官可接收的最小打赏金额，存储为最小单位
		}
		tmp.ChipTemplateIds = v.ChipTemplate
		if err := Service.Merchant.Add(c, &tmp); err != nil {
			return err
		}

		// 创建子商户==================================================================================================
		for _, son := range v.Sons {
			// 创建子商户
			subTmp := request.AddSubMerchantRequest{
				ParentId:     tmp.Username,     // 上级商户ID
				MerchantName: son.MerchantName, // 商户名称
				Username:     son.Username,     // 账号
				Nickname:     son.MerchantName, // 昵称
				Password:     son.Password,     // 密码

				CommissionRate: 90000, // 成数
				RollingRate:    1000,  // 洗码比

				ChipIds: chipIds, // 筹码ID数组

				ChipTemplateIds: tmp.ChipTemplateIds, // 筹码模板ID数组
				RoleId:          2,                   // 角色ID

				PlanFormIds:       []string{template.PlanFormId}, // 系统限红模板ID数组
				DefaultPlanFormId: template.PlanFormId,           // 默认系统限红模板ID
			}
			if err := Service.Merchant.AddChildren(c, &subTmp); err != nil {
				return err
			}
		}

		cStart++
	}
	return nil
}

// AddPlayers 批量添加玩家
func (s *DataInit) AddPlayers(req *request.AddPlayersRequest) error {
	var merchants []model.Merchant
	db := global.DB.Where(fmt.Sprintf("%s != ?", model.MerchantColumns.ParentId), "root").Preload("MerchantCurrency").Order("id asc").Find(&merchants)
	if req.MerchantId != "" {
		db.Where(map[string]interface{}{
			model.MerchantColumns.Username: req.MerchantId,
		})
	}
	db.Find(&merchants)

	var data []request.SavePlayerRequest

	inx := 1

	for _, merchant := range merchants {
		// 筹码
		var chipIds []uint
		chipStr := strings.Split(merchant.ChipsTemplate, ",")
		for _, p := range chipStr {
			var chip int
			chip, err := strconv.Atoi(p) // 转换为 int
			if err != nil {
				return err
			}
			chipIds = append(chipIds, uint(chip))
		}

		for i := req.Start; i <= req.End; i++ {
			tmp := request.SavePlayerRequest{
				MerchantId: merchant.Username,                   // 商户账号
				Username:   fmt.Sprintf("player%03d", inx),      // 用户名,创建的时候必传
				Nickname:   fmt.Sprintf("玩家%03d", inx),          // 昵称
				ParentId:   "root",                              // 上级玩家账号
				PlayerType: 1,                                   // 玩家类型:取值说明在模型里
				MoneyType:  1,                                   // 玩家模式类型:取值说明在模型里
				Password:   fmt.Sprintf("p_5188layer%03d", inx), // 密码,创建的时候必传
				ChipIds:    chipIds,                             // 筹码ID数组
				TemplateId: merchant.TemplateId,                 // 商户限红模板ID
			}
			if req.UsernamePrefix != "" {
				tmp.Username = fmt.Sprintf("%s%03d", req.UsernamePrefix, i) // 用户名,创建的时候必传
				tmp.Password = "123456"                                     // 用户名,创建的时候必传
			}
			if err := Service.Player.Add(&tmp, merchant); err != nil {
				return err
			}
			inx++
			data = append(data, tmp)
		}
	}

	return nil
}

// Deposit 批量存款上分
func (s *DataInit) Deposit(req *request.PlayersDepositRequest) error {
	var wallets []model.PlayerWallet
	global.DB.Find(&wallets)

	for _, wallet := range wallets {
		nowAmount := req.Amount * wallet.Digits
		wallet.CreditSum = wallet.CreditSum + nowAmount
		wallet.Balance = wallet.Balance + nowAmount
		global.DB.Save(wallet)
	}

	return nil
}
