package dtos

type SignInReq struct {
	UsernameOrEmail string `json:"usernameOrEmail" validate:"required"`
	Password        string `json:"password" validate:"required"`
}
