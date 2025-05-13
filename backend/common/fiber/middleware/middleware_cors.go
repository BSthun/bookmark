package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (r *Middleware) Cors() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	// example: "https://www.google.com, https://www.bsthun.com, http://localhost:8080"
	origins := ""
	for i, s := range r.config.Cors {
		origins += *s
		if i < len(r.config.Cors)-1 {
			origins += ", "
		}
	}

	c := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(c)
}
