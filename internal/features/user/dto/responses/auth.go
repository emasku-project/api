package responses

type Login struct {
	Token string `json:"token" validate:"required"`
} // @name LoginRes

type Register struct {
	Token string `json:"token" validate:"required"`
} // @name RegisterRes

type Logout struct {
	Message string `json:"message" validate:"required"`
} // @name LogoutRes
