package endpoint

import (
	"captcha-service/internal/endpoint/action"
	"captcha-service/internal/endpoint/controller"
	"captcha-service/internal/endpoint/image"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("endpoint",
	image.Module,
	fx.Provide(
		//controller.Annotation(image.newCaptchaController),
		controller.Annotation(action.NewActionController),
	),
	fx.Invoke(registerController),
)

type RegisterControllerParam struct {
	fx.In
	App         *fiber.App
	Controllers []controller.Controller `group:"endpoint"`
}

func registerController(param RegisterControllerParam) {
	app, controllers := param.App, param.Controllers
	for _, c := range controllers {
		c.RegisterRoute(app)
	}
}
