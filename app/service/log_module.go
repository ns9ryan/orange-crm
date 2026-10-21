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

type LogModule struct{}

// AdminLoginLogList 管理员登录日志列表
func (s *LogModule) AdminLoginLogList(req *request.AdminLoginLogListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.AdminLoginLogResponse{},
	}
	// 排序
	order := "id desc"
	db := global.DB.Model(&model.AdminLoginLog{}).Order(order).Preload("Admin.Role")

	// 账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.AdminLoginLogColumns.Username: req.Username,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.AdminLoginLogColumns.LoginTime), req.StartTime)
	}

	// 开始时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.AdminLoginLogColumns.LoginTime), req.EndTime)
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
	var adminLoginLogs []model.AdminLoginLog
	if err = db.Find(&adminLoginLogs).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.AdminLoginLogResponse
	for _, log := range adminLoginLogs {
		tmp := response.AdminLoginLogResponse{
			Id:        log.ID,                             // ID
			Username:  log.Username,                       // 账号
			Nickname:  log.Admin.Nickname,                 // 昵称
			Avatar:    log.Admin.Avatar,                   // 头像
			RoleName:  log.Admin.Role.RoleName,            // 角色名称
			LoginIp:   log.LoginIp,                        // 登录IP
			LoginTime: log.LoginTime.Format(time.RFC3339), // 登录时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// MerchantLoginLogList 商户登录日志列表
func (s *LogModule) MerchantLoginLogList(req *request.MerchantLoginLogListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.MerchantLoginLogResponse{},
	}

	// 排序
	order := "id desc"
	db := global.DB.Model(&model.LoginLog{}).Where(map[string]interface{}{
		model.LoginLogColumns.UserType: model.LoginLogMerchant,
	}).Order(order).Preload("Merchant.Role")

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的登录日志
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.LoginLogColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.LoginLogColumns.Username: userInfo.Username,
			})
		}
	}

	// 账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.LoginLogColumns.Username: req.Username,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.LoginLogColumns.LoginTime), req.StartTime)
	}

	// 结束时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.LoginLogColumns.LoginTime), req.EndTime)
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
	var loginLog []model.LoginLog
	if err = db.Find(&loginLog).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.MerchantLoginLogResponse
	for _, log := range loginLog {
		tmp := response.MerchantLoginLogResponse{
			Id:        log.ID,                             // ID
			Username:  log.Username,                       // 账号
			Nickname:  log.Merchant.Nickname,              // 昵称
			RoleName:  log.Merchant.Role.RoleName,         // 角色名称
			LoginIp:   log.LoginIp,                        // 登录IP
			LoginTime: log.LoginTime.Format(time.RFC3339), // 登录时间
			IsOk:      log.IsOk,                           // 是否成功:取值说明在模型里
			Reason:    log.Reason,                         // 失败原因
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// PlayerLoginLogList 会员登录日志列表
func (s *LogModule) PlayerLoginLogList(req *request.PlayerLoginLogListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.PlayerLoginLogResponse{},
	}

	// 排序
	order := "id desc"
	db := global.DB.Model(&model.LoginLog{}).Where(map[string]interface{}{
		model.LoginLogColumns.UserType: model.LoginLogPlayer, // 玩家登录类型
	}).Preload("Player").Order(order)

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户和自己的登录日志
			db.Where(map[string]interface{}{
				model.LoginLogColumns.RootKey: userInfo.ID,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.LoginLogColumns.MerchantId: userInfo.Username,
			})
		}
	}

	// 商户账号
	if req.MerchantId != "" {
		db.Where(map[string]interface{}{
			model.LoginLogColumns.MerchantId: req.MerchantId,
		})
	}

	// 玩家账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.LoginLogColumns.Username: req.Username,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.LoginLogColumns.LoginTime), req.StartTime)
	}

	// 结束时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.LoginLogColumns.LoginTime), req.EndTime)
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
	var loginLog []model.LoginLog
	if err = db.Find(&loginLog).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.PlayerLoginLogResponse
	for _, log := range loginLog {
		tmp := response.PlayerLoginLogResponse{
			Id:        log.ID,                             // ID
			Username:  log.Username,                       // 账号
			Nickname:  log.Player.Nickname,                // 昵称
			LoginIp:   log.LoginIp,                        // 登录IP
			LoginTime: log.LoginTime.Format(time.RFC3339), // 登录时间
			IsOk:      log.IsOk,                           // 是否成功:取值说明在模型里
			Reason:    log.Reason,                         // 失败原因
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// AdminOperationLogList 管理员操作日志列表
func (s *LogModule) AdminOperationLogList(req *request.AdminOperationLogListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.AdminOperationLogListResponse{},
	}
	db := global.DB.Model(&model.AdminOperationLog{}).Order("id desc")

	// 账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.AdminOperationLogColumns.Username: req.Username,
		})
	}

	// 事件类型
	if req.AdminOperationEventType != "" {
		db.Where(map[string]interface{}{
			model.AdminOperationLogColumns.AdminOperationEventType: req.AdminOperationEventType,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.AdminOperationLogColumns.CreatedTime), req.StartTime)
	}

	// 结束时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.AdminOperationLogColumns.CreatedTime), req.EndTime)
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
	var adminOperationLogs []model.AdminOperationLog
	if err = db.Find(&adminOperationLogs).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.AdminOperationLogListResponse
	for _, log := range adminOperationLogs {
		tmp := response.AdminOperationLogListResponse{
			Id:                      log.ID,                               // ID
			Username:                log.Username,                         // 账号
			Module:                  log.Module,                           // 模块名
			Action:                  log.Action,                           // 操作动作
			Content:                 log.Content,                          // 操作内容描述
			AdminOperationEventType: uint(log.AdminOperationEventType),    // 事件类型
			CreatedTime:             log.CreatedTime.Format(time.RFC3339), // 操作时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// MerchantOperationLogList 商户操作日志列表
func (s *LogModule) MerchantOperationLogList(req *request.MerchantOperationLogListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.MerchantOperationLogListResponse{},
	}
	db := global.DB.Model(&model.MerchantOperationLog{}).Order("id desc")

	// 商户端
	if userInfo != nil {
		if userInfo.ParentId == global.MerchantRootName {
			// 顶级商户：可以看到下面所有子商户的操作日志
			// 查询底下的所有子商户的username
			var usernames []string
			global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
				model.MerchantColumns.ParentId: userInfo.Username,
			}).Pluck("username", &usernames)
			// 要加上查询自己的
			usernames = append(usernames, userInfo.Username)
			db.Where(map[string]interface{}{
				model.MerchantOperationLogColumns.Username: usernames,
			})
		} else {
			// 子商户：只能看到自己的
			db.Where(map[string]interface{}{
				model.MerchantOperationLogColumns.Username: userInfo.Username,
			})
		}
	}

	// 账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.MerchantOperationLogColumns.Username: req.Username,
		})
	}

	// 事件类型
	if req.MerchantOperationEventType != "" {
		db.Where(map[string]interface{}{
			model.MerchantOperationLogColumns.MerchantOperationEventType: req.MerchantOperationEventType,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.MerchantOperationLogColumns.CreatedTime), req.StartTime)
	}

	// 结束时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.MerchantOperationLogColumns.CreatedTime), req.EndTime)
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
	var merchantOperationLogs []model.MerchantOperationLog
	if err = db.Find(&merchantOperationLogs).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.MerchantOperationLogListResponse
	for _, log := range merchantOperationLogs {
		tmp := response.MerchantOperationLogListResponse{
			Id:                         log.ID,                               // ID
			Username:                   log.Username,                         // 账号
			Module:                     log.Module,                           // 模块名
			Action:                     log.Action,                           // 操作动作
			Content:                    log.Content,                          // 操作内容描述
			MerchantOperationEventType: uint(log.MerchantOperationEventType), // 事件类型
			CreatedTime:                log.CreatedTime.Format(time.RFC3339), // 操作时间
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// ControlLoginLogList 控端登录日志列表
func (s *LogModule) ControlLoginLogList(req *request.ControlLoginLogListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.ControlLoginLogResponse{},
	}
	// 排序
	order := "id desc"
	db := global.DB.Model(&model.ControlLoginLog{}).Order(order)

	// 账号
	if req.Username != "" {
		db.Where(map[string]interface{}{
			model.ControlLoginLogColumns.Username: req.Username,
		})
	}

	// 开始时间
	if req.StartTime != "" {
		db.Where(fmt.Sprintf("%s >= ?", model.ControlLoginLogColumns.LoginTime), req.StartTime)
	}

	// 开始时间
	if req.EndTime != "" {
		db.Where(fmt.Sprintf("%s <= ?", model.ControlLoginLogColumns.LoginTime), req.EndTime)
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
	var controlLoginLogs []model.ControlLoginLog
	if err = db.Find(&controlLoginLogs).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.ControlLoginLogResponse
	for _, log := range controlLoginLogs {
		tmp := response.ControlLoginLogResponse{
			Id:        log.ID,                             // ID
			Username:  log.Username,                       // 账号
			LoginIp:   log.LoginIp,                        // 登录IP
			LoginTime: log.LoginTime.Format(time.RFC3339), // 登录时间
			IsOk:      log.IsOk,                           // 是否成功:取值说明在模型里
			Reason:    log.Reason,                         // 失败原因
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}
