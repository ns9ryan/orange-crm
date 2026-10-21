package service

import (
	"errors"
	"github.com/gogf/gf/v2/util/guid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"time"
)

type SystemNotice struct{}

// List 系统公告列表
func (s *SystemNotice) List(req *request.SystemNoticeListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.SystemNoticeListResponse{},
	}
	db := global.DB.Model(&model.SystemNotice{})

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
	var systemNotices []model.SystemNotice
	if err = db.Find(&systemNotices).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.SystemNoticeListResponse
	for _, systemNotice := range systemNotices {
		tmp := response.SystemNoticeListResponse{
			Id:          systemNotice.ID,                                // 游戏桌台ID
			NoticeId:    systemNotice.NoticeId,                          // 公告ID
			Title:       systemNotice.Title,                             // 公告标题
			Content:     systemNotice.Content,                           // 公告内容
			StartTime:   systemNotice.StartTime.Format(time.DateTime),   // 开始时间
			EndTime:     systemNotice.EndTime.Format(time.DateTime),     // 结束时间
			IsAll:       systemNotice.IsAll,                             // 是否发给所有商户:取值说明在模型里
			MerchantIds: []string{},                                     // 指定发给的商户
			OperatorId:  systemNotice.OperatorId,                        // 发布公告的账号
			CreatedTime: systemNotice.CreatedTime.Format(time.DateTime), // 创建时间
		}

		// 发送给指定商户
		global.DB.Model(&model.SystemNoticeMerchant{}).Where(map[string]interface{}{
			model.SystemNoticeMerchantColumns.NoticeId: systemNotice.NoticeId,
		}).Pluck(model.SystemNoticeMerchantColumns.MerchantId, &tmp.MerchantIds)
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建系统公告
func (s *SystemNotice) Add(req *request.SaveSystemNoticeRequest, operatorId string) error {
	startTime, err := time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(time.DateTime, req.EndTime)
	if err != nil {
		return err
	}
	noticeId := guid.S() // 公告ID
	systemNotice := model.SystemNotice{
		NoticeId:   noticeId,    // 公告ID
		Title:      req.Title,   // 公告标题
		Content:    req.Content, // 公告内容
		StartTime:  &startTime,  // 开始时间
		EndTime:    &endTime,    // 结束时间
		IsAll:      req.IsAll,   // 是否发给所有商户:取值说明在模型里
		OperatorId: operatorId,  // 发布公告的账号
	}

	var systemNoticeMerchants []model.SystemNoticeMerchant
	var merchantIds []string
	if req.IsAll == 0 {
		// 发给所有商户
		global.DB.Model(&model.Merchant{}).Pluck(model.MerchantColumns.Username, &merchantIds)
	} else {
		// 发送给指定商户
		global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
			model.MerchantColumns.Username: req.MerchantIds,
		}).Pluck(model.MerchantColumns.Username, &merchantIds)
	}
	for _, merchantId := range merchantIds {
		systemNoticeMerchant := model.SystemNoticeMerchant{
			NoticeId:   noticeId,   // 公告ID
			MerchantId: merchantId, // 商户ID
		}
		systemNoticeMerchants = append(systemNoticeMerchants, systemNoticeMerchant)
	}

	tx := global.DB.Begin()

	if err = tx.Create(&systemNotice).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(systemNoticeMerchants) != 0 {
		if err = tx.Create(&systemNoticeMerchants).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Edit 编辑系统公告
func (s *SystemNotice) Edit(req *request.SaveSystemNoticeRequest) error {
	var systemNotice model.SystemNotice
	if errors.Is(global.DB.First(&systemNotice, req.Id).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	startTime, err := time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(time.DateTime, req.EndTime)
	if err != nil {
		return err
	}
	systemNotice.Title = req.Title      // 公告标题
	systemNotice.Content = req.Content  // 公告内容
	systemNotice.StartTime = &startTime // 开始时间
	systemNotice.EndTime = &endTime     // 结束时间
	systemNotice.IsAll = req.IsAll      // 是否发给所有商户:取值说明在模型里

	var systemNoticeMerchants []model.SystemNoticeMerchant
	var merchantIds []string
	if req.IsAll == 0 {
		// 发给所有商户
		global.DB.Model(&model.Merchant{}).Pluck(model.MerchantColumns.Username, &merchantIds)
	} else {
		// 发送给指定商户
		global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
			model.MerchantColumns.Username: req.MerchantIds,
		}).Pluck(model.MerchantColumns.Username, &merchantIds)
	}

	for _, merchantId := range merchantIds {
		systemNoticeMerchant := model.SystemNoticeMerchant{
			NoticeId:   systemNotice.NoticeId, // 公告ID
			MerchantId: merchantId,            // 商户ID
		}
		systemNoticeMerchants = append(systemNoticeMerchants, systemNoticeMerchant)
	}

	tx := global.DB.Begin()

	if err = tx.Save(&systemNotice).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Where(model.SystemNoticeMerchantColumns.NoticeId, systemNotice.NoticeId).Delete(&model.SystemNoticeMerchant{})

	if len(systemNoticeMerchants) != 0 {
		if err = tx.Create(&systemNoticeMerchants).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// Detail 系统公告详情
func (s *SystemNotice) Detail(id uint) (*response.SystemNoticeDetailResponse, error) {
	var systemNotice model.SystemNotice

	if err := global.DB.Where(map[string]interface{}{
		model.SystemNoticeColumns.Id: id,
	}).First(&systemNotice).Error; err != nil {
		return nil, err
	}

	res := response.SystemNoticeDetailResponse{
		Id:          systemNotice.ID,                              // 游戏桌台ID
		NoticeId:    systemNotice.NoticeId,                        // 公告ID
		Title:       systemNotice.Title,                           // 公告标题
		Content:     systemNotice.Content,                         // 公告内容
		StartTime:   systemNotice.StartTime.Format(time.DateTime), // 开始时间
		EndTime:     systemNotice.EndTime.Format(time.DateTime),   // 结束时间
		IsAll:       systemNotice.IsAll,                           // 是否发给所有商户:取值说明在模型里
		MerchantIds: []string{},                                   // 指定发给的商户
		OperatorId:  systemNotice.OperatorId,                      // 发布公告的账号
	}

	// 发送给指定商户
	global.DB.Model(&model.SystemNoticeMerchant{}).Where(map[string]interface{}{
		model.SystemNoticeMerchantColumns.NoticeId: systemNotice.NoticeId,
	}).Pluck(model.SystemNoticeMerchantColumns.MerchantId, &res.MerchantIds)

	return &res, nil
}

// Del 删除系统公告
func (s *SystemNotice) Del(id uint) (title string, err error) {
	var systemNotice model.SystemNotice
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.SystemNoticeColumns.Id: id,
	}).First(&systemNotice).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}
	return systemNotice.Title, global.DB.Delete(&systemNotice).Error
}

// ListForNew 商户新消息
func (s *SystemNotice) ListForNew(userInfo *request.CustomClaims) (res []response.LoginMerchantNoticeListResponse, err error) {
	// 系统消息
	var notices []model.SystemNoticeMerchant
	global.DB.Where(map[string]interface{}{
		model.SystemNoticeMerchantColumns.MerchantId: userInfo.Username,
		model.SystemNoticeMerchantColumns.Read:       0, // 未读
	}).Preload("SystemNotice").Find(&notices)

	// 获取今天的开始时间和结束时间
	currentTime := time.Now()
	todayStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()) // 今天的 00:00:00
	todayEndTime := todayStartTime.Add(24*time.Hour - time.Second)                                                              // 今天的 23:59:59

	for _, notice := range notices {
		// 还没过期的
		if (*notice.SystemNotice.EndTime).After(todayEndTime) == true {
			noticeTemp := response.LoginMerchantNoticeListResponse{
				Id:       notice.ID,                   // ID
				NoticeId: notice.NoticeId,             // 公告ID
				Title:    notice.SystemNotice.Title,   // 公告标题
				Content:  notice.SystemNotice.Content, // 公告内容
			}
			res = append(res, noticeTemp)
		}
	}
	return
}

// ListForSelf 商户收到的系统公共列表
func (s *SystemNotice) ListForSelf(req *request.SystemNoticeListForSelfRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.SystemNoticeListForSelfResponse{},
	}

	// 查询符合条件的公共ID
	db := global.DB.Model(&model.SystemNoticeMerchant{}).Where(map[string]interface{}{
		model.SystemNoticeMerchantColumns.MerchantId: userInfo.Username,
	}).Preload("SystemNotice")

	// 是否已读
	if req.Read != nil {
		db.Where(map[string]interface{}{
			model.SystemNoticeMerchantColumns.Read: req.Read,
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
	var notices []model.SystemNoticeMerchant
	if err = db.Find(&notices).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.SystemNoticeListForSelfResponse
	for _, notice := range notices {
		tmp := response.SystemNoticeListForSelfResponse{
			Id:        notice.ID,                                           // 游戏桌台ID
			NoticeId:  notice.NoticeId,                                     // 公告ID
			Title:     notice.SystemNotice.Title,                           // 公告标题
			Content:   notice.SystemNotice.Content,                         // 公告内容
			StartTime: notice.SystemNotice.StartTime.Format(time.DateTime), // 开始时间
			EndTime:   notice.SystemNotice.EndTime.Format(time.DateTime),   // 结束时间
			Read:      notice.Read,                                         // 是否已读:取值说明在模型里
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// MarkRead 标记已读
func (s *SystemNotice) MarkRead(merchantNotice model.SystemNoticeMerchant, userInfo *request.CustomClaims) error {
	merchantNotice.Read = 1                       // 已读
	merchantNotice.OperatorId = userInfo.Username // 操作的账号
	return global.DB.Save(&merchantNotice).Error
}
