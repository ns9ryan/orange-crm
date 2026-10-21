package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"orange-crm/app/model"
	"orange-crm/app/service"
	"orange-crm/utils"
	"time"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 将响应内容写入自定义缓存和原始响应体
func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 将响应内容写入缓存
	return w.ResponseWriter.Write(b) // 同时写入原始响应
}

// AdminLogMiddleware 自动记录管理员操作日志
func AdminLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte

		var err error
		body, err = io.ReadAll(c.Request.Body)
		if err != nil {
			utils.Response.FailSystem(c, errors.New(fmt.Sprintf("read body from request error:%s", err.Error())))
			c.Abort()
			return
		} else {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		// 设置请求相关上下文信息
		c.Set("request_url", c.Request.RequestURI)
		c.Set("request_method", c.Request.Method)
		c.Set("request_params", string(body))
		c.Set("ip", c.ClientIP())
		c.Set("user_agent", c.GetHeader("User-Agent"))
		c.Set("request_time", time.Now())

		bw := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bw
		now := time.Now()

		// 处理请求
		c.Next()

		// 设置响应内容上下文信息
		c.Set("response_body", bw.body.String())
		c.Set("status_code", c.Writer.Status())
		latency := time.Since(now)
		c.Set("latency", latency)

		if raw, exists := c.Get("admin_log"); exists {
			if logData, ok := raw.(model.AdminOperationLog); ok {
				service.Service.AdminOperationLog.SaveAdminLog(c, logData)
			}
		}
	}
}

// MerchantLogMiddleware 自动记录商户操作日志
func MerchantLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte

		var err error
		body, err = io.ReadAll(c.Request.Body)
		if err != nil {
			utils.Response.FailSystem(c, errors.New(fmt.Sprintf("read body from request error:%s", err.Error())))
			c.Abort()
			return
		} else {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		// 设置请求相关上下文信息
		c.Set("request_url", c.Request.RequestURI)
		c.Set("request_method", c.Request.Method)
		c.Set("request_params", string(body))
		c.Set("ip", c.ClientIP())
		c.Set("user_agent", c.GetHeader("User-Agent"))
		c.Set("request_time", time.Now())

		bw := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bw
		now := time.Now()

		// 处理请求
		c.Next()

		// 设置响应内容上下文信息
		c.Set("response_body", bw.body.String())
		c.Set("status_code", c.Writer.Status())
		latency := time.Since(now)
		c.Set("latency", latency)

		if raw, exists := c.Get("merchant_log"); exists {
			if logData, ok := raw.(model.MerchantOperationLog); ok {
				service.Service.MerchantOperationLog.SaveMerchantLog(c, logData)
			}
		}
	}
}
