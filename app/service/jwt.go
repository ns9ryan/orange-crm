package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
	"strings"
	"time"
)

// LoginAdminToken 获取管理员token
func LoginAdminToken(admin model.Admin) (token string, claims request.CustomClaims, err error) {
	// 过期时间
	expiresDuration := time.Second * time.Duration(global.CONFIG.Jwt.ExpiresTime)

	var roleModel model.Role
	global.DB.Where("id = ?", admin.RoleId).First(&roleModel)
	claims = request.CustomClaims{
		ID:       admin.ID,             // 管理员ID
		Username: admin.Username,       // 账号
		Nickname: admin.Nickname,       // 昵称
		RoleId:   admin.RoleId,         // 角色ID
		IsSuper:  admin.IsSuper,        // 是否超级管理员:取值说明在模型里
		Identity: global.AdminIdentity, // 身份:取值说明在全局变量里
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strings.TrimSpace(global.CONFIG.Jwt.Issuer),         // 发布者
			Audience:  jwt.ClaimStrings{"crm"},                             // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),           // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      // 发行时间
		},
	}

	token, err = CreateToken(claims)

	return
}

// LoginMerchantToken 获取商户token
func LoginMerchantToken(merchant model.Merchant) (token string, claims request.CustomClaims, err error) {
	// 过期时间
	expiresDuration := time.Second * time.Duration(global.CONFIG.Jwt.ExpiresTime)

	claims = request.CustomClaims{
		ID:       merchant.ID,             // 商户ID
		Username: merchant.Username,       // 账号
		Nickname: merchant.Nickname,       // 昵称
		RoleId:   merchant.RoleId,         // 角色ID
		ParentId: merchant.ParentId,       // 上级商户
		Identity: global.MerchantIdentity, // 身份:取值说明在全局变量里
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strings.TrimSpace(global.CONFIG.Jwt.Issuer),         // 发布者
			Audience:  jwt.ClaimStrings{"crm"},                             // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),           // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      // 发行时间
		},
	}

	token, err = CreateToken(claims)

	return
}

// CreateToken 创建token
func CreateToken(claims request.CustomClaims) (string, error) {
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(strings.TrimSpace(global.CONFIG.Jwt.SecretKey)))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*request.CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(strings.TrimSpace(global.CONFIG.Jwt.SecretKey)), nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetAdminInfo 从Gin的Context中获取从jwt解析出来的管理员信息
func GetAdminInfo(c *gin.Context) (*request.CustomClaims, error) {
	userinfo, exist := c.Get("userinfo")
	if !exist {
		return nil, errors.New("登录失效")
	}
	userCache := userinfo.(*request.CustomClaims)
	return userCache, nil
}

// GetMerchantInfo 从Gin的Context中获取从jwt解析出来的商户信息
func GetMerchantInfo(c *gin.Context) (*request.CustomClaims, error) {
	userinfo, exist := c.Get("userinfo")
	if !exist {
		return nil, errors.New("登录失效")
	}
	userCache := userinfo.(*request.CustomClaims)
	return userCache, nil
}
