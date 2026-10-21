package model

import (
	"gorm.io/gorm"
	"orange-crm/global"
)

// templatePlanColumns defines and stores column names for the table or_template_plan.
type templatePlanColumns struct {
	Id          string // ID
	CategorySub string // 游戏子分类
	Name        string // 方案名称
	PlanId      string // 限红方案ID
	MinValue    string // 方案的最小投注值
	MaxValue    string // 方案的最大投注值
	Digits      string // 方案投注的最小单位倍数
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// TemplatePlanColumns holds the columns for the table or_template_plan.
var TemplatePlanColumns = templatePlanColumns{
	Id:          "id",
	CategorySub: "category_sub",
	Name:        "name",
	PlanId:      "plan_id",
	MinValue:    "min_value",
	MaxValue:    "max_value",
	Digits:      "digits",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// TemplatePlan  限红方案表结构体
type TemplatePlan struct {
	ID                uint               `gorm:"primarykey;type:int unsigned;comment:ID"`
	CategorySub       string             `gorm:"size:32;not null;comment:游戏子分类"`
	Name              string             `gorm:"size:100;not null;default:'';comment:方案名称"`
	PlanId            string             `gorm:"uniqueindex;size:36;not null;default:'';comment:限红方案ID"`
	MinValue          uint               `gorm:"not null;default:0;comment:方案的最小投注值"`
	MaxValue          uint               `gorm:"not null;default:0;comment:方案的最大投注值"`
	Digits            uint               `gorm:"type:int unsigned;not null;default:1;comment:方案投注的最小单位倍数"`
	TemplatePlanAreas []TemplatePlanArea `gorm:"foreignKey:PlanId;references:PlanId"`
	GameCategory      GameCategory       `gorm:"foreignKey:CategorySub;references:SubCode"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *TemplatePlan) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='限红方案表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		//filePath := "./data/init/or_template_plan.json"
		//// 打开 JSON 文件
		//file, err := os.Open(filePath)
		//if err != nil {
		//	return fmt.Errorf("failed to open file %s: %v", filePath, err)
		//}
		//defer file.Close()
		//
		//type TemplatePlanJSON struct {
		//	CategorySub string `json:"category_sub"`
		//	Name        string `json:"name"`
		//	PlanId      string `json:"plan_id"`
		//	MinValue    uint   `json:"min_value"`
		//	MaxValue    uint   `json:"max_value"`
		//	Digits      uint   `json:"digits"`
		//}
		//
		//// 解析 JSON 内容
		//var jsonData []TemplatePlanJSON
		//if err = json.NewDecoder(file).Decode(&jsonData); err != nil {
		//	return fmt.Errorf("failed to decode JSON: %v", err)
		//}
		//
		//var records []TemplatePlan
		//for _, item := range jsonData {
		//	records = append(records, TemplatePlan{
		//		CategorySub: item.CategorySub,
		//		Name:        item.Name,
		//		PlanId:      item.PlanId,
		//		MinValue:    item.MinValue,
		//		MaxValue:    item.MaxValue,
		//		Digits:      item.Digits,
		//	})
		//}
		//
		//// 插入数据库
		//if err = global.DB.Create(&records).Error; err != nil {
		//	return fmt.Errorf("failed to insert into table %s: %v", "TemplatePlanArea", err)
		//}

		// 填充数据
		data := []TemplatePlan{
			{CategorySub: "CLASSIC_BACX", Name: "CNY:10-1万", MinValue: 10, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr11z8d9c69d0igjak20024dyy5"},
			{CategorySub: "CLASSIC_BACX", Name: "USD:1-1千", MinValue: 10, MaxValue: 100000, Digits: 100, PlanId: "173r3pr11z8d9c6b3szwp243001k10ee"},
			{CategorySub: "CLASSIC_BACX", Name: "RUB:100-10万", MinValue: 100, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr11z8d9c6nxc5hr584004ztih6"},
			{CategorySub: "CLASSIC_BACX", Name: "VND:10万-200万", MinValue: 400000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr11z8d9c6wet5tjbg500ued2ca"},

			{CategorySub: "DTX", Name: "CNY:10-1万", MinValue: 100, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr190od9c7esiaq23s100bui45r"},
			{CategorySub: "DTX", Name: "USD:1-1千", MinValue: 10, MaxValue: 100000, Digits: 100, PlanId: "173r3pr190od9c7ld0woczc200pi0ntl"},
			{CategorySub: "DTX", Name: "RUB:100-10万", MinValue: 200, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr190od9c7n947xhq03009jkfsu"},
			{CategorySub: "DTX", Name: "VND:10万-200万", MinValue: 200000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr190od9c7o56ffe00400nich1z"},

			{CategorySub: "ROT", Name: "CNY:10-1万", MinValue: 100, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr190od9c7ssgy3w4o5000oa27o"},
			{CategorySub: "ROT", Name: "USD:1-1千", MinValue: 10, MaxValue: 100000, Digits: 100, PlanId: "173r3pr190od9c7w180jckc600h0bw23"},
			{CategorySub: "ROT", Name: "RUB:100-10万", MinValue: 300, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr190od9c8ap17rrfg70067jfh9"},
			{CategorySub: "ROT", Name: "VND:10万-200万", MinValue: 300000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr190od9c8hbs0lbpk800m3oegh"},

			{CategorySub: "SICBO", Name: "CNY:10-1万", MinValue: 100, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr190od9c8oit01wsg900gb9m7l"},
			{CategorySub: "SICBO", Name: "USD:1-1千", MinValue: 10, MaxValue: 100000, Digits: 100, PlanId: "173r3pr190od9c8rfzeleioa006k3o61"},
			{CategorySub: "SICBO", Name: "RUB:100-10万", MinValue: 100, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr190od9c8vqnox13cb00r8c2dg"},
			{CategorySub: "SICBO", Name: "VND:10万-200万", MinValue: 100000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr190od9c90vgxyq2kc00b6969i"},

			{CategorySub: "FTAN", Name: "CNY:10-1万", MinValue: 1000, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr190od9c91hea84ood00322k4y"},
			{CategorySub: "FTAN", Name: "USD:1-1千", MinValue: 100, MaxValue: 100000, Digits: 100, PlanId: "173r3pr190od9c91opbk7bse00mrcnb0"},
			{CategorySub: "FTAN", Name: "RUB:100-10万", MinValue: 10000, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr190od9c91wnv7h9cf00f29o0w"},
			{CategorySub: "FTAN", Name: "VND:10万-200万", MinValue: 10000000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr190od9c9242pq3xwg00u0ptur"},

			{CategorySub: "SDIE", Name: "CNY:10-1万", MinValue: 100, MaxValue: 1000000, Digits: 100, PlanId: "173r3pr190od9c93umyqkj8h00dhddog"},
			{CategorySub: "SDIE", Name: "USD:1-1千", MinValue: 10, MaxValue: 100000, Digits: 100, PlanId: "173r3pr190od9c95d7jyrzci00zgdt5x"},
			{CategorySub: "SDIE", Name: "RUB:100-10万", MinValue: 600, MaxValue: 10000000, Digits: 100, PlanId: "173r3pr190od9c974z1cdkkj003u5mbp"},
			{CategorySub: "SDIE", Name: "VND:10万-200万", MinValue: 800000, MaxValue: 200000000, Digits: 100, PlanId: "173r3pr190od9c98yf1nzqkk00b3703v"},
		}

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete 删除钩子
func (m *TemplatePlan) AfterDelete(tx *gorm.DB) (err error) {
	// 删除区域数据
	tx.Where(map[string]interface{}{
		TemplatePlanAreaColumns.PlanId: m.PlanId,
	}).Delete(&TemplatePlanArea{})
	return
}
