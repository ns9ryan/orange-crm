package model

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"orange-crm/global"
	"os"
)

// templatePlanAreaColumns defines and stores column names for the table or_template_plan_area.
type templatePlanAreaColumns struct {
	Id           string // ID
	CategorySub  string // 游戏子分类
	AreaId       string // 区域ID
	AreaName     string // 区域名称
	MinBetAmount string // 最小投注金额
	MaxBetAmount string // 最大投注金额
	Digits       string // 投注的最小单位倍数
	PlanId       string // 限红方案ID
	PlanName     string // 方案名称
	CreatedTime  string // 创建时间
	UpdatedTime  string // 更新时间
}

// TemplatePlanAreaColumns holds the columns for the table or_template_plan_area.
var TemplatePlanAreaColumns = templatePlanAreaColumns{
	Id:           "id",
	CategorySub:  "category_sub",
	AreaId:       "area_id",
	AreaName:     "area_name",
	MinBetAmount: "min_bet_amount",
	MaxBetAmount: "max_bet_amount",
	Digits:       "digits",
	PlanId:       "plan_id",
	PlanName:     "plan_name",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}

// TemplatePlanArea  限红方案投注区域详情表结构体
type TemplatePlanArea struct {
	ID           uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	CategorySub  string `gorm:"size:32;not null;comment:游戏子分类"`
	AreaId       string `gorm:"size:10;not null;not null;comment:区域ID"`
	AreaName     string `gorm:"size:50;not null;default:'';comment:区域名称"`
	MinBetAmount uint   `gorm:"not null;default:0;comment:最小投注金额"`
	MaxBetAmount uint   `gorm:"not null;default:0;comment:最大投注金额"`
	Digits       uint   `gorm:"type:int unsigned;not null;default:1;comment:投注的最小单位倍数"`
	PlanId       string `gorm:"size:36;not null;default:'';comment:限红方案ID"`
	PlanName     string `gorm:"size:100;not null;default:'';comment:方案名称"`
	BaseModel
}

var PlanAreaDigits uint = 100

// RegisterTable 数据迁移
func (m *TemplatePlanArea) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='限红方案投注区域详情表'").Migrator().CreateTable(m); err != nil {
			return err
		}
		filePath := "./data/init/or_template_plan_area.json"
		// 打开 JSON 文件
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filePath, err)
		}
		defer file.Close()

		type TemplatePlanAreaJSON struct {
			CategorySub  string `json:"category_sub"`
			AreaId       string `json:"area_id"`
			AreaName     string `json:"area_name"`
			MinBetAmount uint   `json:"min_bet_amount"`
			MaxBetAmount uint   `json:"max_bet_amount"`
			Digits       uint   `json:"digits"`
			PlanId       string `json:"plan_id"`
			PlanName     string `json:"plan_name"`
		}

		// 解析 JSON 内容
		var jsonData []TemplatePlanAreaJSON
		if err = json.NewDecoder(file).Decode(&jsonData); err != nil {
			return fmt.Errorf("failed to decode JSON: %v", err)
		}

		var records []TemplatePlanArea
		for _, item := range jsonData {
			records = append(records, TemplatePlanArea{
				CategorySub:  item.CategorySub,
				AreaId:       item.AreaId,
				AreaName:     item.AreaName,
				MinBetAmount: item.MinBetAmount,
				MaxBetAmount: item.MaxBetAmount,
				Digits:       item.Digits,
				PlanId:       item.PlanId,
				PlanName:     item.PlanName,
			})
		}

		// 插入数据库
		if err = global.DB.Create(&records).Error; err != nil {
			return fmt.Errorf("failed to insert into table %s: %v", "TemplatePlanArea", err)
		}
	}
	return nil
}

// RegisterTableCopy 数据迁移
func (m *TemplatePlanArea) RegisterTableCopy() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='限红方案投注区域详情表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 百家乐
		bacx := []TemplatePlanArea{
			{CategorySub: "CLASSIC_BACX", AreaId: "1001", AreaName: "闲", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1002", AreaName: "庄", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1003", AreaName: "和", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1004", AreaName: "闲对子", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1005", AreaName: "庄对子", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1006", AreaName: "免佣的庄", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1007", AreaName: "超级Super", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1008", AreaName: "任意对子", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1009", AreaName: "完美对子", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1010", AreaName: "大", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1011", AreaName: "小", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1012", AreaName: "闲龙宝", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1013", AreaName: "庄龙宝", MinBetAmount: 1000, MaxBetAmount: 100000000, Digits: 100, PlanId: "9ou42208hw0d8vc1ki58stc200nmcti4", PlanName: "10至100万"},

			{CategorySub: "CLASSIC_BACX", AreaId: "1001", AreaName: "闲", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1002", AreaName: "庄", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1003", AreaName: "和", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1004", AreaName: "闲对子", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1005", AreaName: "庄对子", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1006", AreaName: "免佣的庄", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1007", AreaName: "超级Super", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1008", AreaName: "任意对子", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1009", AreaName: "完美对子", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1010", AreaName: "大", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1011", AreaName: "小", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1012", AreaName: "闲龙宝", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
			{CategorySub: "CLASSIC_BACX", AreaId: "1013", AreaName: "庄龙宝", MinBetAmount: 100000, MaxBetAmount: 200000000, Digits: 100, PlanId: "9ou42208hw0d8vc1tjsw4t4300pga4zi", PlanName: "1000至200万"},
		}

		// 龙虎
		dTX := []TemplatePlanArea{
			{CategorySub: "DTX", AreaId: "1014", AreaName: "龙", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid", PlanName: "2至5万"},
			{CategorySub: "DTX", AreaId: "1015", AreaName: "闲", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid", PlanName: "2至5万"},
			{CategorySub: "DTX", AreaId: "1016", AreaName: "和", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid", PlanName: "2至5万"},
			{CategorySub: "DTX", AreaId: "1017", AreaName: "同花和", MinBetAmount: 200, MaxBetAmount: 5000000, Digits: 100, PlanId: "9ou42208hw0d8vc258fhwbg400vbzmid", PlanName: "2至5万"},

			{CategorySub: "DTX", AreaId: "1014", AreaName: "龙", MinBetAmount: 1000, MaxBetAmount: 50000000, Digits: 100, PlanId: "9ou42208hw0d8vc2ic1btoo5006cgbhy", PlanName: "10至50万"},
			{CategorySub: "DTX", AreaId: "1015", AreaName: "闲", MinBetAmount: 1000, MaxBetAmount: 50000000, Digits: 100, PlanId: "9ou42208hw0d8vc2ic1btoo5006cgbhy", PlanName: "10至50万"},
			{CategorySub: "DTX", AreaId: "1016", AreaName: "和", MinBetAmount: 1000, MaxBetAmount: 50000000, Digits: 100, PlanId: "9ou42208hw0d8vc2ic1btoo5006cgbhy", PlanName: "10至50万"},
			{CategorySub: "DTX", AreaId: "1017", AreaName: "同花和", MinBetAmount: 1000, MaxBetAmount: 50000000, Digits: 100, PlanId: "9ou42208hw0d8vc2ic1btoo5006cgbhy", PlanName: "10至50万"},
		}

		// 轮盘
		rot1 := []TemplatePlanArea{
			{CategorySub: "ROT", AreaId: "1100", AreaName: "龙", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1101", AreaName: "闲", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//单数(1区) -- 36       1赔35         ( rgnMain: 0, rgnsub 0--- 35 )
			{CategorySub: "ROT", AreaId: "1100", AreaName: "0", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1101", AreaName: "1", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1102", AreaName: "2", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1103", AreaName: "3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1104", AreaName: "4", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1105", AreaName: "5", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1106", AreaName: "6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1107", AreaName: "7", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1108", AreaName: "8", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1109", AreaName: "9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1110", AreaName: "10", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1111", AreaName: "11", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1112", AreaName: "12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1113", AreaName: "13", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1114", AreaName: "14", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1115", AreaName: "15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1116", AreaName: "16", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1117", AreaName: "17", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1118", AreaName: "18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1119", AreaName: "19", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1120", AreaName: "20", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1121", AreaName: "21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1122", AreaName: "22", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1123", AreaName: "23", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1124", AreaName: "24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1125", AreaName: "25", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1126", AreaName: "26", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1127", AreaName: "27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1128", AreaName: "28", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1129", AreaName: "29", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1130", AreaName: "30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1131", AreaName: "31", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1132", AreaName: "32", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1133", AreaName: "33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1134", AreaName: "34", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1135", AreaName: "35", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1136", AreaName: "36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//双数(2区)  1赔17   ( rgnMain:1   rgnsub: 0 -- 59   压双数 )
			{CategorySub: "ROT", AreaId: "1137", AreaName: "1或2", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1138", AreaName: "2或3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1139", AreaName: "4或5", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1140", AreaName: "5或6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1141", AreaName: "7或8", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1142", AreaName: "8或9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1143", AreaName: "10或11", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1144", AreaName: "11或12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1145", AreaName: "13或14", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1146", AreaName: "14或15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1147", AreaName: "16或17", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1148", AreaName: "17或18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1149", AreaName: "19或20", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1150", AreaName: "20或21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1151", AreaName: "22或23", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1152", AreaName: "23或24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1153", AreaName: "25或26", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1154", AreaName: "26或27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1155", AreaName: "28或29", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1156", AreaName: "29或30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1157", AreaName: "31或32", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1158", AreaName: "32或33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1159", AreaName: "34或35", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1160", AreaName: "35或36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1161", AreaName: "1或4", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1162", AreaName: "4或7", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1163", AreaName: "7或10", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1164", AreaName: "10或13", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1165", AreaName: "13或16", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1166", AreaName: "16或19", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1167", AreaName: "19或22", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1168", AreaName: "22或25", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1169", AreaName: "25或28", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1170", AreaName: "28或31", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1171", AreaName: "31或34", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1172", AreaName: "2或5", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1173", AreaName: "5或8", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1174", AreaName: "8或11", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1175", AreaName: "11或14", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1176", AreaName: "14或17", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1177", AreaName: "17或20", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1178", AreaName: "20或23", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1179", AreaName: "23或26", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1180", AreaName: "26或29", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1181", AreaName: "29或32", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1182", AreaName: "32或35", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1183", AreaName: "3或6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1184", AreaName: "6或9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1185", AreaName: "9或12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1186", AreaName: "12或15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1187", AreaName: "15或18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1188", AreaName: "18或21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1189", AreaName: "21或24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1190", AreaName: "24或27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1191", AreaName: "27或30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1192", AreaName: "30或33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1193", AreaName: "33或36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1194", AreaName: "0或1", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1195", AreaName: "0或2", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1196", AreaName: "0或3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//三数(3区)  1赔11   ( rgnMain:2   rgnsub: 0 -- 13   压三数 )
			{CategorySub: "ROT", AreaId: "1197", AreaName: "压三个数：1到3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1198", AreaName: "压三个数：4到6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1199", AreaName: "压三个数：7到9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1200", AreaName: "压三个数：10到12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1201", AreaName: "压三个数：13到15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1202", AreaName: "压三个数：16到18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1203", AreaName: "压三个数：19到21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1204", AreaName: "压三个数：22到24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1205", AreaName: "压三个数：25到27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1206", AreaName: "压三个数：28到30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1207", AreaName: "压三个数：31到33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1208", AreaName: "压三个数：34到36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1209", AreaName: "压三个数：0到2", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1210", AreaName: "压三个数：0或2或3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//四数(4区)  1赔8 ( rgnMain:3   rgnsub: 0 -- 22   压四个数的 )
			{CategorySub: "ROT", AreaId: "1211", AreaName: "压四个数：0/2/4/5", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1212", AreaName: "压四个数：2/3/5/6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1213", AreaName: "压四个数：4/5/7/8", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1214", AreaName: "压四个数：5/6/8/9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1215", AreaName: "压四个数：7/8/10/11", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1216", AreaName: "压四个数：8/9/11/12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1217", AreaName: "压四个数：10/11/13/14", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1218", AreaName: "压四个数：11/12/14/15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1219", AreaName: "压四个数：13/14/16/17", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1220", AreaName: "压四个数：14/15/17/18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1221", AreaName: "压四个数：16/17/19/20", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1222", AreaName: "压四个数：17/18/20/21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1223", AreaName: "压四个数：19/20/22/23", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1224", AreaName: "压四个数：20/21/23/24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1225", AreaName: "压四个数：22/23/25/26", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1226", AreaName: "压四个数：23/24/26/27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1227", AreaName: "压四个数：25/26/28/29", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1228", AreaName: "压四个数：26/27/29/30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1229", AreaName: "压四个数：28/29/31/32", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1230", AreaName: "压四个数：29/30/32/33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1231", AreaName: "压四个数：31/32/34/35", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1232", AreaName: "压四个数：32/33/35/36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1233", AreaName: "压四个数：0/1/2/3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//六数(5区)	1赔5 ( rgnMain:4   rgnsub: 0 -- 10   //压6个数 )
			{CategorySub: "ROT", AreaId: "1234", AreaName: "压六个数：1到6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1235", AreaName: "压六个数：4到9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1236", AreaName: "压六个数：7到12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1237", AreaName: "压六个数：10到15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1238", AreaName: "压六个数：13到18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1239", AreaName: "压六个数：16到21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1240", AreaName: "压六个数：19到24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1241", AreaName: "压六个数：22到27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1242", AreaName: "压六个数：25到30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1243", AreaName: "压六个数：28到33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1244", AreaName: "压六个数：31到36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//12数(6区)	1赔2 ( rgnMain:5   rgnsub: 0 -- 5  压12个数  )
			{CategorySub: "ROT", AreaId: "1245", AreaName: "压12个数：1rd12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1246", AreaName: "压12个数：2nd12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1247", AreaName: "压12个数：3rd12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1248", AreaName: "压12个数：2to1a", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1249", AreaName: "压12个数：2to1b", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1250", AreaName: "压12个数：2to1c", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//红黑，小大，单双(7区)  1赔1 ( rgnMain:6   rgnsub: 0 -- 5  )
			{CategorySub: "ROT", AreaId: "1251", AreaName: "红", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1252", AreaName: "黑", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1253", AreaName: "小", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1254", AreaName: "大", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1255", AreaName: "单", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1256", AreaName: "双", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//以下8-12五个区，是轮盘物理区下注，包括邻注/轮盘下角注/轮盘孤注/零旁注上角/轮上零旁，示意图如下：
			// 5  24 16 33 1  20 14 31 9  22 18 29 7  28 12 35
			// 10                                            3
			// 23  Tiers   \Orphelins   /Voisins       /Jeu0 26
			// 8                                             0
			// 30 11 36 13 27 6  34 17 25 2  21 4  19 15 32
			//Neighbour邻注(8区)
			//说明：目前这版先实现每个数的左右两个邻居(NeighbourTwo)
			//下注：筹码需要是5的倍数，相当于下了5个单注
			//开奖：如果命中5个中的某个数字，赔率固定是1陪35
			{CategorySub: "ROT", AreaId: "1257", AreaName: "邻注0", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1258", AreaName: "邻注32", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1259", AreaName: "邻注15", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1260", AreaName: "邻注19", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1261", AreaName: "邻注4", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1262", AreaName: "邻注21", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1263", AreaName: "邻注2", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1264", AreaName: "邻注25", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1265", AreaName: "邻注17", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1266", AreaName: "邻注34", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1267", AreaName: "邻注6", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1268", AreaName: "邻注27", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1269", AreaName: "邻注13", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1270", AreaName: "邻注36", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1271", AreaName: "邻注11", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1272", AreaName: "邻注30", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1273", AreaName: "邻注8", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1274", AreaName: "邻注23", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1275", AreaName: "邻注10", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1276", AreaName: "邻注5", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1277", AreaName: "邻注24", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1278", AreaName: "邻注16", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1279", AreaName: "邻注33", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1280", AreaName: "邻注1", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1281", AreaName: "邻注20", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1282", AreaName: "邻注14", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1283", AreaName: "邻注31", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1284", AreaName: "邻注9", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1285", AreaName: "邻注22", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1286", AreaName: "邻注18", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1287", AreaName: "邻注29", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1288", AreaName: "邻注7", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1289", AreaName: "邻注28", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1290", AreaName: "邻注12", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1291", AreaName: "邻注35", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1292", AreaName: "邻注3", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			{CategorySub: "ROT", AreaId: "1293", AreaName: "邻注26", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//Tiers，轮盘下角注(9区)
			//说明：0对面的区域，包括30,11,36,13,27,33,16,24,5,10,23,8
			//下注：筹码需要是6的倍数，其中包括5,8及10,11及13,16及23,24及27,30及33,36六个双数注
			//开奖，如果命中其中的某个数字，赔率固定是1陪17
			{CategorySub: "ROT", AreaId: "1294", AreaName: "轮盘下角注", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//Orphelins，轮盘孤注(10区)
			//说明：轮盘侧边的6,34,17,9,31,14,20,1
			//下注：筹码需要是5的倍数，其中包括了1一个单注，和6,9及14,17及17,20及31,34四个双数注
			//开奖：如果开1则赔率是1陪35，如果不是1则赔率是1陪17
			{CategorySub: "ROT", AreaId: "1295", AreaName: "轮盘孤注", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//Voisins，零旁注上角(11区)
			//说明：轮盘上0边上的25,2,21,4,19,15,32,0,26,3,35,12,28,7,29,18,22
			//下注：筹码需要是9的倍数，其中包括了0,2,3一个三数注(注意要两个筹码)、25,26,28,29一个四数注(注意要两个筹码)、4,7和12,15等5个双数注，加起来一共需要9个筹码
			//开奖：如果开0或2或3赔率是1陪11，如果开25或26或28或29赔率是1陪8，开其他数则赔率是1陪17
			{CategorySub: "ROT", AreaId: "1296", AreaName: "零旁注上角", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
			//Jeu0，轮上零旁(12区)
			//说明：轮盘上0边上的15,32,0,26,3,35,12
			//下注：筹码需要是4的倍数，其中包括了26一个单注，0,3双数注，12,15双数注，32,35双数注
			//开奖：如果开26赔率是1陪35，如果不是26则是1陪17
			{CategorySub: "ROT", AreaId: "1297", AreaName: "轮上零旁", MinBetAmount: 1000, MaxBetAmount: 200000, Digits: 100, PlanId: "9ou4220juo0d8vucdxk1q6w100go7j4w", PlanName: "10至2000"},
		}
		// 轮盘
		rot2 := []TemplatePlanArea{
			//单数(1区) -- 36       1赔35         ( rgnMain: 0, rgnsub 0--- 35 )
			{CategorySub: "ROT", AreaId: "1100", AreaName: "0", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1101", AreaName: "1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1102", AreaName: "2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1103", AreaName: "3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1104", AreaName: "4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1105", AreaName: "5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1106", AreaName: "6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1107", AreaName: "7", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1108", AreaName: "8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1109", AreaName: "9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1110", AreaName: "10", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1111", AreaName: "11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1112", AreaName: "12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1113", AreaName: "13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1114", AreaName: "14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1115", AreaName: "15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1116", AreaName: "16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1117", AreaName: "17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1118", AreaName: "18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1119", AreaName: "19", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1120", AreaName: "20", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1121", AreaName: "21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1122", AreaName: "22", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1123", AreaName: "23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1124", AreaName: "24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1125", AreaName: "25", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1126", AreaName: "26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1127", AreaName: "27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1128", AreaName: "28", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1129", AreaName: "29", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1130", AreaName: "30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1131", AreaName: "31", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1132", AreaName: "32", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1133", AreaName: "33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1134", AreaName: "34", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1135", AreaName: "35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1136", AreaName: "36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//双数(2区)  1赔17   ( rgnMain:1   rgnsub: 0 -- 59   压双数 )
			{CategorySub: "ROT", AreaId: "1137", AreaName: "1或2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1138", AreaName: "2或3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1139", AreaName: "4或5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1140", AreaName: "5或6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1141", AreaName: "7或8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1142", AreaName: "8或9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1143", AreaName: "10或11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1144", AreaName: "11或12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1145", AreaName: "13或14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1146", AreaName: "14或15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1147", AreaName: "16或17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1148", AreaName: "17或18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1149", AreaName: "19或20", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1150", AreaName: "20或21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1151", AreaName: "22或23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1152", AreaName: "23或24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1153", AreaName: "25或26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1154", AreaName: "26或27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1155", AreaName: "28或29", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1156", AreaName: "29或30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1157", AreaName: "31或32", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1158", AreaName: "32或33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1159", AreaName: "34或35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1160", AreaName: "35或36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1161", AreaName: "1或4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1162", AreaName: "4或7", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1163", AreaName: "7或10", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1164", AreaName: "10或13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1165", AreaName: "13或16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1166", AreaName: "16或19", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1167", AreaName: "19或22", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1168", AreaName: "22或25", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1169", AreaName: "25或28", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1170", AreaName: "28或31", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1171", AreaName: "31或34", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1172", AreaName: "2或5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1173", AreaName: "5或8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1174", AreaName: "8或11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1175", AreaName: "11或14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1176", AreaName: "14或17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1177", AreaName: "17或20", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1178", AreaName: "20或23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1179", AreaName: "23或26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1180", AreaName: "26或29", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1181", AreaName: "29或32", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1182", AreaName: "32或35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1183", AreaName: "3或6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1184", AreaName: "6或9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1185", AreaName: "9或12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1186", AreaName: "12或15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1187", AreaName: "15或18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1188", AreaName: "18或21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1189", AreaName: "21或24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1190", AreaName: "24或27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1191", AreaName: "27或30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1192", AreaName: "30或33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1193", AreaName: "33或36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1194", AreaName: "0或1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1195", AreaName: "0或2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1196", AreaName: "0或3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//三数(3区)  1赔11   ( rgnMain:2   rgnsub: 0 -- 13   压三数 )
			{CategorySub: "ROT", AreaId: "1197", AreaName: "压三个数：1到3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1198", AreaName: "压三个数：4到6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1199", AreaName: "压三个数：7到9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1200", AreaName: "压三个数：10到12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1201", AreaName: "压三个数：13到15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1202", AreaName: "压三个数：16到18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1203", AreaName: "压三个数：19到21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1204", AreaName: "压三个数：22到24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1205", AreaName: "压三个数：25到27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1206", AreaName: "压三个数：28到30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1207", AreaName: "压三个数：31到33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1208", AreaName: "压三个数：34到36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1209", AreaName: "压三个数：0到2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1210", AreaName: "压三个数：0或2或3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//四数(4区)  1赔8 ( rgnMain:3   rgnsub: 0 -- 22   压四个数的 )
			{CategorySub: "ROT", AreaId: "1211", AreaName: "压四个数：0/2/4/5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1212", AreaName: "压四个数：2/3/5/6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1213", AreaName: "压四个数：4/5/7/8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1214", AreaName: "压四个数：5/6/8/9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1215", AreaName: "压四个数：7/8/10/11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1216", AreaName: "压四个数：8/9/11/12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1217", AreaName: "压四个数：10/11/13/14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1218", AreaName: "压四个数：11/12/14/15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1219", AreaName: "压四个数：13/14/16/17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1220", AreaName: "压四个数：14/15/17/18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1221", AreaName: "压四个数：16/17/19/20", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1222", AreaName: "压四个数：17/18/20/21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1223", AreaName: "压四个数：19/20/22/23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1224", AreaName: "压四个数：20/21/23/24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1225", AreaName: "压四个数：22/23/25/26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1226", AreaName: "压四个数：23/24/26/27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1227", AreaName: "压四个数：25/26/28/29", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1228", AreaName: "压四个数：26/27/29/30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1229", AreaName: "压四个数：28/29/31/32", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1230", AreaName: "压四个数：29/30/32/33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1231", AreaName: "压四个数：31/32/34/35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1232", AreaName: "压四个数：32/33/35/36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1233", AreaName: "压四个数：0/1/2/3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//六数(5区)	1赔5 ( rgnMain:4   rgnsub: 0 -- 10   //压6个数 )
			{CategorySub: "ROT", AreaId: "1234", AreaName: "压六个数：1到6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1235", AreaName: "压六个数：4到9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1236", AreaName: "压六个数：7到12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1237", AreaName: "压六个数：10到15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1238", AreaName: "压六个数：13到18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1239", AreaName: "压六个数：16到21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1240", AreaName: "压六个数：19到24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1241", AreaName: "压六个数：22到27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1242", AreaName: "压六个数：25到30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1243", AreaName: "压六个数：28到33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1244", AreaName: "压六个数：31到36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//12数(6区)	1赔2 ( rgnMain:5   rgnsub: 0 -- 5  压12个数  )
			{CategorySub: "ROT", AreaId: "1245", AreaName: "压12个数：1rd12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1246", AreaName: "压12个数：2nd12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1247", AreaName: "压12个数：3rd12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1248", AreaName: "压12个数：2to1a", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1249", AreaName: "压12个数：2to1b", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1250", AreaName: "压12个数：2to1c", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//红黑，小大，单双(7区)  1赔1 ( rgnMain:6   rgnsub: 0 -- 5  )
			{CategorySub: "ROT", AreaId: "1251", AreaName: "红", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1252", AreaName: "黑", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1253", AreaName: "小", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1254", AreaName: "大", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1255", AreaName: "单", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1256", AreaName: "双", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//以下8-12五个区，是轮盘物理区下注，包括邻注/轮盘下角注/轮盘孤注/零旁注上角/轮上零旁，示意图如下：
			// 5  24 16 33 1  20 14 31 9  22 18 29 7  28 12 35
			// 10                                            3
			// 23  Tiers   \Orphelins   /Voisins       /Jeu0 26
			// 8                                             0
			// 30 11 36 13 27 6  34 17 25 2  21 4  19 15 32
			//Neighbour邻注(8区)
			//说明：目前这版先实现每个数的左右两个邻居(NeighbourTwo)
			//下注：筹码需要是5的倍数，相当于下了5个单注
			//开奖：如果命中5个中的某个数字，赔率固定是1陪35
			{CategorySub: "ROT", AreaId: "1257", AreaName: "邻注0", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1258", AreaName: "邻注32", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1259", AreaName: "邻注15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1260", AreaName: "邻注19", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1261", AreaName: "邻注4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1262", AreaName: "邻注21", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1263", AreaName: "邻注2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1264", AreaName: "邻注25", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1265", AreaName: "邻注17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1266", AreaName: "邻注34", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1267", AreaName: "邻注6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1268", AreaName: "邻注27", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1269", AreaName: "邻注13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1270", AreaName: "邻注36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1271", AreaName: "邻注11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1272", AreaName: "邻注30", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1273", AreaName: "邻注8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1274", AreaName: "邻注23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1275", AreaName: "邻注10", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1276", AreaName: "邻注5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1277", AreaName: "邻注24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1278", AreaName: "邻注16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1279", AreaName: "邻注33", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1280", AreaName: "邻注1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1281", AreaName: "邻注20", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1282", AreaName: "邻注14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1283", AreaName: "邻注31", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1284", AreaName: "邻注9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1285", AreaName: "邻注22", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1286", AreaName: "邻注18", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1287", AreaName: "邻注29", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1288", AreaName: "邻注7", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1289", AreaName: "邻注28", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1290", AreaName: "邻注12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1291", AreaName: "邻注35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1292", AreaName: "邻注3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			{CategorySub: "ROT", AreaId: "1293", AreaName: "邻注26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//Tiers，轮盘下角注(9区)
			//说明：0对面的区域，包括30,11,36,13,27,33,16,24,5,10,23,8
			//下注：筹码需要是6的倍数，其中包括5,8及10,11及13,16及23,24及27,30及33,36六个双数注
			//开奖，如果命中其中的某个数字，赔率固定是1陪17
			{CategorySub: "ROT", AreaId: "1294", AreaName: "轮盘下角注", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//Orphelins，轮盘孤注(10区)
			//说明：轮盘侧边的6,34,17,9,31,14,20,1
			//下注：筹码需要是5的倍数，其中包括了1一个单注，和6,9及14,17及17,20及31,34四个双数注
			//开奖：如果开1则赔率是1陪35，如果不是1则赔率是1陪17
			{CategorySub: "ROT", AreaId: "1295", AreaName: "轮盘孤注", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//Voisins，零旁注上角(11区)
			//说明：轮盘上0边上的25,2,21,4,19,15,32,0,26,3,35,12,28,7,29,18,22
			//下注：筹码需要是9的倍数，其中包括了0,2,3一个三数注(注意要两个筹码)、25,26,28,29一个四数注(注意要两个筹码)、4,7和12,15等5个双数注，加起来一共需要9个筹码
			//开奖：如果开0或2或3赔率是1陪11，如果开25或26或28或29赔率是1陪8，开其他数则赔率是1陪17
			{CategorySub: "ROT", AreaId: "1296", AreaName: "零旁注上角", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
			//Jeu0，轮上零旁(12区)
			//说明：轮盘上0边上的15,32,0,26,3,35,12
			//下注：筹码需要是4的倍数，其中包括了26一个单注，0,3双数注，12,15双数注，32,35双数注
			//开奖：如果开26赔率是1陪35，如果不是26则是1陪17
			{CategorySub: "ROT", AreaId: "1297", AreaName: "轮上零旁", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "9ou4220juo0d8vucf20k8xs200hiuc6y", PlanName: "100至10万"},
		}

		// 筛宝
		sicbo := []TemplatePlanArea{
			{CategorySub: "SICBO", AreaId: "1500", AreaName: "大", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1501", AreaName: "小", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1502", AreaName: "单", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1503", AreaName: "双", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			// 特定双数
			{CategorySub: "SICBO", AreaId: "1504", AreaName: "牌九式12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1505", AreaName: "牌九式13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1506", AreaName: "牌九式14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1507", AreaName: "牌九式15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1508", AreaName: "牌九式16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1509", AreaName: "牌九式23", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1510", AreaName: "牌九式24", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1511", AreaName: "牌九式25", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1512", AreaName: "牌九式26", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1513", AreaName: "牌九式34", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1514", AreaName: "牌九式35", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1515", AreaName: "牌九式36", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1516", AreaName: "牌九式45", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1517", AreaName: "牌九式46", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1518", AreaName: "牌九式56", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			// 特定三数
			{CategorySub: "SICBO", AreaId: "1519", AreaName: "围骰1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1520", AreaName: "围骰2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1521", AreaName: "围骰3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1522", AreaName: "围骰4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1523", AreaName: "围骰5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1524", AreaName: "围骰6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1525", AreaName: "全围", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			// 点数和
			{CategorySub: "SICBO", AreaId: "1526", AreaName: "和值4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1527", AreaName: "和值5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1528", AreaName: "和值6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1529", AreaName: "和值7", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1530", AreaName: "和值8", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1531", AreaName: "和值9", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1532", AreaName: "和值10", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1533", AreaName: "和值11", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1534", AreaName: "和值12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1535", AreaName: "和值13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1536", AreaName: "和值14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1537", AreaName: "和值15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1538", AreaName: "和值16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1539", AreaName: "和值17", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			// 一对组合
			{CategorySub: "SICBO", AreaId: "1540", AreaName: "对子1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1541", AreaName: "对子2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1542", AreaName: "对子3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1543", AreaName: "对子4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1544", AreaName: "对子5", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1545", AreaName: "对子6", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			// 数字
			{CategorySub: "SICBO", AreaId: "1546", AreaName: "单点1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1547", AreaName: "单点12", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1548", AreaName: "单点13", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1549", AreaName: "单点14", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1550", AreaName: "单点15", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
			{CategorySub: "SICBO", AreaId: "1551", AreaName: "单点16", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr16qwd98lxh9mv1ow100ln3goa", PlanName: "100至10万"},
		}

		// 番摊
		ftan := []TemplatePlanArea{
			{CategorySub: "FTAN", AreaId: "1600", AreaName: "小", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1601", AreaName: "大", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1602", AreaName: "单", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1603", AreaName: "双", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1604", AreaName: "1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1605", AreaName: "2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1606", AreaName: "3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1607", AreaName: "4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1608", AreaName: "1-2角", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1609", AreaName: "2-3角", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1610", AreaName: "3-4角", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1611", AreaName: "4-1角", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1612", AreaName: "1念2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1613", AreaName: "1念3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1614", AreaName: "1念4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1615", AreaName: "2念1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1616", AreaName: "2念3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1617", AreaName: "2念4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1618", AreaName: "3念1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1619", AreaName: "3念2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1620", AreaName: "3念4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1621", AreaName: "4念1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1622", AreaName: "4念2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1623", AreaName: "4念3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1624", AreaName: "三门3-2-1", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1625", AreaName: "三门2-1-4", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1626", AreaName: "三门4-3-2", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
			{CategorySub: "FTAN", AreaId: "1627", AreaName: "三门1-4-3", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr8ug0d98llbyrp86w100mo5qhl", PlanName: "100至10万"},
		}

		// 色碟
		sdie := []TemplatePlanArea{
			{CategorySub: "SDIE", AreaId: "1700", AreaName: "单", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1701", AreaName: "双", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1702", AreaName: "小", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1703", AreaName: "大", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1704", AreaName: "0红4白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1705", AreaName: "4红0白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1706", AreaName: "1红3白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1707", AreaName: "3红1白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1708", AreaName: "2红2白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
			{CategorySub: "SDIE", AreaId: "1709", AreaName: "0红4白或4红0白", MinBetAmount: 10000, MaxBetAmount: 10000000, Digits: 100, PlanId: "173r3pr14gsd98lt7fakjz4100sr0b66", PlanName: "100至10万"},
		}

		// 填充数据
		data := slice.Concat(bacx, dTX, rot1, rot2, sicbo, ftan, sdie)

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
