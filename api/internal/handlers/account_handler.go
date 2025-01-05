package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/security"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) FindAccount(c *gin.Context) {
	userSession, err := security.GetUserSession(c)
	utils.PanicIfErr(err)

	var account dtos.AccountRes

	utils.PanicIfErr(copier.Copy(&account, userSession.User))

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {}

func (h *AccountHandler) FindAccountDevices(c *gin.Context) {
	userSession, err := security.GetUserSession(c)
	utils.PanicIfErr(err)

	devices, err := h.accountService.FindAccountDevices(c.Request.Context(), userSession.User)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"devices": *devices,
	})
}

func (h *AccountHandler) RemoveAccountDevice(c *gin.Context) {}
