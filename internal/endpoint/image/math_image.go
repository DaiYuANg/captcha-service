package image

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-module/base64Captcha"
	"github.com/golang-module/base64Captcha/driver"
	"github.com/golang-module/base64Captcha/store"
	"image/color"
)

func (i *ImageCaptchaController) mathImage(ctx fiber.Ctx) error {
	base64Driver := driver.NewDriverMath(driver.DriverMath{
		Width:           240,                                 // 宽度
		Height:          60,                                  // 高度
		NoiseCount:      2,                                   // 点数量
		ShowLineOptions: 0,                                   // 显示线条
		BgColor:         &color.RGBA{R: 0, G: 0, B: 0, A: 0}, // 背景颜色
	})

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
