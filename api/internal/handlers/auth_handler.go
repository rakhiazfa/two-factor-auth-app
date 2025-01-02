package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dtos.SignInReq
	utils.PanicIfErr(c.ShouldBind(&req))

	refreshToken, accessToken, err := h.authService.SignIn(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dtos.SignUpReq
	utils.PanicIfErr(c.ShouldBind(&req))

	err := h.authService.SignUp(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created a new account.",
	})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed out",
	})
}
