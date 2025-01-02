package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

func RequiresAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil),
			)
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil),
			)
		}

		accessToken := tokenParts[1]

		claims, err := utils.VerifyAccessToken(accessToken)
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		sub, err := claims.GetSubject()
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		userId, err := uuid.Parse(sub)
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		var user entities.User

		if err := db.Raw("SELECT * FROM users WHERE id = ?", userId).First(&user).Error; err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		c.Set("user", &user)

		c.Next()
	}
}
