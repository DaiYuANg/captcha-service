package image

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-module/base64Captcha"
	"github.com/golang-module/base64Captcha/driver"
	"github.com/golang-module/base64Captcha/store"
)

func (i *CaptchaController) wordImage(ctx fiber.Ctx) error {
	base64Driver := driver.DefaultDriverLetter

	captcha := base64Captcha.NewCaptcha(base64Driver, store.DefaultStoreMemory)
	a, b, _, err := captcha.Generate()
	if err != nil {
		return err
	}
	res, err := captcha.Driver.DrawCaptcha(b)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"token": a,
		"data":  res.Encoder(),
	})
}
