package endpoint

import "github.com/gofiber/fiber/v3"

type viewController struct {
}

func newViewController() *viewController {
	return &viewController{}
}
func (v *viewController) RegisterRoute(app *fiber.App) {
	app.Get("/", v.index)
}

func (v *viewController) index(ctx fiber.Ctx) error {
	return ctx.Render(
		"number_image",
		fiber.Map{
			"Title": "数字验证码",
		},
	)
}
