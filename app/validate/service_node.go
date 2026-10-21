package validate

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"gorm.io/gorm"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/global"
)

type ServiceNode struct{}

// SaveValidate 创建、编辑服务节点验证
func (v *ServiceNode) SaveValidate(req *request.SaveServiceNodeRequest) error {
	// 自定义校验
	ruleFuncMap := map[string]gvalid.RuleFunc{
		"checkId": v.checkId,
	}
	if err := g.Validator().Bail().RuleFuncMap(ruleFuncMap).Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// ListValidate 服务节点列表验证
func (v *ServiceNode) ListValidate(req *request.ServiceNodeListRequest) error {
	if err := g.Validator().Bail().Data(req).Run(global.I18N.Context); err != nil {
		return err
	}
	return nil
}

// 验证公告是否存在
func (v *ServiceNode) checkId(ctx context.Context, in gvalid.RuleFuncInput) error {
	if in.Value.Uint() != 0 {
		var notice = model.ServiceNode{ID: in.Value.Uint()}
		if err := global.DB.First(&notice).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return gerror.Newf(`{#service_node.id.no.exist}`)
		}
	}
	return nil
}
