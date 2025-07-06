package http

import (
	"captcha-service/view"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"net/http"
	"runtime/debug"
)

var Module = fx.Module("http",
	fx.Provide(
		newTemplateEngine,
		newServer,
	),
	middlewareModule,
	fx.Invoke(httpLifecycle),
)

func newTemplateEngine() *html.Engine {
	return html.NewFileSystem(http.FS(view.View), ".html")
}

func newServer(engine *html.Engine) *fiber.App {
	info, ok := debug.ReadBuildInfo()
	header := lo.Ternary(ok, "X-Captcha-"+info.Main.Version, "X-Captcha")
	app := fiber.New(
		fiber.Config{
			ViewsLayout:       "layout",
			Views:             engine,
			PassLocalsToViews: true,
			Immutable:         true,
			StreamRequestBody: true,
			ServerHeader:      header,
		},
	)

	return app
}
