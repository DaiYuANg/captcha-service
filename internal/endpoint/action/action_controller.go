package action

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"image"
)

type CaptchaController struct {
	font    *truetype.Font
	bgImage []image.Image
}

func (action *CaptchaController) RegisterRoute(app *fiber.App) {
	group := app.Group("/action")
	group.Get("/click", action.click)
	group.Get("/drag/drop", action.dragDrop)
}

func NewActionController() *CaptchaController {
	font, err := fzshengsksjw.GetFont()
	if err != nil {
		panic(err)
	}
	images, err := imagesv2.GetImages()
	if err != nil {
		panic(err)
	}
	return &CaptchaController{
		font:    font,
		bgImage: images,
	}
}
