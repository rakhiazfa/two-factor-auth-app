package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/middlewares"
	"gorm.io/gorm"
)

func initAccountRoutes(r *gin.RouterGroup, handler *handlers.AccountHandler, db *gorm.DB) {
	group := r.Group("/account", middlewares.RequiresAuth(db))

	group.GET("", handler.FindAccount)
	group.PUT("", handler.UpdateAccount)

	group.GET("/devices", handler.FindAccountDevices)
	group.DELETE("/devices/:id", handler.RemoveAccountDevice)
}
