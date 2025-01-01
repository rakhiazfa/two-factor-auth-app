package dtos

type SignUpReq struct {
	Name                 string `json:"name" validate:"required,max=100"`
	Username             string `json:"username" validate:"required,username,max=100"`
	Email                string `json:"email" validate:"required,email,max=100"`
	Password             string `json:"password" validate:"required,min=8,max=100"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}
