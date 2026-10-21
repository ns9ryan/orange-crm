package utils

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// ParseDecimal 解析浮点数,返回整: 数部分,小数部分,换算成整数,小数倍率,错误, 示例: "3.14150000" =>
func ParseDecimal(numberStr string) (int64, float64, int64, int64, error) {
	// 先转成 float64（自动去掉无意义的 0）得到 3.1415
	floatValue, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 转回字符串（去掉无意义 0）得到 "3.1415"
	strValue := strconv.FormatFloat(floatValue, 'f', -1, 64)

	// 得到 ["3","1415"]
	parts := strings.Split(strValue, ".")
	// 整数部分 3 int64
	intPart, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var decimalDigits int64 = 1 // 小数倍率,10000
	var decimalAsInt int64 = 0  // 小数部分的数字表示,1415

	if len(parts) > 1 {
		decimalPlaces := len(parts[1])                         // 计算小数位数
		decimalDigits = int64(math.Pow10(decimalPlaces))       // 例如 10000
		decimalAsInt, err = strconv.ParseInt(parts[1], 10, 64) // 例如 1415
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	return intPart, float64(decimalAsInt) / float64(decimalDigits), intPart*decimalDigits + decimalAsInt, decimalDigits, nil
}

// ParseDecimalWithZero 解析浮点数,返回整: 数部分,小数部分,换算成整数,小数倍率,错误, 示例: "3.14150000" =>
func ParseDecimalWithZero(numberStr string) (int64, float64, int64, int64, error) {
	// 得到 ["3","14150000"]
	parts := strings.Split(numberStr, ".")
	// 整数部分 3 int64
	intPart, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var decimalDigits int64 = 1 // 小数倍率,10000
	var decimalAsInt int64 = 0  // 小数部分的数字表示,1415

	if len(parts) > 1 {
		decimalPlaces := len(parts[1])                         // 计算小数位数
		decimalDigits = int64(math.Pow10(decimalPlaces))       // 例如 10000
		decimalAsInt, err = strconv.ParseInt(parts[1], 10, 64) // 例如 1415
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	return intPart, float64(decimalAsInt) / float64(decimalDigits), intPart*decimalDigits + decimalAsInt, decimalDigits, nil
}

// SumStringInt 将多个字符串表示的整数累加，并返回字符串结果
func SumStringInt(values ...string) string {
	sum := big.NewInt(0)
	tmp := new(big.Int)
	for _, val := range values {
		if _, ok := tmp.SetString(val, 10); ok { // 第二个参数 10 表示十进制
			sum.Add(sum, tmp)
		}
	}
	return sum.String()
}

// SumString 将多个字符串数字累加，并返回字符串结果
func SumString(values ...string) string {
	sum := big.NewFloat(0)
	tmp := new(big.Float)
	for _, val := range values {
		if f, ok := tmp.SetString(val); ok {
			sum.Add(sum, f)
		}
	}
	return sum.Text('f', 6) // 保留6位小数（根据你实际需求改）
}
