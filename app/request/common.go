package request

import "github.com/golang-jwt/jwt/v5"

// CustomClaims 自定义jwt的Claims
type CustomClaims struct {
	ID                   uint   `json:"id"`
	Username             string `json:"username"`  // 账号
	Nickname             string `json:"nickname"`  // 昵称
	RoleId               uint   `json:"role_id"`   // 角色ID
	IsSuper              uint   `json:"is_super"`  // 是否超级管理员:取值说明在模型里	- 只有管理员有
	ParentId             string `json:"parent_id"` // 上级商户	- 只有商户才有
	Identity             uint   `json:"identity"`  // 身份:取值说明在全局变量里
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// PageRequest 公共分页请求参数
type PageRequest struct {
	Paging   uint `json:"paging" v:"in:0,1"`                        // 是否分页:0否 1是
	Page     *int `json:"page" v:"required-if:paging,1|min:1"`      // 当前页码，当paging=1时必填
	PageSize *int `json:"page_size" v:"required-if:paging,1|min:1"` // 分页数量，当paging=1时必填
}

var SortTypeEnum = []string{
	"DESC",
	"ASC",
}

// SortRequest 公共排序请求参数
type SortRequest struct {
	SortField *string `json:"sort_field" v:"length:2,30"` // 排序字段
	SortType  *string `json:"sort_type" v:"length:2,30"`  // 排序类型
}

// IdRequest 公共ID请求参数
type IdRequest struct {
	Id uint `json:"id" v:"required"` // ID
}

// UploadImgRequest 通用上传图片请求参数
type UploadImgRequest struct {
	EventType string `form:"event_type" v:"checkUploadImgEventType"` // 事件类型
}

// UploadChunkRequest 通用上传切片请求参数
type UploadChunkRequest struct {
	EventType string `form:"event_type" v:"checkMethod"` // 事件类型
}
