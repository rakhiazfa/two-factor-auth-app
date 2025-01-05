package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/middlewares"
	"gorm.io/gorm"
)

func initAuthRoutes(r *gin.RouterGroup, handler *handlers.AuthHandler, db *gorm.DB) {
	group := r.Group("/auth")

	group.POST("/sign-in", handler.SignIn)
	group.POST("/sign-up", handler.SignUp)
	group.POST("/sign-out", middlewares.RequiresAuth(db), handler.SignOut)

	group.POST("/sessions/:sessionId/send-option-numbers", handler.SendOptionNumbers2FA)
	group.POST("/sessions/:sessionId/verify-option", middlewares.RequiresAuth(db), handler.Verify2FANumber)
}
