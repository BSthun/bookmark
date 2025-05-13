package middleware

import (
	"bookmark-backend/type/share"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

func (r *Middleware) Jwt(guard bool) fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte(*r.config.Secret),
		TokenLookup: "cookie:login",
		ContextKey:  "l",
		Claims:      &share.UserClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if !guard {
				return c.Next()
			}
			return gut.Err(false, "JWT validation failure", err)
		},
	}

	return jwtware.New(conf)
}
