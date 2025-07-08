package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"runtime/debug"
)

var Module = fx.Module("http",
	fx.Provide(
		newServer,
	),
	middlewareModule,
	fx.Invoke(httpLifecycle),
)

func newServer() *fiber.App {
	info, ok := debug.ReadBuildInfo()
	header := lo.Ternary(ok, "X-Captcha-"+info.Main.Version, "X-Captcha")
	app := fiber.New(
		fiber.Config{
			ViewsLayout:       "layout",
			PassLocalsToViews: true,
			Immutable:         true,
			StreamRequestBody: true,
			ServerHeader:      header,
			ReduceMemoryUsage: true,
		},
	)

	return app
}
