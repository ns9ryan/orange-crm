package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"orange-crm/app/model"
	"orange-crm/global"
	"time"
)

type GameNotifyLog struct{}

// RequestHeader 请求的消息头,每个结构体都有这个头
type RequestHeader struct {
	RequestId uint `json:"RequestId"` // 请求的ID,每次请求这个ID值应不一样
	MsgId     uint `json:"MsgId"`     // 消息Id
}

// ResponseMsg 统一回复的消息体
type ResponseMsg struct {
	ResponseId uint `json:"ResponseId"` // 哪个请求ID对应的回复
	MsgId      uint `json:"MsgId"`      // 处理哪条消息的回复
	Success    uint `json:"Success"`    // 处理的结果 1：成功 0：失败
	ErrCode    int  `json:"ErrCode"`    // 错误码 ，上面的Err值之一

}
type NotifyResponse struct {
	Ret  uint        `json:"Ret"`
	Data ResponseMsg `json:"Data"`
}

// {
//     "Data": {
//         "ResponseId": 1234,
//         "MsgId": 1001,
//         "Sucess": 1,
//         "ErrCode": 0
//     },
//     "Ret": 1
// }

func (s *GameNotifyLog) sendNotifyRequest(c *gin.Context, urlPath string, data interface{}) model.GameNotifyLog {
	fullUrl := fmt.Sprintf("http://%s:%d%s", global.CONFIG.GameNotify.Host, global.CONFIG.GameNotify.Port, urlPath)
	fullLog := model.GameNotifyLog{
		Method: "POST",
		Url:    fullUrl,
		Host:   global.CONFIG.GameNotify.Host, // 请求主机名，如 192.168.0.71
		Port:   global.CONFIG.GameNotify.Port, // 端口号，如 9002
	}
	// 请求数据
	requestData, _ := json.Marshal(data)
	fullLog.RequestData = string(requestData)

	start := time.Now()
	var r *gclient.Response
	var err error
	defer func() {
		if r != nil {
			r.Close()
		}
	}()

	if r, err = g.Client().ContentJson().Post(c, fullUrl, data); err != nil {
		fullLog.Remark = fmt.Sprintf("请求失败:%v", err)
		fullLog.DurationMs = uint(time.Since(start).Milliseconds())
		return fullLog
	}

	fullLog.DurationMs = uint(time.Since(start).Milliseconds()) // 响应耗时，单位毫秒

	// 解析返回值
	var notifyResponse NotifyResponse
	content := r.ReadAll()
	if err = json.Unmarshal(content, &notifyResponse); err != nil {
		fullLog.Remark = fmt.Sprintf("返回值：%s,解析返回值错误:%v", string(content), err)
		return fullLog
	}

	fullLog.Ret = notifyResponse.Ret                    // 返回编码:0失败,1成功
	fullLog.ResponseId = notifyResponse.Data.ResponseId // 哪个请求ID对应的回复
	fullLog.MsgId = notifyResponse.Data.MsgId           // 处理哪条消息的回复
	fullLog.Success = notifyResponse.Data.Success       // 处理的结果:0失败,1成功
	fullLog.ErrCode = notifyResponse.Data.ErrCode       // 错误码

	// 响应数据
	responseData, _ := json.Marshal(notifyResponse)
	fullLog.ResponseData = string(responseData)

	return fullLog
}

// =======================================================商户=======================================================================

// MerchantItem 用于通知新建，更新商户的结构体
type MerchantItem struct {
	RequestHeader

	MerchantId string `json:"MerchantId"` // 商户ID
	ParentId   string `json:"ParentId"`   // 它上级商户的Id
	Status     int    `json:"Status"`     // 商户状态
	BetPause   int    `json:"BetPause"`   // 商户投注状态
}

// SendHttpCreateMerchant 发送创建商户通知
func (s *GameNotifyLog) SendHttpCreateMerchant(c *gin.Context, merchant model.Merchant) {
	if global.CONFIG.GameNotify.Host == "" || global.CONFIG.GameNotify.Port == 0 {
		return
	}
	data := &MerchantItem{
		RequestHeader: RequestHeader{
			RequestId: 1,
			MsgId:     global.CrmMsgId_CreateMerchant,
		},
		MerchantId: merchant.Username,
		ParentId:   merchant.ParentId,
		Status:     int(merchant.Status),
		BetPause:   int(merchant.BetPaused),
	}
	log := s.sendNotifyRequest(c, "/core/createmerchant", data)
	_ = global.DB.Select("*").Create(&log).Error
}

// SendUpdateMerchantHttp 发送更新商户通知
func (s *GameNotifyLog) SendUpdateMerchantHttp(c *gin.Context, merchant model.Merchant) {
	if global.CONFIG.GameNotify.Host == "" || global.CONFIG.GameNotify.Port == 0 {
		return
	}
	data := &MerchantItem{
		RequestHeader: RequestHeader{
			RequestId: 1,
			MsgId:     global.CrmMsgId_UpdateMerchant,
		},
		MerchantId: merchant.Username,
		ParentId:   merchant.ParentId,
		Status:     int(merchant.Status),
		BetPause:   int(merchant.BetPaused),
	}
	log := s.sendNotifyRequest(c, "/core/updatemerchant", data)
	_ = global.DB.Select("*").Create(&log).Error
}

// =========================================================模板=====================================================================

// PlanIdItem 模板ID
type PlanIdItem struct {
	RequestHeader

	Opt      int    `json:"Opt"` // 操作动作  0:新建 1:更新 2:删除
	PlanId   string `json:"PlanId"`
	MinValue uint   `json:"MinValue"`
	MaxValue uint   `json:"MaxValue"`
	Digits   uint   `json:"Digits"`
}

// SendPlanOptHttp 发送限红方案操作通知
func (s *GameNotifyLog) SendPlanOptHttp(c *gin.Context, plan model.TemplatePlan, Opt int) {
	if global.CONFIG.GameNotify.Host == "" || global.CONFIG.GameNotify.Port == 0 {
		return
	}
	data := &PlanIdItem{
		RequestHeader: RequestHeader{
			RequestId: 1,
			MsgId:     global.CrmMsgId_OperatePlanId,
		},
		Opt:      Opt,           // 操作动作  0:新建 1:更新 2:删除
		PlanId:   plan.PlanId,   // 限红方案ID
		MinValue: plan.MinValue, // 方案的最小投注值
		MaxValue: plan.MaxValue, // 方案的最大投注值
		Digits:   plan.Digits,   // 方案投注的最小单位倍数
	}
	log := s.sendNotifyRequest(c, "/core/operateplanid", data)
	_ = global.DB.Select("*").Create(&log).Error
}

// ======================================================游戏配置========================================================================

// GameConfigItem 游戏配置更新
type GameConfigItem struct {
	RequestHeader

	Opt          int    `json:"Opt"`          // 操作动作  0:新建 1:更新 2:删除
	GameId       string `json:"GameId"`       // 游戏ID
	DeskNo       uint   `json:"DeskNo"`       // 桌号
	TotalBetTime uint   `json:"TotalBetTime"` // 下注时间，以秒为单位
	VideoUrl     string `json:"VideoUrl"`     // 视讯游戏地址
	GameTitle    string `json:"GameTitle"`    // 游戏标题

}

// SendGameConfigOptHttp 发送游戏配置操作通知
func (s *GameNotifyLog) SendGameConfigOptHttp(c *gin.Context, gameTable model.GameTable, Opt int) {
	if global.CONFIG.GameNotify.Host == "" || global.CONFIG.GameNotify.Port == 0 {
		return
	}
	data := &GameConfigItem{
		RequestHeader: RequestHeader{
			RequestId: 1,
			MsgId:     global.CrmMsgId_OperateGameConfig,
		},
		Opt:          Opt,                     // 操作动作  0:新建 1:更新 2:删除
		GameId:       gameTable.GameTableCode, // 游戏ID
		DeskNo:       gameTable.DeskNo,        // 桌号
		TotalBetTime: gameTable.BetInterval,   // 下注时间，以秒为单位
		VideoUrl:     gameTable.StreamUrl,     // 视讯游戏地址
		GameTitle:    gameTable.Description,   // 游戏标题
	}
	log := s.sendNotifyRequest(c, "/core/operategameconfig", data)
	_ = global.DB.Select("*").Create(&log).Error
}

// ==============================================================================================================================

// PlayerItem 更新玩家的数据
type PlayerItem struct {
	RequestHeader

	PlayerId string `json:"PlayerId"`
	ParentId string `json:"ParentId"`

	Status   int `json:"Status"`
	BetPause int `json:"BetPause"`
}
