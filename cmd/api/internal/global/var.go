package global

import (
	"tiktok-api/internal/model/config"
	"tiktok-common/logger"
	"tiktok-common/token"
)

var (
	Settings config.Config // Public配置
	Maker    token.Maker
	Logger   *logger.Log // 日志
)
