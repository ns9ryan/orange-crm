package validate

import (
	"github.com/gogf/gf/v2/frame/g"
	"orange-crm/app/request"
	"orange-crm/global"
)

type LogModule struct{}

// AdminLoginLogListValidate 管理员登录日志列表验证
func (v *LogModule) AdminLoginLogListValidate(req *request.AdminLoginLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// MerchantLoginLogListValidate 商户登录日志列表验证
func (v *LogModule) MerchantLoginLogListValidate(req *request.MerchantLoginLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// PlayerLoginLogListValidate 会员登录日志列表验证
func (v *LogModule) PlayerLoginLogListValidate(req *request.PlayerLoginLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// AdminOperationLogListValidate 管理员操作日志列表验证
func (v *LogModule) AdminOperationLogListValidate(req *request.AdminOperationLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// MerchantOperationLogListValidate 商户操作日志列表验证
func (v *LogModule) MerchantOperationLogListValidate(req *request.MerchantOperationLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ControlLoginLogListValidate 控端登录日志列表验证
func (v *LogModule) ControlLoginLogListValidate(req *request.ControlLoginLogListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}
