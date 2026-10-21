package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"orange-crm/global"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path    // /hello/xxxx
		raw := c.Request.URL.RawQuery // ?后面的值

		// 处理请求
		c.Next()

		// 执行的时间
		cost := time.Since(start)
		// 如果花费的时间大于1分钟，取最小单位秒
		if cost > time.Minute {
			cost = cost.Truncate(time.Second)
		}

		if raw != "" {
			path = path + "?" + raw
		}

		sprintf := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %#v%s\n",
			start.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			cost,
			c.ClientIP(),
			c.Request.Method,
			path,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)

		filename := fmt.Sprintf("./%s/%s/server_gin.log", global.CONFIG.Zap.Dir, time.Now().Format("20060102"))

		lumberJackLogger := &lumberjack.Logger{
			Filename: filename, // 日志文件的位置
		}
		lumberJackLogger.Write([]byte(sprintf))
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool

				// 目标服务不可达（连接拒绝）
				// 请求超时（模拟连接超时）
				// 主机不存在（DNS 解析失败）
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}

					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.LOG.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					global.LOG.Panic("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.LOG.Panic("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// GormLogger gorm日志
func GormLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := fmt.Sprintf("./%s/%s/server_gorm.log", global.CONFIG.Zap.Dir, time.Now().Format("20060102"))
		lumberJackLogger := &lumberjack.Logger{
			Filename: filename, // 日志文件的位置
		}
		logLevel := logger.LogLevel(global.CONFIG.Server.GormLogLevel)
		global.DB.Config.Logger = logger.New(log.New(lumberJackLogger, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      false,
		})
	}
}
