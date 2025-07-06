package cmd

import (
	"captcha-service/internal/config"
	"captcha-service/internal/endpoint"
	"captcha-service/internal/http"
	"captcha-service/internal/logger"
	"captcha-service/internal/store"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func container() *fx.App {
	return fx.New(
		config.Module,
		logger.Module,
		store.Module,
		endpoint.Module,
		http.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			fxLogger := &fxevent.ZapLogger{Logger: log}
			fxLogger.UseLogLevel(zapcore.DebugLevel)
			return fxLogger
		}),
	)
}
