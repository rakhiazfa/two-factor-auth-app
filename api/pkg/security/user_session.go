package security

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
)

func GetUserSession(c *gin.Context) (*entities.TwoFactorAuthSession, error) {
	rawUserSession, exists := c.Get("userSession")

	if !exists {
		return nil, fmt.Errorf("unable to extract user session from request context")
	}

	log.Println(rawUserSession)

	userSession, ok := rawUserSession.(*entities.TwoFactorAuthSession)
	if !ok {
		return nil, fmt.Errorf("invalid user session type in request context")
	}

	return userSession, nil
}
