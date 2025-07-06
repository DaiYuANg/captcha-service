package endpoint

import (
	"captcha-service/internal/endpoint/action"
	"captcha-service/internal/endpoint/image"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoute(app *fiber.App)
}

func controller[T any](a T) interface{} {
	return fx.Annotate(
		a,
		fx.As(new(Controller)),
		fx.ResultTags(`group:"endpoint"`),
	)
}

var Module = fx.Module("endpoint",
	fx.Provide(
		controller(image.NewCaptchaController),
		controller(action.NewActionController),
		controller(newViewController),
	),
	fx.Invoke(registerController),
)

type RegisterControllerParam struct {
	fx.In
	App         *fiber.App
	Controllers []Controller `group:"endpoint"`
}

func registerController(param RegisterControllerParam) {
	app, controllers := param.App, param.Controllers
	for _, controller := range controllers {
		controller.RegisterRoute(app)
	}
}
