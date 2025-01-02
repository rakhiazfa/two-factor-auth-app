package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AccountHandler struct{}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) FindAccount(c *gin.Context) {
	user, err := utils.ExtractUserFormContext(c)
	utils.PanicIfErr(err)

	var account dtos.AccountRes

	utils.PanicIfErr(copier.Copy(&account, user))

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}
