package http

import (
	"captcha-service/internal/config"
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type LifecycleDependency struct {
	fx.In
	Lc     fx.Lifecycle
	App    *fiber.App
	Config *config.Config
	Logger *zap.SugaredLogger
}

func httpLifecycle(dep LifecycleDependency) {
	lc, app, log, cfg := dep.Lc, dep.App, dep.Logger, dep.Config
	lc.Append(fx.StartStopHook(
		func() {
			go func() {
				localAddress := "http://127.0.0.1:" + cfg.Http.GetPort()
				log.Debugf("Http Listening on %s", localAddress)
				lo.Must0(app.Listen(
					":"+cfg.Http.GetPort(),
					fiber.ListenConfig{
						EnablePrintRoutes: true,
						ShutdownTimeout:   1000,
						EnablePrefork:     cfg.Http.Prefork,
					},
				), "sproxy start fail")
			}()
		},
		func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	),
	)
}
