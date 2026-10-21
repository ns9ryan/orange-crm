package initialize

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"go.uber.org/zap"
	"orange-crm/config"
	"orange-crm/global"
	"orange-crm/utils"
)

func I18n() config.I18n {
	i18n := gi18n.New()

	lang := global.CONFIG.Server.Lang
	if len(lang) != 0 {
		lang = "zh"
	}
	i18n.SetLanguage(lang)

	i18n.SetPath(utils.GetDirByName("i18n"))

	global.LOG.Error("path:", zap.Any("i18n", utils.GetDirByName("i18n")))

	return config.I18n{
		Manager: i18n,
		Context: gi18n.WithLanguage(context.TODO(), lang),
	}
}
