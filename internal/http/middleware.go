package http

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/contrib/monitor"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/fx"
	"time"
)

var middlewareModule = fx.Module("middleware", fx.Invoke(
	registerRequestid,
	registerLogger,
	registerMonitor,
	registerHealthcheck,
	registerCompress,
	registerFavicon,
	registerHelmet,
	registerLimiter,
))

func registerLogger(app *fiber.App) {
	_ = fiberzap.NewLogger(fiberzap.LoggerConfig{
		ExtraKeys: []string{"request_id"},
	})
	app.Use(logger.New(logger.Config{}))
}

func registerHealthcheck(app *fiber.App) {
	// Provide a minimal config for startup check
	app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
}

func registerHelmet(app *fiber.App) {
	app.Use(helmet.New())
}

func registerLimiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}

func registerFavicon(app *fiber.App) {
	app.Use(favicon.New())
}

func registerCompress(app *fiber.App) {
	app.Use(compress.New())
}

func registerRequestid(app *fiber.App) {
	app.Use(requestid.New())
}

func registerMonitor(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
