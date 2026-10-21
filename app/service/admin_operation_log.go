package service

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/model"
	"orange-crm/global"
	"orange-crm/utils"
)

type AdminOperationLog struct{}

// SaveAdminLog 保存管理员操作日志（异步写入）
func (s *AdminOperationLog) SaveAdminLog(c *gin.Context, partial model.AdminOperationLog) {
	userInfo, err := GetAdminInfo(c)
	if err != nil {
		return
	}

	fullLog := model.AdminOperationLog{
		Username:                userInfo.Username,                                                // 账号
		Module:                  model.AdminOperationLogEventMap[partial.AdminOperationEventType], // 模块名
		Action:                  partial.Action,                                                   // 操作动作
		Content:                 partial.Content,                                                  // 操作内容描述
		IsSuccess:               partial.IsSuccess,                                                // 是否成功:取值说明在模型里
		ErrorMessage:            partial.ErrorMessage,                                             // 错误信息（失败时记录）
		AdminOperationEventType: partial.AdminOperationEventType,                                  // 事件类型

		Latency:       utils.GetCtxDurationMillis(c, "latency"),
		RequestUrl:    utils.GetCtxString(c, "request_url"),                                // 请求地址
		RequestMethod: utils.GetCtxString(c, "request_method"),                             // 请求方法
		RequestParams: utils.TruncateString(utils.GetCtxString(c, "request_params"), 4096), // 请求参数
		ResponseData:  utils.TruncateString(utils.GetCtxString(c, "response_body"), 5120),  // 响应数据
		Ip:            utils.GetCtxString(c, "ip"),                                         // 请求 IP
		UserAgent:     utils.GetCtxString(c, "user_agent"),                                 // 请求 User-Agent
		StatusCode:    uint(c.GetInt("status_code")),                                       // HTTP 状态码
	}

	// 异步写入日志，避免阻塞接口请求
	go func() {
		_ = global.DB.Select("*").Create(&fullLog).Error
	}()
}
