package dtos

type SignInReq struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}
