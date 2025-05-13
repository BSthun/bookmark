package middleware

import (
	"bookmark-backend/type/share"
	"bookmark-backend/util/network"
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func SentryRecover(hub *sentry.Hub, c *fiber.Ctx, e *error) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(
			context.WithValue(context.Background(), sentry.RequestContextKey, c),
			err,
		)
		if eventID != nil {
			hub.Flush(2 * time.Second)
		}
		*e = fmt.Errorf("%v", err)
	}
}

func Sentry() fiber.Handler {
	return func(c *fiber.Ctx) (e error) {
		hub := sentry.CurrentHub().Clone()
		scope := hub.Scope()
		scope.SetRequest(network.ConvertRequest(c.Context()))
		scope.SetRequestBody(c.Body())

		// * configure scope
		hub.ConfigureScope(func(scope *sentry.Scope) {
			if c.Locals("l") != nil {
				claims := c.Locals("l").(*jwt.Token).Claims.(*share.UserClaims)
				scope.SetUser(sentry.User{
					ID:        *claims.UserId,
					Email:     "",
					IPAddress: c.Get("X-Forwarded-For", c.IP()),
					Username:  "",
					Name:      "",
					Data:      nil,
				})
			}
		})

		sentryCtx := context.TODO()
		sentryCtx = sentry.SetHubOnContext(sentryCtx, hub)

		// * start a transaction
		span := sentry.StartSpan(sentryCtx, "http.server")
		span.Description = fmt.Sprintf("%s (%s)", c.Path(), c.Method())
		sentryCtx = context.WithValue(sentryCtx, "span", span)
		c.Locals("sentry", sentryCtx)

		defer SentryRecover(hub, c, &e)
		return c.Next()
	}
}
