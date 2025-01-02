package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/middlewares"
	"gorm.io/gorm"
)

func InitRoutes(
	db *gorm.DB,
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
) *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(middlewares.Recovery())

	apiGroup := r.Group("/api")

	initAuthRoutes(apiGroup, authHandler, db)
	initAccountRoutes(apiGroup, accountHandler, db)

	return r
}
