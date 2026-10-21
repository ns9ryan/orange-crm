package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strings"
)

type Dealer struct{}

// GenerateEmployeeNumber 生成工号
func (s *Dealer) GenerateEmployeeNumber() (employeeNumber string, err error) {
	var dealer model.Dealer
	err = global.DB.Last(&dealer).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	inx := 1
	if dealer.ID != 0 {
		inx = int(dealer.ID) + 1
	}
	employeeNumber = fmt.Sprintf("D%07d", inx)
	return employeeNumber, nil
	//src := rand.NewSource(time.Now().UnixNano()) // 使用 Unix 纳秒时间作为种子
	//r := rand.New(src)
	//employeeNumber = utils.GenerateRandomString(r, 8)
	//var dealer model.Dealer
	//for {
	//	err = global.DB.Where(map[string]interface{}{
	//		model.DealerColumns.EmployeeNumber: employeeNumber,
	//	}).First(&dealer).Error
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		break
	//	} else if err != nil {
	//		return "", err
	//	}
	//}
	//uuid, err := random.UUIdV4()
	//if err != nil {
	//	return
	//}
	//return uuid, nil
}

// List 荷官列表
func (s *Dealer) List(req *request.DealerListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.DealerListResponse{},
	}
	db := global.DB.Model(&model.Dealer{})

	// 总条数
	db.Count(&res.Total)

	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var dealers []model.Dealer
	if err = db.Find(&dealers).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.DealerListResponse
	for _, dealer := range dealers {
		tmp := response.DealerListResponse{
			Id:             dealer.ID,               // 荷官ID
			Nickname:       dealer.Nickname,         // 昵称
			EmployeeNumber: dealer.EmployeeNumber,   // 工号
			BalanceAmount:  dealer.BalanceAmount,    // 账户余额
			CurrencyCode:   dealer.CurrencyCode,     // 币种代码
			CurrencyName:   dealer.CurrencyName,     // 币种名称
			CurrencySymbol: dealer.CurrencySymbol,   // 币种符号
			Digits:         dealer.Digits,           // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
			DealerType:     uint(dealer.DealerType), // 类型:取值说明在模型里
			Level:          uint(dealer.Level),      // 等级
			LeadName:       dealer.LeadName,         // 上级名称
			Status:         dealer.Status,           // 状态:取值说明在模型里
			Avatar:         dealer.Avatar,           // 头像
			Qrcode:         dealer.Qrcode,           // qrcode
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建荷官
func (s *Dealer) Add(req *request.SaveDealerRequest) error {
	dealerCurrencyCode := global.CONFIG.Server.DealerCurrencyCode
	if len(dealerCurrencyCode) == 0 {
		dealerCurrencyCode = "USD"
	}

	var systemCurrency model.SystemCurrency
	if errors.Is(global.DB.Where(map[string]interface{}{
		model.SystemCurrencyColumns.Code: dealerCurrencyCode,
	}).First(&systemCurrency).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	dealerModel := model.Dealer{
		Nickname:       strings.TrimSpace(req.Nickname),       // 昵称
		EmployeeNumber: strings.TrimSpace(req.EmployeeNumber), // 工号
		CurrencyCode:   systemCurrency.Code,                   // 币种代码
		CurrencyName:   systemCurrency.Name,                   // 币种名称
		CurrencySymbol: systemCurrency.Symbol,                 // 币种符号
		Digits:         systemCurrency.Digits,                 // 币种的最小单位倍数
		DealerType:     model.DealerType(req.DealerType),      // 类型:取值说明在模型里
		Level:          model.DealerLevel(req.Level),          // 等级
		Status:         req.Status,                            // 状态:取值说明在模型里
		Avatar:         strings.TrimSpace(req.Avatar),         // 头像
		LeadName:       strings.TrimSpace(req.LeadName),       // 上级名称
	}
	// 生成qrcode
	if _, err := utils.MakeQrcode(dealerModel.EmployeeNumber); err != nil {
		return err
	}
	dealerModel.Qrcode = fmt.Sprintf("/%s/%s", global.CONFIG.Upload.Qrcode.SavePath, fmt.Sprintf("%s.png", dealerModel.EmployeeNumber))

	return global.DB.Create(&dealerModel).Error
}

// Edit 编辑荷官
func (s *Dealer) Edit(req *request.SaveDealerRequest) (err error) {
	var dealerModel model.Dealer
	err = global.DB.Where("id = ?", req.Id).First(&dealerModel).Error
	if err != nil {
		return err
	}

	dealerModel.Nickname = strings.TrimSpace(req.Nickname)    // 昵称
	dealerModel.DealerType = model.DealerType(req.DealerType) // 类型:取值说明在模型里
	dealerModel.Level = model.DealerLevel(req.Level)          // 等级
	dealerModel.Status = req.Status                           // 状态:取值说明在模型里
	dealerModel.LeadName = strings.TrimSpace(req.LeadName)    // 上级名称
	dealerModel.Avatar = strings.TrimSpace(req.Avatar)        // 头像

	return global.DB.Save(&dealerModel).Error
}

// Detail 荷官详情
func (s *Dealer) Detail(id uint) (*response.DealerDetailResponse, error) {
	var dealer model.Dealer

	err := global.DB.Model(&model.Dealer{}).First(&dealer, id).Error

	res := response.DealerDetailResponse{
		Id:             dealer.ID,               // 荷官ID
		Nickname:       dealer.Nickname,         // 昵称
		EmployeeNumber: dealer.EmployeeNumber,   // 工号
		BalanceAmount:  dealer.BalanceAmount,    // 账户余额
		CurrencyCode:   dealer.CurrencyCode,     // 币种代码
		CurrencyName:   dealer.CurrencyName,     // 币种名称
		CurrencySymbol: dealer.CurrencySymbol,   // 币种符号
		Digits:         dealer.Digits,           // 币种的最小单位倍数,例如人民币是到分实际计算是按元为单位,这个值就要为100
		DealerType:     uint(dealer.DealerType), // 类型:取值说明在模型里
		Level:          uint(dealer.Level),      // 等级
		LeadName:       dealer.LeadName,         // 上级名称
		Status:         dealer.Status,           // 状态:取值说明在模型里
		Avatar:         dealer.Avatar,           // 头像
		Qrcode:         dealer.Qrcode,           // qrcode
	}
	return &res, err
}

// Del 删除荷官
func (s *Dealer) Del(id uint) (employeeNumber string, err error) {
	var dealerModel model.Dealer
	if errors.Is(global.DB.First(&dealerModel, id).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}

	return dealerModel.EmployeeNumber, global.DB.Delete(&dealerModel).Error
}
