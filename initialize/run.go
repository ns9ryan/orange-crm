package initialize

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"go.uber.org/zap"
	"net/http"
	"orange-crm/global"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func Run() {
	initGlobal()

	address := fmt.Sprintf(":%d", global.CONFIG.Server.Addr)

	// 初始化Server
	s := initServer(address, global.Router)

	go func() {
		global.LOG.Info("访问地址", zap.String("url", "http://localhost"+address))
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Error("Gin 启动失败", zap.Error(err))
		}
	}()

	//global.LOG.Error(s.ListenAndServe().Error())

	// 注册关闭函数
	global.HttpServer = s

	// 如果是独立运行，则监听退出信号并优雅关闭
	if _, ok := os.LookupEnv("RUN_AS_STANDALONE"); ok {
		//global.LOG.Error("RUN_AS_STANDALONE：", zap.Any("data", os.Getenv("RUN_AS_STANDALONE")))
		waitForSignalAndShutdown()
	}
}

// 项目初始化
func initGlobal() {
	// 初始化Viper
	global.VP = Viper()
	// 初始化日志
	global.LOG = Zap()
	zap.ReplaceGlobals(global.LOG) // 替换全局记录器 后续 可以使用 zap.L()
	// 初始化数据库
	global.DB = Gorm()
	// 初始化redis
	global.Redis = Redis()
	// 初始化语言包
	global.I18N = I18n()
	// 初始化casbin
	global.CASBIN = Casbin()
	// 数据迁移
	RegisterTables()
	// 初始化gin
	global.Router = Gin()
	// HTTP 客户端
	global.GClient = g.Client()

	// 日志打印示例
	//global.LOG.Debug("this is debug")
	//global.LOG.Info("this is info")
	//global.LOG.Warn("this is warn")
	//global.LOG.Error("this is error")
	//global.LOG.DPanic("this is dPanic")
	//global.LOG.Panic("this is panic")
	//global.LOG.Fatal("this is fatal")
}

// 初始化Server
func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,          // 监听的 TCP 地址
		Handler:        router,           // Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求s
		ReadTimeout:    10 * time.Minute, // ReadTimeout：允许读取的最大时间
		WriteTimeout:   10 * time.Minute, // WriteTimeout：允许写入的最大时间
		MaxHeaderBytes: 1 << 20,          // MaxHeaderBytes：请求头的最大字节数
	}
}

func waitForSignalAndShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	global.LOG.Info("收到退出信号，准备优雅关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := global.HttpServer.Shutdown(ctx); err != nil {
		global.LOG.Error("服务关闭异常", zap.Error(err))
	} else {
		global.LOG.Info("服务已优雅退出")
	}
}
