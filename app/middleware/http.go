package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

// HttpLogger 请求日志
func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params []byte
		var body []byte
		// 请求方法
		method := c.Request.Method
		if method == http.MethodGet {
			// 读取"?"后面的部分
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)

			// 分割"&"
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			params, _ = json.Marshal(&m)
		} else {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Error(err))
			} else {
				// 请求体只能读取一次所以要重写赋值
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		utils.FF(string(body))
		utils.FF(params)
		//start := time.Now()
		//path := c.Request.URL.Path    // /hello/xxxx
		//raw := c.Request.URL.RawQuery // ?后面的值

		// 处理请求
		c.Next()

		//// 执行的时间
		//cost := time.Since(start)
		//// 如果花费的时间大于1分钟，取最小单位秒
		//if cost > time.Minute {
		//	cost = cost.Truncate(time.Second)
		//}
		//
		//if raw != "" {
		//	path = path + "?" + raw
		//}
		//
		//sprintf := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %#v%s\n",
		//	start.Format("2006/01/02 - 15:04:05"),
		//	c.Writer.Status(),
		//	cost,
		//	c.ClientIP(),
		//	c.Request.Method,
		//	path,
		//	c.Errors.ByType(gin.ErrorTypePrivate).String(),
		//)
		//
		//filename := fmt.Sprintf("./%s/%s/server_gin.log", global.CONFIG.Zap.Dir, time.Now().Format("20060102"))
		//
		//lumberJackLogger := &lumberjack.Logger{
		//	Filename: filename, // 日志文件的位置
		//}
		//lumberJackLogger.Write([]byte(sprintf))
	}
}
