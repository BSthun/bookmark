package fiber

import (
	"bookmark-backend/common/config"
	"context"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Init(lc fx.Lifecycle, config *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:                 HandleError,
		Prefork:                      false,
		StrictRouting:                true,
		AppName:                      "Bookmark",
		StreamRequestBody:            true,
		DisablePreParseMultipartForm: true,
		Network:                      "tcp",
	})

	// * lifecycle
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen(*config.WebListen)
				if err != nil {
					gut.Fatal("unable to listen", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			// * shutdown
			_ = app.Shutdown()
			return nil
		},
	})

	return app
}
