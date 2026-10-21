package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
	"time"
)

type Admin struct{}

// LoginAdmin 登录
func (s *Admin) LoginAdmin(req *request.LoginAdminRequest, c *gin.Context) (*response.LoginAdminResponse, error) {
	// 查询用户
	var adminModel model.Admin
	global.DB.Where("username = ?", req.Username).Preload("Role.DefaultMenu").First(&adminModel)
	if adminModel.ID == 0 {
		global.LOG.Info("用户名或密码错误")
		return nil, errors.New(utils.TT("http.login.admin.no.exit"))
	}

	// 是否禁止登录
	if adminModel.Status == global.CannotLogin {
		global.LOG.Info("用户禁止登录")
		return nil, errors.New(utils.TT("http.login.admin.prohibited"))
	}

	// 密码检查
	if ok := utils.BcryptCheck(req.Password, adminModel.Password); !ok {
		global.LOG.Info("密码错误")
		return nil, errors.New(utils.TT("http.login.admin.password.error"))
	}

	token, claims, err := LoginAdminToken(adminModel)
	if err != nil {
		global.LOG.Error("获取token失败：", zap.Error(err))
		return nil, errors.New(utils.TT("http.login.admin.token.get.error") + ":" + err.Error())
	}

	// 更新登录时间
	currentTime := time.Now()
	adminModel.LoginTime = &currentTime
	// 更新登录IP地址
	adminModel.LoginIp = c.ClientIP()
	global.DB.Save(&adminModel)

	// 登录日志
	log := model.AdminLoginLog{
		Username:  adminModel.Username,  // 账号
		LoginTime: adminModel.LoginTime, // 登录时间
		LoginIp:   adminModel.LoginIp,   // 登录IP
	}
	global.DB.Create(&log)

	res := &response.LoginAdminResponse{
		Id:          adminModel.ID,
		RoleId:      adminModel.RoleId,
		Username:    adminModel.Username,
		Nickname:    adminModel.Nickname,
		Avatar:      adminModel.Avatar,
		Status:      adminModel.Status,
		Token:       token,
		DefaultPath: adminModel.Role.DefaultMenu.Path,
		ExpiresAt:   utils.TimestampToDate(claims.RegisteredClaims.ExpiresAt.Unix(), ""),
	}

	// 转json字符串
	resJson, err := json.Marshal(res)
	if err != nil {
		global.LOG.Error("登录信息转json失败：", zap.Error(err))
		return nil, errors.New(utils.TT("http.login.admin.token.json.error") + ":" + err.Error())
	}

	// 保存到redis
	expiresDuration := time.Second * time.Duration(global.CONFIG.Jwt.ExpiresTime)
	err = global.Redis.Set(context.Background(), token, string(resJson), expiresDuration).Err()
	if err != nil {
		global.LOG.Error("用户登录时存入token到redis失败：", zap.Error(err))
		return nil, errors.New(utils.TT("http.login.admin.token.save.error") + ":" + err.Error())
	}

	return res, nil
}

// List 管理员列表
func (s *Admin) List(req *request.AdminListRequest) (*response.ListData, error) {
	var total int64
	var err error
	var list []response.AdminListResponse
	var res response.ListData

	db := global.DB.Model(&model.Admin{}).Preload("Role")
	db.Count(&total)

	if total <= 0 {
		res.List = []response.AdminListResponse{}
		return &res, nil
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	var adminModels []model.Admin
	err = db.Find(&adminModels).Error

	for _, adminModel := range adminModels {
		var loginTimeStr string
		if adminModel.LoginTime != nil {
			loginTimeStr = adminModel.LoginTime.Format("2006-01-02 15:04:05")
		}
		tmp := response.AdminListResponse{
			Id:        adminModel.ID,            // 管理员ID
			Username:  adminModel.Username,      // 账号
			Nickname:  adminModel.Nickname,      // 昵称
			Avatar:    adminModel.Avatar,        // 头像
			Status:    adminModel.Status,        // 是否允许登录:取值说明在模型里
			RoleId:    adminModel.RoleId,        // 角色ID
			RoleName:  adminModel.Role.RoleName, // 角色名称
			LoginTime: loginTimeStr,             // 登录时间
			LoginIp:   adminModel.LoginIp,       // 登录IP
		}
		list = append(list, tmp)
	}

	res.Total = total
	if list != nil {
		res.List = list
	} else {
		res.List = []response.AdminListResponse{}
	}

	return &res, err
}

// Add 创建管理员
func (s *Admin) Add(req *request.SaveAdminRequest) error {
	tx := global.DB.Begin()

	adminModel := model.Admin{
		Username: strings.TrimSpace(req.Username), // 账号
		Nickname: strings.TrimSpace(req.Nickname), // 昵称
		Avatar:   strings.TrimSpace(req.Avatar),   // 头像
		RoleId:   req.RoleId,                      // 角色ID
		Status:   req.Status,                      // 是否允许登录:取值说明在模型里
	}
	adminModel.Password = utils.BcryptHash(strings.TrimSpace(req.Password)) // 密码

	tx.Create(&adminModel)

	// 根据角色ID创建权限
	var apiIds []uint
	global.DB.Model(&model.RoleApi{}).Where("role_id = ?", req.RoleId).Pluck("api_id", &apiIds)

	var apiModel []model.Api
	global.DB.Where("id in ?", apiIds).Find(&apiModel)

	customCasbin := NewCasbinCustom(global.CASBIN)
	sub := "admin_" + gconv.String(adminModel.ID)

	if err := customCasbin.UpdatePolicy(sub, apiModel); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Edit 编辑管理员
func (s *Admin) Edit(req *request.SaveAdminRequest) error {
	var adminModel model.Admin
	if errors.Is(global.DB.Where("id = ?", req.Id).First(&adminModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	adminModel.Nickname = strings.TrimSpace(req.Nickname) // 昵称
	adminModel.Avatar = strings.TrimSpace(req.Avatar)     // 头像
	adminModel.Status = req.Status                        // 是否允许登录:取值说明在模型里
	adminModel.RoleId = req.RoleId                        // 角色ID

	// 如果密码不为空，重新设置密码
	if len(strings.TrimSpace(req.Password)) != 0 {
		adminModel.Password = utils.BcryptHash(strings.TrimSpace(req.Password))
	}

	tx := global.DB.Begin()

	tx.Save(&adminModel)

	// 根据角色ID创建权限
	var apiIds []uint
	global.DB.Model(&model.RoleApi{}).Where("role_id = ?", req.RoleId).Pluck("api_id", &apiIds)

	var apiModel []model.Api
	global.DB.Where("id in ?", apiIds).Find(&apiModel)

	customCasbin := NewCasbinCustom(global.CASBIN)
	sub := "admin_" + gconv.String(adminModel.ID)

	if err := customCasbin.UpdatePolicy(sub, apiModel); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Detail 管理员详情
func (s *Admin) Detail(id uint) (*response.AdminDetailResponse, error) {
	var adminModel model.Admin

	err := global.DB.Where("id = ?", id).Preload("Role").First(&adminModel).Error

	var loginTimeStr string
	if adminModel.LoginTime != nil {
		loginTimeStr = adminModel.LoginTime.Format("2006-01-02 15:04:05")
	}

	res := response.AdminDetailResponse{
		Id:        adminModel.ID,            // 管理员ID
		Username:  adminModel.Username,      // 账号
		Nickname:  adminModel.Nickname,      // 昵称
		Avatar:    adminModel.Avatar,        // 头像
		Status:    adminModel.Status,        // 是否允许登录:取值说明在模型里
		RoleId:    adminModel.RoleId,        // 角色ID
		RoleName:  adminModel.Role.RoleName, // 角色名称
		LoginTime: loginTimeStr,             // 登录时间
		LoginIp:   adminModel.LoginIp,       // 登录IP
	}
	return &res, err
}

// Del 删除管理员
func (s *Admin) Del(id uint) (username string, err error) {
	var adminModel model.Admin
	if errors.Is(global.DB.Where("id = ?", id).First(&adminModel).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	if err = global.DB.Select("Roles").Delete(&adminModel).Error; err != nil {
		global.LOG.Error("删除失败：", zap.Error(err))
		return "", err
	}

	customCasbin := NewCasbinCustom(global.CASBIN)
	sub := "admin_" + gconv.String(adminModel.ID)
	customCasbin.RemovePolicy(0, sub)

	return adminModel.Username, nil
}

// EditCurrentPassword 编辑当前登录的管理员密码
func (s *Admin) EditCurrentPassword(req *request.EditCurrentPasswordRequest, adminId uint) error {
	var adminModel model.Admin
	if errors.Is(global.DB.Where("id = ?", adminId).First(&adminModel).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	// 密码检查
	if ok := utils.BcryptCheck(req.OldPassword, adminModel.Password); !ok {
		return errors.New(utils.TT("http.login.admin.password.error"))
	}

	adminModel.Password = utils.BcryptHash(strings.TrimSpace(req.NewPassword))
	return global.DB.Save(&adminModel).Error
}

// ResetPassword 重置管理员密码
func (s *Admin) ResetPassword(id uint) (*response.ResetAdminPasswordResponse, error) {
	var adminModel model.Admin
	if errors.Is(global.DB.Where("id = ?", id).First(&adminModel).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	resetPassword := global.CONFIG.Server.ResetPassword
	if len(resetPassword) == 0 {
		resetPassword = "123456"
	}

	adminModel.Password = utils.BcryptHash(strings.TrimSpace(resetPassword))

	if err := global.DB.Save(&adminModel).Error; err != nil {
		return nil, err
	}
	res := response.ResetAdminPasswordResponse{
		Id:       adminModel.ID,       // 管理员ID
		Username: adminModel.Username, // 账号
		Password: resetPassword,       // 密码
	}
	return &res, nil
}

// GetAdminInfo 获取管理员信息
func (s *Admin) GetAdminInfo(id uint) (*response.GetAdminInfoResponse, error) {
	// 查询用户
	var adminModel model.Admin
	global.DB.Where("id = ?", id).Preload("Role.DefaultMenu").First(&adminModel)
	if err := global.DB.Where("id = ?", id).Preload("Role").First(&adminModel).Error; err != nil {
		return nil, err
	}

	var loginTimeStr string
	if adminModel.LoginTime != nil {
		loginTimeStr = adminModel.LoginTime.Format("2006-01-02 15:04:05")
	}

	res := &response.GetAdminInfoResponse{
		Id:        adminModel.ID,            // 管理员ID
		Username:  adminModel.Username,      // 账号
		Nickname:  adminModel.Nickname,      // 昵称
		Avatar:    adminModel.Avatar,        // 头像
		Status:    adminModel.Status,        // 是否允许登录:取值说明在模型里
		RoleId:    adminModel.RoleId,        // 角色ID
		RoleName:  adminModel.Role.RoleName, // 角色名称
		LoginTime: loginTimeStr,             // 登录时间
		LoginIp:   adminModel.LoginIp,       // 登录IP
	}

	return res, nil
}
