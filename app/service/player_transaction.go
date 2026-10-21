package service

import (
	"fmt"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"time"
)

type PlayerTransaction struct{}

// CreditAndDebitList 上分/下分记录列表
func (s *PlayerTransaction) CreditAndDebitList(req *request.CreditAndDebitListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.CreditAndDebitListResponse{},
	}
	db := global.DB.Model(&model.PlayerTransaction{}).Where(map[string]interface{}{
		model.PlayerTransactionColumns.TransactionType: []model.TransactionType{
			// 存款上分
			model.TransactionDeposit,
			// 取现下分
			model.TransactionWithdraw,
		},
		// 只查找余额字段
		model.PlayerTransactionColumns.WalletField: model.PlayerWalletColumns.Balance, // 账户余额
	})

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的玩家
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.PlayerTransactionColumns.MerchantId: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.PlayerTransactionColumns.MerchantId: userInfo.Username,
			})
		}
	} else {
		// 总网端才能搜索商户
		// 商户账号
		if req.MerchantId != "" {
			db.Where(map[string]interface{}{
				model.PlayerTransactionColumns.MerchantId: req.MerchantId,
			})
		}
	}

	// 会员账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.PlayerTransactionColumns.Username: req.Username,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.PlayerTransactionColumns.CreatedTime), req.StartTime)
	}

	// 开始时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.PlayerTransactionColumns.CreatedTime), req.EndTime)
	}

	// 总条数
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
	var playerTransactions []model.PlayerTransaction
	if err = db.Find(&playerTransactions).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.CreditAndDebitListResponse
	for _, playerTransaction := range playerTransactions {
		tmp := response.CreditAndDebitListResponse{
			ID:              playerTransaction.ID,                                // ID
			RootKey:         playerTransaction.RootKey,                           // 顶级Key
			MerchantId:      playerTransaction.MerchantId,                        // 对应商户表的username
			Username:        playerTransaction.Username,                          // 玩家账号,每个玩家的唯一值
			Currency:        playerTransaction.Currency,                          // 币种
			Digits:          playerTransaction.Digits,                            // 币种的最小单位倍数
			TransactionType: uint(playerTransaction.TransactionType),             // 交易类型:取值说明在模型里
			WalletField:     playerTransaction.WalletField,                       // 钱包字段
			Amount:          playerTransaction.Amount,                            // 变动金额,存储为最小单位
			OrangeTransId:   playerTransaction.OrangeTransId,                     // 系统产生的订单号,全局唯一
			MerchantTransId: playerTransaction.MerchantTransId,                   // 商户产生的订单号,api商户传过来的订单号
			Remark:          playerTransaction.Remark,                            // 备注
			CreatedTime:     playerTransaction.CreatedTime.Format(time.DateTime), // 创建时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}
