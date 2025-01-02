package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
)

func ExtractUserFormContext(c *gin.Context) (*entities.User, error) {
	rawUser, exists := c.Get("user")

	if !exists {
		return nil, NewHttpError(http.StatusInternalServerError, "Unable to extract user from request context", nil)
	}

	user, ok := rawUser.(*entities.User)
	if !ok {
		return nil, NewHttpError(http.StatusInternalServerError, "Invalid user type in request context", nil)
	}

	return user, nil
}
