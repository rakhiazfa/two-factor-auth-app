package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				switch err := r.(type) {
				case *utils.HttpError:
					handleHttpError(c, err)
					return
				case *utils.UniqueFieldError:
					handleUniqueFieldError(c, err)
					return
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error":   "Internal Server Error",
						"message": err.(error).Error(),
					})
				}
			}
		}()

		c.Next()
	}
}

func handleHttpError(c *gin.Context, err *utils.HttpError) {
	var validationErrors validator.ValidationErrors

	if errors.As(err.Reason, &validationErrors) {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"errors": utils.FormatValidationErrors(validationErrors),
		})
	} else {
		var errReason *string

		if err.Reason != nil {
			errReason = utils.ToPointer(err.Reason.Error())
		}

		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"message": err.Message,
			"reason":  errReason,
		})
	}
}

func handleUniqueFieldError(c *gin.Context, err *utils.UniqueFieldError) {
	c.AbortWithStatusJSON(err.StatusCode, gin.H{
		"field":   err.Field,
		"message": err.Message,
	})
}
