package image

import (
	"captcha-service/internal/model"
	"captcha-service/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type CaptchaController struct {
	store *store.Store[*model.ImageCaptchaModel]
}

type NewCaptchaControllerParameter struct {
	fx.In
	Store *store.Store[*model.ImageCaptchaModel] `name:"image"`
}

func newCaptchaController(parameter NewCaptchaControllerParameter) *CaptchaController {
	return &CaptchaController{
		store: parameter.Store,
	}
}

func (i *CaptchaController) RegisterRoute(app *fiber.App) {
	group := app.Group("image")
	generate := group.Group("/generate")
	verify := group.Group("verify")

	generate.Get("/number", i.numberImage)
	verify.Post("/number", i.verifyNumber)

	group.Get("/image/word", i.wordImage)
	group.Get("/image/math", i.mathImage)
}
