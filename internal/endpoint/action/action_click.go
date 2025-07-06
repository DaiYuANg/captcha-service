package action

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v3"
	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
)

func (action *ActionCaptchaController) click(ctx fiber.Ctx) error {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
	)

	var chars []string
	for i := 0; i < 8; i++ {
		animal := gofakeit.Animal()
		runes := []rune(animal) // 处理可能的 Unicode（emoji、中文等）
		if len(runes) >= 2 {
			chars = append(chars, string(runes[:2]))
		} else {
			chars = append(chars, string(runes)) // 不足 2 个字符也加入
		}
	}

	builder.SetResources(
		click.WithChars(chars),                         // 每次变化
		click.WithFonts([]*truetype.Font{action.font}), // 静态
		click.WithBackgrounds(action.bgImage),
	)

	textCapt := builder.Make()

	captData, err := textCapt.Generate()

	if err != nil {
		return err
	}

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		return err
	}
	tBase64, err = captData.GetThumbImage().ToBase64()
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"m": mBase64,
		"t": tBase64,
	})
}
