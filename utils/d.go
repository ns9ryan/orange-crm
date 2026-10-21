package utils

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"go.uber.org/zap"
	"orange-crm/global"
)

// DD 通用日志工具
func DD(data any) {
	global.LOG.Debug("调试", zap.Any("data", data))
}

// TT 国际化
func TT(msg string, params ...any) string {
	if len(params) == 0 {
		return global.I18N.Manager.T(global.I18N.Context, msg)
	} else {
		return global.I18N.Manager.Tf(global.I18N.Context, msg, params...)
	}
}

// RandNumeralOrLetter 生成给定长度的随机字符串（数字+字母)
func RandNumeralOrLetter(length int) string {
	return random.RandNumeralOrLetter(length)
}

func FF(data any) {
	fmt.Printf("data:%+v\n", data)
}

// Btoi 布尔转int ==================================
func Btoi(b bool) uint {
	if b {
		return 1
	}
	return 0
}
