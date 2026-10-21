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
)

type ServiceNode struct{}

// List 服务节点列表
func (s *ServiceNode) List(req *request.ServiceNodeListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.ServiceNodeListResponse{},
	}
	db := global.DB.Model(&model.ServiceNode{}).Order("id desc")

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
	var serviceNodes []model.ServiceNode
	if err = db.Find(&serviceNodes).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.ServiceNodeListResponse
	for _, serviceNode := range serviceNodes {
		tmp := response.ServiceNodeListResponse{
			Id:          serviceNode.ID,          // ID
			Ip:          serviceNode.Ip,          // IP地址
			Port:        serviceNode.Port,        // 端口
			OnlineCount: serviceNode.OnlineCount, // 当前在线玩家数
			Disabled:    serviceNode.Disabled,    // 是否禁用:取值说明在模型里
			Note:        serviceNode.Note,        // 备注
		}
		list = append(list, tmp)
	}
	res.List = list

	return
}

// Add 创建服务节点
func (s *ServiceNode) Add(req *request.SaveServiceNodeRequest) error {
	serviceNode := model.ServiceNode{
		Ip:       req.Ip,       // IP地址
		Port:     req.Port,     // 端口
		Disabled: req.Disabled, // 是否禁用:取值说明在模型里
		Note:     req.Note,     // 备注
	}
	return global.DB.Create(&serviceNode).Error
}

// Edit 编辑服务节点
func (s *ServiceNode) Edit(req *request.SaveServiceNodeRequest) (beforeAddr string, err error) {
	var serviceNode model.ServiceNode
	if errors.Is(global.DB.First(&serviceNode, req.Id).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}
	beforeAddr = fmt.Sprintf("%s:%d", serviceNode.Ip, serviceNode.Port)
	serviceNode.Ip = req.Ip             // IP地址
	serviceNode.Port = req.Port         // 端口
	serviceNode.Disabled = req.Disabled // 是否禁用:取值说明在模型里
	serviceNode.Note = req.Note         // 备注
	return beforeAddr, global.DB.Save(&serviceNode).Error
}

// Detail 服务节点详情
func (s *ServiceNode) Detail(id uint) (*response.ServiceNodeDetailResponse, error) {
	var serviceNode model.ServiceNode

	err := global.DB.Model(&model.ServiceNode{}).Where("id = ?", id).First(&serviceNode).Error

	res := response.ServiceNodeDetailResponse{
		Id:          serviceNode.ID,          // ID
		Ip:          serviceNode.Ip,          // IP地址
		Port:        serviceNode.Port,        // 端口
		OnlineCount: serviceNode.OnlineCount, // 当前在线玩家数
		Disabled:    serviceNode.Disabled,    // 是否禁用:取值说明在模型里
		Note:        serviceNode.Note,        // 备注
	}

	return &res, err
}

// Del 删除服务节点
func (s *ServiceNode) Del(id uint) (addr string, err error) {
	var serviceNode model.ServiceNode
	if errors.Is(global.DB.Where("id = ?", id).First(&serviceNode).Error, gorm.ErrRecordNotFound) {
		return "", gorm.ErrRecordNotFound
	}
	addr = fmt.Sprintf("%s:%d", serviceNode.Ip, serviceNode.Port)
	return addr, global.DB.Delete(&serviceNode).Error
}
