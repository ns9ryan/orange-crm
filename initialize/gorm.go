package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"orange-crm/app/model"
	"orange-crm/global"
	"os"
	"time"
)

// Gorm 初始化gorm
func Gorm() *gorm.DB {
	config := global.CONFIG.Mysql

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      config.Dsn(), // DSN data source name
		DefaultStringSize:        256,          // string 类型字段的默认长度
		DisableDatetimePrecision: true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		DisableForeignKeyConstraintWhenMigrating: true, // 设置为 true 后，GORM 不会生成外键约束，只处理逻辑关联。
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.CONFIG.Mysql.Prefix, // 表名前缀，`User` 的表名应该是 `prefix_users`
			SingularTable: true,                       // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `prefix_user`
		},
		//Logger: logger.Default.LogMode(logger.Silent),
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值，如果 SQL 执行时间 超过 1 秒，就会被标记为“慢 SQL”。
				LogLevel:      logger.Info,            // 日志级别
				Colorful:      true,                   // 关闭日志颜色显示
				//IgnoreRecordNotFoundError: true,           // 当查询结果为空时，不记录 ErrRecordNotFound 错误
				//ParameterizedQueries:      true,           // 不在日志中显示 SQL 参数，有助于保护敏感数据
			},
		),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// RegisterTables 注册表
func RegisterTables() {
	models := []struct {
		name string
		fn   func() error
	}{
		// 系统数据
		{"Lang", model.Model.Lang.RegisterTable},                     // 语言表
		{"SystemChip", model.Model.SystemChip.RegisterTable},         // 系统筹码表
		{"SystemCurrency", model.Model.SystemCurrency.RegisterTable}, // 系统币种表
		{"BetArea", model.Model.BetArea.RegisterTable},               // 下注区域表
		{"GameCategory", model.Model.GameCategory.RegisterTable},     // 游戏分类表
		{"SystemTip", model.Model.SystemTip.RegisterTable},           // 系统打赏配置表

		// 平台管理
		{"Currency", model.Model.Currency.RegisterTable},                                       // 币种表
		{"CurrencyExchangeRate", model.Model.CurrencyExchangeRate.RegisterTable},               // 币种汇率表
		{"CurrencyExchangeRateHistory", model.Model.CurrencyExchangeRateHistory.RegisterTable}, // 币种汇率历史表
		{"Dealer", model.Model.Dealer.RegisterTable},                                           // 荷官表
		{"ControlAccount", model.Model.ControlAccount.RegisterTable},                           // 控端账号表
		{"SystemNotice", model.Model.SystemNotice.RegisterTable},                               // 系统公告表
		{"SystemNoticeMerchant", model.Model.SystemNoticeMerchant.RegisterTable},               // 系统公告与商户关联表
		{"MerchantNotice", model.Model.MerchantNotice.RegisterTable},                           // 商户公告表
		{"MerchantNoticePlayer", model.Model.MerchantNoticePlayer.RegisterTable},               // 商户公告与玩家关联表
		{"ServiceNode", model.Model.ServiceNode.RegisterTable},                                 // 服务节点池表

		// 日志管理
		{"AdminLoginLog", model.Model.AdminLoginLog.RegisterTable},               // 管理员登录日志表
		{"AdminLoginLog", model.Model.AdminOperationLog.RegisterTable},           // 管理员操作记录表
		{"MerchantOperationLog", model.Model.MerchantOperationLog.RegisterTable}, // 商户操作记录表
		{"LoginLog", model.Model.LoginLog.RegisterTable},                         // 商户+会员登录日志表
		{"ControlLoginLog", model.Model.ControlLoginLog.RegisterTable},           // 控端登录日志表

		// 游戏管理
		{"GameTable", model.Model.GameTable.RegisterTable},                       // 游戏桌台表
		{"BetRecord", model.Model.BetRecord.RegisterTable},                       // 投注明细表
		{"AnswerBacx", model.Model.AnswerBacx.RegisterTable},                     // 开答表-百家乐
		{"AnswerDtx", model.Model.AnswerDtx.RegisterTable},                       // 开答表-龙虎
		{"AnswerFtan", model.Model.AnswerFtan.RegisterTable},                     // 开答表-番摊
		{"AnswerRot", model.Model.AnswerRot.RegisterTable},                       // 开答表-轮盘
		{"AnswerSicbo", model.Model.AnswerSicbo.RegisterTable},                   // 开答表-骰宝
		{"AnswerSdie", model.Model.AnswerSdie.RegisterTable},                     // 开答表-色碟
		{"TemplatePlan", model.Model.TemplatePlan.RegisterTable},                 // 限红方案表
		{"TemplatePlanArea", model.Model.TemplatePlanArea.RegisterTable},         // 限红方案投注区域详情表
		{"TemplatePlanForm", model.Model.TemplatePlanForm.RegisterTable},         // 限红模板表
		{"TemplatePlanFormArea", model.Model.TemplatePlanFormArea.RegisterTable}, // 限红模板投注区域表

		// 商户管理
		{"Merchant", model.Model.Merchant.RegisterTable},                                         // 商户表
		{"MerchantCurrency", model.Model.MerchantCurrency.RegisterTable},                         // 商户币种表
		{"MerchantTemplateBetlimit", model.Model.MerchantTemplateBetlimit.RegisterTable},         // 商户限红模板表
		{"MerchantTemplateBetlimitArea", model.Model.MerchantTemplateBetlimitArea.RegisterTable}, // 商户限红模板投注区域表
		{"MerchantDayaccount", model.Model.MerchantDayaccount.RegisterTable},                     // 商户输赢表

		// 玩家管理
		{"Player", model.Model.Player.RegisterTable},                       // 玩家表
		{"PlayerWallet", model.Model.PlayerWallet.RegisterTable},           // 玩家钱包表
		{"PlayerTransaction", model.Model.PlayerTransaction.RegisterTable}, // 玩家交易流水表
		{"PlayerDayaccount", model.Model.PlayerDayaccount.RegisterTable},   // 玩家输赢表

		// 权限管理
		{"Admin", model.Model.Admin.RegisterTable},           // 管理员表
		{"Api", model.Model.Api.RegisterTable},               // Api表
		{"Menu", model.Model.Menu.RegisterTable},             // 菜单表
		{"Role", model.Model.Role.RegisterTable},             // 角色表
		{"RoleApi", model.Model.RoleApi.RegisterTable},       // 角色api表
		{"RoleMenu", model.Model.RoleMenu.RegisterTable},     // 角色菜单表
		{"CasbinRule", model.Model.CasbinRule.RegisterTable}, // casbin表

		{"GameNotifyLog", model.Model.GameNotifyLog.RegisterTable}, // 游戏通知接口请求日志表
		// =======================================================================================================================todo
		{"BetAccountRecord", model.Model.BetAccountRecord.RegisterTable}, // 投注帐目表
		{"TipRecord", model.Model.TipRecord.RegisterTable},               // 打赏明细表
	}

	for _, m := range models {
		if err := m.fn(); err != nil {
			global.LOG.Error(fmt.Sprintf("注册 %s 失败", m.name), zap.Error(err))
			panic(fmt.Errorf("注册 %s 失败: %v", m.name, err))
		}
	}
}
