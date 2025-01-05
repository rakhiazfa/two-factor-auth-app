package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

func RequiresAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := extractTokenFromHeader(c)
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

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

		sessionId, err := uuid.Parse(sub)
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		userSession, err := validateUserSession(c, db, sessionId)
		if err != nil {
			utils.PanicIfErr(
				utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err),
			)
		}

		c.Set("userSession", userSession)

		c.Next()
	}
}

func extractTokenFromHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		return "", fmt.Errorf("token is required")
	}

	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid token")
	}

	return tokenParts[1], nil
}

func validateUserSession(c *gin.Context, db *gorm.DB, sessionId uuid.UUID) (*entities.TwoFactorAuthSession, error) {
	var userSession entities.TwoFactorAuthSession

	if err := db.WithContext(c.Request.Context()).Preload("User").Preload("UserDevice").First(&userSession, "id = ?", sessionId).Error; err != nil {
		return nil, fmt.Errorf("session not found")
	}

	if userSession.ExpiresAt.Before(time.Now()) || !userSession.Verified {
		return nil, fmt.Errorf("session has expired or session has not been verified")
	}

	return &userSession, nil
}
