package action

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"image"
)

type ActionCaptchaController struct {
	font    *truetype.Font
	bgImage []image.Image
}

func (action *ActionCaptchaController) RegisterRoute(app *fiber.App) {
	group := app.Group("/action")
	group.Get("/click", action.click)
	group.Get("/drag/drop", action.dragDrop)
}

func NewActionController() *ActionCaptchaController {
	font, err := fzshengsksjw.GetFont()
	if err != nil {
		panic(err)
	}
	images, err := imagesv2.GetImages()
	if err != nil {
		panic(err)
	}
	return &ActionCaptchaController{
		font:    font,
		bgImage: images,
	}
}
