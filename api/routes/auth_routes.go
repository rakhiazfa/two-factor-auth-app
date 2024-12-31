package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
)

func initAuthRoutes(r *gin.RouterGroup, handler *handlers.AuthHandler) {
	group := r.Group("/auth")

	group.POST("/sign-in", handler.SignIn)
	group.POST("/sign-up", handler.SignUp)
}
