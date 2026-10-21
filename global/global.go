package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/duanhf2012/origin/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"orange-crm/config"
)

// global 定义全局变量
var (
	CONFIG     config.Common // 系统配置
	VP         *viper.Viper  // viper
	LOG        *zap.Logger   // zap
	DB         *gorm.DB
	Redis      *redis.Client
	Router     *gin.Engine
	I18N       config.I18n
	CASBIN     *casbin.CachedEnforcer
	HttpServer *http.Server
	CrmModule  service.Module
	GClient    *gclient.Client // HTTP 客户端
)

const (
	Charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 基本字符集
	CanLogin     = 1                                                                // 可以登录
	CannotLogin  = 0                                                                // 禁止登录
	CanEdit      = 1                                                                // 可编辑
	CannotEdit   = 0                                                                // 不可编辑
	CanDelete    = 1                                                                // 可删除
	CannotDelete = 0                                                                // 不可删除
	CanUse       = 1                                                                // 可用
	CannotUse    = 0                                                                // 禁用

	UploadAdminAvatar           = "upload_admin_avatar"
	UploadDealerAvatar          = "upload_dealer_avatar"
	UploadMerchantAccountAvatar = "upload_merchant_account_avatar"
	UploadPlayerAvatar          = "upload_player_avatar"
	UploadControlAccountAvatar  = "upload_control_account_avatar"

	AdminIdentity    = 1 // 管理员身份
	MerchantIdentity = 2 // 商户身份

	CanAssignMerchant    = 1 // 不可以分配给商户
	CannotAssignMerchant = 0 // 可分配给商户

	// 时间格式
	TIME_STANDARD    = "2006-01-02 15:04:05"              // 标准格式 (秒)
	TIME_YMD         = "2006-01-02"                       // 年-月-日
	TIME_YM          = "2006-01"                          // 年-月
	TIME_YMD_HM      = "2006-01-02 15:04"                 // 年-月-日 时:分
	TIME_STANDARD_MS = "2006-01-02 15:04:05.000"          // ✅ 标准格式 (毫秒)
	TIME_RFC3339_MS  = "2006-01-02T15:04:05.000Z07:00"    // RFC3339 (毫秒)
	TIME_STANDARD_US = "2006-01-02 15:04:05.000000"       // ✅ 标准格式 (微秒)
	TIME_RFC3339_US  = "2006-01-02T15:04:05.000000Z07:00" // ✅ RFC3339 (微秒)

	AdminLogName    = "admin_log"
	MerchantLogName = "merchant_log"

	MerchantRootName = "root" // 商户顶级名称
	PlayerRootName   = "root" // 会员顶级名称
)

var (
	UploadImgEventType = []string{UploadAdminAvatar, UploadDealerAvatar, UploadMerchantAccountAvatar, UploadPlayerAvatar, UploadControlAccountAvatar}
)

const (
	ToNodeId       = "admin"
	ToServiceName  = "AdminService"
	NotifyMerchant = "RPC_NotifyMerchant" // 商户通知
)

const (
	CrmMsgId_Define            = 1000                // 消息ID起始用
	CrmMsgId_CreateMerchant    = CrmMsgId_Define + 1 // 创建商户消息
	CrmMsgId_UpdateMerchant    = CrmMsgId_Define + 2 // 更新商户的消息
	CrmMsgId_CreatePlayer      = CrmMsgId_Define + 3 // 创建玩家的消息
	CrmMsgId_UpdatePlayer      = CrmMsgId_Define + 4 // 更新玩家的消息，包括限红模板ID
	CrmMsgId_OperatePlanId     = CrmMsgId_Define + 5 // 操作限红模板
	CrmMsgId_OperateGameConfig = CrmMsgId_Define + 6 // 修改游戏配置
	CrmMsgId_KickOutPlayer     = CrmMsgId_Define + 7 // 把玩家踢出去
)

const (
	ErrParam   = 101 // 参数错误
	ErrIdError = 102 // 消息ID错误
	ErrCluster = 103 // 服务器内部出错
)
