package image

import (
	"captcha-service/internal/endpoint/controller"
	"captcha-service/internal/model"
	"captcha-service/internal/store"
	"go.uber.org/fx"
)

var Module = fx.Module("image",
	fx.Provide(
		fx.Annotate(
			newStore,
			fx.ResultTags(`name:"image"`),
		),
		controller.Annotation(newCaptchaController),
	),
)

func newStore(parameter store.NewStoreParameter) *store.Store[*model.ImageCaptchaModel] {
	return store.NewStore[*model.ImageCaptchaModel](parameter)
}
