package image

import (
	"captcha-service/internal/model"
	"context"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-module/base64Captcha/driver"
	"image/color"
	"strconv"
)

func (i *ImageCaptchaController) numberImage(ctx fiber.Ctx) error {
	ctx.Query("width", "240")
	base64Driver := driver.NewDriverDigit(driver.DriverDigit{
		Width:      240,
		Height:     60,
		Length:     6,
		NoiseCount: 10,
		Source:     strconv.Itoa(gofakeit.Number(100000, 999999)),
		BgColor:    &color.RGBA{R: 0, G: 0, B: 0, A: 0}, // 背景颜色
	})

	id, content, s := base64Driver.GenerateCaptcha()
	cache := model.ImageCaptchaModel{
		ID:      id,
		Content: content,
		Answer:  s,
	}
	marshal, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	err = i.store.Set(context.Background(), id, marshal)
	if err != nil {
		return err
	}

	res, err := base64Driver.DrawCaptcha(content)

	if err != nil {
		return err
	}

	return model.OK(ctx, fiber.Map{
		"token": id,
		"data":  res.Encoder(),
	})
}

func (i *ImageCaptchaController) verifyNumber(ctx fiber.Ctx) error {
	return nil
}
