package publicEndpoint

import (
	"bookmark-backend/type/payload"
	"bookmark-backend/type/response"
	"bookmark-backend/type/share"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

func (r *Handler) HandleLoginCallback(c *fiber.Ctx) error {
	// * Parse the request body into an OauthCallback payload
	body := new(payload.OauthCallback)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "unable to parse body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * exchange code for token
	token, err := r.Oauth2Config.Exchange(c.Context(), *body.Code)
	if err != nil {
		return gut.Err(false, "failed to exchange code for token", err)
	}

	// * get user info
	userInfo, err := r.OidcProvider.UserInfo(c.Context(), oauth2.StaticTokenSource(token))
	if err != nil {
		return gut.Err(false, "failed to get user info", err)
	}

	// * parse user claims
	oidcClaims := new(share.OidcClaims)
	if err := userInfo.Claims(oidcClaims); err != nil {
		return gut.Err(false, "failed to parse user claims", err)
	}

	// * generate jwt token
	claims := &share.UserClaims{
		UserId:    oidcClaims.Id,
		Firstname: oidcClaims.Firstname,
		Lastname:  oidcClaims.Lastname,
		Picture:   oidcClaims.Picture,
		Email:     oidcClaims.Email,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err := jwtToken.SignedString([]byte(*r.config.Secret))
	if err != nil {
		return gut.Err(false, "failed to sign jwt token", err)
	}

	// * set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "login",
		Value:    signedJwtToken,
		MaxAge:   86400,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.JSON(response.Success(&payload.OauthCallbackResponse{
		Login: &signedJwtToken,
	}))
}
