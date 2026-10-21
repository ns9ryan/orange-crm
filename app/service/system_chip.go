package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"orange-crm/app/model"
	"orange-crm/app/request"
	"orange-crm/app/response"
	"orange-crm/global"
	"orange-crm/utils"
	"strconv"
)

type SystemChip struct{}

// List 系统筹码列表
func (s *SystemChip) List(req *request.SystemChipListRequest) (res *response.ListData, err error) {
	res = &response.ListData{
		Total: 0,
		List:  []response.SystemChipResponse{},
	}
	db := global.DB.Model(&model.SystemChip{})

	// 总条数
	db.Count(&res.Total)
	// 没有数据直接返回
	if res.Total <= 0 {
		return
	}

	// 分页
	if req.Paging > 0 {
		db.Scopes(utils.Paginate(*req.Page, *req.PageSize))
	}

	// 获取数据
	var chips []model.SystemChip
	if err = db.Find(&chips).Error; err != nil {
		return
	}

	// 列表赋值
	var list []response.SystemChipResponse
	for _, chip := range chips {
		tmp := response.SystemChipResponse{
			Id:     chip.ID,                                                                    // 筹码ID
			Name:   chip.Name,                                                                  // 名称
			Amount: strconv.FormatFloat(float64(chip.Value)/float64(chip.Digits), 'f', -1, 64), // 面额
			Status: chip.Status,                                                                // 状态:取值说明在模型里
		}

		list = append(list, tmp)
	}
	res.List = list

	return
}

// GetSystemChipBox 系统筹码下拉框
func (s *SystemChip) GetSystemChipBox() (list []response.SystemChipResponse, err error) {
	// 先从redis读取数据
	data, err := g.Redis().Do(gctx.New(), "GET", "system:system_chip")
	if err != nil {
		return nil, err
	}
	if err = data.Struct(&list); err != nil {
		return nil, err
	}
	// 如果存在直接返回
	if len(list) != 0 {
		return list, err
	}

	// 从数据库种读取数据
	var chips []model.SystemChip
	if err = global.DB.Find(&chips).Error; err != nil {
		return
	}

	// 赋值
	for _, chip := range chips {
		tmp := response.SystemChipResponse{
			Id:     chip.ID,                                                                    // 筹码ID
			Name:   chip.Name,                                                                  // 名称
			Amount: strconv.FormatFloat(float64(chip.Value)/float64(chip.Digits), 'f', -1, 64), // 面额
			Status: chip.Status,                                                                // 状态:取值说明在模型里
		}

		list = append(list, tmp)
	}

	// 存入redis
	_, err = g.Redis().Do(gctx.New(), "set", "system:system_chip", list)

	return list, err
}
