package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AuthHandler struct {
	validator   *utils.Validator
	authService *services.AuthService
}

func NewAuthHandler(validator *utils.Validator, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		validator:   validator,
		authService: authService,
	}
}

func (handler *AuthHandler) SignIn(c *gin.Context) {
	var req dtos.SignInReq

	utils.PanicIfErr(c.ShouldBind(&req))
	utils.PanicIfErr(handler.validator.Validate(req))

	token, err := handler.authService.SignIn(req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (handler *AuthHandler) SignUp(c *gin.Context) {
	var req dtos.SignUpReq

	utils.PanicIfErr(c.ShouldBind(&req))
	utils.PanicIfErr(handler.validator.Validate(req))

	handler.authService.SignUp(req)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created a new account.",
	})
}
