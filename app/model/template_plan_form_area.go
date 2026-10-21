package model

import (
	"encoding/json"
	"fmt"
	"orange-crm/global"
	"os"
)

// templatePlanFormAreaColumns defines and stores column names for the table or_template_plan_form_area.
type templatePlanFormAreaColumns struct {
	Id           string // ID
	PlanFormId   string // 限红模板ID
	Name         string // 模板名称
	GameId       string // 游戏ID
	AreaId       string // 区域ID
	MinBetAmount string // 最小投注金额
	MaxBetAmount string // 最大投注金额
	Digits       string // 投注的最小单位倍数
	PlanId       string // 限红方案ID
	CreatedTime  string // 创建时间
	UpdatedTime  string // 更新时间
}

// TemplatePlanFormAreaColumns holds the columns for the table or_template_plan_form_area.
var TemplatePlanFormAreaColumns = templatePlanFormAreaColumns{
	Id:           "id",
	PlanFormId:   "plan_form_id",
	Name:         "name",
	GameId:       "game_id",
	AreaId:       "area_id",
	MinBetAmount: "min_bet_amount",
	MaxBetAmount: "max_bet_amount",
	Digits:       "digits",
	PlanId:       "plan_id",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}

// TemplatePlanFormArea  限红模板投注区域表结构体
type TemplatePlanFormArea struct {
	ID           uint         `gorm:"primarykey;type:int unsigned;comment:ID"`
	PlanFormId   string       `gorm:"size:36;not null;default:'';comment:限红模板ID"`
	Name         string       `gorm:"size:100;not null;default:'';comment:模板名称"`
	GameId       string       `gorm:"size:36;not null;default:'';comment:游戏ID"`
	AreaId       string       `gorm:"size:10;not null;not null;comment:区域ID"`
	MinBetAmount uint         `gorm:"not null;default:0;comment:最小投注金额"`
	MaxBetAmount uint         `gorm:"not null;default:0;comment:最大投注金额"`
	Digits       uint         `gorm:"type:int unsigned;not null;default:1;comment:投注的最小单位倍数"`
	PlanId       string       `gorm:"size:36;not null;default:'';comment:限红方案ID"`
	TemplatePlan TemplatePlan `gorm:"foreignKey:PlanId;references:PlanId"`
	BaseModel
}

// RegisterTable 数据迁移
func (m *TemplatePlanFormArea) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='限红模板投注区域表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 导入 JSON 文件中的数据
		if err := m.ImportAndProcessJSON(); err != nil {
			return err
		}

		// 填充数据
		//data := []TemplatePlanFormArea{
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1001", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1002", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1003", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1004", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1005", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1006", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1007", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1008", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1009", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1010", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1011", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1012", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX001", AreaId: "1013", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1001", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1002", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1003", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1004", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1005", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1006", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1007", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1008", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1009", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1010", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1011", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1012", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX002", AreaId: "1013", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1001", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1002", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1003", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1004", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1005", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1006", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1007", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1008", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1009", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1010", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1011", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1012", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "BACX003", AreaId: "1013", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX001", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX001", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX001", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX001", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX002", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX002", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX002", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX002", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX003", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX003", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX003", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX003", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX004", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX004", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX004", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX004", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX005", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX005", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX005", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX005", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX006", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX006", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX006", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX006", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX007", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX007", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX007", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX007", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX008", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX008", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX008", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX008", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX009", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX009", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX009", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX009", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX010", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX010", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX010", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX010", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX011", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX011", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX011", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX011", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX012", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX012", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX012", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX012", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX013", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX013", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX013", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX013", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX014", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX014", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX014", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX014", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX015", AreaId: "1014", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX015", AreaId: "1015", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX015", AreaId: "1016", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "DTX015", AreaId: "1017", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1100", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1101", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1102", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1103", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1104", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1105", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1106", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1107", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1108", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1109", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1110", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1111", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1112", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1113", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1114", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1115", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1116", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1117", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1118", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1119", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1120", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1121", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1122", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1123", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1124", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1125", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1126", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1127", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1128", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1129", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1130", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1131", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1132", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1133", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1134", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1135", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1136", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1137", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1138", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1139", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1140", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1141", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1142", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1143", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1144", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1145", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1146", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1147", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1148", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1149", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1150", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1151", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1152", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1153", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1154", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1155", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1156", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1157", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1158", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1159", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1160", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1161", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1162", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1163", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1164", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1165", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1166", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1167", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1168", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1169", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1170", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1171", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1172", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1173", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1174", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1175", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1176", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1177", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1178", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1179", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1180", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1181", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1182", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1183", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1184", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1185", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1186", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1187", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1188", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1189", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1190", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1191", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1192", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1193", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1194", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1195", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1196", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1197", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1198", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1199", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1200", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1201", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1202", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1203", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1204", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1205", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1206", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1207", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1208", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1209", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1210", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1211", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1212", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1213", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1214", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1215", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1216", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1217", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1218", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1219", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1220", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1221", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1222", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1223", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1224", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1225", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1226", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1227", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1228", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1229", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1230", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1231", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1232", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1233", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1234", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1235", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1236", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1237", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1238", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1239", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1240", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1241", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1242", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1243", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1244", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1245", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1246", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1247", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1248", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1249", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1250", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1251", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1252", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1253", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1254", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1255", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1256", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1257", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1258", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1259", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1260", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1261", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1262", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1263", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1264", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1265", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1266", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1267", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1268", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1269", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1270", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1271", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1272", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1273", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1274", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1275", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1276", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1277", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1278", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1279", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1280", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1281", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1282", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1283", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1284", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1285", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1286", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1287", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1288", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1289", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1290", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1291", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1292", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1293", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1294", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1295", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1296", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT001", AreaId: "1297", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1100", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1101", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1102", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1103", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1104", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1105", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1106", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1107", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1108", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1109", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1110", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1111", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1112", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1113", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1114", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1115", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1116", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1117", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1118", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1119", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1120", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1121", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1122", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1123", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1124", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1125", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1126", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1127", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1128", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1129", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1130", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1131", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1132", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1133", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1134", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1135", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1136", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1137", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1138", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1139", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1140", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1141", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1142", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1143", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1144", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1145", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1146", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1147", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1148", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1149", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1150", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1151", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1152", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1153", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1154", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1155", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1156", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1157", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1158", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1159", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1160", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1161", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1162", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1163", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1164", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1165", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1166", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1167", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1168", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1169", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1170", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1171", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1172", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1173", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1174", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1175", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1176", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1177", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1178", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1179", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1180", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1181", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1182", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1183", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1184", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1185", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1186", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1187", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1188", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1189", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1190", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1191", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1192", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1193", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1194", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1195", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1196", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1197", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1198", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1199", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1200", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1201", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1202", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1203", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1204", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1205", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1206", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1207", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1208", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1209", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1210", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1211", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1212", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1213", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1214", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1215", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1216", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1217", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1218", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1219", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1220", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1221", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1222", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1223", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1224", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1225", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1226", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1227", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1228", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1229", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1230", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1231", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1232", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1233", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1234", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1235", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1236", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1237", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1238", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1239", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1240", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1241", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1242", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1243", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1244", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1245", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1246", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1247", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1248", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1249", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1250", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1251", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1252", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1253", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1254", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1255", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1256", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1257", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1258", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1259", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1260", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1261", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1262", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1263", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1264", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1265", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1266", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1267", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1268", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1269", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1270", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1271", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1272", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1273", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1274", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1275", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1276", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1277", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1278", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1279", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1280", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1281", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1282", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1283", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1284", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1285", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1286", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1287", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1288", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1289", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1290", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1291", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1292", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1293", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1294", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1295", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1296", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//	{PlanFormId: "9ou422016lod8w4p9449mg4200ta0ujt", Name: "模板:百家乐10-100万|龙虎2-5万|轮盘10-2千|其他100-10万", GameId: "ROT002", AreaId: "1297", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w"},
		//}

		//if err := global.DB.Create(&data).Error; err != nil {
		//	return err
		//}
	}
	return nil
}

// ImportAndProcessJSON 从 JSON 文件导入数据到数据库，并支持对每条数据执行额外处理逻辑
func (m *TemplatePlanFormArea) ImportAndProcessJSON() error {
	filePath := "./data/init/or_template_plan_form_area.json"
	// 打开 JSON 文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	type TemplatePlanFormAreaJSON struct {
		PlanFormId   string `json:"plan_form_id"`
		Name         string `json:"name"`
		GameId       string `json:"game_id"`
		AreaId       string `json:"area_id"`
		MinBetAmount uint   `json:"min_bet_amount"`
		MaxBetAmount uint   `json:"max_bet_amount"`
		Digits       uint   `json:"digits"`
		PlanId       string `json:"plan_id"`
	}

	// 解析 JSON 内容
	var jsonData []TemplatePlanFormAreaJSON
	if err = json.NewDecoder(file).Decode(&jsonData); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	var records []TemplatePlanFormArea
	for _, item := range jsonData {
		records = append(records, TemplatePlanFormArea{
			PlanFormId:   item.PlanFormId,
			Name:         item.Name,
			GameId:       item.GameId,
			AreaId:       item.AreaId,
			MinBetAmount: item.MinBetAmount,
			MaxBetAmount: item.MaxBetAmount,
			Digits:       item.Digits,
			PlanId:       item.PlanId,
		})
	}

	// 插入数据库
	if err = global.DB.Create(&records).Error; err != nil {
		return fmt.Errorf("failed to insert into table %s: %v", "TemplatePlanFormArea", err)
	}

	return nil
}
