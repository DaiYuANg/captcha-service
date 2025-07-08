package image

import (
	"captcha-service/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-module/base64Captcha/driver"
	"image/color"
	"strconv"
)

func (i *CaptchaController) numberImage(ctx fiber.Ctx) error {
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

	err := i.store.Set(id, &cache)
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

type VerifyData struct {
	ID    string `json:"id"`
	Input string `json:"input"`
}

func (i *CaptchaController) verifyNumber(ctx fiber.Ctx) error {
	// 解析请求体
	var data VerifyData
	if err := ctx.Bind().Body(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "请求参数错误")
	}

	// 从 store 取验证码数据
	cache, err := i.store.Get(data.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "验证码不存在或已过期")
	}

	if data.Input == "" || data.Input != cache.Answer {
		return fiber.NewError(fiber.StatusBadRequest, "验证码错误")
	}

	// 验证通过后，删除缓存，防止重复使用
	_ = i.store.Delete(data.ID)

	// 返回成功响应
	return model.OK(ctx, fiber.Map{
		"success": true,
	})
}
