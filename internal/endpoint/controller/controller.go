package controller

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoute(app *fiber.App)
}

func Annotation[T any](a T) interface{} {
	return fx.Annotate(
		a,
		fx.As(new(Controller)),
		fx.ResultTags(`group:"endpoint"`),
	)
}
