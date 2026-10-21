package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"time"
)

type MerchantNotice struct{}

// List 商户公告列表
func (s *MerchantNotice) List(req *request.MerchantNoticeListRequest, userInfo *request.CustomClaims) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.MerchantNoticeListResponse{},
	}
	db := global.DB.Model(&model.MerchantNotice{})

	// 主商户可以看到自己发布的公告
	if userInfo.ParentId == global.MerchantRootName {
		db.Where(map[string]interface{}{
			model.MerchantNoticeColumns.MerchantId: userInfo.Username,
		})
	} else {
		// 子商户 可以看到主商户发布的公告
		db.Where(map[string]interface{}{
			model.MerchantNoticeColumns.MerchantId: userInfo.ParentId,
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
	var merchantNotices []model.MerchantNotice
	if err = db.Find(&merchantNotices).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.MerchantNoticeListResponse
	for _, notice := range merchantNotices {
		tmp := response.MerchantNoticeListResponse{
			Id:          notice.ID,                               // 游戏桌台ID
			NoticeId:    notice.NoticeId,                         // 公告ID
			Title:       notice.Title,                            // 公告标题
			Content:     notice.Content,                          // 公告内容
			StartTime:   notice.StartTime.Format(time.RFC3339),   // 开始时间
			EndTime:     notice.EndTime.Format(time.RFC3339),     // 结束时间
			IsAll:       notice.IsAll,                            // 是否发给所有商户:取值说明在模型里
			PlayerIds:   []string{},                              // 指定发给的会员
			OperatorId:  notice.OperatorId,                       // 发布公告的账号
			CreatedTime: notice.CreatedTime.Format(time.RFC3339), // 创建时间
		}

		// 发送给指定会员
		global.DB.Model(&model.MerchantNoticePlayer{}).Where(map[string]interface{}{
			model.MerchantNoticePlayerColumns.NoticeId: notice.NoticeId,
		}).Pluck(model.MerchantNoticePlayerColumns.PlayerId, &tmp.PlayerIds)
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建商户公告
func (s *MerchantNotice) Add(req *request.SaveMerchantNoticeRequest, userInfo *request.CustomClaims) error {
	startTime, err := time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(time.DateTime, req.EndTime)
	if err != nil {
		return err
	}

	// Time: 当前时间
	now := time.Now().Format("20060102")

	var count int64
	global.DB.Model(&model.MerchantNotice{}).Count(&count)
	countIndex := count + 1
	noticeId := fmt.Sprintf("MN%s%d%03d", now, userInfo.ID, countIndex) // 公告ID

	notice := model.MerchantNotice{
		NoticeId:   noticeId,          // 公告ID
		MerchantId: userInfo.Username, // 商户ID
		Title:      req.Title,         // 公告标题
		Content:    req.Content,       // 公告内容
		StartTime:  &startTime,        // 开始时间
		EndTime:    &endTime,          // 结束时间
		IsAll:      req.IsAll,         // 是否发给所有商户:取值说明在模型里
		OperatorId: userInfo.Username, // 发布公告的账号
	}

	var noticePlayers []model.MerchantNoticePlayer
	var playerIds []string
	if req.IsAll == 0 {
		// 发给商户底下的所有会员
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.MerchantId: userInfo.Username,
		}).Pluck(model.PlayerColumns.Username, &playerIds)
	} else {
		// 发送给指定商户会员
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.Username:   req.PlayerIds,
			model.PlayerColumns.MerchantId: userInfo.Username,
		}).Pluck(model.PlayerColumns.Username, &playerIds)
		if len(playerIds) != len(req.PlayerIds) {
			return gerror.Newf(`{#play.id.no.exist}`) // 玩家不存在
		}
	}
	for _, playerId := range playerIds {
		merchantNoticePlayer := model.MerchantNoticePlayer{
			NoticeId:   noticeId,          // 公告ID
			MerchantId: userInfo.Username, // 商户ID
			PlayerId:   playerId,          // 玩家ID
		}
		noticePlayers = append(noticePlayers, merchantNoticePlayer)
	}

	tx := global.DB.Begin()

	if err = tx.Create(&notice).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(noticePlayers) != 0 {
		if err = tx.Create(&noticePlayers).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Edit 编辑商户公告
func (s *MerchantNotice) Edit(req *request.SaveMerchantNoticeRequest, userInfo *request.CustomClaims) error {
	var notice model.MerchantNotice
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantNoticeColumns.Id:         req.Id,
		model.MerchantNoticeColumns.MerchantId: userInfo.Username,
	}).First(&notice, req.Id).Error, gorm.ErrRecordNotFound) {
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
	notice.Title = req.Title      // 公告标题
	notice.Content = req.Content  // 公告内容
	notice.StartTime = &startTime // 开始时间
	notice.EndTime = &endTime     // 结束时间
	notice.IsAll = req.IsAll      // 是否发给所有商户:取值说明在模型里

	var noticePlayers []model.MerchantNoticePlayer
	var playerIds []string
	if req.IsAll == 0 {
		// 发给商户底下的所有会员
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.MerchantId: userInfo.Username,
		}).Pluck(model.PlayerColumns.Username, &playerIds)
	} else {
		// 发送给指定商户会员
		global.DB.Model(&model.Player{}).Where(map[string]interface{}{
			model.PlayerColumns.Username:   req.PlayerIds,
			model.PlayerColumns.MerchantId: userInfo.Username,
		}).Pluck(model.PlayerColumns.Username, &playerIds)
	}

	for _, playerId := range playerIds {
		merchantNoticePlayer := model.MerchantNoticePlayer{
			NoticeId:   notice.NoticeId,   // 公告ID
			MerchantId: userInfo.Username, // 商户ID
			PlayerId:   playerId,          // 玩家ID
		}
		noticePlayers = append(noticePlayers, merchantNoticePlayer)
	}

	tx := global.DB.Begin()

	if err = tx.Save(&notice).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Where(model.MerchantNoticePlayerColumns.NoticeId, notice.NoticeId).Delete(&model.MerchantNoticePlayer{})

	if len(noticePlayers) != 0 {
		if err = tx.Create(&noticePlayers).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// Detail 商户公告详情
func (s *MerchantNotice) Detail(id uint, userInfo *request.CustomClaims) (*response.MerchantNoticeDetailResponse, error) {
	var notice model.MerchantNotice

	db := global.DB.Model(&model.MerchantNotice{}).Where(map[string]interface{}{
		model.MerchantNoticeColumns.Id: id,
	})

	// 主商户可以看到自己发布的公告
	if userInfo.ParentId == global.MerchantRootName {
		db.Where(map[string]interface{}{
			model.MerchantNoticeColumns.MerchantId: userInfo.Username,
		})
	} else {
		// 子商户 可以看到主商户发布的公告
		db.Where(map[string]interface{}{
			model.MerchantNoticeColumns.MerchantId: userInfo.ParentId,
		})
	}

	if err := db.First(&notice).Error; err != nil {
		return nil, err
	}

	res := response.MerchantNoticeDetailResponse{
		Id:         notice.ID,                             // 游戏桌台ID
		NoticeId:   notice.NoticeId,                       // 公告ID
		Title:      notice.Title,                          // 公告标题
		Content:    notice.Content,                        // 公告内容
		StartTime:  notice.StartTime.Format(time.RFC3339), // 开始时间
		EndTime:    notice.EndTime.Format(time.RFC3339),   // 结束时间
		IsAll:      notice.IsAll,                          // 是否发给所有会员:取值说明在模型里
		PlayerIds:  []string{},                            // 指定发给的会员
		OperatorId: notice.OperatorId,                     // 发布公告的账号
	}

	// 指定发给的会员
	global.DB.Model(&model.MerchantNoticePlayer{}).Where(map[string]interface{}{
		model.MerchantNoticePlayerColumns.NoticeId: notice.NoticeId,
	}).Pluck(model.MerchantNoticePlayerColumns.PlayerId, &res.PlayerIds)

	return &res, nil
}

// Del 删除商户公告
func (s *MerchantNotice) Del(id uint, userInfo *request.CustomClaims) (title string, err error) {
	var merchantNotice model.MerchantNotice
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.MerchantNoticeColumns.Id:         id,
		model.MerchantNoticeColumns.MerchantId: userInfo.Username,
	}).First(&merchantNotice).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}
	return merchantNotice.Title, global.DB.Delete(&merchantNotice).Error
}
