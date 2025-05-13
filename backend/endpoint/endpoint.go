package endpoint

import (
	"bookmark-backend/common/config"
	"bookmark-backend/common/fiber/middleware"
	"bookmark-backend/common/swagger"
	"bookmark-backend/endpoint/profile"
	"bookmark-backend/endpoint/public"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
)

func Bind(
	app *fiber.App,
	config *config.Config,
	middleware *middleware.Middleware,
	swagger swagger.Handler,
	publicEndpoint *publicEndpoint.Handler,
	handler *profileEndpoint.Handler,
) {
	// * swagger
	app.Get("/swagger/*", swagger)

	// * api
	api := app.Group("api")
	api.Use(middleware.Cors())

	// * public group
	public := api.Group("public")
	public.Post("/login/redirect", publicEndpoint.HandleLoginRedirect)
	public.Post("/login/callback", publicEndpoint.HandleLoginCallback)

	// * profile group
	profile := api.Group("profile", middleware.Jwt(true))
	profile.Post("/state", handler.HandleState)

	// * spa
	app.Static("/", *config.WebRoot)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(filepath.Join(*config.WebRoot, "index.html"))
	})

	// * not found
	app.Use(HandleNotFound)
}

func HandleNotFound(c *fiber.Ctx) error {
	return fiber.ErrNotFound
}
