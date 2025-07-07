package image

import (
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/gofiber/fiber/v3"
)

type CaptchaController struct {
	store *cache.Cache[[]byte]
}

func NewCaptchaController(store *cache.Cache[[]byte]) *CaptchaController {
	return &CaptchaController{
		store: store,
	}
}

func (i *CaptchaController) RegisterRoute(app *fiber.App) {
	group := app.Group("image")
	generate := group.Group("/generate")
	verify := group.Group("verify")

	generate.Get("/number", i.numberImage)
	verify.Get("/number", i.verifyNumber)

	group.Get("/image/word", i.wordImage)
	group.Get("/image/math", i.mathImage)
}
