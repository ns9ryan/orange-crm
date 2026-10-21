package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"net/http"
)

// 常用返回消息
const (
	MsgApiSuccess      = `{#success}`
	MsgApiFail         = `{#fail}`
	MsgLoginInvalid    = `{#login_invalid}`
	MsgRequestErr      = `{#request_err}`
	MsgNoPermission    = `{#no_permission}`
	MsgEditableErr     = `{#editable_err}`
	MsgDeletableErr    = `{#deletable_err}`
	MsgValidateErr     = `{#validate_err}`
	MsgValidateScanErr = `{#validate_err_scan}`
	MsgIdentityInvalid = `{#identity_invalid}`
)

// ResponseData 统一返回数据结构
type responseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// response 统一返回结构体
type response struct{}

// Response 对外统一提供返回
var Response = &response{}

// 统一返回json
func (r *response) jsonResponse(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, responseData{
		Code: code,
		Msg:  TT(msg), // 调用统一的翻译函数
		Data: data,
	})
}

// Success 请求成功（状态码 200）
func (r *response) Success(c *gin.Context, msg string, data any) {
	if msg == "" {
		msg = MsgApiSuccess
	}
	r.jsonResponse(http.StatusOK, msg, data, c)
}

// SuccessSimple 请求成功，仅返回默认消息
func (r *response) SuccessSimple(c *gin.Context) {
	r.Success(c, "", nil)
}

// SuccessWithData 请求成功，附带数据
func (r *response) SuccessWithData(c *gin.Context, data any) {
	r.Success(c, "", data)
}

// Fail 请求失败（状态码可选，默认 400）
func (r *response) Fail(c *gin.Context, code int, msg string, data any) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	if msg == "" {
		msg = MsgApiFail
	}
	r.jsonResponse(code, msg, data, c)
}

// FailBadRequest 非法请求
func (r *response) FailBadRequest(c *gin.Context, err error) {
	r.Fail(c, http.StatusBadRequest, gerror.Wrap(err, MsgRequestErr).Error(), nil)
}

// FailValidate 参数错误
func (r *response) FailValidate(c *gin.Context, err error) {
	r.Fail(c, http.StatusBadRequest, gerror.Wrap(err, MsgValidateErr).Error(), nil)
}

// FailSystem 系统错误
func (r *response) FailSystem(c *gin.Context, err error) {
	r.Fail(c, http.StatusBadRequest, gerror.Wrap(err, MsgApiFail).Error(), nil)
}

// FailUnauthorized 未登录
func (r *response) FailUnauthorized(c *gin.Context) {
	r.Fail(c, http.StatusNoContent, MsgLoginInvalid, nil)
}

// FailForbidden 无权限
func (r *response) FailForbidden(c *gin.Context) {
	r.Fail(c, http.StatusBadRequest, MsgNoPermission, nil)
}

// FailIdentity 身份不对
func (r *response) FailIdentity(c *gin.Context) {
	r.Fail(c, http.StatusNoContent, MsgIdentityInvalid, nil)
}
