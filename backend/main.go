package main

import (
	"bookmark-backend/common/config"
	"bookmark-backend/common/fiber"
	"bookmark-backend/common/fiber/middleware"
	"bookmark-backend/common/swagger"
	"bookmark-backend/endpoint"
	"bookmark-backend/endpoint/profile"
	"bookmark-backend/endpoint/public"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			fiber.Init,
			swagger.Init,
			middleware.Init,
			publicEndpoint.Handle,
			profileEndpoint.Handle,
		),
		fx.Invoke(
			endpoint.Bind,
		),
	).Run()
}
