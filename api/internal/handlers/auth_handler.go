package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/security"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AuthHandler struct {
	authService          *services.AuthService
	twoFactorAuthService *services.TwoFactorAuthService
}

func NewAuthHandler(authService *services.AuthService, twoFactorAuthService *services.TwoFactorAuthService) *AuthHandler {
	return &AuthHandler{
		authService:          authService,
		twoFactorAuthService: twoFactorAuthService,
	}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dtos.SignInReq
	utils.PanicIfErr(c.ShouldBind(&req))

	twoFactorAuthSessionId, correctNumber, err := h.authService.SignIn(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"sessionId":     twoFactorAuthSessionId,
		"correctNumber": correctNumber,
	})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dtos.SignUpReq
	utils.PanicIfErr(c.ShouldBind(&req))

	refreshToken, accessToken, err := h.authService.SignUp(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Successfully created a new account.",
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed out",
	})
}

func (h *AuthHandler) SendOptionNumbers2FA(c *gin.Context) {
	sessionId, err := uuid.Parse(c.Param("sessionId"))
	utils.PanicIfErr(err)

	err = h.twoFactorAuthService.SendOptionNumbers2FA(c.Request.Context(), sessionId)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully send number options",
	})
}

func (h *AuthHandler) Verify2FANumber(c *gin.Context) {
	sessionId, err := uuid.Parse(c.Param("sessionId"))
	utils.PanicIfErr(err)

	userSession, err := security.GetUserSession(c)
	utils.PanicIfErr(err)

	var req dtos.Verify2FAOptionReq
	utils.PanicIfErr(c.ShouldBind(&req))

	err = h.twoFactorAuthService.Verify2FANumber(c.Request.Context(), userSession.UserDevice, sessionId, &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully verify number options",
	})
}
