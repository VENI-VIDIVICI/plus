package bootstrap

import (
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
)

func SetupLogger() {
	logger.InitLogger(config.Get("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.Get("log.type"),
		config.Get("log.level"),
	)
}
