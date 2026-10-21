package model

import (
	"github.com/duke-git/lancet/v2/slice"
	"orange-crm/global"
)

// betAreaColumns defines and stores column names for the table or_bet_area.
type betAreaColumns struct {
	Id               string // ID
	GameCategoryCode string // 游戏分类
	AreaId           string // 区域ID
	AreaKey          string // 区域标识
	AreaName         string // 区域名称
	Description      string // 下注区域说明
	Odds             string // 赔率
}

// BetAreaColumns holds the columns for the table or_bet_area.
var BetAreaColumns = betAreaColumns{
	Id:               "id",
	GameCategoryCode: "game_category_code",
	AreaId:           "area_id",
	AreaKey:          "area_key",
	AreaName:         "area_name",
	Description:      "description",
	Odds:             "odds",
}

// BetArea  下注区域表结构体
type BetArea struct {
	ID               uint   `gorm:"primarykey;type:int unsigned;comment:ID"`
	GameCategoryCode string `gorm:"size:32;not null;comment:游戏分类"`
	AreaId           string `gorm:"uniqueindex;size:10;not null;not null;comment:区域ID"`
	AreaKey          string `gorm:"size:50;not null;default:'';comment:区域标识"`
	AreaName         string `gorm:"size:50;not null;default:'';comment:区域名称"`
	Description      string `gorm:"size:100;not null;default:'';comment:下注区域说明"`
	Odds             string `gorm:"size:100;not null;default:'';comment:赔率"`
	MinBetAmount     uint   `gorm:"not null;default:0;comment:最小投注金额"`
	MaxBetAmount     uint   `gorm:"not null;default:0;comment:最大投注金额"`
	Digits           uint   `gorm:"type:int unsigned;not null;default:1;comment:投注的最小单位倍数"`
}

// RegisterTable 数据迁移
func (m *BetArea) RegisterTable() error {
	if !global.DB.Migrator().HasTable(m) {
		if err := global.DB.Set("gorm:table_options", "COMMENT='下注区域表'").Migrator().CreateTable(m); err != nil {
			return err
		}

		// 百家乐
		bacx := []BetArea{
			{GameCategoryCode: "BACX", AreaId: "1001", AreaKey: "Bet_Bacx_Player", AreaName: "闲", Description: "闲家点数比庄家高获胜", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1002", AreaKey: "Bet_Bacx_Banker", AreaName: "庄", Description: "庄家点数比闲家高获胜", Odds: "1:0.95", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1003", AreaKey: "Bet_Bacx_Tie", AreaName: "和", Description: "庄家和闲家点数相同", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1004", AreaKey: "Bet_Bacx_PPair", AreaName: "闲对子", Description: "闲家前两张牌为对子", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 45000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1005", AreaKey: "Bet_Bacx_BPair", AreaName: "庄对子", Description: "庄家前两张牌为对子", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 45000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1006", AreaKey: "Bet_Bacx_BankerNocomm", AreaName: "庄免佣", Description: "庄家胜出但不收取5%佣金", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1007", AreaKey: "Bet_Bacx_SuperSix", AreaName: "幸运六", Description: "庄家以6点胜出", Odds: "1:12", MinBetAmount: 100, MaxBetAmount: 2500, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1008", AreaKey: "Bet_Bacx_AnyPair", AreaName: "任意对子", Description: "庄家或闲家前两张牌为对子", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1009", AreaKey: "Bet_Bacx_PerfectPair", AreaName: "完美对子", Description: "庄家或闲家前两张牌为同花色对子", Odds: "1:25", MinBetAmount: 100, MaxBetAmount: 20000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1010", AreaKey: "Bet_Bacx_Big", AreaName: "大", Description: "总发牌数为5或6张", Odds: "1:0.5", MinBetAmount: 1000, MaxBetAmount: 350000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1011", AreaKey: "Bet_Bacx_Small", AreaName: "小", Description: "总发牌数为4张", Odds: "1:1.5", MinBetAmount: 1000, MaxBetAmount: 350000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1012", AreaKey: "Bet_Bacx_PlayerBonus", AreaName: "闲龙宝", Description: "闲家胜出且与庄家点数差距大（如自然赢或大胜）", Odds: "1:1~1:30", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "BACX", AreaId: "1013", AreaKey: "Bet_Bacx_BankerBonus", AreaName: "庄龙宝", Description: "庄家胜出且与闲家点数差距大（如自然赢或大胜）", Odds: "1:1~1:30", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
		}

		// 龙虎
		dTX := []BetArea{
			{GameCategoryCode: "DTX", AreaId: "1014", AreaKey: "Bet_Dtx_Dragon", AreaName: "龙", Description: "龙牌点数大于虎牌点数", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "DTX", AreaId: "1015", AreaKey: "Bet_Dtx_Tiger", AreaName: "闲", Description: "虎牌点数大于龙牌点数", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "DTX", AreaId: "1016", AreaKey: "Bet_Dtx_Tie", AreaName: "和", Description: "龙牌点数等于虎牌点数", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "DTX", AreaId: "1017", AreaKey: "Bet_Dtx_TiePair", AreaName: "同花和", Description: "龙牌与虎牌点数相同且花色相同", Odds: "1:50", MinBetAmount: 100, MaxBetAmount: 2500, Digits: 100},
		}

		// 轮盘
		rOT := []BetArea{
			// 单数(1区) -- 36       1赔35         ( rgnMain: 0, rgnsub 0--- 35 )
			{GameCategoryCode: "ROT", AreaId: "1100", AreaKey: "Bet_Rot_Single0", AreaName: "直注0", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1101", AreaKey: "Bet_Rot_Single1", AreaName: "直注1", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1102", AreaKey: "Bet_Rot_Single2", AreaName: "直注2", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1103", AreaKey: "Bet_Rot_Single3", AreaName: "直注3", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1104", AreaKey: "Bet_Rot_Single4", AreaName: "直注4", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1105", AreaKey: "Bet_Rot_Single5", AreaName: "直注5", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1106", AreaKey: "Bet_Rot_Single6", AreaName: "直注6", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1107", AreaKey: "Bet_Rot_Single7", AreaName: "直注7", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1108", AreaKey: "Bet_Rot_Single8", AreaName: "直注8", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1109", AreaKey: "Bet_Rot_Single9", AreaName: "直注9", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1110", AreaKey: "Bet_Rot_Single10", AreaName: "直注10", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1111", AreaKey: "Bet_Rot_Single11", AreaName: "直注11", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1112", AreaKey: "Bet_Rot_Single12", AreaName: "直注12", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1113", AreaKey: "Bet_Rot_Single13", AreaName: "直注13", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1114", AreaKey: "Bet_Rot_Single14", AreaName: "直注14", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1115", AreaKey: "Bet_Rot_Single15", AreaName: "直注15", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1116", AreaKey: "Bet_Rot_Single16", AreaName: "直注16", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1117", AreaKey: "Bet_Rot_Single17", AreaName: "直注17", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1118", AreaKey: "Bet_Rot_Single18", AreaName: "直注18", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1119", AreaKey: "Bet_Rot_Single19", AreaName: "直注19", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1120", AreaKey: "Bet_Rot_Single20", AreaName: "直注20", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1121", AreaKey: "Bet_Rot_Single21", AreaName: "直注21", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1122", AreaKey: "Bet_Rot_Single22", AreaName: "直注22", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1123", AreaKey: "Bet_Rot_Single23", AreaName: "直注23", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1124", AreaKey: "Bet_Rot_Single24", AreaName: "直注24", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1125", AreaKey: "Bet_Rot_Single25", AreaName: "直注25", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1126", AreaKey: "Bet_Rot_Single26", AreaName: "直注26", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1127", AreaKey: "Bet_Rot_Single27", AreaName: "直注27", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1128", AreaKey: "Bet_Rot_Single28", AreaName: "直注28", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1129", AreaKey: "Bet_Rot_Single29", AreaName: "直注29", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1130", AreaKey: "Bet_Rot_Single30", AreaName: "直注30", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1131", AreaKey: "Bet_Rot_Single31", AreaName: "直注31", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1132", AreaKey: "Bet_Rot_Single32", AreaName: "直注32", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1133", AreaKey: "Bet_Rot_Single33", AreaName: "直注33", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1134", AreaKey: "Bet_Rot_Single34", AreaName: "直注34", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1135", AreaKey: "Bet_Rot_Single35", AreaName: "直注35", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1136", AreaKey: "Bet_Rot_Single36", AreaName: "直注36", Description: "", Odds: "1:35", MinBetAmount: 100, MaxBetAmount: 14000, Digits: 100},
			//双数(2区)  1赔17   ( rgnMain:1   rgnsub: 0 -- 59   压双数 )
			{GameCategoryCode: "ROT", AreaId: "1137", AreaKey: "Bet_Rot_Double1or2", AreaName: "分注1,2", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1138", AreaKey: "Bet_Rot_Double2or3", AreaName: "分注2,3", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1139", AreaKey: "Bet_Rot_Double4or5", AreaName: "分注4,5", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1140", AreaKey: "Bet_Rot_Double5or6", AreaName: "分注5,6", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1141", AreaKey: "Bet_Rot_Double7or8", AreaName: "分注7,8", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1142", AreaKey: "Bet_Rot_Double8or9", AreaName: "分注8,9", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1143", AreaKey: "Bet_Rot_Double10or11", AreaName: "分注10,11", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1144", AreaKey: "Bet_Rot_Double11or12", AreaName: "分注11,12", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1145", AreaKey: "Bet_Rot_Double13or14", AreaName: "分注13,14", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1146", AreaKey: "Bet_Rot_Double14or15", AreaName: "分注14,15", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1147", AreaKey: "Bet_Rot_Double16or17", AreaName: "分注16,17", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1148", AreaKey: "Bet_Rot_Double17or18", AreaName: "分注17,18", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1149", AreaKey: "Bet_Rot_Double19or20", AreaName: "分注19,20", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1150", AreaKey: "Bet_Rot_Double20or21", AreaName: "分注20,21", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1151", AreaKey: "Bet_Rot_Double22or23", AreaName: "分注22,23", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1152", AreaKey: "Bet_Rot_Double23or24", AreaName: "分注23,24", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1153", AreaKey: "Bet_Rot_Double25or26", AreaName: "分注25,26", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1154", AreaKey: "Bet_Rot_Double26or27", AreaName: "分注26,27", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1155", AreaKey: "Bet_Rot_Double28or29", AreaName: "分注28,29", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1156", AreaKey: "Bet_Rot_Double29or30", AreaName: "分注29,30", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1157", AreaKey: "Bet_Rot_Double31or32", AreaName: "分注31,32", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1158", AreaKey: "Bet_Rot_Double32or33", AreaName: "分注32,33", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1159", AreaKey: "Bet_Rot_Double34or35", AreaName: "分注34,35", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1160", AreaKey: "Bet_Rot_Double35or36", AreaName: "分注35,36", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1161", AreaKey: "Bet_Rot_Double1or4", AreaName: "分注1,4", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1162", AreaKey: "Bet_Rot_Double4or7", AreaName: "分注4,7", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1163", AreaKey: "Bet_Rot_Double7or10", AreaName: "分注7,10", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1164", AreaKey: "Bet_Rot_Double10or13", AreaName: "分注10,13", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1165", AreaKey: "Bet_Rot_Double13or16", AreaName: "分注13,16", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1166", AreaKey: "Bet_Rot_Double16or19", AreaName: "分注16,19", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1167", AreaKey: "Bet_Rot_Double19or22", AreaName: "分注19,22", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1168", AreaKey: "Bet_Rot_Double22or25", AreaName: "分注22,25", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1169", AreaKey: "Bet_Rot_Double25or28", AreaName: "分注25,28", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1170", AreaKey: "Bet_Rot_Double28or31", AreaName: "分注28,31", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1171", AreaKey: "Bet_Rot_Double31or34", AreaName: "分注31,34", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1172", AreaKey: "Bet_Rot_Double2or5", AreaName: "分注2,5", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1173", AreaKey: "Bet_Rot_Double5or8", AreaName: "分注5,8", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1174", AreaKey: "Bet_Rot_Double8or11", AreaName: "分注8,11", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1175", AreaKey: "Bet_Rot_Double11or14", AreaName: "分注11,14", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1176", AreaKey: "Bet_Rot_Double14or17", AreaName: "分注14,17", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1177", AreaKey: "Bet_Rot_Double17or20", AreaName: "分注17,20", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1178", AreaKey: "Bet_Rot_Double20or23", AreaName: "分注20,23", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1179", AreaKey: "Bet_Rot_Double23or26", AreaName: "分注23,26", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1180", AreaKey: "Bet_Rot_Double26or29", AreaName: "分注26,29", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1181", AreaKey: "Bet_Rot_Double29or32", AreaName: "分注29,32", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1182", AreaKey: "Bet_Rot_Double32or35", AreaName: "分注32,35", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1183", AreaKey: "Bet_Rot_Double3or6", AreaName: "分注3,6", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1184", AreaKey: "Bet_Rot_Double6or9", AreaName: "分注6,9", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1185", AreaKey: "Bet_Rot_Double9or12", AreaName: "分注9,12", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1186", AreaKey: "Bet_Rot_Double12or15", AreaName: "分注12,15", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1187", AreaKey: "Bet_Rot_Double15or18", AreaName: "分注15,18", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1188", AreaKey: "Bet_Rot_Double18or21", AreaName: "分注18,21", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1189", AreaKey: "Bet_Rot_Double21or24", AreaName: "分注21,24", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1190", AreaKey: "Bet_Rot_Double24or27", AreaName: "分注24,27", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1191", AreaKey: "Bet_Rot_Double27or30", AreaName: "分注27,30", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1192", AreaKey: "Bet_Rot_Double30or33", AreaName: "分注30,33", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1193", AreaKey: "Bet_Rot_Double33or36", AreaName: "分注33,36", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1194", AreaKey: "Bet_Rot_Double0or1", AreaName: "分注0,1", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1195", AreaKey: "Bet_Rot_Double0or2", AreaName: "分注0,2", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1196", AreaKey: "Bet_Rot_Double0or3", AreaName: "分注0,3", Description: "", Odds: "1:17", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			//三数(3区)  1赔11   ( rgnMain:2   rgnsub: 0 -- 13   压三数 )
			{GameCategoryCode: "ROT", AreaId: "1197", AreaKey: "Bet_Rot_Three1to3", AreaName: "街注1,2,3", Description: "压三个数：1到3", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1198", AreaKey: "Bet_Rot_Three4to6", AreaName: "街注4,5,6", Description: "压三个数：4到6", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1199", AreaKey: "Bet_Rot_Three7to9", AreaName: "街注7,,89", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1200", AreaKey: "Bet_Rot_Three10to12", AreaName: "街注10,11,12", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1201", AreaKey: "Bet_Rot_Three13to15", AreaName: "街注13,14,15", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1202", AreaKey: "Bet_Rot_Three16to18", AreaName: "街注16,17,18", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1203", AreaKey: "Bet_Rot_Three19to21", AreaName: "街注19,20,21", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1204", AreaKey: "Bet_Rot_Three22to24", AreaName: "街注22,23,24", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1205", AreaKey: "Bet_Rot_Three25to27", AreaName: "街注25,26,27", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1206", AreaKey: "Bet_Rot_Three28to30", AreaName: "街注28,29,30", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1207", AreaKey: "Bet_Rot_Three31to33", AreaName: "街注31,32,33", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1208", AreaKey: "Bet_Rot_Three34to36", AreaName: "街注34,35,36", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1209", AreaKey: "Bet_Rot_Three0to2", AreaName: "三数0,1,2", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1210", AreaKey: "Bet_Rot_Three0or2or3", AreaName: "三数0,2,3", Description: "", Odds: "1:11", MinBetAmount: 500, MaxBetAmount: 450, Digits: 100},
			//四数(4区)  1赔8 ( rgnMain:3   rgnsub: 0 -- 22   压四个数的 )
			{GameCategoryCode: "ROT", AreaId: "1211", AreaKey: "Bet_Rot_Four1or2or4or5", AreaName: "角注0,2,4,5", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1212", AreaKey: "Bet_Rot_Four2or3or5or6", AreaName: "角注2,3,5,6", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1213", AreaKey: "Bet_Rot_Four4or5or7or8", AreaName: "角注4,5,7,8", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1214", AreaKey: "Bet_Rot_Four5or6or8or9", AreaName: "角注5,6,8,9", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1215", AreaKey: "Bet_Rot_Four7or8or10or11", AreaName: "角注7,8,10,11", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1216", AreaKey: "Bet_Rot_Four8or9or11or12", AreaName: "角注8,9,11,12", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1217", AreaKey: "Bet_Rot_Four10or11or13or14", AreaName: "角注10,11,13,14", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1218", AreaKey: "Bet_Rot_Four11or12or14or15", AreaName: "角注11,12,14,15", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1219", AreaKey: "Bet_Rot_Four13or14or16or17", AreaName: "角注13,14,16,17", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1220", AreaKey: "Bet_Rot_Four14or15or17or18", AreaName: "角注14,15,17,18", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1221", AreaKey: "Bet_Rot_Four16or17or19or20", AreaName: "角注16,17,19,20", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1222", AreaKey: "Bet_Rot_Four17or18r20or21", AreaName: "角注17,18,20,21", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1223", AreaKey: "Bet_Rot_Four19or2or22or23", AreaName: "角注19,20,22,23", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1224", AreaKey: "Bet_Rot_Four20or21or23or24", AreaName: "角注20,21,23,24", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1225", AreaKey: "Bet_Rot_Four22or23or25or26", AreaName: "角注22,23,25,26", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1226", AreaKey: "Bet_Rot_Four23or24or26or27", AreaName: "角注23,24,26,27", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1227", AreaKey: "Bet_Rot_Four25or26or28or29", AreaName: "角注25,26,28,29", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1228", AreaKey: "Bet_Rot_Four26or27or29or30", AreaName: "角注26,27,29,30", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1229", AreaKey: "Bet_Rot_Four28or29or31or32", AreaName: "角注28,29,31,32", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1230", AreaKey: "Bet_Rot_Four29or30or32or33", AreaName: "角注29,30,32,33", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1231", AreaKey: "Bet_Rot_Four31or32or34or35", AreaName: "角注31,32,34,35", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1232", AreaKey: "Bet_Rot_Four32or33or35or36", AreaName: "角注32,33,35,36", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1233", AreaKey: "Bet_Rot_Four0or1or2or3", AreaName: "四个号码：0,1,2,3", Description: "", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			//六数(5区)	1赔5 ( rgnMain:4   rgnsub: 0 -- 10   //压6个数 )
			{GameCategoryCode: "ROT", AreaId: "1234", AreaKey: "Bet_Rot_Six1to6", AreaName: "线注1,2,3,4,5,6", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1235", AreaKey: "Bet_Rot_Six4to9", AreaName: "线注4,5,6,7,8,9", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1236", AreaKey: "Bet_Rot_Six7to12", AreaName: "线注7,8,9,10,11,12", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1237", AreaKey: "Bet_Rot_Six10to15", AreaName: "线注10,11,12,13,14,15", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1238", AreaKey: "Bet_Rot_Six13to18", AreaName: "线注13,14,15,16,17,18", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1239", AreaKey: "Bet_Rot_Six16to21", AreaName: "线注16,17,18,19,20,21", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1240", AreaKey: "Bet_Rot_Six19to24", AreaName: "线注19,20,21,22,23,24", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1241", AreaKey: "Bet_Rot_Six22to27", AreaName: "线注22,23,24,25,26,27", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1242", AreaKey: "Bet_Rot_Six25to30", AreaName: "线注25,26,27,28,29,30", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1243", AreaKey: "Bet_Rot_Six28to33", AreaName: "线注28,29,30,31,32,33", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1244", AreaKey: "Bet_Rot_Six31to36", AreaName: "线注31,32,33,34,35,36", Description: "", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			//12数(6区)	1赔2 ( rgnMain:5   rgnsub: 0 -- 5  压12个数  )
			{GameCategoryCode: "ROT", AreaId: "1245", AreaKey: "Bet_Rot_Twelve1rd12", AreaName: "列1", Description: "压12个数：1rd12", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1246", AreaKey: "Bet_Rot_Twelve2nd12", AreaName: "列2", Description: "压12个数：2nd12", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1247", AreaKey: "Bet_Rot_Twelve3rd12", AreaName: "列3", Description: "压12个数：3rd12", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1248", AreaKey: "Bet_Rot_Twelve2to1a", AreaName: "打一", Description: "压12个数：2to1a", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1249", AreaKey: "Bet_Rot_Twelve2to1b", AreaName: "打二", Description: "压12个数：2to1b", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1250", AreaKey: "Bet_Rot_Twelve2to1c", AreaName: "打三", Description: "压12个数：2to1c", Odds: "1:2", MinBetAmount: 1000, MaxBetAmount: 250000, Digits: 100},
			//红黑，小大，单双(7区)  1赔1 ( rgnMain:6   rgnsub: 0 -- 5  )
			{GameCategoryCode: "ROT", AreaId: "1251", AreaKey: "Bet_Rot_EighteenRed", AreaName: "红", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1252", AreaKey: "Bet_Rot_EighteenBlack", AreaName: "黑", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1253", AreaKey: "Bet_Rot_EighteenSmall", AreaName: "小", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1254", AreaKey: "Bet_Rot_EighteenBig", AreaName: "大", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1255", AreaKey: "Bet_Rot_EighteenSingle", AreaName: "单", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1256", AreaKey: "Bet_Rot_EighteenDouble", AreaName: "双", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
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
			{GameCategoryCode: "ROT", AreaId: "1257", AreaKey: "Bet_Rot_NeighbourTwo0", AreaName: "邻注0", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1258", AreaKey: "Bet_Rot_NeighbourTwo32", AreaName: "邻注32", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1259", AreaKey: "Bet_Rot_NeighbourTwo15", AreaName: "邻注15", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1260", AreaKey: "Bet_Rot_NeighbourTwo19", AreaName: "邻注19", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1261", AreaKey: "Bet_Rot_NeighbourTwo4", AreaName: "邻注4", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1262", AreaKey: "Bet_Rot_NeighbourTwo21", AreaName: "邻注21", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1263", AreaKey: "Bet_Rot_NeighbourTwo2", AreaName: "邻注2", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1264", AreaKey: "Bet_Rot_NeighbourTwo25", AreaName: "邻注25", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1265", AreaKey: "Bet_Rot_NeighbourTwo17", AreaName: "邻注17", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1266", AreaKey: "Bet_Rot_NeighbourTwo34", AreaName: "邻注34", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1267", AreaKey: "Bet_Rot_NeighbourTwo6", AreaName: "邻注6", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1268", AreaKey: "Bet_Rot_NeighbourTwo27", AreaName: "邻注27", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1269", AreaKey: "Bet_Rot_NeighbourTwo13", AreaName: "邻注13", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1270", AreaKey: "Bet_Rot_NeighbourTwo36", AreaName: "邻注36", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1271", AreaKey: "Bet_Rot_NeighbourTwo11", AreaName: "邻注11", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1272", AreaKey: "Bet_Rot_NeighbourTwo30", AreaName: "邻注30", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1273", AreaKey: "Bet_Rot_NeighbourTwo8", AreaName: "邻注8", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1274", AreaKey: "Bet_Rot_NeighbourTwo23", AreaName: "邻注23", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1275", AreaKey: "Bet_Rot_NeighbourTwo10", AreaName: "邻注10", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1276", AreaKey: "Bet_Rot_NeighbourTwo5", AreaName: "邻注5", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1277", AreaKey: "Bet_Rot_NeighbourTwo24", AreaName: "邻注24", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1278", AreaKey: "Bet_Rot_NeighbourTwo16", AreaName: "邻注16", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1279", AreaKey: "Bet_Rot_NeighbourTwo33", AreaName: "邻注33", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1280", AreaKey: "Bet_Rot_NeighbourTwo1", AreaName: "邻注1", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1281", AreaKey: "Bet_Rot_NeighbourTwo20", AreaName: "邻注20", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1282", AreaKey: "Bet_Rot_NeighbourTwo14", AreaName: "邻注14", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1283", AreaKey: "Bet_Rot_NeighbourTwo31", AreaName: "邻注31", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1284", AreaKey: "Bet_Rot_NeighbourTwo9", AreaName: "邻注9", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1285", AreaKey: "Bet_Rot_NeighbourTwo22", AreaName: "邻注22", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1286", AreaKey: "Bet_Rot_NeighbourTwo18", AreaName: "邻注18", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1287", AreaKey: "Bet_Rot_NeighbourTwo29", AreaName: "邻注29", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1288", AreaKey: "Bet_Rot_NeighbourTwo7", AreaName: "邻注7", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1289", AreaKey: "Bet_Rot_NeighbourTwo28", AreaName: "邻注28", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1290", AreaKey: "Bet_Rot_NeighbourTwo12", AreaName: "邻注12", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1291", AreaKey: "Bet_Rot_NeighbourTwo35", AreaName: "邻注35", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1292", AreaKey: "Bet_Rot_NeighbourTwo3", AreaName: "邻注3", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			{GameCategoryCode: "ROT", AreaId: "1293", AreaKey: "Bet_Rot_NeighbourTwo26", AreaName: "邻注26", Description: "", Odds: "1:35", MinBetAmount: 500, MaxBetAmount: 70000, Digits: 100},
			//Tiers，轮盘下角注(9区)
			//说明：0对面的区域，包括30,11,36,13,27,33,16,24,5,10,23,8
			//下注：筹码需要是6的倍数，其中包括5,8及10,11及13,16及23,24及27,30及33,36六个双数注
			//开奖，如果命中其中的某个数字，赔率固定是1陪17
			{GameCategoryCode: "ROT", AreaId: "1294", AreaKey: "Bet_Rot_Tiers", AreaName: "轮盘下角注", Description: "", Odds: "1:17", MinBetAmount: 3000, MaxBetAmount: 180000, Digits: 100},
			//Orphelins，轮盘孤注(10区)
			//说明：轮盘侧边的6,34,17,9,31,14,20,1
			//下注：筹码需要是5的倍数，其中包括了1一个单注，和6,9及14,17及17,20及31,34四个双数注
			//开奖：如果开1则赔率是1陪35，如果不是1则赔率是1陪17
			{GameCategoryCode: "ROT", AreaId: "1295", AreaKey: "Bet_Rot_Orphelins", AreaName: "轮盘孤注", Description: "", Odds: "1:17", MinBetAmount: 2500, MaxBetAmount: 350000, Digits: 100},
			//Voisins，零旁注上角(11区)
			//说明：轮盘上0边上的25,2,21,4,19,15,32,0,26,3,35,12,28,7,29,18,22
			//下注：筹码需要是9的倍数，其中包括了0,2,3一个三数注(注意要两个筹码)、25,26,28,29一个四数注(注意要两个筹码)、4,7和12,15等5个双数注，加起来一共需要9个筹码
			//开奖：如果开0或2或3赔率是1陪11，如果开25或26或28或29赔率是1陪8，开其他数则赔率是1陪17
			{GameCategoryCode: "ROT", AreaId: "1296", AreaKey: "Bet_Rot_Voisins", AreaName: "零旁注上角", Description: "", Odds: "1:8~1:17", MinBetAmount: 4500, MaxBetAmount: 585000, Digits: 100},
			//Jeu0，轮上零旁(12区)
			//说明：轮盘上0边上的15,32,0,26,3,35,12
			//下注：筹码需要是4的倍数，其中包括了26一个单注，0,3双数注，12,15双数注，32,35双数注
			//开奖：如果开26赔率是1陪35，如果不是26则是1陪17
			{GameCategoryCode: "ROT", AreaId: "1297", AreaKey: "Bet_Rot_Jeu0", AreaName: "轮上零旁", Description: "", Odds: "1:35~1:17", MinBetAmount: 400, MaxBetAmount: 56000, Digits: 100},
		}

		// 筛宝
		sICBO := []BetArea{
			{GameCategoryCode: "SICBO", AreaId: "1500", AreaKey: "Bet_Sicbo_One_Big", AreaName: "大", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1501", AreaKey: "Bet_Sicbo_One_Small", AreaName: "小", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1502", AreaKey: "Bet_Sicbo_One_Odd", AreaName: "单", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1503", AreaKey: "Bet_Sicbo_One_Even", AreaName: "双", Description: "", Odds: "1:1", MinBetAmount: 2000, MaxBetAmount: 500000, Digits: 100},
			// 特定双数
			{GameCategoryCode: "SICBO", AreaId: "1504", AreaKey: "Bet_Sicbo_TwoNum_1And2", AreaName: "牌九式12", Description: "特定双数1和2", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1505", AreaKey: "Bet_Sicbo_TwoNum_1And3", AreaName: "牌九式13", Description: "特定双数1和3", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1506", AreaKey: "Bet_Sicbo_TwoNum_1And4", AreaName: "牌九式14", Description: "特定双数1和4", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1507", AreaKey: "Bet_Sicbo_TwoNum_1And5", AreaName: "牌九式15", Description: "特定双数1和5", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1508", AreaKey: "Bet_Sicbo_TwoNum_1And6", AreaName: "牌九式16", Description: "特定双数1和6", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1509", AreaKey: "Bet_Sicbo_TwoNum_2And3", AreaName: "牌九式23", Description: "特定双数2和3", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1510", AreaKey: "Bet_Sicbo_TwoNum_2And4", AreaName: "牌九式24", Description: "特定双数2和4", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1511", AreaKey: "Bet_Sicbo_TwoNum_2And5", AreaName: "牌九式25", Description: "特定双数2和5", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1512", AreaKey: "Bet_Sicbo_TwoNum_2And6", AreaName: "牌九式26", Description: "特定双数2和6", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1513", AreaKey: "Bet_Sicbo_TwoNum_3And4", AreaName: "牌九式34", Description: "特定双数3和4", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1514", AreaKey: "Bet_Sicbo_TwoNum_3And5", AreaName: "牌九式35", Description: "特定双数3和5", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1515", AreaKey: "Bet_Sicbo_TwoNum_3And6", AreaName: "牌九式36", Description: "特定双数3和6", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1516", AreaKey: "Bet_Sicbo_TwoNum_4And5", AreaName: "牌九式45", Description: "特定双数4和5", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1517", AreaKey: "Bet_Sicbo_TwoNum_4And6", AreaName: "牌九式46", Description: "特定双数4和6", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1518", AreaKey: "Bet_Sicbo_TwoNum_5And6", AreaName: "牌九式56", Description: "特定双数5和6", Odds: "1:5", MinBetAmount: 500, MaxBetAmount: 100000, Digits: 100},
			// 特定三数
			{GameCategoryCode: "SICBO", AreaId: "1519", AreaKey: "Bet_Sicbo_ThreeSame_1All", AreaName: "围骰1", Description: "特定三数都是1", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1520", AreaKey: "Bet_Sicbo_ThreeSame_2All", AreaName: "围骰2", Description: "特定三数都是2", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1521", AreaKey: "Bet_Sicbo_ThreeSame_3All", AreaName: "围骰3", Description: "特定三数都是3", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1522", AreaKey: "Bet_Sicbo_ThreeSame_4All", AreaName: "围骰4", Description: "特定三数都是4", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1523", AreaKey: "Bet_Sicbo_ThreeSame_5All", AreaName: "围骰5", Description: "特定三数都是5", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1524", AreaKey: "Bet_Sicbo_ThreeSame_6All", AreaName: "围骰6", Description: "特定三数都是6", Odds: "1:180", MinBetAmount: 100, MaxBetAmount: 3000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1525", AreaKey: "Bet_Sicbo_ThreeSame_Any", AreaName: "全围", Description: "任意三数", Odds: "1:30", MinBetAmount: 100, MaxBetAmount: 17000, Digits: 100},
			// 点数和
			{GameCategoryCode: "SICBO", AreaId: "1526", AreaKey: "Bet_Sicbo_SumPoint_4", AreaName: "和值4", Description: "点数和4", Odds: "1:60", MinBetAmount: 100, MaxBetAmount: 8500, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1527", AreaKey: "Bet_Sicbo_SumPoint_5", AreaName: "和值5", Description: "点数和5", Odds: "1:30", MinBetAmount: 100, MaxBetAmount: 17000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1528", AreaKey: "Bet_Sicbo_SumPoint_6", AreaName: "和值6", Description: "点数和6", Odds: "1:18", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1529", AreaKey: "Bet_Sicbo_SumPoint_7", AreaName: "和值7", Description: "点数和7", Odds: "1:12", MinBetAmount: 500, MaxBetAmount: 40000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1530", AreaKey: "Bet_Sicbo_SumPoint_8", AreaName: "和值8", Description: "点数和8", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 5000000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1531", AreaKey: "Bet_Sicbo_SumPoint_9", AreaName: "和值9", Description: "点数和9", Odds: "1:6", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1532", AreaKey: "Bet_Sicbo_SumPoint_10", AreaName: "和值10", Description: "点数和10", Odds: "1:6", MinBetAmount: 500, MaxBetAmount: 85000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1533", AreaKey: "Bet_Sicbo_SumPoint_11", AreaName: "和值11", Description: "点数和11", Odds: "1:6", MinBetAmount: 500, MaxBetAmount: 85000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1534", AreaKey: "Bet_Sicbo_SumPoint_12", AreaName: "和值12", Description: "点数和12", Odds: "1:6", MinBetAmount: 500, MaxBetAmount: 85000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1535", AreaKey: "Bet_Sicbo_SumPoint_13", AreaName: "和值13", Description: "点数和13", Odds: "1:8", MinBetAmount: 500, MaxBetAmount: 65000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1536", AreaKey: "Bet_Sicbo_SumPoint_14", AreaName: "和值14", Description: "点数和14", Odds: "1:12", MinBetAmount: 500, MaxBetAmount: 40000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1537", AreaKey: "Bet_Sicbo_SumPoint_15", AreaName: "和值15", Description: "点数和15", Odds: "1:18", MinBetAmount: 500, MaxBetAmount: 30000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1538", AreaKey: "Bet_Sicbo_SumPoint_16", AreaName: "和值16", Description: "点数和16", Odds: "1:30", MinBetAmount: 100, MaxBetAmount: 17000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1539", AreaKey: "Bet_Sicbo_SumPoint_17", AreaName: "和值17", Description: "点数和17", Odds: "1:60", MinBetAmount: 100, MaxBetAmount: 8500, Digits: 100},
			// 一对组合
			{GameCategoryCode: "SICBO", AreaId: "1540", AreaKey: "Bet_Sicbo_PairSame_1", AreaName: "对子1", Description: "一对组合点数1", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1541", AreaKey: "Bet_Sicbo_PairSame_2", AreaName: "对子2", Description: "一对组合点数2", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1542", AreaKey: "Bet_Sicbo_PairSame_3", AreaName: "对子3", Description: "一对组合点数3", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1543", AreaKey: "Bet_Sicbo_PairSame_4", AreaName: "对子4", Description: "一对组合点数4", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1544", AreaKey: "Bet_Sicbo_PairSame_5", AreaName: "对子5", Description: "一对组合点数5", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1545", AreaKey: "Bet_Sicbo_PairSame_6", AreaName: "对子6", Description: "一对组合点数6", Odds: "1:10", MinBetAmount: 500, MaxBetAmount: 50000, Digits: 100},
			// 数字
			{GameCategoryCode: "SICBO", AreaId: "1546", AreaKey: "Bet_Sicbo_NumPoint_1", AreaName: "单点1", Description: "数字点数1", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1547", AreaKey: "Bet_Sicbo_NumPoint_2", AreaName: "单点12", Description: "数字点数2", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1548", AreaKey: "Bet_Sicbo_NumPoint_3", AreaName: "单点13", Description: "数字点数3", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1549", AreaKey: "Bet_Sicbo_NumPoint_4", AreaName: "单点14", Description: "数字点数4", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1550", AreaKey: "Bet_Sicbo_NumPoint_5", AreaName: "单点15", Description: "数字点数5", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
			{GameCategoryCode: "SICBO", AreaId: "1551", AreaKey: "Bet_Sicbo_NumPoint_6", AreaName: "单点16", Description: "数字点数6", Odds: "1:1~1:3", MinBetAmount: 500, MaxBetAmount: 170000, Digits: 100},
		}

		// 番摊
		Ftan := []BetArea{
			{GameCategoryCode: "FTAN", AreaId: "1600", AreaKey: "Bet_Ftan_Small", AreaName: "小", Description: "开[1,2]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1601", AreaKey: "Bet_Ftan_Big", AreaName: "大", Description: "开[3,4]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1602", AreaKey: "Bet_Ftan_Single", AreaName: "单", Description: "开[1,3]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1603", AreaKey: "Bet_Ftan_Double", AreaName: "双", Description: "开[2,4]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1604", AreaKey: "Bet_Ftan_1Fan", AreaName: "1", Description: "开[1]赢", Odds: "1:2.85", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1605", AreaKey: "Bet_Ftan_2Fan", AreaName: "2", Description: "开[2]赢", Odds: "1:2.85", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1606", AreaKey: "Bet_Ftan_3Fan", AreaName: "3", Description: "开[3]赢", Odds: "1:2.85", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1607", AreaKey: "Bet_Ftan_4Fan", AreaName: "4", Description: "开[4]赢", Odds: "1:2.85", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1608", AreaKey: "Bet_Ftan_12Jiao", AreaName: "1-2角", Description: "开[1,2]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1609", AreaKey: "Bet_Ftan_23Jiao", AreaName: "2-3角", Description: "开[2,3]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1610", AreaKey: "Bet_Ftan_34Jiao", AreaName: "3-4角", Description: "开[3,4]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1611", AreaKey: "Bet_Ftan_41Jiao", AreaName: "4-1角", Description: "开[4,1]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1612", AreaKey: "Bet_Ftan_1Nian2", AreaName: "1念2", Description: "开[1]赢，开[2]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1613", AreaKey: "Bet_Ftan_1Nian3", AreaName: "1念3", Description: "开[1]赢，开[3]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1614", AreaKey: "Bet_Ftan_1Nian4", AreaName: "1念4", Description: "开[1]赢，开[4]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1615", AreaKey: "Bet_Ftan_2Nian1", AreaName: "2念1", Description: "开[2]赢，开[1]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1616", AreaKey: "Bet_Ftan_2Nian3", AreaName: "2念3", Description: "开[2]赢，开[3]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1617", AreaKey: "Bet_Ftan_2Nian4", AreaName: "2念4", Description: "开[2]赢，开[4]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1618", AreaKey: "Bet_Ftan_3Nian1", AreaName: "3念1", Description: "开[3]赢，开[1]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1619", AreaKey: "Bet_Ftan_3Nian2", AreaName: "3念2", Description: "开[3]赢，开[2]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1620", AreaKey: "Bet_Ftan_3Nian4", AreaName: "3念4", Description: "开[3]赢，开[4]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1621", AreaKey: "Bet_Ftan_4Nian1", AreaName: "4念1", Description: "开[4]赢，开[1]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1622", AreaKey: "Bet_Ftan_4Nian2", AreaName: "4念2", Description: "开[4]赢，开[2]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1623", AreaKey: "Bet_Ftan_4Nian3", AreaName: "4念3", Description: "开[4]赢，开[3]和", Odds: "1.90:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1624", AreaKey: "Bet_Ftan_321ThreeRed", AreaName: "三门3-2-1", Description: "开[3,2,1]赢", Odds: "1:0.316667", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1625", AreaKey: "Bet_Ftan_214ThreeRed", AreaName: "三门2-1-4", Description: "开[2,1,4]赢", Odds: "1:0.316667", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1626", AreaKey: "Bet_Ftan_432ThreeRed", AreaName: "三门4-3-2", Description: "开[4,3,2]赢", Odds: "1:0.316667", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FTAN", AreaId: "1627", AreaKey: "Bet_Ftan_143ThreeRed", AreaName: "三门1-4-3", Description: "开[1,4,3]赢", Odds: "1:0.316667", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
		}

		// 色碟
		Sdie := []BetArea{
			{GameCategoryCode: "SDIE", AreaId: "1700", AreaKey: "Bet_Sdie_Single", AreaName: "单", Description: "开[3红和1白,1红3白]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1701", AreaKey: "Bet_Sdie_Double", AreaName: "双", Description: "开[0红4白,2紅2白,4紅0白]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1702", AreaKey: "Bet_Sdie_Small", AreaName: "小", Description: "开[0红4白,1红3白]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1703", AreaKey: "Bet_Sdie_Big", AreaName: "大", Description: "开[3红1白,4红0白]赢", Odds: "1:0.95", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1704", AreaKey: "Bet_Sdie_0Red_4White", AreaName: "0红4白", Description: "开[0红4白]赢", Odds: "1:12", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1705", AreaKey: "Bet_Sdie_4Red_0White", AreaName: "4红0白", Description: "开[4红0白]赢", Odds: "1:12", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1706", AreaKey: "Bet_Sdie_1Red_3White", AreaName: "1红3白", Description: "开[1红3白]赢", Odds: "1:2.6", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1707", AreaKey: "Bet_Sdie_3Red_1White", AreaName: "3红1白", Description: "开[3红1白]赢", Odds: "1:2.6", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1708", AreaKey: "Bet_Sdie_2Red_2White", AreaName: "2红2白", Description: "开[2红2白]赢", Odds: "1:1.5", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "SDIE", AreaId: "1709", AreaKey: "Bet_Sdie_4White_Or_4Red", AreaName: "0红4白或4红0白", Description: "开[0红4白,4红0白]赢", Odds: "1:6.5", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
		}

		// 鱼虾蟹
		Fsc := []BetArea{
			{GameCategoryCode: "FSC", AreaId: "1800", AreaKey: "Bet_Fsc_Single_Point1", AreaName: "鱼", Description: "鱼", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FSC", AreaId: "1801", AreaKey: "Bet_Fsc_Single_Point2", AreaName: "虾", Description: "虾", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FSC", AreaId: "1802", AreaKey: "Bet_Fsc_Single_Point3", AreaName: "葫芦", Description: "葫芦", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FSC", AreaId: "1803", AreaKey: "Bet_Fsc_Single_Point4", AreaName: "鹿", Description: "鹿", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FSC", AreaId: "1804", AreaKey: "Bet_Fsc_Single_Point5", AreaName: "蟹", Description: "蟹", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
			{GameCategoryCode: "FSC", AreaId: "1805", AreaKey: "Bet_Fsc_Single_Point6", AreaName: "鸡", Description: "鸡", Odds: "1:1", MinBetAmount: 10, MaxBetAmount: 10000, Digits: 100},
		}
		// 填充数据
		data := slice.Concat(bacx, dTX, rOT, sICBO, Ftan, Sdie, Fsc)

		if err := global.DB.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}
