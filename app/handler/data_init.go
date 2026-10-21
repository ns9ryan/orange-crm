package handler

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/service"
	"orange-crm/app/validate"
	"orange-crm/global"
	"orange-crm/utils"
	"time"
)

type DataInit struct{}

// AddMerchants 批量添加商户
func (h *DataInit) AddMerchants(c *gin.Context) {
	var req request.AddMerchantsRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.DataInit.AddMerchantsValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.DataInit.AddMerchants(c, &req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// AddPlayers 批量添加玩家
func (h *DataInit) AddPlayers(c *gin.Context) {
	var req request.AddPlayersRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.DataInit.AddPlayersValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.DataInit.AddPlayers(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// Deposit 批量存款上分
func (h *DataInit) Deposit(c *gin.Context) {
	var req request.PlayersDepositRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	if err := validate.Validate.DataInit.PlayersDepositValidate(&req); err != nil {
		utils.Response.FailValidate(c, err)
		return
	}

	if err := service.Service.DataInit.Deposit(&req); err != nil {
		utils.Response.FailSystem(c, err)
		return
	}

	utils.Response.SuccessSimple(c)
}

// EditLoginLog 更新登录日志
func (h *DataInit) EditLoginLog(c *gin.Context) {
	var usernames []string
	global.DB.Model(&model.LoginLog{}).Where(map[string]interface{}{
		model.LoginLogColumns.UserType: model.LoginLogPlayer,
		model.LoginLogColumns.RootKey:  0,
	}).Group(model.LoginLogColumns.Username).Pluck(model.LoginLogColumns.Username, &usernames)
	for _, username := range usernames {
		var player model.Player
		if err := global.DB.Where(map[string]interface{}{
			model.PlayerColumns.Username: username,
		}).First(&player, model.Player{}).Error; err != nil {
			continue
		}
		global.DB.Model(&model.LoginLog{}).Where(map[string]interface{}{
			model.LoginLogColumns.Username: username,
			model.LoginLogColumns.UserType: model.LoginLogPlayer,
			model.LoginLogColumns.RootKey:  0,
		}).Updates(model.LoginLog{RootKey: player.RootKey, MerchantId: player.MerchantId})
	}
	utils.Response.SuccessSimple(c)
}

// EditPlayerTransaction 更新玩家流水表的商户字段，后面增加的字段没有赋值
func (h *DataInit) EditPlayerTransaction(c *gin.Context) {
	var usernames []string
	global.DB.Model(&model.PlayerTransaction{}).Where(map[string]interface{}{
		model.PlayerTransactionColumns.RootKey: 0,
	}).Group(model.PlayerTransactionColumns.Username).Pluck(model.PlayerTransactionColumns.Username, &usernames)

	for _, username := range usernames {
		var player model.Player
		if err := global.DB.Where(map[string]interface{}{
			model.PlayerColumns.Username: username,
		}).First(&player, model.Player{}).Error; err != nil {
			continue
		}
		global.DB.Model(&model.PlayerTransaction{}).Where(map[string]interface{}{
			model.PlayerTransactionColumns.Username: username,
			model.PlayerTransactionColumns.RootKey:  0,
		}).Updates(model.PlayerTransaction{RootKey: player.RootKey, MerchantId: player.MerchantId})
	}
	utils.Response.SuccessSimple(c)
}

// EditPlayerWallet 更新玩家钱包表的商户字段，后面增加的字段没有赋值
func (h *DataInit) EditPlayerWallet(c *gin.Context) {
	var usernames []string
	global.DB.Model(&model.PlayerWallet{}).Where(map[string]interface{}{
		model.PlayerWalletColumns.RootKey: 0,
	}).Group(model.PlayerWalletColumns.Username).Pluck(model.PlayerWalletColumns.Username, &usernames)

	for _, username := range usernames {
		var player model.Player
		if err := global.DB.Where(map[string]interface{}{
			model.PlayerColumns.Username: username,
		}).First(&player, model.Player{}).Error; err != nil {
			continue
		}
		global.DB.Model(&model.PlayerWallet{}).Where(map[string]interface{}{
			model.PlayerWalletColumns.Username: username,
			model.PlayerWalletColumns.RootKey:  0,
		}).Updates(model.PlayerTransaction{RootKey: player.RootKey})
	}
	utils.Response.SuccessSimple(c)
}

// EditMerchantTips 更新商户的打赏表
func (h *DataInit) EditMerchantTips(c *gin.Context) {
	var usernames []string

	global.DB.Model(&model.Merchant{}).Where(map[string]interface{}{
		model.MerchantColumns.ParentId: global.MerchantRootName,
	}).Group(model.MerchantColumns.Username).Pluck(model.MerchantColumns.Username, &usernames)

	for _, username := range usernames {
		var merchantCurrency model.MerchantCurrency
		if err := global.DB.Where(map[string]interface{}{
			model.MerchantCurrencyColumns.Username: username,
		}).First(&merchantCurrency).Error; err != nil {
			continue
		}

		var systemChips []model.SystemTip
		if err := global.DB.Where(map[string]interface{}{
			model.SystemTipColumns.Currency: merchantCurrency.Currency,
		}).Find(&systemChips).Error; err != nil {
			continue
		}

		if len(merchantCurrency.TipOptions) == 0 {
			var tipConfig []uint
			for _, systemChip := range systemChips {
				tipConfig = append(tipConfig, systemChip.Value)
			}

			if len(tipConfig) == 0 {
				continue
			}

			slice.Sort(tipConfig)
			if len(tipConfig) < 1 {
				continue
			}
			first := tipConfig[0]
			last := tipConfig[len(tipConfig)-1]
			if first < merchantCurrency.MinTip {
				merchantCurrency.MinTip = first
			}

			if last > merchantCurrency.MaxTip {
				merchantCurrency.MaxTip = last
			}

			merchantCurrency.TipOptions = slice.Join(tipConfig, ",")

			if err := global.DB.Save(&merchantCurrency).Error; err != nil {
				continue
			}
		}
	}

	utils.Response.SuccessSimple(c)
}

// MakeTokenRequest 批量添加玩家请求参数
type MakeTokenRequest struct {
	MerchantCode string `json:"merchant_code" v:"required"`
}

type MakeTokenResponse struct {
	MerchantCode string `json:"merchant_code"`
	OrToken      string `json:"or_token"`     // 商户api调用验证or_token
	OrTokenKey   string `json:"or_token_key"` // 生成token用的商户唯一key，永久有效
}

// MakeToken 生成商户token
func (h *DataInit) MakeToken(c *gin.Context) {
	var req MakeTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Response.FailBadRequest(c, err)
		return
	}

	orTokenKey := grand.S(32, false) // 生成token用的商户唯一key，永久有效
	var orToken string
	var err error
	// 商户api调用验证or_token
	orToken, err = h.makeOrToken(req.MerchantCode, []byte(orTokenKey))
	if err != nil {
		utils.Response.FailSystem(c, err)
		return
	}
	res := &MakeTokenResponse{
		MerchantCode: req.MerchantCode, // ID
		OrToken:      orToken,          // 商户api调用验证or_token
		OrTokenKey:   orTokenKey,       // 生成token用的商户唯一key，永久有效
	}
	utils.Response.SuccessWithData(c, res)
}

// 生成or_token
func (h *DataInit) makeOrToken(merchantCode string, orTokenKey []byte) (string, error) {
	// 过期时间（纳秒）
	var expiresDuration time.Duration
	if global.CONFIG.Jwt.MerchantTokenExpiresTime == 0 {
		expiresDuration = 30 * 24 * time.Hour // 30天，单位是纳秒
	} else {
		expiresDuration = time.Duration(global.CONFIG.Jwt.MerchantTokenExpiresTime) * time.Second // 把秒转换成纳秒
	}
	claims := jwt.MapClaims{
		"merchant_code": merchantCode,                           // 商户号
		"ext":           time.Now().Add(expiresDuration).Unix(), // 过期时间
		"iat":           time.Now().Unix(),                      // 发行时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(orTokenKey)
}
