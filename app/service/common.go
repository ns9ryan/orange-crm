package service

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"mime/multipart"
	"orange-crm/app/model"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"path/filepath"
	"strconv"
	"time"
)

type Common struct{}

// GetLangBox 语言下拉框
func (s *Common) GetLangBox() (list []response.LangResponse, err error) {
	// 先从redis读取数据
	data, err := g.Redis().Do(gctx.New(), "GET", "system:lang")
	if err != nil {
		return nil, err
	}
	if err = data.Struct(&list); err != nil {
		return nil, err
	}
	// 如果存在直接返回
	if len(list) != 0 {
		return list, err
	}

	// 从数据库种读取数据
	if err = global.DB.Model(&model.Lang{}).Find(&list).Error; err != nil {
		return nil, err
	}
	// 存入redis
	_, err = g.Redis().Do(gctx.New(), "set", "system:lang", list)

	return list, err
}

// GetSystemConfig 获取系统配置数据
func (s *Common) GetSystemConfig() (data *response.SystemConfigResponse, err error) {
	data = &response.SystemConfigResponse{
		LangList:           []response.LangResponse{},           // 系统语言
		SystemCurrencyList: []response.SystemCurrencyResponse{}, // 系统币种
		BetAreaList:        []response.BetAreaResponse{},        // 下注区域
		SystemChipList:     []response.SystemChipResponse{},     // 系统筹码
		GameCategoryList:   []response.GameCategoryResponse{},   // 游戏分类
	}

	if data.LangList, err = s.GetLangBox(); err != nil {
		return nil, err
	}

	if data.SystemCurrencyList, err = Service.SystemCurrency.GetSystemCurrencyBox(); err != nil {
		return nil, err
	}

	if data.BetAreaList, err = Service.BetArea.GetBetAreaBox(); err != nil {
		return nil, err
	}

	if data.SystemChipList, err = Service.SystemChip.GetSystemChipBox(); err != nil {
		return nil, err
	}

	if data.GameCategoryList, err = Service.GameCategory.GetGameCategoryBox(); err != nil {
		return nil, err
	}

	return data, nil
}

// UploadAvatar 通用上传头像
func (s *Common) UploadAvatar(file *multipart.FileHeader, subDir string, ctx *gin.Context) (string, error) {
	// 文件大小
	if file.Size > global.CONFIG.Upload.Avatar.Size {
		return "", gerror.Newf(`{#upload.avatar.too.large}`)
	}

	// 文件后缀
	fileExt := filepath.Ext(file.Filename)
	if exist := slice.Contain(global.CONFIG.Upload.Avatar.Ext, fileExt); exist == false {
		return "", gerror.Newf(`{#upload.avatar.format.incorrect}`)
	}

	// 根据当前时间戳生成一个新的文件名
	fileName := strconv.FormatInt(time.Now().Unix(), 10)

	//新的文件名
	fileFullName := fileName + fileExt

	fileDir := fmt.Sprintf("%s%s", global.CONFIG.Upload.Avatar.SavePath, subDir)

	if err := utils.CreateDirectory(fileDir); err != nil {
		return "", err
	}

	if err := ctx.SaveUploadedFile(file, fileDir+fileFullName); err != nil {
		return "", err
	}

	return fileDir + fileFullName, nil
}

// UploadImg 通用上传图片
func (s *Common) UploadImg(file *multipart.FileHeader, subDir string, ctx *gin.Context) (string, error) {
	// 文件大小
	if file.Size > global.CONFIG.Upload.Avatar.Size {
		return "", gerror.Newf(`{#upload.img.too.large}`)
	}

	// 文件后缀
	fileExt := filepath.Ext(file.Filename)
	if exist := slice.Contain(global.CONFIG.Upload.Avatar.Ext, fileExt); exist == false {
		return "", gerror.Newf(`{#upload.img.format.incorrect}`)
	}

	// 根据当前时间戳生成一个新的文件名
	fileName := strconv.FormatInt(time.Now().Unix(), 10)

	//新的文件名
	fileFullName := fileName + fileExt

	fileDir := fmt.Sprintf("%s%s", global.CONFIG.Upload.Img.SavePath, subDir)

	if err := utils.CreateDirectory(fileDir); err != nil {
		return "", err
	}

	if err := ctx.SaveUploadedFile(file, fileDir+fileFullName); err != nil {
		return "", err
	}

	return fileDir + fileFullName, nil
}
