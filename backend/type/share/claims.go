package share

type OidcClaims struct {
	Id        *string `json:"sub"`
	Firstname *string `json:"name"`
	Lastname  *string `json:"family_name"`
	Picture   *string `json:"picture"`
	Email     *string `json:"email"`
}

type UserClaims struct {
	UserId    *string `json:"userId"`
	Firstname *string `json:"name"`
	Lastname  *string `json:"lastname"`
	Picture   *string `json:"picture"`
	Email     *string `json:"email"`
}

func (r *UserClaims) Valid() error {
	return nil
}
