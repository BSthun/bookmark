package payload

type Profile struct {
	UserId    *string `json:"userId"`
	Firstname *string `json:"name"`
	Lastname  *string `json:"lastname"`
	Picture   *string `json:"picture"`
	Email     *string `json:"email"`
}
