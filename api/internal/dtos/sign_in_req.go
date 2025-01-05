package dtos

type SignInReq struct {
	UsernameOrEmail string `json:"usernameOrEmail" validate:"required"`
	Password        string `json:"password" validate:"required"`
	DeviceType      string `json:"deviceType" validate:"required,max=100"`
	DeviceName      string `json:"deviceName" validate:"required,max=100"`
	DeviceToken     string `json:"deviceToken" validate:"required,max=255"`
}
