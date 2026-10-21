package service

import (
	"errors"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type ControlAccount struct{}

// GetLevelBox 获取控端账号等级选择框
func (s *ControlAccount) GetLevelBox() []response.ControlAccountLevelBox {
	var controlAccountLevelList []response.ControlAccountLevelBox
	for _, v := range model.ControlAccountLevelMap {
		tmp := response.ControlAccountLevelBox{
			Name:  v.Name,
			Value: uint(v.Value),
		}
		controlAccountLevelList = append(controlAccountLevelList, tmp)
	}
	return controlAccountLevelList
}

// List 控端账号列表
func (s *ControlAccount) List(req *request.ControlAccountListRequest) (*response.ListData, error) {
	var total int64
	var list []response.ControlAccountListResponse
	var res response.ListData
	var err error

	db := global.DB.Model(&model.ControlAccount{})

	db.Count(&total)

	if total <= 0 {
		res.List = []response.ControlAccountListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var controlAccounts []model.ControlAccount
	if err = db.Find(&controlAccounts).Error; err != nil {
		return &res, err
	}

	for _, controlAccount := range controlAccounts {
		var loginTimeStr string
		if controlAccount.LoginTime != nil {
			loginTimeStr = controlAccount.LoginTime.Format("2006-01-02 15:04:05")
		}
		tmp := response.ControlAccountListResponse{
			Id:        controlAccount.ID,          // 控端账号ID
			Username:  controlAccount.Username,    // 账号
			Nickname:  controlAccount.Nickname,    // 昵称
			Level:     uint(controlAccount.Level), // 等级
			Status:    controlAccount.Status,      // 是否允许登录:取值说明在模型里
			LoginTime: loginTimeStr,               // 上次登录时间
			LoginIp:   controlAccount.LoginIp,     // 登录IP
		}
		list = append(list, tmp)
	}

	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.ControlAccountListResponse{}
	}

	return &res, err
}

// Add 创建控端账号
func (s *ControlAccount) Add(req *request.SaveControlAccountRequest) error {
	controlAccountModel := model.ControlAccount{
		Username: strings.ToLower(strings.TrimSpace(req.Username)), // 账号
		Nickname: strings.TrimSpace(req.Nickname),                  // 昵称
		Level:    model.ControlAccountLevel(req.Level),             // 等级
		Status:   req.Status,                                       // 是否允许登录:取值说明在模型里
	}
	controlAccountModel.Password = utils.BcryptHash(strings.TrimSpace(req.Password)) // 密码
	return global.DB.Create(&controlAccountModel).Error
}

// Edit 编辑控端账号
func (s *ControlAccount) Edit(req *request.SaveControlAccountRequest) (err error) {
	var controlAccountModel model.ControlAccount
	err = global.DB.First(&controlAccountModel, req.Id).Error
	if err != nil {
		return err
	}

	controlAccountModel.Nickname = strings.TrimSpace(req.Nickname)   // 昵称
	controlAccountModel.Level = model.ControlAccountLevel(req.Level) // 等级
	controlAccountModel.Status = req.Status                          // 是否允许登录:取值说明在模型里

	// 如果密码不为空，重新设置密码
	if len(strings.TrimSpace(req.Password)) != 0 {
		controlAccountModel.Password = utils.BcryptHash(strings.TrimSpace(req.Password))
	}

	return global.DB.Save(&controlAccountModel).Error
}

// Detail 控端账号详情
func (s *ControlAccount) Detail(id uint) (*response.ControlAccountDetailResponse, error) {
	var controlAccount model.ControlAccount

	err := global.DB.Model(&model.ControlAccount{}).Where("id = ?", id).First(&controlAccount).Error

	var loginTimeStr string
	if controlAccount.LoginTime != nil {
		loginTimeStr = controlAccount.LoginTime.Format("2006-01-02 15:04:05")
	}

	res := response.ControlAccountDetailResponse{
		Id:        controlAccount.ID,          // 控端账号ID
		Username:  controlAccount.Username,    // 账号
		Nickname:  controlAccount.Nickname,    // 昵称
		Level:     uint(controlAccount.Level), // 等级
		Status:    controlAccount.Status,      // 是否允许登录:取值说明在模型里
		LoginTime: loginTimeStr,               // 上次登录时间
		LoginIp:   controlAccount.LoginIp,     // 登录IP
	}
	return &res, err
}

// Del 删除控端账号
func (s *ControlAccount) Del(id uint) (username string, err error) {
	var controlAccountModel model.ControlAccount
	if errors.Is(global.DB.Where("id = ?", id).First(&controlAccountModel).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	return controlAccountModel.Username, global.DB.Delete(&controlAccountModel).Error
}

// ResetPassword 重置控端账号密码
func (s *ControlAccount) ResetPassword(id uint) (*response.ResetControlAccountPasswordResponse, error) {
	var ControlAccountModel model.ControlAccount
	if errors.Is(global.DB.Where("id = ?", id).First(&ControlAccountModel).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	resetPassword := global.CONFIG.Server.ResetPassword
	if len(resetPassword) == 0 {
		resetPassword = "123456"
	}

	ControlAccountModel.Password = utils.BcryptHash(strings.TrimSpace(resetPassword))

	if err := global.DB.Save(&ControlAccountModel).Error; err != nil {
		return nil, err
	}
	res := response.ResetControlAccountPasswordResponse{
		Id:       ControlAccountModel.ID, // 控端账号ID
		Password: resetPassword,          // 密码
	}
	return &res, nil
}

// GetBox 获取控端账号下拉框数据
func (s *ControlAccount) GetBox() (list []response.ControlAccountBoxResponse, err error) {
	var controlAccounts []model.ControlAccount
	if err = global.DB.Find(&controlAccounts).Error; err != nil {
		return
	}

	// 赋值
	for _, controlAccount := range controlAccounts {
		tmp := response.ControlAccountBoxResponse{
			Id:       controlAccount.ID,       // ID
			Username: controlAccount.Username, // 账号
		}

		list = append(list, tmp)
	}
	return list, err
}
