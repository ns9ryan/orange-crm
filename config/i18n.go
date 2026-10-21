package config

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type I18n struct {
	Manager *gi18n.Manager
	Context context.Context
}
