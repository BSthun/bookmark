package payload

type OauthRedirect struct {
	RedirectUrl *string `json:"redirectUrl" validate:"required"`
}
type OauthCallback struct {
	Code  *string `json:"code" validate:"required"`
	State *string `json:"state" validate:"required"`
}

type OauthCallbackResponse struct {
	Login *string `json:"login"`
}
