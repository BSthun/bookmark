package profileEndpoint

import (
	"bookmark-backend/type/payload"
	"bookmark-backend/type/response"
	"bookmark-backend/type/share"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleState(c *fiber.Ctx) error {
	// * login state
	l := c.Locals("l").(*jwt.Token).Claims.(*share.UserClaims)

	// * response
	return c.JSON(response.Success(&payload.Profile{
		UserId:    l.UserId,
		Firstname: l.Firstname,
		Lastname:  l.Lastname,
		Picture:   l.Picture,
		Email:     l.Email,
	}))
}
