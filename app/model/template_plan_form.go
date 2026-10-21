package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"orange-crm/global"
	"os"
)

// templatePlanFormColumns defines and stores column names for the table or_template_plan_form.
type templatePlanFormColumns struct {
	Id         string // ID
	PlanFormId string // 限红模板ID
	Name       string // 模板名称
	//Content     string // 详细内容
	CreatedTime string // 创建时间
	UpdatedTime string // 更新时间
}

// TemplatePlanFormColumns holds the columns for the table or_template_plan_form.
var TemplatePlanFormColumns = templatePlanFormColumns{
	Id:         "id",
	PlanFormId: "plan_form_id",
	Name:       "name",
	//Content:     "content",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// TemplatePlanForm  限红模板表结构体
type TemplatePlanForm struct {
	ID         uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	PlanFormId string `gorm:"uniqueindex;size:36;default:'';comment:限红模板ID"`
	Name       string `gorm:"size:100;not null;default:'';comment:模板名称"`
	//Content               json.RawMessage        `gorm:"type:json;comment:详细内容"`
	TemplatePlanFormAreas []TemplatePlanFormArea `gorm:"foreignKey:PlanFormId;references:PlanFormId"`
	BaseModel
}

// AfterDelete 删除钩子
func (m *TemplatePlanForm) AfterDelete(tx *gorm.DB) (err error) {
	// 删除区域数据
	tx.Where(map[string]interface{}{
		TemplatePlanFormAreaColumns.PlanFormId: m.PlanFormId,
	}).Delete(&TemplatePlanFormArea{})
	return
}

// GameTemplatePlan 游戏限红模板数据
type GameTemplatePlan struct {
	GameId                string                 `json:"game_id"`                  // 游戏ID
	GameTemplatePlanAreas []GameTemplatePlanArea `json:"game_template_plan_areas"` // 限红区域
}

// GameTemplatePlanArea 游戏限红方案区域数据
type GameTemplatePlanArea struct {
	AreaId       string `json:"area_id"`        // 区域ID
	MinBetAmount uint   `json:"min_bet_amount"` // 最小投注金额
	MaxBetAmount uint   `json:"max_bet_amount"` // 最大投注金额
	Digits       uint   `json:"digits"`         // 投注的最小单位倍数
}

// RegisterTable 数据迁移
func (m *TemplatePlanForm) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='限红模板表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		filePath := "./data/init/or_template_plan_form.json"
		// 打开 JSON 文件
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filePath, err)
		}
		defer file.Close()

		type TemplatePlanFormJSON struct {
			PlanFormId string `json:"plan_form_id"` // 限红模板ID
			Name       string `json:"name"`         // 模板名称
			//Content    string `json:"content"`      // 详细内容
		}

		// 解析 JSON 内容
		var jsonData []TemplatePlanFormJSON
		if err = json.NewDecoder(file).Decode(&jsonData); err != nil {
			return fmt.Errorf("failed to decode JSON: %v", err)
		}

		var records []TemplatePlanForm
		for _, item := range jsonData {
			//content, _ := json.Marshal(item.Content)
			records = append(records, TemplatePlanForm{
				PlanFormId: item.PlanFormId,
				Name:       item.Name,
				//Content:    content,
			})
		}

		// 插入数据库
		if err = global.DB.Create(&records).Error; err != nil {
			return fmt.Errorf("failed to insert into table %s: %v", "TemplatePlanForm", err)
		}

		// 导入 JSON 文件中的数据
		//if err := utils.ImportAndProcessJSON[TemplatePlanForm](global.DB, "./data/init/or_template_plan_form.json", "TemplatePlanForm", nil); err != nil {
		//	return err
		//}
	}
	return nil
}
